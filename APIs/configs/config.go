package configs

import "fmt"

var configs *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JWTExperesIn  int
}

func LoadConfig(path string) (*conf, error) {
	return nil, fmt.Errorf("Fail to load configs")
}
