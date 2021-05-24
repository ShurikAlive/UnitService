package UnitApp

import (
	UnitDB "UnitService/pkg/Unit/infrastructure/DB"
	. "UnitService/pkg/Unit/model"
	"errors"
	uuid "github.com/nu7hatch/gouuid"
)

type UnitInf struct {
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

type UnitEditInf struct {
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

func SerializationUnitInfData(unitDB UnitDB.UnitDB) (UnitInf) {
	unitInfData := UnitInf{
		unitDB.Id,
		unitDB.Name,
		unitDB.ForceName,
		unitDB.Hp,
		unitDB.Initiative,
		unitDB.Bs,
		unitDB.Fs,
		unitDB.AdditionalRule,
	}

	return unitInfData
}

func SerializationUnitInputData(unitDB UnitDB.UnitDB) (UnitInputData) {
	unitInData := UnitInputData{
		unitDB.Id,
		unitDB.Name,
		unitDB.ForceName,
		unitDB.Hp,
		unitDB.Initiative,
		unitDB.Bs,
		unitDB.Fs,
		unitDB.AdditionalRule,
	}

	return unitInData
}

func SerializationUnitInputDataFromUnitInf(unitInf UnitInf) (UnitInputData) {
	unitInData := UnitInputData{
		unitInf.Id,
		unitInf.Name,
		unitInf.ForceName,
		unitInf.Hp,
		unitInf.Initiative,
		unitInf.Bs,
		unitInf.Fs,
		unitInf.AdditionalRule,
	}

	return unitInData
}

func SerializationUnitInf(unit Unit) (UnitInf) {
	unitInf := UnitInf{
		unit.Id,
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule,
	}

	return unitInf
}

func ConvertUnitEditInfToUnitInf(id string, unitEditInf UnitEditInf) (UnitInf) {
	unitInf := UnitInf {
		id,
		unitEditInf.Name,
		unitEditInf.ForceName,
		unitEditInf.Hp,
		unitEditInf.Initiative,
		unitEditInf.Bs,
		unitEditInf.Fs,
		unitEditInf.AdditionalRule,
	}

	return unitInf
}

func SerialitheUnitInputDB(unit Unit) (UnitDB.UnitInputDB) {
	unitInputDB := UnitDB.UnitInputDB{
		unit.Id,
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule,
	}

	return unitInputDB;
}

func SerialitheRequiredParameters(unit Unit) (UnitDB.RequiredParameters) {
	requiredParameters := UnitDB.RequiredParameters{
		unit.Name,
		unit.ForceName,
	}

	return requiredParameters
}

func GenerateId() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	id := u.String()
	return id, nil
}

func UnitIdExist(db UnitDB.IUnitDB, id string) (bool) {
	unitFromDB, err := db.GetUnitInDBById(id)
	if err != nil {
		return false
	}

	unitDB := UnitDB.UnitDB{}

	if unitFromDB.Id == unitDB.Id {
		return false
	}

	return true
}

func UnitExist(db UnitDB.IUnitDB, unit Unit) (bool) {
	unitInputDB := SerialitheRequiredParameters(unit)
	unitFromDB, err := db.GetUnitInDBByRequiredParameters(unitInputDB)
	if err != nil {
		return false
	}

	unitDB := UnitDB.UnitDB{}

	if unitFromDB.Id == unitDB.Id {
		return false
	}

	return true
}

func GetUnitById(db UnitDB.IUnitDB, id string) (UnitInf, error) {
	unitFromDB, err := db.GetUnitInDBById(id)
	if err != nil {
		return UnitInf {}, err
	}
	unitInf := SerializationUnitInfData(unitFromDB)
	return unitInf, nil
}

func GetAllUnit(db UnitDB.IUnitDB) ([]UnitInf, error) {
	var units = []UnitInf{}
	unitsDB, err := db.GetAllUnits()
	if err != nil {
		return units, err
	}

	for i := 0; i < len(unitsDB); i++ {
		unitDB := unitsDB[i]
		unitInf := SerializationUnitInfData(unitDB)
		units = append(units, unitInf)
	}

	return units, nil
}

func AddNewUnit(db UnitDB.IUnitDB, unitInfo UnitEditInf) (string, error) {
	id, err := GenerateId()
	if err != nil {
		return "", err
	}
	unitInf := ConvertUnitEditInfToUnitInf(id, unitInfo)
	unitInData := SerializationUnitInputDataFromUnitInf(unitInf)
	unit, err := CreateUnit(unitInData)
	if err != nil {
		return "", err
	}
	if UnitExist(db, unit) {
		return "", errors.New("This unit Exist!")
	}
	unitInputDB := SerialitheUnitInputDB(unit)
	insertedId, err := db.InsertNewUnit(unitInputDB)
	if err != nil {
		return "", err
	}
	return insertedId, nil
}

func UpdateUnitInf(db UnitDB.IUnitDB, id string, unitEditInf UnitEditInf) (string, error) {
	unitInf := ConvertUnitEditInfToUnitInf(id, unitEditInf)
	unitInData := SerializationUnitInputDataFromUnitInf(unitInf)
	unit, err := CreateUnit(unitInData)
	if err != nil {
		return "", err
	}
	if !UnitIdExist(db, id) {
		return "", errors.New("Unit id not found!")
	}
	unitInputDB := SerialitheUnitInputDB(unit)
	updateId, err := db.UpdateUnit(unitInputDB)
	if err != nil {
		return "", err
	}
	return updateId, nil
}

func DeleteByIdInf(db UnitDB.IUnitDB, id string) (string, error) {
	if !UnitIdExist(db, id) {
		return "", errors.New("Unit id not found!")
	}
	deleteId, err := db.DeleteUnit(id)
	if err != nil {
		return "", err
	}
	return deleteId, nil
}