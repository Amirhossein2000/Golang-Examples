package main

import (
	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "tci_log"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert a tweet
	//if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
	//	"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
	//	panic(err)
	//}

	if err := session.Query(`insert into ipdr( username, mac , time , protonum , invalidip , invalidport , validip , validport , dstip , dstport , loc1 , loc2 , apn , imsi  , imei , operatorname , servicenum , collectorname ) values( '4137634390', '28285dc772a0', 1557916756,  '6',  '',  '', '1.1.1.1',  '46198', '172.65.37.170',  '443',  '',  '',  '', ' ',  '', 'TCI', 3,  'AzarbayjanEas t' )`).Exec(); err != nil {
		panic(err)
	}

	//var id gocql.UUID
	//var text string

	/* Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search */
	//if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
	//	"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
	//	panic(err)
	//}
	//fmt.Println("Tweet:", id, text)
	//
	//// list all tweets
	//iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
	//for iter.Scan(&id, &text) {
	//	fmt.Println("Tweet:", id, text)
	//}
	//if err := iter.Close(); err != nil {
	//	panic(err)
	//}
}
