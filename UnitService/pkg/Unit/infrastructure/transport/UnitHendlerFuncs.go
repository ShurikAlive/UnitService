package transport

import (
	"UnitService/cmd/DB"
	App "UnitService/pkg/Unit/app"
	MySqlDB "UnitService/pkg/Unit/infrastructure/DB"
	Model "UnitService/pkg/Unit/model"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
)

type UnitServer struct {
	app        App.UnitApp
	formatter   JsonFormatter
}

func CreateUnitServer(connection *DB.Connection) (*UnitServer) {
	unitServer := new(UnitServer)
	db := MySqlDB.NewUnitDB(connection)
	unitServer.app = App.CreateUnitApp(db)
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
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	b, err := s.formatter.ConvertAllUnitAppDataToJSON(units)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, b)
}

func (s *UnitServer) UnitPost(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	unitEdit, err := s.formatter.ConvertJsonToUnitEditAppData(b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := s.app.AddNewUnit(unitEdit)

	if err != nil {
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	idJSON:= s.formatter.ConvertIdToJSON(id)

	fmt.Fprintf(w,idJSON)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (s *UnitServer) UnitUnitIdDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["unitId"]

	deleteId, err := s.app.DeleteById(id)

	if err != nil {
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	idJSON:= s.formatter.ConvertIdToJSON(deleteId)

	fmt.Fprintf(w,idJSON)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (s *UnitServer) UnitUnitIdGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["unitId"]

	unit, err:= s.app.GetUnitById(id)

	if err != nil {
		http.Error(w, err.Error(), s.getErrorCode(err))
		return
	}

	b, err := s.formatter.ConvertUnitAppDataToJSON(unit)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	JSONResponse(w, b)
}

func (s *UnitServer) UnitUnitIdPut(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["unitId"]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer r.Body.Close()

	unitEdit, err := s.formatter.ConvertJsonToUnitEditAppData(b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updateId, err := s.app.UpdateUnit(id, unitEdit)

	if err != nil {
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
