package transport

import (
	Roster "UnitService/pkg/unit/infrastructure/transport/services/roster"
	App "UnitService/pkg/unit/app"
	MySqlDB "UnitService/pkg/unit/infrastructure/db"
	Model "UnitService/pkg/unit/model"
	DB "UnitService/pkg/common/infrastructure"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"io"
	"io/ioutil"
	"net/http"
)

type UnitServer struct {
	app        App.UnitApp
	formatter   JsonFormatter
}

func CreateUnitServer(connection *DB.Connection, channelRebbitMQ *amqp.Channel, queueRebbitMQ amqp.Queue) (*UnitServer) {
	unitServer := new(UnitServer)
	db := MySqlDB.NewUnitDB(connection)
	rosterService := Roster.CreateRosterServices(channelRebbitMQ, queueRebbitMQ)
	unitServer.app = App.CreateUnitApp(db, rosterService)
	unitServer.formatter = CreateJSONFormatter()
	return unitServer
}

func (s *UnitServer) getErrorCode(err error) int {
	code := http.StatusInternalServerError
	switch err {
	case App.ErrorUnitExist:
		code = http.StatusBadRequest
	case App.ErrorUnitIdExist:
		code = http.StatusBadRequest

	case MySqlDB.ErrorEmptyConnection:
		code = http.StatusBadRequest
	case MySqlDB.ErrorInitConnection:
		code = http.StatusBadRequest
	case MySqlDB.ErrorRecordNotFound:
		code = http.StatusBadRequest

	case Model.InvalidUnitInitiative:
		code = http.StatusBadRequest
	case Model.InvalidUnitBs:
		code = http.StatusBadRequest
	case Model.InvalidUnitFs:
		code = http.StatusBadRequest
	case Model.InvalidUnitHp:
		code = http.StatusBadRequest
	case Model.InvalidUnitForceName:
		code = http.StatusBadRequest
	case Model.InvalidUnitId:
		code = http.StatusBadRequest
	case Model.InvalidUnitName:
		code = http.StatusBadRequest
	}

	return code
}

func (s *UnitServer) UnitGet(w http.ResponseWriter, r *http.Request) {
	units, err:= s.app.GetAllUnit()

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	b, err := s.formatter.ConvertAllUnitAppDataToJSON(units)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, b)
}

func (s *UnitServer) UnitPost(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	unitEdit, err := s.formatter.ConvertJsonToUnitEditAppData(b)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := s.app.AddNewUnit(unitEdit)

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

func (s *UnitServer) UnitIdDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["unitId"]

	deleteId, err := s.app.DeleteById(id)

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

func (s *UnitServer) UnitIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["unitId"]

	unit, err:= s.app.GetUnitById(id)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	b, err := s.formatter.ConvertUnitAppDataToJSON(unit)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, b)
}

func (s *UnitServer) UnitIdPut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["unitId"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	unitEdit, err := s.formatter.ConvertJsonToUnitEditAppData(b)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateId, err := s.app.UpdateUnit(id, unitEdit)

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

func JSONResponse(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type","application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, string(json))
	if err != nil {
		log.WithField("err", err).Error("write response error")
	}
}
