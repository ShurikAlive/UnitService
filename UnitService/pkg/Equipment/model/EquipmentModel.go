package EquipmentModel

import "errors"

var InvalidEquipmentId = errors.New("incorrect id Equipment")
var InvalidEquipmentName = errors.New("Incorrect name Equipment")
var InvalidEquipmentLimitOnUnit = errors.New("Incorrect Limit On Unit Equipment")
var InvalidEquipmentLimitOnTeam = errors.New("Incorrect Limit On Team Equipment")
var InvalidEquipmentAmmo = errors.New("Incorrect Ammo Equipment")
var InvalidEquipmentCost = errors.New("Incorrect Cost Equipment")

type Equipment struct {
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

type EquipmentInputData struct {
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

func ConvertEquipmentInputDataToUnit(equipmentInData EquipmentInputData) Equipment {
	unit := Equipment{
		equipmentInData.Id,
		equipmentInData.Name,
		equipmentInData.LimitOnUnit,
		equipmentInData.LimitOnTeam,
		equipmentInData.SoldarRole,
		equipmentInData.Rule,
		equipmentInData.Ammo,
		equipmentInData.Cost,
	}

	return unit
}

func IsEmpty(param string) bool {
	return param == ""
}

func IsNotNaturalNumber(param int32) bool {
	return param <= 0
}

func CreateEquipment(equipmentInData EquipmentInputData) (Equipment, error) {
	if IsEmpty(equipmentInData.Id) {
		return Equipment{}, InvalidEquipmentId
	}

	if IsEmpty(equipmentInData.Name) {
		return Equipment{}, InvalidEquipmentName
	}

	if IsNotNaturalNumber(equipmentInData.LimitOnUnit) || (equipmentInData.LimitOnUnit != -1) {
		return Equipment{}, InvalidEquipmentLimitOnUnit
	}

	if IsNotNaturalNumber(equipmentInData.LimitOnTeam) || (equipmentInData.LimitOnTeam != -1) {
		return Equipment{}, InvalidEquipmentLimitOnTeam
	}

	if IsNotNaturalNumber(equipmentInData.Ammo) || (equipmentInData.Ammo != -1) {
		return Equipment{}, InvalidEquipmentAmmo
	}

	if IsNotNaturalNumber(equipmentInData.Cost) {
		return Equipment{}, InvalidEquipmentCost
	}

	return ConvertEquipmentInputDataToUnit(equipmentInData), nil
}