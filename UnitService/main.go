package main

import (
	DB "UnitService/pkg/DB/MySQLDB"
	"UnitService/pkg/Swagger"
	configs "UnitService/pkg/config"
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	serverParameters := initServerParameters()
	initLogFile()
	err := initDB()
	if err != nil {
		return
	}
	defer DB.DisconectDB()
	err = DB.MakeMigrationDB()
	if err != nil {
		return
	}
	serverUrl := serverParameters.ServeRESTAddress
	killSignalChan := getKillSignalChan()
	srv := startServer(serverUrl)
	waitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())
}

func initDB() (error) {
	db := DB.GetDBInstance()
	if db == nil {
		return errors.New("Error init DB")
	}
	return nil
}

func initServerParameters() (*configs.Config) {
	serverParameters, err := configs.ParseEnv()
	if err != nil {
		log.Fatal("Cannot init server parameters!")
	}
	return serverParameters
}

func initLogFile() {
	log.SetFormatter(&log.JSONFormatter{})
	file, err := os.OpenFile("my.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err == nil {
		log.SetOutput(file)
		defer file.Close()
	}
}

func startServer(serverUrl string) *http.Server {
	router := Swagger.NewRouter()
	log.Fatal(http.ListenAndServe(serverUrl, router))
	srv := &http.Server{Addr: serverUrl, Handler: router}
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	return srv
}

func getKillSignalChan() chan os.Signal {
	osKillSignalChan := make(chan os.Signal, 1)
	signal.Notify(osKillSignalChan, os.Interrupt, syscall.SIGTERM)
	return osKillSignalChan
}

func waitForKillSignal(killSignalChan <-chan os.Signal) {
	killSignal := <-killSignalChan
	switch killSignal {
	case os.Interrupt:
		log.Info("got SIGINT...")
	case syscall.SIGTERM:
		log.Info("got SIGTERM...")
	}
}
