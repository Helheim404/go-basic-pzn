package database

var connection string

func init() {
	connection = "mysql connection"
}

func GetDatabase() string {
	return connection
}
