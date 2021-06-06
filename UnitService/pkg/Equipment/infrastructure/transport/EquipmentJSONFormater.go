package EquipmentTransport

import (
	App "UnitService/pkg/equipment/app"
	"encoding/json"
)

type JsonFormatter struct {

}

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

func CreateJSONFormatter() JsonFormatter {
	return JsonFormatter{}
}


func (formatter *JsonFormatter) createAllEquipmentsJson(equipmentsApp []App.EquipmentAppData) []EquipmentJSON {
	var equipments []EquipmentJSON
	for i := 0; i < len(equipmentsApp); i++ {
		equipmentApp := equipmentsApp[i]
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

func (formatter *JsonFormatter) createAllEquipmentJson(equipmentApp App.EquipmentAppData) EquipmentJSON {
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

func (formatter *JsonFormatter) createEquipmentEditAppData(equipment EditEquipmentJSON) App.EditEquipmentAppData {
	equipmentApp := App.EditEquipmentAppData{
		equipment.Name,
		equipment.LimitOnUnit,
		equipment.LimitOnTeam,
		equipment.SoldarRole,
		equipment.Rule,
		equipment.Ammo,
		equipment.Cost,
	}

	return equipmentApp
}

func (formatter *JsonFormatter) ConvertIdToJSON (id string) string {
	return "\"" + id + "\""
}

func (formatter *JsonFormatter) ConvertAllEquipmentAppDataToJSON (equipments []App.EquipmentAppData)  ([]byte, error) {
	equipmentsJson := formatter.createAllEquipmentsJson(equipments)
	b, err := json.Marshal(equipmentsJson)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (formatter *JsonFormatter) ConvertEquipmentAppDataToJSON (equipment App.EquipmentAppData)  ([]byte, error) {
	equipmentJson := formatter.createAllEquipmentJson(equipment)
	b, err := json.Marshal(equipmentJson)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (formatter *JsonFormatter) ConvertJsonToEquipmentEditAppData (equipmentJson []byte) (App.EditEquipmentAppData, error) {
	var msg EditEquipmentJSON
	err := json.Unmarshal(equipmentJson, &msg)
	if err != nil {
		return App.EditEquipmentAppData{}, err
	}

	return formatter.createEquipmentEditAppData(msg), nil
}

