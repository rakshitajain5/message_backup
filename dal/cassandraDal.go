package dal

import(
	"github.com/gocql/gocql"
)

var session *gocql.Session

func initiateCassandra(){
	cluster := gocql.NewCluster("172.23.115.14","172.23.115.15","172.23.115.16")
	cluster.Keyspace = "messagemicroservice"
	cluster.ProtoVersion = 4
	//cluster.Timeout = 600*time.Microsecond
	cluster.Consistency = gocql.One

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
}


func init(){
	initiateCassandra()
}

func PushinCass(batch *gocql.Batch) error{

	err := session.ExecuteBatch(batch)
	return err
}
