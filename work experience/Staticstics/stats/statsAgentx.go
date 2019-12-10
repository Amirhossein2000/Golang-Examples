package stats

import (
	"bytes"
	"log"
	"strconv"
	"time"

	"github.com/orian/counters"
	"github.com/orian/counters/global"
	"github.com/posteo/go-agentx"
	"github.com/posteo/go-agentx/pdu"
	"github.com/posteo/go-agentx/value"
)

// OIDPREFIX is unique for this product of parspooyesh
const OIDPREFIX = "1.3.6.1.4.1.51670.4" // NOTE: this prefix regstered for parspooyesh and 4 is rebin-rad-sniffer

type statAgentx struct {
	net     string
	addr    string
	started bool
	client  *agentx.Client
}

type agentxHandler struct {
	counters []counters.Counter
	peerIPs  []string
	prefix   value.OID
}

var agent *statAgentx

// Get tries to find the provided oid and returns the corresponding value.
func (a *agentxHandler) Get(oid value.OID) (value.OID, pdu.VariableType, interface{}, error) {
	log.Println("Agent request oid %s", oid)

	if len(oid.CommonPrefix(a.prefix)) != len(a.prefix) {
		log.Println("Invalid agentx request oid: %s prefix: %s", oid, a.prefix)
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}

	lastSubID := int(oid[len(oid)-1])

	if len(oid) == len(a.prefix)+1 {
		// normal counters

		if lastSubID < 0 || lastSubID > len(a.counters)-1 {
			log.Println("Invalid agentx request oid: %s prefix: %s: Bad Last Index", oid, a.prefix)
			return nil, pdu.VariableTypeNoSuchObject, nil, nil
		}

		return oid, pdu.VariableTypeCounter64, uint64(a.counters[lastSubID].Value()), nil
	}

	log.Println("Invalid agentx request oid: %s prefix: %s: Bad OID Length", oid, a.prefix)
	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}

// GetNext tries to find the value that follows the provided oid and returns it.
func (a *agentxHandler) GetNext(from value.OID, includeFrom bool, to value.OID) (value.OID, pdu.VariableType, interface{}, error) {
	if a.counters == nil {
		return nil, pdu.VariableTypeNoSuchObject, nil, nil
	}

	fromOID, toOID := from.String(), to.String()

	lastSubID := 0
	if len(from) == len(a.prefix)+1 {
		lastSubID = int(from[len(from)-1])
	}

	for idx := lastSubID; idx <= len(a.counters); idx++ {
		oid := a.prefix.String() + "." + strconv.Itoa(idx)
		if oidWithin(oid, fromOID, includeFrom, toOID) {
			return a.Get(value.MustParseOID(oid))
		}
	}

	return nil, pdu.VariableTypeNoSuchObject, nil, nil
}

func oidWithin(oid string, from string, includeFrom bool, to string) bool {
	oidBytes, fromBytes, toBytes := []byte(oid), []byte(from), []byte(to)

	fromCompare := bytes.Compare(fromBytes, oidBytes)
	toCompare := bytes.Compare(toBytes, oidBytes)

	return (fromCompare == -1 || (fromCompare == 0 && includeFrom)) && (toCompare == 1)
}

func initHandler() *agentxHandler {
	prefix := value.MustParseOID(OIDPREFIX)
	cntrs := []counters.Counter{
		global.GetCounter("sniffer_packets_total"),
		global.GetCounter("sniffer_packets_droped"),
		global.GetCounter("producer_packets_wrote"),
	}

	return &agentxHandler{prefix: prefix, counters: cntrs}
}

// InitAgentx for initiate agentx and start it.
func InitAgentx() {
	net := "unix"
	addr := "/var/agentx/master"

	if net == "" || addr == "" {
		log.Println("Disabling Agent X support, net: %s addr: %s", net, addr)
		return
	}

	client := &agentx.Client{
		Net:               net,
		Address:           addr,
		Timeout:           1 * time.Minute,
		ReconnectInterval: 1 * time.Second,
	}

	agent = &statAgentx{net: net, addr: addr, client: client}

	go agent.start()

}

// ShutdownAgentx gracefully shutdown Agentx
func ShutdownAgentx() {
	if agent != nil && agent.started {
		agent.client.Close()
	}
}

func (s *statAgentx) start() {
	client := s.client

	if err := client.Open(); err != nil {
		log.Println("Could not open agentx connection net: %s addr: %s err: %s", s.net, s.addr, err)
		return
	}

	session, err := client.Session()
	if err != nil {
		log.Println("Could not open agentx session net: %s addr: %s err: %s", s.net, s.addr, err)
		return
	}

	handler := initHandler()
	session.Handler = handler

	log.Println("Starting statistic agentx on %s:%s", s.net, s.addr)
	s.started = true

	if err := session.Register(127, handler.prefix); err != nil {
		log.Println("Could not open agentx session net: %s addr: %s err: %s", s.net, s.addr, err)
	}

}
