package dal

import(
	"github.com/gocql/gocql"
	"message_backup/resources"
)

var session *gocql.Session

func initiateCassandra(){
	cluster := gocql.NewCluster("172.23.16.14", "172.23.16.15", "172.23.16.16")
	cluster.Keyspace = "messagemicroservice"
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
