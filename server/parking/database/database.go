package database

type Database struct {
	host string
	port int
}

func NewDatabase(host string, port int) *Database {
	return &Database{
		host: host,
		port: port,
	}
}

func (d Database) CheckUserLogin(login, password string) (int, error) {
	return 1, nil
}
