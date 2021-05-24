package EquipmentTransport

import (
	"UnitService/DB"
	EquipmentDB "UnitService/pkg/Equipment/infrastructure/DB"
	EquipmentApp "UnitService/pkg/Equipment/app"
	"encoding/json"
)

type EquipmentJSON struct {
	// ID equipment
	Id string `json:"id"`
	// FULL NAME equipment
	Name string `json:"name"`
	// limit equipment on one unit. -1 - unlimit
	LimitOnUnit int32 `json:"limitOnUnit"`
	// limit equipment on one team. -1 - unlimit
	LimitOnTeam int32 `json:"limitOnTeam"`
	// The role of a soldier available when selecting ammunition.
	SoldarRole string `json:"soldarRole"`
	// game rule equipment
	Rule string `json:"rule"`
	// limit equipment on game. -1 - unlimit
	Ammo int32 `json:"ammo"`
	// cost equipment in game points
	Cost int32 `json:"cost"`
}

type EditEquipmentJSON struct {
	// FULL NAME equipment
	Name string `json:"name"`
	// limit equipment on one unit. -1 - unlimit
	LimitOnUnit int32 `json:"limitOnUnit"`
	// limit equipment on one team. -1 - unlimit
	LimitOnTeam int32 `json:"limitOnTeam"`
	// The role of a soldier available when selecting ammunition.
	SoldarRole string `json:"soldarRole"`
	// game rule equipment
	Rule string `json:"rule"`
	// limit equipment on game. -1 - unlimit
	Ammo int32 `json:"ammo"`
	// cost equipment in game points
	Cost int32 `json:"cost"`
}

func ConvertEquipmentAppToEquipmentJson(equipmentApp EquipmentApp.EquipmentApp) (EquipmentJSON) {
	equipment := EquipmentJSON{
		equipmentApp.Id,
		equipmentApp.Name,
		equipmentApp.LimitOnUnit,
		equipmentApp.LimitOnTeam,
		equipmentApp.SoldarRole,
		equipmentApp.Rule,
		equipmentApp.Ammo,
		equipmentApp.Cost,
	}

	return equipment
}

func ConvertEditEquipmentJsonToEquipmentEditApp(editEquipmentJSON EditEquipmentJSON) (EquipmentApp.EditEquipmentApp) {
	equipment := EquipmentApp.EditEquipmentApp{
		editEquipmentJSON.Name,
		editEquipmentJSON.LimitOnUnit,
		editEquipmentJSON.LimitOnTeam,
		editEquipmentJSON.SoldarRole,
		editEquipmentJSON.Rule,
		editEquipmentJSON.Ammo,
		editEquipmentJSON.Cost,
	}

	return equipment
}

func ConvertAllEquipmentsToEquipmentsJson(allEquipments []EquipmentApp.EquipmentApp) ([]EquipmentJSON) {
	equipments := []EquipmentJSON {}
	for i := 0; i < len(allEquipments); i++ {
		equipmentApp := allEquipments[i]
		equipment := EquipmentJSON {
			equipmentApp.Id,
			equipmentApp.Name,
			equipmentApp.LimitOnUnit,
			equipmentApp.LimitOnTeam,
			equipmentApp.SoldarRole,
			equipmentApp.Rule,
			equipmentApp.Ammo,
			equipmentApp.Cost,
		}

		equipments = append(equipments, equipment)
	}

	return equipments
}

func DeleteById(connection *DB.Connection, id string) (string, error) {
	mySQLDB := EquipmentDB.CreateMySQLDB(connection)
	deleteId, err := EquipmentApp.DeleteByIdApp(&mySQLDB, id)
	if err != nil {
		return "", err
	}

	return deleteId, nil
}

func GetJSONEquipmentById(connection *DB.Connection, id string) ([]byte, error) {
	mySQLDB := EquipmentDB.CreateMySQLDB(connection)
	equipmentApp, err := EquipmentApp.GetEquipmentById(&mySQLDB, id)
	if err != nil {
		return nil, err
	}

	units := ConvertEquipmentAppToEquipmentJson(equipmentApp)
	b, err := json.Marshal(units)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func UpdateEquipment(connection *DB.Connection, id string, equipmentJson []byte) (string, error) {
	var msg EditEquipmentJSON
	err := json.Unmarshal(equipmentJson, &msg)
	if err != nil {
		return "", err
	}

	equipmentEditApp:= ConvertEditEquipmentJsonToEquipmentEditApp(msg)
	mySQLDB := EquipmentDB.CreateMySQLDB(connection)
	idEdit, err := EquipmentApp.UpdateEquipmentApp(&mySQLDB, id, equipmentEditApp)
	if err != nil {
		return "", err
	}

	return idEdit, nil
}

func GetJSONAllEquipment(connection *DB.Connection) ([]byte, error) {
	mySQLDB := EquipmentDB.CreateMySQLDB(connection)
	allEquipments, err := EquipmentApp.GetAllEquipment(&mySQLDB)
	if err != nil {
		return nil, err
	}

	equipments := ConvertAllEquipmentsToEquipmentsJson(allEquipments)
	b, err := json.Marshal(equipments)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func AddEquipment(connection *DB.Connection, equipmentJson []byte) (string, error) {
	var msg EditEquipmentJSON
	err := json.Unmarshal(equipmentJson, &msg)
	if err != nil {
		return "", err
	}

	equipmentEditApp:= ConvertEditEquipmentJsonToEquipmentEditApp(msg)
	mySQLDB := EquipmentDB.CreateMySQLDB(connection)
	id, err := EquipmentApp.AddNewEquipment(&mySQLDB, equipmentEditApp)
	if err != nil {
		return "", err
	}

	return id, nil
}
