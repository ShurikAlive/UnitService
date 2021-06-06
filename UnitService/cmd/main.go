package main

import (
	"UnitService/pkg/common/infrastructure"
	Equipment "UnitService/pkg/equipment/infrastructure/transport"
	Unit "UnitService/pkg/unit/infrastructure/transport"
	Swagger "UnitService/swagger/go"
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
	initLog()

	db, err := infrastructure.InitDB(serverParameters.DBType, serverParameters.DBUsername, serverParameters.DBPassword, serverParameters.DBName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.DisconectDB()
	err = db.MakeMigrationDB(serverParameters.DBMigrationsPath)
	if err != nil {
		log.Fatal(err)
		return
	}

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

func initLog() {
	log.SetFormatter(&log.JSONFormatter{})
}

func InitUnitHendlerFunc(router *mux.Router, connection *infrastructure.Connection) (*mux.Router) {
	unitServer := Unit.CreateUnitServer(connection)

	unitHandlerFuncs := map[string]http.HandlerFunc {
		"UnitGet" : unitServer.UnitGet,
		"UnitPost" : unitServer.UnitPost,
		"UnitUnitIdDelete" : unitServer.UnitIdDelete,
		"UnitUnitIdPut" : unitServer.UnitIdPut,
		"UnitUnitIdGet" : unitServer.UnitIdGet,
	}

	for name, unitHendlerFunc := range unitHandlerFuncs {
		router.GetRoute(name).Handler(unitHendlerFunc)
	}

	return router
}

func InitEquipmentHendlerFunc(router *mux.Router, connection *infrastructure.Connection) (*mux.Router) {
	equipmentServer := Equipment.CreateEquipmentServer(connection)

	equipmentHandlerFuncs := map[string]http.HandlerFunc {
		"EquipmentEquipmentIdDelete" : equipmentServer.EquipmentIdDelete,
		"EquipmentEquipmentIdGet" : equipmentServer.EquipmentIdGet,
		"EquipmentEquipmentIdPut" : equipmentServer.EquipmentIdPut,
		"EquipmentGet" : equipmentServer.EquipmentGet,
		"EquipmentPost" : equipmentServer.EquipmentPost,
	}

	for name, equipmentHendlerFunc := range equipmentHandlerFuncs {
		router.GetRoute(name).Handler(equipmentHendlerFunc)
	}

	return router
}

func startServer(connection *infrastructure.Connection, serverUrl string) *http.Server {
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



