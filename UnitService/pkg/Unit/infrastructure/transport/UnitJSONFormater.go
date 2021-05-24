package transport

import (
	"UnitService/DB"
	. "UnitService/pkg/Unit/app"
	MySQLDB "UnitService/pkg/Unit/infrastructure/DB"
	"encoding/json"
)

type UnitJSON struct {
	// ID unit
	Id string `json:"id"`
	// FULL NAME unit
	Name string `json:"name"`
	// force name unit
	ForceName string `json:"forceName"`
	// count heals point unit
	Hp int32 `json:"hp"`
	// initiative unit
	Initiative int32 `json:"initiative"`
	// ability to shoot unit
	Bs int32 `json:"bs"`
	// ability to fight unit
	Fs int32 `json:"fs"`
	// Additionat ability soldes
	AdditionalRule string `json:"additionalRule"`
}

type EditUnitJSON struct {
	// FULL NAME unit
	Name string `json:"name"`
	// force name unit
	ForceName string `json:"forceName"`
	// count heals point unit
	Hp int32 `json:"hp"`
	// initiative unit
	Initiative int32 `json:"initiative"`
	// ability to shoot unit
	Bs int32 `json:"bs"`
	// ability to fight unit
	Fs int32 `json:"fs"`
	// Additionat ability soldes
	AdditionalRule string `json:"additionalRule"`
}

func SerializationUnit(unitInf UnitInf) (UnitJSON) {
	unit := UnitJSON{
		unitInf.Id,
		unitInf.Name,
		unitInf.ForceName,
		unitInf.Hp,
		unitInf.Initiative,
		unitInf.Bs,
		unitInf.Fs,
		unitInf.AdditionalRule,
	}

	return unit
}

func SerializationAllUnit(unitsInf []UnitInf) ([]UnitJSON) {
	units := []UnitJSON {}
	for i := 0; i < len(unitsInf); i++ {
		unitInf := unitsInf[i]
		unit := UnitJSON{
			unitInf.Id,
			unitInf.Name,
			unitInf.ForceName,
			unitInf.Hp,
			unitInf.Initiative,
			unitInf.Bs,
			unitInf.Fs,
			unitInf.AdditionalRule,
		}

		units = append(units, unit)
	}

	return units
}

func SerializationEditUnit(unitEditJson EditUnitJSON) (UnitEditInf) {
	unitEdit := UnitEditInf {
		unitEditJson.Name,
		unitEditJson.ForceName,
		unitEditJson.Hp,
		unitEditJson.Initiative,
		unitEditJson.Bs,
		unitEditJson.Fs,
		unitEditJson.AdditionalRule,
	}

	return unitEdit
}

func GetJSONUnitById(connection *DB.Connection, unitId string) ([]byte, error) {
	mySQLDB := MySQLDB.CreateMySQLDB(connection)
	unitInf, err := GetUnitById(&mySQLDB, unitId)
	if err != nil {
		return nil, err
	}

	units := SerializationUnit(unitInf)
	b, err := json.Marshal(units)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func GetJSONAllUnitById(connection *DB.Connection) ([]byte, error) {
	mySQLDB := MySQLDB.CreateMySQLDB(connection)
	allUnits, err := GetAllUnit(&mySQLDB)
	if err != nil {
		return nil, err
	}

	units := SerializationAllUnit(allUnits)
	b, err := json.Marshal(units)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func AddUnit(connection *DB.Connection, unitJson []byte) (string, error) {
	var msg EditUnitJSON
	err := json.Unmarshal(unitJson, &msg)
	if err != nil {
		return "", err
	}

	unitEditInf:= SerializationEditUnit(msg)
	mySQLDB := MySQLDB.CreateMySQLDB(connection)
	id, err := AddNewUnit(&mySQLDB, unitEditInf)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateUnit(connection *DB.Connection, id string, unitJson []byte) (string, error) {
	var msg EditUnitJSON
	err := json.Unmarshal(unitJson, &msg)
	if err != nil {
		return "", err
	}

	unitEditInf:= SerializationEditUnit(msg)
	mySQLDB := MySQLDB.CreateMySQLDB(connection)
	idEdit, err := UpdateUnitInf(&mySQLDB, id, unitEditInf)
	if err != nil {
		return "", err
	}

	return idEdit, nil
}

func DeleteById(connection *DB.Connection, id string) (string, error) {
	mySQLDB := MySQLDB.CreateMySQLDB(connection)
	deleteId, err := DeleteByIdInf(&mySQLDB, id)
	if err != nil {
		return "", err
	}

	return deleteId, nil
}