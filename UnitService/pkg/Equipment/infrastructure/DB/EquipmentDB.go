package EquipmentMySQLDB

type EquipmentDB struct {
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

type EquipmentInputDB struct {
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

type RequiredParameters struct {
	// FULL NAME unit
	Name string
	// cost equipment in game points
	Cost int32
}

type IEquipmentDB interface {
	GetEquipmentInDBById(id string) (EquipmentDB, error)
	GetEquipmentInDBByRequiredParameters(equipmentParams RequiredParameters) (EquipmentDB, error)
	GetAllEquipments() ([]EquipmentDB, error)
	InsertNewEquipment(equipment EquipmentInputDB) (string, error)
	UpdateEquipment(equipment EquipmentInputDB) (string, error)
	DeleteEquipment(id string) (string, error)
}

