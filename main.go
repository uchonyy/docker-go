package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"net/http"
	"flag"
)

var conf *Config
func main(){
	configFile := flag.String("conf", "docker-go.yaml", "Path to config file")
	flag.Parse()
	log.Info("Init app")
	config, err := LoadConfig(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	conf = config
	e :=echo.New()
	e.HideBanner = true
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/", index)
	log.Info("Starting http server at :8080")
	log.Fatal(e.Start(":8080"))
}

func index (c echo.Context) error  {
	return c.String( http.StatusOK, "docker-go app is started: " + conf.Value)
}

type Config struct {
	Value string
}

func LoadConfig(configFile string) (*Config, error){
	config := &Config{}
	if err := configor.Load(config, configFile); err != nil {
		return nil, err
	}
	return config, nil
}