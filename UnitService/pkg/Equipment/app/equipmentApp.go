package EquipmentApp

import (
	EquipmentDB "UnitService/pkg/Equipment/infrastructure/DB"
	EquipmentModel "UnitService/pkg/Equipment/model"
	"errors"
	uuid "github.com/nu7hatch/gouuid"
)

var ErrorUnitNotFound = errors.New("Unit id not found!")

type EquipmentApp struct {
	// ID equipment
	Id string
	// FULL NAME equipment
	Name string
	// limit equipment on one unit. -1 - unlimit
	LimitOnUnit int32
	// limit equipment on one team. -1 - unlimit
	LimitOnTeam int32
	// The role of a soldier available when selecting ammunition.
	SoldarRole string
	// game rule equipment
	Rule string
	// limit equipment on game. -1 - unlimit
	Ammo int32
	// cost equipment in game points
	Cost int32
}

type EditEquipmentApp struct {
	// FULL NAME equipment
	Name string
	// limit equipment on one unit. -1 - unlimit
	LimitOnUnit int32
	// limit equipment on one team. -1 - unlimit
	LimitOnTeam int32
	// The role of a soldier available when selecting ammunition.
	SoldarRole string
	// game rule equipment
	Rule string
	// limit equipment on game. -1 - unlimit
	Ammo int32
	// cost equipment in game points
	Cost int32
}

func GenerateId() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	id := u.String()
	return id, nil
}

func ConvertEquipmentDBToEquipmentApp (equipmentDB EquipmentDB.EquipmentDB) EquipmentApp {
	equipment := EquipmentApp {
		equipmentDB.Id,
		equipmentDB.Name,
		equipmentDB.LimitOnUnit,
		equipmentDB.LimitOnTeam,
		equipmentDB.SoldarRole,
		equipmentDB.Rule,
		equipmentDB.Ammo,
		equipmentDB.Cost,
	}

	return equipment
}

func ConvertEditEquipmentAppToEquipmentApp (id string, equipmentEdit EditEquipmentApp) EquipmentApp {
	equipment := EquipmentApp {
		id,
		equipmentEdit.Name,
		equipmentEdit.LimitOnUnit,
		equipmentEdit.LimitOnTeam,
		equipmentEdit.SoldarRole,
		equipmentEdit.Rule,
		equipmentEdit.Ammo,
		equipmentEdit.Cost,
	}

	return equipment
}

func ConvertEquipmentAppToEquipmentInputData(equipment EquipmentApp) EquipmentModel.EquipmentInputData {
	equipmentInput := EquipmentModel.EquipmentInputData {
		equipment.Id,
		equipment.Name,
		equipment.LimitOnUnit,
		equipment.LimitOnTeam,
		equipment.SoldarRole,
		equipment.Rule,
		equipment.Ammo,
		equipment.Cost,
	}

	return equipmentInput
}

func ConvertEquipmentToUnitInputDB(equipment EquipmentModel.Equipment) EquipmentDB.EquipmentInputDB {
	equipmentInDB := EquipmentDB.EquipmentInputDB {
		equipment.Id,
		equipment.Name,
		equipment.LimitOnUnit,
		equipment.LimitOnTeam,
		equipment.SoldarRole,
		equipment.Rule,
		equipment.Ammo,
		equipment.Cost,
	}

	return equipmentInDB
}

func EquipmentIdExist(db EquipmentDB.IEquipmentDB, id string) bool {
	equipmentFromDB, err := db.GetEquipmentInDBById(id)
	if err != nil {
		return false
	}

	if equipmentFromDB.Id == "" {
		return false
	}

	return true
}

func ConvertEquipmentToEquipmentRequiredParameters(equipment EquipmentModel.Equipment) EquipmentDB.RequiredParameters {
	requiredParameters := EquipmentDB.RequiredParameters {
		equipment.Name,
		equipment.Cost,
	}

	return requiredParameters
}

func EquipmentExist(db EquipmentDB.IEquipmentDB, equipment EquipmentModel.Equipment) (bool) {
	equipmentInputDB := ConvertEquipmentToEquipmentRequiredParameters(equipment)
	unitFromDB, err := db.GetEquipmentInDBByRequiredParameters(equipmentInputDB)
	if err != nil {
		return false
	}

	if unitFromDB.Id == "" {
		return false
	}

	return true
}

func DeleteByIdApp(db EquipmentDB.IEquipmentDB, id string) (string, error) {
	if !EquipmentIdExist(db, id) {
		return "", ErrorUnitNotFound
	}
	deleteId, err := db.DeleteEquipment(id)
	if err != nil {
		return "", err
	}
	return deleteId, nil
}


func GetEquipmentById(db EquipmentDB.IEquipmentDB, id string) (EquipmentApp, error) {
	equipmentFromDB, err := db.GetEquipmentInDBById(id)
	if err != nil {
		return EquipmentApp {}, err
	}
	equipmentApp := ConvertEquipmentDBToEquipmentApp(equipmentFromDB)
	return equipmentApp, nil
}

func AddNewEquipment(db EquipmentDB.IEquipmentDB, equipmentInfo EditEquipmentApp) (string, error) {
	id, err := GenerateId()
	if err != nil {
		return "", err
	}
	equipmentApp := ConvertEditEquipmentAppToEquipmentApp(id, equipmentInfo)
	equipmentInData := ConvertEquipmentAppToEquipmentInputData(equipmentApp)
	equipment, err := EquipmentModel.CreateEquipment(equipmentInData)
	if err != nil {
		return "", err
	}
	if EquipmentExist(db, equipment) {
		return "", ErrorUnitNotFound
	}
	equipmentInputDB := ConvertEquipmentToUnitInputDB(equipment)
	insertedId, err := db.InsertNewEquipment(equipmentInputDB)
	if err != nil {
		return "", err
	}
	return insertedId, nil
}


func UpdateEquipmentApp(db EquipmentDB.IEquipmentDB, id string, equipmentInfo EditEquipmentApp) (string, error) {
	equipmentApp := ConvertEditEquipmentAppToEquipmentApp(id, equipmentInfo)
	equipmentInData := ConvertEquipmentAppToEquipmentInputData(equipmentApp)
	equipment, err := EquipmentModel.CreateEquipment(equipmentInData)
	if err != nil {
		return "", err
	}
	if !EquipmentIdExist(db, id) {
		return "", ErrorUnitNotFound
	}
	equipmentInputDB := ConvertEquipmentToUnitInputDB(equipment)
	updateId, err := db.UpdateEquipment(equipmentInputDB)
	if err != nil {
		return "", err
	}
	return updateId, nil
}

func GetAllEquipment(db EquipmentDB.IEquipmentDB) ([]EquipmentApp, error) {
	var equipments = []EquipmentApp{}
	equipmentsDB, err := db.GetAllEquipments()
	if err != nil {
		return equipments, err
	}

	for i := 0; i < len(equipmentsDB); i++ {
		equipmentDB := equipmentsDB[i]
		equipmentApp := ConvertEquipmentDBToEquipmentApp(equipmentDB)
		equipments = append(equipments, equipmentApp)
	}

	return equipments, nil
}

