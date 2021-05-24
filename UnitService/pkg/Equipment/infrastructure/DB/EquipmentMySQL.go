package EquipmentMySQLDB

import (
	"UnitService/DB"
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


func (db *MySQLDB) GetEquipmentInDBById(id string) (EquipmentDB, error) {
	if db.Connection.Db == nil {
		return EquipmentDB{}, errorEmptyConnection
	}

	query := "SELECT * FROM unit_db.equipments where id = '" + id + "';";
	rows, err := db.Connection.Db.Query(query)
	if err != nil {
		return EquipmentDB{}, err
	}
	defer rows.Close()

	equipmentDB := EquipmentDB{}
	for rows.Next() {
		err = rows.Scan(&equipmentDB.Id,
			&equipmentDB.Name,
			&equipmentDB.LimitOnUnit,
			&equipmentDB.LimitOnTeam,
			&equipmentDB.SoldarRole,
			&equipmentDB.Rule,
			&equipmentDB.Ammo,
			&equipmentDB.Cost)
		if err != nil {
			return EquipmentDB{}, err
		}
	}

	if equipmentDB.Id == "" {
		return EquipmentDB{}, errorRecordNotFound
	}

	return equipmentDB, nil
}

func (db *MySQLDB) GetEquipmentInDBByRequiredParameters(equipmentParams RequiredParameters) (EquipmentDB, error) {
	if db.Connection.Db == nil {
		return EquipmentDB{}, errorEmptyConnection
	}

	query := "SELECT * FROM unit_db.equipments where name = '" + equipmentParams.Name + "' AND Cost = " + string(equipmentParams.Cost) + ";";
	rows, err := db.Connection.Db.Query(query)
	if err != nil {
		return EquipmentDB{}, err
	}
	defer rows.Close()

	equipmentDB := EquipmentDB{}
	for rows.Next() {
		err = rows.Scan(&equipmentDB.Id,
			&equipmentDB.Name,
			&equipmentDB.LimitOnUnit,
			&equipmentDB.LimitOnTeam,
			&equipmentDB.SoldarRole,
			&equipmentDB.Rule,
			&equipmentDB.Ammo,
			&equipmentDB.Cost)
		if err != nil {
			return EquipmentDB{}, err
		}
	}

	if equipmentDB.Id == "" {
		return EquipmentDB{}, errorRecordNotFound
	}

	return equipmentDB, nil
}

func (db *MySQLDB) GetAllEquipments() ([]EquipmentDB, error) {
	if db.Connection.Db == nil {
		return []EquipmentDB{}, errorEmptyConnection
	}

	rows, err := db.Connection.Db.Query("SELECT * FROM unit_db.equipments;")
	if err != nil {
		return []EquipmentDB{}, err
	}
	defer rows.Close()

	equipments := []EquipmentDB{}

	for rows.Next() {
		equipmentDB := EquipmentDB{}
		err = rows.Scan(&equipmentDB.Id,
			&equipmentDB.Name,
			&equipmentDB.LimitOnUnit,
			&equipmentDB.LimitOnTeam,
			&equipmentDB.SoldarRole,
			&equipmentDB.Rule,
			&equipmentDB.Ammo,
			&equipmentDB.Cost)
		if err != nil {
			return []EquipmentDB{}, err
		}
		equipments = append(equipments, equipmentDB)
	}

	return equipments, nil
}

func (db *MySQLDB) InsertNewEquipment(equipment EquipmentInputDB) (string, error) {
	if db.Connection.Db == nil {
		return "", errorEmptyConnection
	}

	_, err := db.Connection.Db.Exec("INSERT INTO `unit_db`.`equipments` 	(`id`, `Name`, `LimitOnUnit`,	`LimitOnTeam`, `SoldarRole`, `Rule`, `Ammo`, `Cost`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		equipment.Id,
		equipment.Name,
		equipment.LimitOnUnit,
		equipment.LimitOnTeam,
		equipment.SoldarRole,
		equipment.Rule,
		equipment.Ammo,
		equipment.Cost)

	if err != nil {
		return "", err
	}

	return equipment.Id, nil
}

func (db *MySQLDB) UpdateEquipment(equipment EquipmentInputDB) (string, error) {
	if db.Connection.Db == nil {
		return "", errorEmptyConnection
	}

	_, err := db.Connection.Db.Exec("UPDATE `unit_db`.`equipments` SET `Name` = ?, `LimitOnUnit` = ?, `LimitOnTeam` = ?, `SoldarRole` = ?, `Rule` = ?, `Ammo` = ?, `Cost` = ? WHERE id = ?;",
		equipment.Name,
		equipment.LimitOnUnit,
		equipment.LimitOnTeam,
		equipment.SoldarRole,
		equipment.Rule,
		equipment.Ammo,
		equipment.Cost,
		equipment.Id)

	if err != nil {
		return "", err
	}

	return equipment.Id, nil
}

func (db *MySQLDB) DeleteEquipment(id string) (string, error) {
	if db.Connection.Db == nil {
		return "", errorEmptyConnection
	}

	_, err := db.Connection.Db.Exec("DELETE FROM `unit_db`.`equipments` WHERE id = ?;", id)

	if err != nil {
		return "", err
	}

	return id, nil
}
