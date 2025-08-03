package config

type Config struct {
	HTTPServer HTTPServer ``
	Database   Database
}

type HTTPServer struct {
	Address string
}

type Database struct {
	URL string
}
