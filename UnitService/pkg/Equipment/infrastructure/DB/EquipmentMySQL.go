package EquipmentMySQLDB

import (
	App "UnitService/pkg/equipment/app"
	Model "UnitService/pkg/equipment/model"
	DB "UnitService/pkg/common/infrastructure"
	"errors"
)

var ErrorEmptyConnection  = errors.New("error DB connection")
var ErrorInitConnection  = errors.New("error Init Connection DB")
var ErrorRecordNotFound = errors.New("record Not Found")

type MySQLDB struct {
	Connection *DB.Connection
}

func CreateMySQLDB(connection *DB.Connection) App.EquipmentRepository {
	if connection.Db == nil {
		return nil
	}
	return &MySQLDB{connection}
}


func (db *MySQLDB) GetEquipmentById(id string) (Model.Equipment, error) {
	query := "SELECT * FROM unit_db.equipments where id = '" + id + "';";
	rows, err := db.Connection.Db.Query(query)
	if err != nil {
		return Model.Equipment{}, err
	}
	defer rows.Close()

	equipmentDB := Model.Equipment{}
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
			return Model.Equipment{}, err
		}
	}

	if equipmentDB.Id == "" {
		return Model.Equipment{}, ErrorRecordNotFound
	}

	return equipmentDB, nil
}

func (db *MySQLDB) GetEquipmentByRequiredParameters(equipmentParams App.RequiredParameters) (Model.Equipment, error) {
	query := "SELECT * FROM unit_db.equipments where name = '" + equipmentParams.Name + "' AND Cost = " + string(equipmentParams.Cost) + ";";
	rows, err := db.Connection.Db.Query(query)
	if err != nil {
		return Model.Equipment{}, err
	}
	defer rows.Close()

	equipmentDB := Model.Equipment{}
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
			return Model.Equipment{}, err
		}
	}

	if equipmentDB.Id == "" {
		return Model.Equipment{}, ErrorRecordNotFound
	}

	return equipmentDB, nil
}

func (db *MySQLDB) GetAllEquipments() ([]Model.Equipment, error) {
	rows, err := db.Connection.Db.Query("SELECT * FROM unit_db.equipments;")
	if err != nil {
		return []Model.Equipment{}, err
	}
	defer rows.Close()

	equipments := []Model.Equipment{}

	for rows.Next() {
		equipmentDB := Model.Equipment{}
		err = rows.Scan(&equipmentDB.Id,
			&equipmentDB.Name,
			&equipmentDB.LimitOnUnit,
			&equipmentDB.LimitOnTeam,
			&equipmentDB.SoldarRole,
			&equipmentDB.Rule,
			&equipmentDB.Ammo,
			&equipmentDB.Cost)
		if err != nil {
			return []Model.Equipment{}, err
		}
		equipments = append(equipments, equipmentDB)
	}

	return equipments, nil
}

func (db *MySQLDB) InsertNewEquipment(equipment Model.Equipment) (string, error) {
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

func (db *MySQLDB) UpdateEquipment(equipment Model.Equipment) (string, error) {
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
	_, err := db.Connection.Db.Exec("DELETE FROM `unit_db`.`equipments` WHERE id = ?;", id)

	if err != nil {
		return "", err
	}

	return id, nil
}
