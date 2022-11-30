package singleton

import (
	"fmt"
	"sync"
)

type Configuration interface {
	GetHost() string
	GetPort() string
}

// One
var once sync.Once
var config Configuration

type singletonConfiguration struct {
	host string
	port string
}

func (s *singletonConfiguration) GetHost() string {
	return s.host
}
func (s *singletonConfiguration) GetPort() string {
	return s.port
}

func GetSingletonConfiguration() Configuration {
	once.Do(func() {
		config = &singletonConfiguration{
			host: "host",
			port: "port",
		}
	})
	return config
}

func main() {
	s := GetSingletonConfiguration()
	fmt.Println(s.GetHost())
	fmt.Println(s.GetPort())
}
