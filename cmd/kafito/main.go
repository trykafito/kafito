package main

import (
	"flag"
	"fmt"

	"github.com/trykafito/kafito/config"
	"github.com/trykafito/kafito/internal/server"
	"github.com/trykafito/kafito/pkg/database"
	"github.com/trykafito/kafito/pkg/logger"
)

var configPath = flag.String("c", "./config.json", "config file path")

func main() {
	flag.Parse()

	c, err := config.Read(*configPath)
	if err != nil {
		logger.Panic(err)
	}

	if err := database.Connect(c.Mongo.Host, c.Mongo.DB, c.Mongo.User, c.Mongo.Password); err != nil {
		logger.Panic(err)
	}

	fmt.Println(`
 _  __      __ _ _        
| |/ /__ _ / _(_) |_ ___  
| ' // _' | |_| | __/ _ \ 
|   \ (_| |  _| | || (_) |
|_|\_\__,_|_| |_|\__\___/ 
	`)

	if err := server.Start(c.Port, c.SecretKey); err != nil {
		logger.Panic(err)
	}
}
