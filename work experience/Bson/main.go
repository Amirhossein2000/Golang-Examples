package main

import (
	"fmt"
	"io"
)

type DemuxOut interface {
	Write([]byte) (int, error)
	Close() error
	Sum64() (uint64, bool)
}

type Demultiplexer struct {
	In io.Reader
	//TODO wrap up these three into a structure
	outs               map[string]DemuxOut
	lengths            map[string]int64
	currentNamespace   string
	buf                [MaxBSONSize]byte
	NamespaceChan      chan string
	NamespaceErrorChan chan error
	NamespaceStatus    map[string]int
}

func main() {
	reader, err := getArchiveReader("hW.gz")
	if err != nil {
		panic(err)
	}
	fmt.Println(FileValidation(reader))

	parser := Parser{In: reader}
	parser.ReadAllBlocks()

}

func (demux *Demultiplexer) Run() error {
	parser := Parser{In: demux.In}
	err := parser.ReadAllBlocks(demux)
	if len(demux.outs) > 0 {
		log.Logvf(log.Always, "demux finishing when there are still outs (%v)", len(demux.outs))
	}

	log.Logvf(log.DebugLow, "demux finishing (err:%v)", err)
	return err
}

// HeaderBSON is part of the ParserConsumer interface and receives headers from parser.
// Its main role is to implement opens and EOFs of the embedded stream.
func (demux *Demultiplexer) HeaderBSON(buf []byte) error {
	colHeader := NamespaceHeader{}
	err := bson.Unmarshal(buf, &colHeader)
	if err != nil {
		return newWrappedError("header bson doesn't unmarshal as a collection header", err)
	}
	log.Logvf(log.DebugHigh, "demux namespaceHeader: %v", colHeader)
	if colHeader.Collection == "" {
		return newError("collection header is missing a Collection")
	}
	demux.currentNamespace = colHeader.Database + "." + colHeader.Collection
	if _, ok := demux.outs[demux.currentNamespace]; !ok {
		if demux.NamespaceStatus[demux.currentNamespace] != NamespaceUnopened {
			return newError("namespace header for already opened namespace")
		}
		demux.NamespaceStatus[demux.currentNamespace] = NamespaceOpened
		if demux.NamespaceChan != nil {
			demux.NamespaceChan <- demux.currentNamespace
			err := <-demux.NamespaceErrorChan
			if err == io.EOF {
				// if the Prioritizer sends us back an io.EOF then it's telling us that
				// it's finishing and doesn't need any more namespace announcements.
				close(demux.NamespaceChan)
				demux.NamespaceChan = nil
				return nil
			}
			if err != nil {
				return newWrappedError("failed arranging a consumer for new namespace", err)
			}
		}
	}
	if colHeader.EOF {
		if rcr, ok := demux.outs[demux.currentNamespace].(*RegularCollectionReceiver); ok {
			rcr.err = io.EOF
		}
		demux.outs[demux.currentNamespace].Close()
		demux.NamespaceStatus[demux.currentNamespace] = NamespaceClosed
		length := int64(demux.lengths[demux.currentNamespace])
		crcUInt64, ok := demux.outs[demux.currentNamespace].Sum64()
		if ok {
			crc := int64(crcUInt64)
			if crc != colHeader.CRC {
				return fmt.Errorf("CRC mismatch for namespace %v, %v!=%v",
					demux.currentNamespace,
					crc,
					colHeader.CRC,
				)
			}
			log.Logvf(log.DebugHigh,
				"demux checksum for namespace %v is correct (%v), %v bytes",
				demux.currentNamespace, crc, length)
		} else {
			log.Logvf(log.DebugHigh,
				"demux checksum for namespace %v was not calculated.",
				demux.currentNamespace)
		}
		delete(demux.outs, demux.currentNamespace)
		delete(demux.lengths, demux.currentNamespace)
		// in case we get a BSONBody with this block,
		// we want to ensure that that causes an error
		demux.currentNamespace = ""
	}
	return nil
}

// End is part of the ParserConsumer interface and receives the end of archive notification.
func (demux *Demultiplexer) End() error {
	log.Logvf(log.DebugHigh, "demux End")
	var err error
	if len(demux.outs) != 0 {
		openNss := []string{}
		for ns := range demux.outs {
			openNss = append(openNss, ns)
			if rcr, ok := demux.outs[ns].(*RegularCollectionReceiver); ok {
				rcr.err = newError("archive io error")
			}
			demux.outs[ns].Close()
		}
		err = newError(fmt.Sprintf("archive finished but contained files were unfinished (%v)", openNss))
	} else {
		for ns, status := range demux.NamespaceStatus {
			if status != NamespaceClosed {
				err = newError(fmt.Sprintf("archive finished before all collections were seen (%v)", ns))
			}
		}
	}

	if demux.NamespaceChan != nil {
		close(demux.NamespaceChan)
	}

	return err
}

// BodyBSON is part of the ParserConsumer interface and receives BSON bodies from the parser.
// Its main role is to dispatch the body to the Read() function of the current DemuxOut.
func (demux *Demultiplexer) BodyBSON(buf []byte) error {
	if demux.currentNamespace == "" {
		return newError("collection data without a collection header")
	}

	demux.lengths[demux.currentNamespace] += int64(len(buf))

	out, ok := demux.outs[demux.currentNamespace]
	if !ok {
		return newError("no demux consumer currently consuming namespace " + demux.currentNamespace)
	}
	_, err := out.Write(buf)
	return err
}

// Open installs the DemuxOut as the handler for data for the namespace ns
func (demux *Demultiplexer) Open(ns string, out DemuxOut) {
	// In the current implementation where this is either called before the demultiplexing is running
	// or while the demutiplexer is inside of the NamespaceChan NamespaceErrorChan conversation
	// I think that we don't need to lock outs, but I suspect that if the implementation changes
	// we may need to lock when outs is accessed
	log.Logvf(log.DebugHigh, "demux Open")
	if demux.outs == nil {
		demux.outs = make(map[string]DemuxOut)
		demux.lengths = make(map[string]int64)
	}
	demux.outs[ns] = out
	demux.lengths[ns] = 0
}
