package main

import (
	"UnitService/DB"
	Swagger "UnitService/Swagger/go"
	Equipment "UnitService/pkg/Equipment/infrastructure/transport"
	Unit "UnitService/pkg/Unit/infrastructure/transport"
	"context"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	serverParameters := initServerParameters()
	initLogFile()

	db, err := DB.InitDB()
	if err != nil {
		return
	}
	defer db.DisconectDB()

	serverUrl := serverParameters.ServeRESTAddress
	killSignalChan := getKillSignalChan()
	srv := startServer(db, serverUrl)
	waitForKillSignal(killSignalChan)
	srv.Shutdown(context.Background())
}

func initServerParameters() (*Config) {
	serverParameters, err := ParseEnv()
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

func InitUnitHendlerFunc(router *mux.Router, connection *DB.Connection) (*mux.Router) {
	unitServer := Unit.CreateUnitServer(connection)

	unitHandlerFuncs := map[string]http.HandlerFunc {
		"UnitGet" : unitServer.UnitGet,
		"UnitPost" : unitServer.UnitPost,
		"UnitUnitIdDelete" : unitServer.UnitUnitIdDelete,
		"UnitUnitIdPut" : unitServer.UnitUnitIdPut,
		"UnitUnitIdGet" : unitServer.UnitUnitIdGet,
	}

	for name, unitHendlerFunc := range unitHandlerFuncs {
		router.GetRoute(name).Handler(unitHendlerFunc)
	}

	return router
}

func InitEquipmentHendlerFunc(router *mux.Router, connection *DB.Connection) (*mux.Router) {
	equipmentServer := Equipment.CreateEquipmentServer(connection)

	equipmentHandlerFuncs := map[string]http.HandlerFunc {
		"EquipmentEquipmentIdDelete" : equipmentServer.EquipmentEquipmentIdDelete,
		"EquipmentEquipmentIdGet" : equipmentServer.EquipmentEquipmentIdGet,
		"EquipmentEquipmentIdPut" : equipmentServer.EquipmentEquipmentIdPut,
		"EquipmentGet" : equipmentServer.EquipmentGet,
		"EquipmentPost" : equipmentServer.EquipmentPost,
	}

	for name, equipmentHendlerFunc := range equipmentHandlerFuncs {
		router.GetRoute(name).Handler(equipmentHendlerFunc)
	}

	return router
}

func startServer(connection *DB.Connection, serverUrl string) *http.Server {
	router := Swagger.NewRouter()
	router = InitUnitHendlerFunc(router, connection)
	router = InitEquipmentHendlerFunc(router, connection)
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



