package EquipmentApp

import (
	Model "UnitService/pkg/equipment/model"
	"errors"
	uuid "github.com/nu7hatch/gouuid"
)

var ErrorEquipmentNotFound = errors.New("equipment id not found!")

type EquipmentRepository interface {
	Lock() error
	Unlock() error

	GetEquipmentById(id string) (Model.Equipment, error)
	GetEquipmentByRequiredParameters(equipmentParams RequiredParameters) (Model.Equipment, error)
	GetAllEquipments() ([]Model.Equipment, error)
	InsertNewEquipment(equipment Model.Equipment) (string, error)
	UpdateEquipment(equipment Model.Equipment) (string, error)
	DeleteEquipment(id string) (string, error)
}

type RosterRepository interface {
	SendEvent(typeEvent string, idRecord string) error
}

type EquipmentApp struct {
	db EquipmentRepository
	roster RosterRepository
}

type RequiredParameters struct {
	// FULL NAME equipment
	Name string
	// cost equipment in game points
	Cost int32
}

type EquipmentAppData struct {
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

type EditEquipmentAppData struct {
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

func generateId() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}

	id := u.String()
	return id, nil
}

func CreateEquipmentApp(db EquipmentRepository, roster RosterRepository) EquipmentApp {
	return EquipmentApp{db, roster}
}

func (app *EquipmentApp) createEquipmentAppById (id string, equipmentEdit EditEquipmentAppData) EquipmentAppData {
	equipment := EquipmentAppData {
		Id:          id,
		Name:        equipmentEdit.Name,
		LimitOnUnit: equipmentEdit.LimitOnUnit,
		LimitOnTeam: equipmentEdit.LimitOnTeam,
		SoldarRole:  equipmentEdit.SoldarRole,
		Rule:        equipmentEdit.Rule,
		Ammo:        equipmentEdit.Ammo,
		Cost:        equipmentEdit.Cost,
	}

	return equipment
}

func (app *EquipmentApp) createEquipmentInputData(equipment EquipmentAppData) Model.EquipmentInputData {
	equipmentInput := Model.EquipmentInputData {
		Id:          equipment.Id,
		Name:        equipment.Name,
		LimitOnUnit: equipment.LimitOnUnit,
		LimitOnTeam: equipment.LimitOnTeam,
		SoldarRole:  equipment.SoldarRole,
		Rule:        equipment.Rule,
		Ammo:        equipment.Ammo,
		Cost:        equipment.Cost,
	}

	return equipmentInput
}

func (app *EquipmentApp) createEquipmentApp(equipment Model.Equipment) EquipmentAppData {
	equipmentApp := EquipmentAppData {
		Id:          equipment.Id,
		Name:        equipment.Name,
		LimitOnUnit: equipment.LimitOnUnit,
		LimitOnTeam: equipment.LimitOnTeam,
		SoldarRole:  equipment.SoldarRole,
		Rule:        equipment.Rule,
		Ammo:        equipment.Ammo,
		Cost:        equipment.Cost,
	}

	return equipmentApp
}

func (app *EquipmentApp) equipmentIdExist(id string) bool {
	equipmentFromDB, err := app.db.GetEquipmentById(id)
	if err != nil {
		return false
	}

	if equipmentFromDB.Id == "" {
		return false
	}

	return true
}

func (app *EquipmentApp) createEquipmentRequiredParameters(equipment Model.Equipment) RequiredParameters {
	requiredParameters := RequiredParameters {
		equipment.Name,
		equipment.Cost,
	}

	return requiredParameters
}

func (app *EquipmentApp) EquipmentExist(equipment Model.Equipment) bool {
	equipmentInputDB := app.createEquipmentRequiredParameters(equipment)
	equipmentFromDB, err := app.db.GetEquipmentByRequiredParameters(equipmentInputDB)
	if err != nil {
		return false
	}

	if equipmentFromDB.Id == "" {
		return false
	}

	return true
}

func (app *EquipmentApp) assertEquipmentIdNotExist(id string) error {
	if !app.equipmentIdExist(id) {
		return ErrorEquipmentNotFound
	}
	return nil
}

func (app *EquipmentApp) DeleteByIdApp(id string) (string, error) {
	err := app.assertEquipmentIdNotExist(id)
	if err != nil {
		return "", err
	}
	app.db.Lock()
	deleteId, err := app.db.DeleteEquipment(id)
	app.db.Unlock()
	if err != nil {
		return "", err
	}
	err = app.roster.SendEvent("DELETE", deleteId)
	if err != nil {
		return "", err
	}
	return deleteId, nil
}


func (app *EquipmentApp) GetEquipmentById(id string) (EquipmentAppData, error) {
	equipmentFromDB, err := app.db.GetEquipmentById(id)
	if err != nil {
		return EquipmentAppData {}, err
	}
	equipmentApp := app.createEquipmentApp(equipmentFromDB)
	return equipmentApp, nil
}

func (app *EquipmentApp) assertEquipmentExist(equipment Model.Equipment) error {
	if app.EquipmentExist(equipment) {
		return ErrorEquipmentNotFound
	}
	return nil
}

func (app *EquipmentApp) AddNewEquipment(equipmentInfo EditEquipmentAppData) (string, error) {
	id, err := generateId()
	if err != nil {
		return "", err
	}
	equipmentApp := app.createEquipmentAppById(id, equipmentInfo)
	equipmentInData := app.createEquipmentInputData(equipmentApp)
	equipment, err := Model.CreateEquipment(equipmentInData)
	if err != nil {
		return "", err
	}
	err = app.assertEquipmentExist(equipment)
	if err != nil {
		return "", err
	}
	insertedId, err := app.db.InsertNewEquipment(equipment)
	if err != nil {
		return "", err
	}
	return insertedId, nil
}

func (app *EquipmentApp) UpdateEquipmentApp(id string, equipmentInfo EditEquipmentAppData) (string, error) {
	equipmentApp := app.createEquipmentAppById(id, equipmentInfo)
	equipmentInData := app.createEquipmentInputData(equipmentApp)
	equipment, err := Model.CreateEquipment(equipmentInData)
	if err != nil {
		return "", err
	}
	err = app.assertEquipmentIdNotExist(id)
	if err != nil {
		return "", err
	}
	app.db.Lock()
	updateId, err := app.db.UpdateEquipment(equipment)
	app.db.Unlock()
	if err != nil {
		return "", err
	}
	err = app.roster.SendEvent("UPDATE", updateId)
	if err != nil {
		return "", err
	}
	return updateId, nil
}

func (app *EquipmentApp) GetAllEquipment() ([]EquipmentAppData, error) {
	var equipments []EquipmentAppData
	equipmentsDB, err := app.db.GetAllEquipments()
	if err != nil {
		return equipments, err
	}

	for i := 0; i < len(equipmentsDB); i++ {
		equipmentDB := equipmentsDB[i]
		equipmentApp := app.createEquipmentApp(equipmentDB)
		equipments = append(equipments, equipmentApp)
	}

	return equipments, nil
}

