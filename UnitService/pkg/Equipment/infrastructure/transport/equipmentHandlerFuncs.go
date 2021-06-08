package EquipmentTransport

import (
	DB "UnitService/pkg/common/infrastructure"
	App "UnitService/pkg/equipment/app"
	MySqlDB "UnitService/pkg/equipment/infrastructure/db"
	Model "UnitService/pkg/equipment/model"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"io"
	"io/ioutil"
	"net/http"
	Roster "UnitService/pkg/equipment/infrastructure/transport/services/roster"
)

type EquipmentServer struct {
	app        App.EquipmentApp
	formatter   JsonFormatter
}

func CreateEquipmentServer(connection *DB.Connection, channelRebbitMQ *amqp.Channel, queueRebbitMQ amqp.Queue) *EquipmentServer {

	equipmentServer := new(EquipmentServer)
	db := MySqlDB.CreateMySQLDB(connection)
	rosterService := Roster.CreateRosterServices(channelRebbitMQ, queueRebbitMQ)
	equipmentServer.app = App.CreateEquipmentApp(db, rosterService)
	equipmentServer.formatter = CreateJSONFormatter()
	return equipmentServer
}

func (s *EquipmentServer) getErrorCode(err error) int {
	code := http.StatusInternalServerError
	switch err {
	case App.ErrorEquipmentNotFound:
		code = http.StatusBadRequest

	case MySqlDB.ErrorEmptyConnection:
		code = http.StatusBadRequest
	case MySqlDB.ErrorInitConnection:
		code = http.StatusBadRequest
	case MySqlDB.ErrorRecordNotFound:
		code = http.StatusBadRequest

	case Model.InvalidEquipmentCost:
		code = http.StatusBadRequest
	case Model.InvalidEquipmentAmmo:
		code = http.StatusBadRequest
	case Model.InvalidEquipmentLimitOnTeam:
		code = http.StatusBadRequest
	case Model.InvalidEquipmentId:
		code = http.StatusBadRequest
	case Model.InvalidEquipmentLimitOnUnit:
		code = http.StatusBadRequest
	case Model.InvalidEquipmentName:
		code = http.StatusBadRequest
	}

	return code
}

func (s *EquipmentServer) EquipmentIdDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["equipmentId"]

	deleteId, err := s.app.DeleteByIdApp(id)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	idJSON:= s.formatter.ConvertIdToJSON(deleteId)

	fmt.Fprintf(w,idJSON)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (s *EquipmentServer) EquipmentIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["equipmentId"]

	equipment, err:= s.app.GetEquipmentById(id)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	b, err := s.formatter.ConvertEquipmentAppDataToJSON(equipment)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, b)
}

func (s *EquipmentServer) EquipmentIdPut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["equipmentId"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	equipmentEdit, err := s.formatter.ConvertJsonToEquipmentEditAppData(b)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateId, err := s.app.UpdateEquipmentApp(id, equipmentEdit)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	idJSON:= s.formatter.ConvertIdToJSON(updateId)

	fmt.Fprintf(w,idJSON)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (s *EquipmentServer) EquipmentGet(w http.ResponseWriter, r *http.Request) {
	equipments, err:= s.app.GetAllEquipment()

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	b, err := s.formatter.ConvertAllEquipmentAppDataToJSON(equipments)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, b)
}

func (s *EquipmentServer) EquipmentPost(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	equipmentEdit, err := s.formatter.ConvertJsonToEquipmentEditAppData(b)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := s.app.AddNewEquipment(equipmentEdit)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	idJSON:= s.formatter.ConvertIdToJSON(id)

	fmt.Fprintf(w,idJSON)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func JSONResponse(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, string(json))
	if err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
