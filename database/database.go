package database

type DB interface {
	Connection()
	Disconnect()
	// Insert()
	// Select(string) map[string]interface{}
}
