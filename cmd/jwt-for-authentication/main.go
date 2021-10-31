package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nogopy/jwt-for-authentication/config/database"
	env "github.com/nogopy/jwt-for-authentication/config/environment"
	"github.com/nogopy/jwt-for-authentication/config/infrastructure"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
}

func main() {
	log.Infof("Starting server on env => " + env.GetConfiguration().AppEnv)

	t := time.Now()
	fmt.Println("Location : ", t.Location(), " Time : ", t) // local time

	db := database.NewDatabase()
	conn, err := db.DB()

	if env.GetConfiguration().AppEnv != "LOCAL" {
		conn.SetMaxOpenConns(100)
		conn.SetMaxIdleConns(20)
		conn.SetConnMaxLifetime(300 * time.Second)
	}

	if err != nil {
		panic(err)
	}
	defer conn.Close()

	if env.GetConfiguration().AppEnv != "LOCAL" {
		gin.SetMode(gin.ReleaseMode)
	}

	ginEngine := infrastructure.NewGinEngine()

	infrastructure.SetupHttpServer(db, ginEngine)

	err = ginEngine.Run()

	if err != nil {
		panic(err)
	}
}
