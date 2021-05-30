package transport

import (
	. "UnitService/pkg/Unit/app"
	"encoding/json"
)

type JsonFormatter struct {

}

type UnitJSON struct {
	// ID unit
	Id string `json:"id"`
	// FULL NAME unit
	Name string `json:"name"`
	// force name unit
	ForceName string `json:"forceName"`
	// count heals point unit
	Hp int32 `json:"hp"`
	// initiative unit
	Initiative int32 `json:"initiative"`
	// ability to shoot unit
	Bs int32 `json:"bs"`
	// ability to fight unit
	Fs int32 `json:"fs"`
	// Additionat ability soldes
	AdditionalRule string `json:"additionalRule"`
}

type EditUnitJSON struct {
	// FULL NAME unit
	Name string `json:"name"`
	// force name unit
	ForceName string `json:"forceName"`
	// count heals point unit
	Hp int32 `json:"hp"`
	// initiative unit
	Initiative int32 `json:"initiative"`
	// ability to shoot unit
	Bs int32 `json:"bs"`
	// ability to fight unit
	Fs int32 `json:"fs"`
	// Additionat ability soldes
	AdditionalRule string `json:"additionalRule"`
}

func CreateJSONFormatter() JsonFormatter {
	return JsonFormatter{}
}

func (formatter *JsonFormatter) convertAllUnitAppToAllUnitJson(unitsApp []UnitAppData) ([]UnitJSON) {
	units := []UnitJSON {}
	for i := 0; i < len(unitsApp); i++ {
		unitApp := unitsApp[i]
		unit := UnitJSON{
			unitApp.Id,
			unitApp.Name,
			unitApp.ForceName,
			unitApp.Hp,
			unitApp.Initiative,
			unitApp.Bs,
			unitApp.Fs,
			unitApp.AdditionalRule,
		}

		units = append(units, unit)
	}

	return units
}

func (formatter *JsonFormatter) convertUnitAppToAllUnitJson(unitApp UnitAppData) UnitJSON {
	unit := UnitJSON{
		unitApp.Id,
		unitApp.Name,
		unitApp.ForceName,
		unitApp.Hp,
		unitApp.Initiative,
		unitApp.Bs,
		unitApp.Fs,
		unitApp.AdditionalRule,
	}

	return unit
}

func (formatter *JsonFormatter) convertEditUnitJSONToUnitEditAppData(unit EditUnitJSON) UnitEditAppData {
	unitApp := UnitEditAppData{
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule,
	}

	return unitApp
}

func (formatter *JsonFormatter) ConvertIdToJSON (id string) string {
	return "\"" + id + "\""
}

func (formatter *JsonFormatter) ConvertAllUnitAppDataToJSON (units []UnitAppData)  ([]byte, error) {
	unitsJson := formatter.convertAllUnitAppToAllUnitJson(units)
	b, err := json.Marshal(unitsJson)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (formatter *JsonFormatter) ConvertUnitAppDataToJSON (unit UnitAppData)  ([]byte, error) {
	unitJson := formatter.convertUnitAppToAllUnitJson(unit)
	b, err := json.Marshal(unitJson)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func (formatter *JsonFormatter) ConvertJsonToUnitEditAppData (unitJson []byte) (UnitEditAppData, error) {
	var msg EditUnitJSON
	err := json.Unmarshal(unitJson, &msg)
	if err != nil {
		return UnitEditAppData{}, err
	}

	return formatter.convertEditUnitJSONToUnitEditAppData(msg), nil
}






