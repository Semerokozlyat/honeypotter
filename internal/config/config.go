package config

type Config struct {
	HTTPServer     HTTPServer
	PacketCapturer PacketCapturer
	Database       Database
}

type HTTPServer struct {
	Address string
}

type Database struct {
	URL string
}

type PacketCapturer struct {
	InterfaceName string
}
