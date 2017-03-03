package dal

import(
	"github.com/gocql/gocql"
	"message_backup/resources"
)

var session *gocql.Session

func initiateCassandra(){
	cluster := gocql.NewCluster(resources.CASSANDRA_SERVERS)
	cluster.Keyspace = "demo"
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}

func PushinCass(batch *gocql.Batch) error{
	initiateCassandra()
	err := session.ExecuteBatch(batch)
	return err
}
