

type dbInterface interface {
	connectDatabase(connectionString string) (*dbInterface, error)
	query(query string) []byte
	closeConnection()
}