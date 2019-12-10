package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"math"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64
	var errCount uint64
	var qrun uint64
	// connect to the cluster
	cluster := gocql.NewCluster("172.16.4.5")
	cluster.NumConns = 4
	cluster.Keyspace = "tci_log"
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = time.Second

	//if err := session.Query(`insert into ipdr( username, mac , time , protonum , invalidip , invalidport , validip , validport , dstip , dstport , loc1 , loc2 , apn , imsi  , imei , operatorname , servicenum , collectorname ) values( '4137634390', '28285dc772a0', 1557916756,  '6',  '',  '', '1.1.1.1',  '46198', '172.65.37.170',  '443',  '',  '',  '', ' ',  '', 'TCI', 3,  'AzarbayjanEas t' );`).Exec(); err != nil {
	//	log.Println(err)
	//}

	//ipChannel := make(chan string, 1024)
	//go Hosts("192.179.1.0/8", ipChannel)
	//go Hosts("169.179.1.0/8", ipChannel)
	//go Hosts("178.179.1.0/8", ipChannel)

	//fmt.Println(SubnetCount("192.179.1.0/8"))
	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 4; i++ {
		go func() {

			defer session.Close()

			for {
				//strQuery := "BEGIN BATCH "
				//for i := 0; i < 32; i++ {
				//strQuery += ` insert into  ipdr2(id, username, mac , time , protonum , invalidip , invalidport , validip , validport , dstip , dstport , loc1 , loc2 , apn , imsi , imei , operatorname , servicenum , collectorname ) values(uuid(), '4137634390', '28285dc772a0', 1557916756,  '6',  '',  '', '84.47.255.134',  '46198', '172.65.37.170',  '443',  '',  '',  '', '',  '', 'TCI', 3,  'AzarbayjanEas t') `
				//strQuery=` insert into ipdr( username, mac , time , protonum , invalidip , invalidport , validip , validport , dstip , dstport , loc1 , loc2 , apn , imsi , imei , operatorname , servicenum , collectorname ) values( '4137634390', '28285dc772a0', `+fmt.Sprint(time.Now().UnixNano())+`,  '6',  '',  '', '84.47.255.134',  '46198', '`+<-ipChannel+`"',  '443',  '',  '',  '', '',  '', 'abc', 3,  'AzarbayjanEast') `
				//strQuery=` insert into ipdr( username, mac , time , protonum , invalidip , invalidport , validip , validport , dstip , dstport , loc1 , loc2 , apn , imsi , imei , operatorname , servicenum , collectorname ) values(?, ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ?)`, "4137634390", "28285dc772a0", time.Now().UnixNano(),  "6",  "",  "", "84.47.255.134",  "46198", <-ipChannel,  "443",  "", "",  "", "",  "", "abc", 3,  "AzarbayjanEast")
				//strQuery="insert into  ipdr( username, mac , time , protonum , invalidip , invalidport , validip , validport , dstip , dstport , loc1 , loc2 , apn , imsi , imei , operatorname , servicenum , collectorname ) values( '4137634390', '28285dc772a0', 2557916756,  '6',  '',  '', '84.47.255.134',  '46198', '172.65.37.170',  '443',  '',  '',  '', '',  '', 'abc', 3,  'AzarbayjanEast' ) "
				//atomic.AddUint64(&ops, 1)
				//}
				//strQuery += " APPLY BATCH ;"
				//if err := session.Query(`insert into ipdr( username, mac , time , protonum , invalidip , invalidport , validip , validport , dstip , dstport , loc1 , loc2 , apn , imsi , imei , operatorname , servicenum , collectorname ) values(?, ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ?)`, "4137634390", "28285dc772a0", time.Now().UnixNano(),  "6",  "",  "", "1.1.1.1",  "46198", "172.65.37.170",  "443",  "", "",  "", "",  "", "abc", 3,  "AzarbayjanEast").Exec(); err != nil {
				//	atomic.AddUint64(&errCount, 1)
				//	log.Println(err)
				//}

				if err := session.Query(`insert into ipdr(partitionKey,id, username, mac , protonum , invalidip , invalidport , validip , validport , dstip , dstport , loc1 , loc2 , apn , imsi , imei , operatorname) values(? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ?)`,
					"partitionKey", gocql.TimeUUID(), "Username", "Mac", "ProtoNum", "InvalidIp", "InvalidPort", "1.1.1.1", "ValidPort", "DstIp", "DstPort", "Loc1", "Loc2", "Apn", "Imsi", "Imei", "OperatorName").Exec(); err != nil {
					atomic.AddUint64(&errCount, 1)
					log.Println(err)
				}

				atomic.AddUint64(&qrun, 1)
			}
		}()
	}

	timer2 := time.NewTimer(100 * time.Second)
	<-timer2.C
	fmt.Println(atomic.LoadUint64(&ops))
	fmt.Println(atomic.LoadUint64(&errCount))
	fmt.Println(atomic.LoadUint64(&qrun))
}

func Hosts(cidr string, cs chan<- string) {
	ip, ipnet, _ := net.ParseCIDR(cidr)

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		cs <- ip.String()
	}
}
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func SubnetCount(ipStrRange string) int {
	rangeNumber, _ := strconv.Atoi(strings.Split(ipStrRange, "/")[1])
	return int(math.Pow(2, float64(32-rangeNumber)))
}
