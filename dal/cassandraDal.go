package dal

import(
	"github.com/gocql/gocql"
	"time"
)

var session *gocql.Session

func initiateCassandra(){
	cluster := gocql.NewCluster("172.23.16.14", "172.23.16.15", "172.23.16.16")
	cluster.Keyspace = "messagemicroservice"
	cluster.ProtoVersion = 4
	cluster.Timeout = 10000 * time.Millisecond
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}

func init() {
	initiateCassandra()
}

func PushinCass(batch *gocql.Batch) error{
	err := session.ExecuteBatch(batch)
	return err
}

func QueryExecute(query gocql.Query) error{
	err := query.Exec()
	return err
}
