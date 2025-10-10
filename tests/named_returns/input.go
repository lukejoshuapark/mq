package namedreturns

type Database interface {
	Query(sql string) (result string, err error)
	Execute(cmd string) (rowsAffected int, err error)
	Connect() (conn *Connection, err error)
}

type Connection struct {
	Host string
	Port int
}
