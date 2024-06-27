package configuration

type Configuration struct {
	Database *Database
	Server   *Server
}

type Database struct {
	Host         string `json:"host"`
	Port         uint16 `json:"port"`
	DatabaseName string `json:"databaseName"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	SSLMode      string `json:"sslMode"`
	TimeZone     string `json:"timeZone"`
}

type Server struct {
	Port uint16 `json:"port"`
}
