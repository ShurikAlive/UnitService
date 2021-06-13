package EquipmentModel

import "errors"

var InvalidEquipmentId = errors.New("incorrect id Equipment")
var InvalidEquipmentName = errors.New("incorrect name Equipment")
var InvalidEquipmentLimitOnUnit = errors.New("incorrect Limit On Unit Equipment")
var InvalidEquipmentLimitOnTeam = errors.New("incorrect Limit On Team Equipment")
var InvalidEquipmentAmmo = errors.New("incorrect Ammo Equipment")
var InvalidEquipmentCost = errors.New("incorrect Cost Equipment")
var InvalidEquipmentRule = errors.New("incorrect Rule Equipment")

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

func createUnit(equipmentInData EquipmentInputData) Equipment {
	equipment := Equipment{
		equipmentInData.Id,
		equipmentInData.Name,
		equipmentInData.LimitOnUnit,
		equipmentInData.LimitOnTeam,
		equipmentInData.SoldarRole,
		equipmentInData.Rule,
		equipmentInData.Ammo,
		equipmentInData.Cost,
	}

	return equipment
}

func isEmpty(param string) bool {
	return param == ""
}

func isNotNaturalNumber(param int32) bool {
	return param <= 0
}

func assertIdEmptiness(id string) error {
	if isEmpty(id) {
		return InvalidEquipmentId
	}
	return nil
}

func assertNameEmptiness(name string) error {
	if isEmpty(name) {
		return InvalidEquipmentName
	}
	return nil
}

func assertRuleEmptiness(rule string) error {
	if isEmpty(rule) {
		return InvalidEquipmentRule
	}
	return nil
}

func assertLimitOnUnitEmptiness(limitOnUnit int32) error {
	if isNotNaturalNumber(limitOnUnit) && (limitOnUnit != -1) {
		return InvalidEquipmentLimitOnUnit
	}
	return nil
}

func assertLimitOnTeamEmptiness(limitOnTeam int32) error {
	if isNotNaturalNumber(limitOnTeam) && (limitOnTeam != -1) {
		return InvalidEquipmentLimitOnTeam
	}
	return nil
}

func assertAmmoEmptiness(ammo int32) error {
	if isNotNaturalNumber(ammo) && (ammo != -1) {
		return InvalidEquipmentAmmo
	}
	return nil
}

func assertCostEmptiness(cost int32) error {
	if isNotNaturalNumber(cost) {
		return InvalidEquipmentCost
	}
	return nil
}

func CreateEquipment(equipmentInData EquipmentInputData) (Equipment, error) {
	err := assertIdEmptiness(equipmentInData.Id)
	if err != nil {
		return Equipment{}, err
	}

	err = assertNameEmptiness(equipmentInData.Name)
	if err != nil {
		return Equipment{}, err
	}

	err = assertRuleEmptiness(equipmentInData.Rule)
	if err != nil {
		return Equipment{}, err
	}

	err = assertLimitOnUnitEmptiness(equipmentInData.LimitOnUnit)
	if err != nil {
		return Equipment{}, err
	}

	err = assertLimitOnTeamEmptiness(equipmentInData.LimitOnTeam)
	if err != nil {
		return Equipment{}, err
	}

	err = assertAmmoEmptiness(equipmentInData.Ammo)
	if err != nil {
		return Equipment{}, err
	}

	err = assertCostEmptiness(equipmentInData.Cost)
	if err != nil {
		return Equipment{}, err
	}

	return createUnit(equipmentInData), nil
}