package UnitMySQLDB

import (
	DB "UnitService/pkg/common/infrastructure"
	App "UnitService/pkg/unit/app"
	Model "UnitService/pkg/unit/model"
	"errors"
)

var ErrorEmptyConnection  = errors.New("error DB connection")
var ErrorInitConnection  = errors.New("error Init Connection DB")
var ErrorRecordNotFound = errors.New("record Not Found")

var lockName = "UNIT_LOCK"

type MySQLDB struct {
	Connection *DB.Connection
}

func NewUnitDB(connection *DB.Connection) App.UnitRepository {
	if connection.Db == nil {
		return nil
	}
	return &MySQLDB{connection}
}

func (db *MySQLDB) Lock() error {
	_, err := db.Connection.Db.Exec("SELECT GET_LOCK(?, -1)", lockName)
	if err != nil {
		return err
	}
	return nil
}

func (db *MySQLDB) Unlock() error {
	_, err := db.Connection.Db.Exec("SELECT RELEASE_LOCK(?)", lockName)
	if err != nil {
		return err
	}
	return nil
}

func (db *MySQLDB) GetUnitById(id string) (Model.Unit, error) {
	query := "SELECT * FROM unit_db.units where id = '" + id + "';";
	rows, err := db.Connection.Db.Query(query)
	if err != nil {
		return Model.Unit{}, err
	}
	defer rows.Close()

	unit := Model.Unit{}
	for rows.Next() {
		err = rows.Scan(&unit.Id,
			&unit.Name,
			&unit.ForceName,
			&unit.Hp,
			&unit.Initiative,
			&unit.Bs,
			&unit.Fs,
			&unit.AdditionalRule)
		if err != nil {
			return Model.Unit{}, err
		}
	}

	if unit.Id == "" {
		return Model.Unit{}, ErrorRecordNotFound
	}

	return unit, nil
}

func (db *MySQLDB) GetUnitByRequiredParameters(unitParams App.RequiredParameters) (Model.Unit, error) {
	query := "SELECT * FROM unit_db.units where name = '" + unitParams.Name + "' AND ForceName = '" + unitParams.ForceName + "';";
	rows, err := db.Connection.Db.Query(query)
	if err != nil {
		return Model.Unit{}, err
	}
	defer rows.Close()

	unit := Model.Unit{}
	for rows.Next() {
		err = rows.Scan(&unit.Id,
			&unit.Name,
			&unit.ForceName,
			&unit.Hp,
			&unit.Initiative,
			&unit.Bs,
			&unit.Fs,
			&unit.AdditionalRule)
		if err != nil {
			return Model.Unit{}, err
		}
	}

	if unit.Id == "" {
		return Model.Unit{}, ErrorRecordNotFound
	}

	return unit, nil
}

func (db *MySQLDB) GetAllUnits() ([]Model.Unit, error) {
	rows, err := db.Connection.Db.Query("SELECT * FROM unit_db.units;")
	if err != nil {
		return []Model.Unit{}, err
	}
	defer rows.Close()

	var units []Model.Unit

	for rows.Next() {
		unit := Model.Unit{}
		err = rows.Scan(&unit.Id,
			&unit.Name,
			&unit.ForceName,
			&unit.Hp,
			&unit.Initiative,
			&unit.Bs,
			&unit.Fs,
			&unit.AdditionalRule)
		if err != nil {
			return []Model.Unit{}, err
		}

		units = append(units, unit)
	}

	return units, nil
}

func (db *MySQLDB) InsertNewUnit(unit Model.Unit) (string, error) {
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

func (db *MySQLDB) UpdateUnit(unit Model.Unit) (string, error) {
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
	_, err := db.Connection.Db.Exec("DELETE FROM `unit_db`.`units` WHERE id = ?;", id)

	if err != nil {
		return "", err
	}

	return id, nil
}

