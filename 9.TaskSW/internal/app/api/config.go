package api

//General instance for API server of REST application
type Congig struct {
	//Port
	BindAddr string `toml:"bind_addr"`
	//Logger level
	LoggerLevel string `toml:"logger_level"`
}

func NewConfig() *Congig {
	return &Congig{
		BindAddr:    "8080",
		LoggerLevel: "debug",
	}
}
