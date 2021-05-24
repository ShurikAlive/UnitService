package UnitMySQLDB

import (
	DB "UnitService/DB"
	"errors"
)

var errorEmptyConnection  = errors.New("Error DB connection")
var errorInitConnection  = errors.New("Error Init Connection DB")
var errorRecordNotFound = errors.New("Record Not Found")

type MySQLDB struct {
	Connection *DB.Connection
}

func CreateMySQLDB(connection *DB.Connection) (MySQLDB) {
	return MySQLDB{connection}
}

func (db *MySQLDB) GetUnitInDBById(id string) (UnitDB, error) {
	if db.Connection.Db == nil {
		return UnitDB{}, errorEmptyConnection
	}

	query := "SELECT * FROM unit_db.units where id = '" + id + "';";
	rows, err := db.Connection.Db.Query(query)
	if err != nil {
		return UnitDB{}, err
	}
	defer rows.Close()

	unitDB := UnitDB{}
	for rows.Next() {
		err = rows.Scan(&unitDB.Id,
			&unitDB.Name,
			&unitDB.ForceName,
			&unitDB.Hp,
			&unitDB.Initiative,
			&unitDB.Bs,
			&unitDB.Fs,
			&unitDB.AdditionalRule)
		if err != nil {
			return UnitDB{}, err
		}
	}

	if unitDB.Id == "" {
		return UnitDB{}, errorRecordNotFound
	}

	return unitDB, nil
}

func (db *MySQLDB) GetUnitInDBByRequiredParameters(unitParams RequiredParameters) (UnitDB, error) {
	if db.Connection.Db == nil {
		return UnitDB{}, errorEmptyConnection
	}

	query := "SELECT * FROM unit_db.units where name = '" + unitParams.Name + "' AND ForceName = '" + unitParams.ForceName + "';";
	rows, err := db.Connection.Db.Query(query)
	if err != nil {
		return UnitDB{}, err
	}
	defer rows.Close()

	unitDB := UnitDB{}
	for rows.Next() {
		err = rows.Scan(&unitDB.Id,
			&unitDB.Name,
			&unitDB.ForceName,
			&unitDB.Hp,
			&unitDB.Initiative,
			&unitDB.Bs,
			&unitDB.Fs,
			&unitDB.AdditionalRule)
		if err != nil {
			return UnitDB{}, err
		}
	}

	if unitDB.Id == "" {
		return UnitDB{}, errorRecordNotFound
	}

	return unitDB, nil
}

func (db *MySQLDB) GetAllUnits() ([]UnitDB, error) {
	if db.Connection.Db == nil {
		return []UnitDB{}, errorEmptyConnection
	}

	rows, err := db.Connection.Db.Query("SELECT * FROM unit_db.units;")
	if err != nil {
		return []UnitDB{}, err
	}
	defer rows.Close()

	units := []UnitDB{}

	for rows.Next() {
		unitDB := UnitDB{}
		err = rows.Scan(&unitDB.Id,
			&unitDB.Name,
			&unitDB.ForceName,
			&unitDB.Hp,
			&unitDB.Initiative,
			&unitDB.Bs,
			&unitDB.Fs,
			&unitDB.AdditionalRule)
		if err != nil {
			return []UnitDB{}, err
		}
		units = append(units, unitDB)
	}

	return units, nil
}

func (db *MySQLDB) InsertNewUnit(unit UnitInputDB) (string, error) {
	if db.Connection.Db == nil {
		return "", errorEmptyConnection
	}

	_, err := db.Connection.Db.Exec("INSERT INTO `unit_db`.`units` 	(`id`, `Name`, `ForceName`,	`Hp`, `Initiative`, `Bs`, `Fs`, `AdditionalRule`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		unit.Id,
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule)

	if err != nil {
		return "", err
	}

	return unit.Id, nil
}

func (db *MySQLDB) UpdateUnit(unit UnitInputDB) (string, error) {
	if db.Connection.Db == nil {
		return "", errorEmptyConnection
	}

	_, err := db.Connection.Db.Exec("UPDATE `unit_db`.`units` SET `Name` = ?, `ForceName` = ?, `Hp` = ?, `Initiative` = ?, `Bs` = ?, `Fs` = ?, `AdditionalRule` = ? WHERE id = ?;",
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule,
		unit.Id)

	if err != nil {
		return "", err
	}

	return unit.Id, nil
}

func (db *MySQLDB) DeleteUnit(id string) (string, error) {
	if db.Connection.Db == nil {
		return "", errorEmptyConnection
	}

	_, err := db.Connection.Db.Exec("DELETE FROM `unit_db`.`units` WHERE id = ?;", id)

	if err != nil {
		return "", err
	}

	return id, nil
}

