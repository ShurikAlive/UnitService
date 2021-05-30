package UnitApp

import (
	. "UnitService/pkg/Unit/model"
	"errors"
	uuid "github.com/nu7hatch/gouuid"
)

var ErrorUnitExist = errors.New("this unit Exist!")
var ErrorUnitIdExist = errors.New("unit id not found!")

type IUnitDB interface {
	GetUnitInDBById(id string) (Unit, error)
	GetUnitInDBByRequiredParameters(unitParams RequiredParameters) (Unit, error)
	GetAllUnits() ([]Unit, error)
	InsertNewUnit(unit Unit) (string, error)
	UpdateUnit(unit Unit) (string, error)
	DeleteUnit(id string) (string, error)
}

type UnitApp struct {
	db IUnitDB
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

func CreateUnitApp(db IUnitDB) UnitApp {
	return UnitApp{db}
}

func (app *UnitApp) convertUnitToUnitAppData(unit Unit) UnitAppData {
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

func (app *UnitApp) convertUnitEditEditAppDataToUnitAppData(id string, unitEdit UnitEditAppData) UnitAppData {
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

func (app *UnitApp) convertUnitAppDataToUnitInputData(unitApp UnitAppData) UnitInputData {
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

func (app *UnitApp) convertUnitToRequiredParameters(unit Unit) RequiredParameters {
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
	unit, err := app.db.GetUnitInDBById(id)
	return (err == nil) && (unit.Id != "")
}

func (app *UnitApp) unitExist(unit Unit) bool {
	unitInput := app.convertUnitToRequiredParameters(unit)
	unit, err := app.db.GetUnitInDBByRequiredParameters(unitInput)
	return (err == nil) && (unit.Id != "")
}

func (app *UnitApp) GetUnitById(id string) (UnitAppData, error) {
	unit, err := app.db.GetUnitInDBById(id)
	if err != nil {
		return UnitAppData {}, err
	}
	unitApp := app.convertUnitToUnitAppData(unit)
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
		unitInf := app.convertUnitToUnitAppData(unitDB)
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
	unitApp := app.convertUnitEditEditAppDataToUnitAppData(id, unitInfo)
	unitInData := app.convertUnitAppDataToUnitInputData(unitApp)
	unit, err := CreateUnit(unitInData)
	if err != nil {
		return "", err
	}
	err = app.assertEquipmentExist(unit)
	if err != nil {
		return "", err
	}
	insertedId, err := app.db.InsertNewUnit(unit)
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
	unitApp := app.convertUnitEditEditAppDataToUnitAppData(id, unitEditInf)
	unitInData := app.convertUnitAppDataToUnitInputData(unitApp)
	unit, err := CreateUnit(unitInData)
	if err != nil {
		return "", err
	}
	err = app.assertEquipmentNotExist(id)
	if err != nil {
		return "", err
	}
	updateId, err := app.db.UpdateUnit(unit)
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
	deleteId, err := app.db.DeleteUnit(id)
	if err != nil {
		return "", err
	}
	return deleteId, nil
}