package Unit

import (
	"errors"
	uuid "github.com/nu7hatch/gouuid"
	"UnitService/pkg/Unit/app/DB"
	. "UnitService/pkg/Unit/model"
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

type AllUnitsInf struct {
	Units        []UnitInf
	IdsWithError []string // Идентификаторы Unit которые лежат в БД, но не соответствуют бизнес логике
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

func SerializationUnitInputData(unitDB DB.UnitDB) (UnitInputData) {
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

func SerialitheUnitInputDB(unit Unit) (DB.UnitInputDB) {
	unitInputDB := DB.UnitInputDB {
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

func GenerateId() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	id := u.String()
	return id, nil
}

func GetUnitById(id string) (UnitInf, error) {
	unitFromDB, err := DB.GetUnitInDBById(id)
	if err != nil {
		return UnitInf {}, err
	}
	unitInData := SerializationUnitInputData(unitFromDB)
	unit, err := CreateUnit(unitInData)
	if err != nil {
		return UnitInf {}, err
	}
	unitInf := SerializationUnitInf(unit)
	return unitInf, nil
}

func GetAllUnit() (AllUnitsInf, error) {
	unitsDB, err := DB.GetAllUnits()
	if err != nil {
		return AllUnitsInf {}, err
	}

	allUnits := AllUnitsInf{}
	for i := 0; i < len(unitsDB); i++ {
		unitDB := unitsDB[i]
		unitInData := SerializationUnitInputData(unitDB)
		unit, err := CreateUnit(unitInData)
		if err != nil {
			allUnits.IdsWithError = append(allUnits.IdsWithError, unitDB.Id)
		} else {
			unitInf := SerializationUnitInf(unit)
			allUnits.Units = append(allUnits.Units, unitInf)
		}
	}

	return allUnits, nil
}

func AddNewUnit(unitInfo UnitEditInf) (string, error) {
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
	unitInputDB := SerialitheUnitInputDB(unit)
	insertedId, err := DB.InsertNewUnit(unitInputDB)
	if err != nil {
		return "", err
	}
	return insertedId, nil
}

func UpdateUnitInf(id string, unitEditInf UnitEditInf) (string, error) {
	unitInf := ConvertUnitEditInfToUnitInf(id, unitEditInf)
	unitInData := SerializationUnitInputDataFromUnitInf(unitInf)
	unit, err := CreateUnit(unitInData)
	if err != nil {
		return "", err
	}
	if !DB.UnitExist(id) {
		return "", errors.New("Unit not found!")
	}
	unitInputDB := SerialitheUnitInputDB(unit)
	updateId, err := DB.UpdateUnit(unitInputDB)
	if err != nil {
		return "", err
	}
	return updateId, nil
}

func DeleteByIdInf(id string) (string, error) {
	if !DB.UnitExist(id) {
		return "", errors.New("Unit not found!")
	}
	deleteId, err := DB.DeleteUnit(id)
	if err != nil {
		return "", err
	}
	return deleteId, nil
}