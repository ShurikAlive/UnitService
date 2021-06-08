package UnitApp

import (
	. "UnitService/pkg/unit/model"
	"errors"
	uuid "github.com/nu7hatch/gouuid"
	"sync"
)

var ErrorUnitExist = errors.New("this unit Exist!")
var ErrorUnitIdExist = errors.New("unit id not found!")

type UnitRepository interface {
	GetUnitById(id string) (Unit, error)
	GetUnitByRequiredParameters(unitParams RequiredParameters) (Unit, error)
	GetAllUnits() ([]Unit, error)
	InsertNewUnit(unit Unit) (string, error)
	UpdateUnit(unit Unit) (string, error)
	DeleteUnit(id string) (string, error)
}

type RosterRepository interface {
	SendEvent(typeEvent string, idRecord string) error
}

type UnitApp struct {
	db UnitRepository
	roster RosterRepository
	mutex *sync.Mutex
}

type RequiredParameters struct {
	// FULL NAME unit
	Name string
	// force name unit
	ForceName string
}

type UnitAppData struct {
	// ID unit
	Id string
	// FULL NAME unit
	Name string
	// force name unit
	ForceName string
	// count heals point unit
	Hp int32
	// initiative unit
	Initiative int32
	// ability to shoot unit
	Bs int32
	// ability to fight unit
	Fs int32
	// Additionat ability soldes
	AdditionalRule string
}

type UnitEditAppData struct {
	// FULL NAME unit
	Name string
	// force name unit
	ForceName string
	// count heals point unit
	Hp int32
	// initiative unit
	Initiative int32
	// ability to shoot unit
	Bs int32
	// ability to fight unit
	Fs int32
	// Additionat ability soldes
	AdditionalRule string
}

func CreateUnitApp(db UnitRepository, roster RosterRepository) UnitApp {
	var mutex = &sync.Mutex{}
	return UnitApp{db, roster,mutex}
}

func (app *UnitApp) createUnitAppData(unit Unit) UnitAppData {
	unitApp := UnitAppData{
		unit.Id,
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule,
	}

	return unitApp
}

func (app *UnitApp) createUnitAppDataById(id string, unitEdit UnitEditAppData) UnitAppData {
	UnitAppData := UnitAppData {
		id,
		unitEdit.Name,
		unitEdit.ForceName,
		unitEdit.Hp,
		unitEdit.Initiative,
		unitEdit.Bs,
		unitEdit.Fs,
		unitEdit.AdditionalRule,
	}

	return UnitAppData
}

func (app *UnitApp) createUnitInputData(unitApp UnitAppData) UnitInputData {
	unitInData := UnitInputData{
		unitApp.Id,
		unitApp.Name,
		unitApp.ForceName,
		unitApp.Hp,
		unitApp.Initiative,
		unitApp.Bs,
		unitApp.Fs,
		unitApp.AdditionalRule,
	}

	return unitInData
}

func (app *UnitApp) createRequiredParameters(unit Unit) RequiredParameters {
	requiredParameters := RequiredParameters{
		unit.Name,
		unit.ForceName,
	}

	return requiredParameters
}

func (app *UnitApp) generateId() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	id := u.String()
	return id, nil
}

func (app *UnitApp) unitIdExist(id string) bool {
	unit, err := app.db.GetUnitById(id)
	return (err == nil) && (unit.Id != "")
}

func (app *UnitApp) unitExist(unit Unit) bool {
	unitInput := app.createRequiredParameters(unit)
	unit, err := app.db.GetUnitByRequiredParameters(unitInput)
	return (err == nil) && (unit.Id != "")
}

func (app *UnitApp) GetUnitById(id string) (UnitAppData, error) {
	unit, err := app.db.GetUnitById(id)
	if err != nil {
		return UnitAppData {}, err
	}
	unitApp := app.createUnitAppData(unit)
	return unitApp, nil
}

func (app *UnitApp) GetAllUnit() ([]UnitAppData, error) {
	var units []UnitAppData
	unitsDB, err := app.db.GetAllUnits()
	if err != nil {
		return units, err
	}

	for i := 0; i < len(unitsDB); i++ {
		unitDB := unitsDB[i]
		unitInf := app.createUnitAppData(unitDB)
		units = append(units, unitInf)
	}

	return units, nil
}

func (app *UnitApp) assertEquipmentExist(unit Unit) error {
	if app.unitExist(unit) {
		return ErrorUnitExist
	}
	return nil
}

func (app *UnitApp) AddNewUnit(unitInfo UnitEditAppData) (string, error) {
	id, err := app.generateId()
	if err != nil {
		return "", err
	}
	unitApp := app.createUnitAppDataById(id, unitInfo)
	unitInData := app.createUnitInputData(unitApp)
	unit, err := CreateUnit(unitInData)
	if err != nil {
		return "", err
	}
	err = app.assertEquipmentExist(unit)
	if err != nil {
		return "", err
	}
	app.mutex.Lock()
	insertedId, err := app.db.InsertNewUnit(unit)
	app.mutex.Unlock()
	if err != nil {
		return "", err
	}
	return insertedId, nil
}

func (app *UnitApp) assertEquipmentNotExist(id string) error {
	if !app.unitIdExist(id) {
		return ErrorUnitExist
	}
	return nil
}

func (app *UnitApp) UpdateUnit(id string, unitEditInf UnitEditAppData) (string, error) {
	unitApp := app.createUnitAppDataById(id, unitEditInf)
	unitInData := app.createUnitInputData(unitApp)
	unit, err := CreateUnit(unitInData)
	if err != nil {
		return "", err
	}
	err = app.assertEquipmentNotExist(id)
	if err != nil {
		return "", err
	}
	app.mutex.Lock()
	updateId, err := app.db.UpdateUnit(unit)
	app.mutex.Unlock()
	if err != nil {
		return "", err
	}
	err = app.roster.SendEvent("UPDATE", updateId)
	if err != nil {
		return "", err
	}
	return updateId, nil
}

func (app *UnitApp) DeleteById(id string) (string, error) {
	err := app.assertEquipmentNotExist(id)
	if err != nil {
		return "", err
	}
	app.mutex.Lock()
	deleteId, err := app.db.DeleteUnit(id)
	app.mutex.Unlock()
	if err != nil {
		return "", err
	}
	err = app.roster.SendEvent("DELETE", deleteId)
	if err != nil {
		return "", err
	}
	return deleteId, nil
}