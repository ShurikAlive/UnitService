package Unit

import (
	. "UnitService/pkg/Unit/infrastructure"
	"encoding/json"
)

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

func SerializationUnit(unitInf UnitInf) (UnitJSON) {
	unit := UnitJSON{
		unitInf.Id,
		unitInf.Name,
		unitInf.ForceName,
		unitInf.Hp,
		unitInf.Initiative,
		unitInf.Bs,
		unitInf.Fs,
		unitInf.AdditionalRule,
	}

	return unit
}

func SerializationAllUnit(unitsInf []UnitInf) ([]UnitJSON) {
	units := []UnitJSON {}
	for i := 0; i < len(unitsInf); i++ {
		unitInf := unitsInf[i]
		unit := UnitJSON{
			unitInf.Id,
			unitInf.Name,
			unitInf.ForceName,
			unitInf.Hp,
			unitInf.Initiative,
			unitInf.Bs,
			unitInf.Fs,
			unitInf.AdditionalRule,
		}

		units = append(units, unit)
	}

	return units
}

func SerializationEditUnit(unitEditJson EditUnitJSON) (UnitEditInf) {
	unitEdit := UnitEditInf {
		unitEditJson.Name,
		unitEditJson.ForceName,
		unitEditJson.Hp,
		unitEditJson.Initiative,
		unitEditJson.Bs,
		unitEditJson.Fs,
		unitEditJson.AdditionalRule,
	}

	return unitEdit
}

func GetJSONUnitById(unitId string) ([]byte, error) {
	unitInf, err := GetUnitById(unitId)
	if err != nil {
		return nil, err
	}

	units := SerializationUnit(unitInf)
	b, err := json.Marshal(units)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func GetJSONAllUnitById() ([]byte, error) {
	allUnits, err := GetAllUnit()
	if err != nil {
		return nil, err
	}

	if len(allUnits.IdsWithError) > 0 {
		// TODO ЗАПИСЫВАЕМ В ЛОГ ТАКИЕ ID's ВОЗМОЖНО ГДЕТО ИНФОРМИРУЕМ АДМИНОВ!!!
	}

	units := SerializationAllUnit(allUnits.Units)
	b, err := json.Marshal(units)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func AddUnit(unitJson []byte) (string, error) {
	var msg EditUnitJSON
	err := json.Unmarshal(unitJson, &msg)
	if err != nil {
		return "", err
	}

	unitEditInf:= SerializationEditUnit(msg)
	id, err := AddNewUnit(unitEditInf)
	if err != nil {
		return "", err
	}

	return id, nil
}

func UpdateUnit(id string, unitJson []byte) (string, error) {
	var msg EditUnitJSON
	err := json.Unmarshal(unitJson, &msg)
	if err != nil {
		return "", err
	}

	unitEditInf:= SerializationEditUnit(msg)
	idEdit, err := UpdateUnitInf(id, unitEditInf)
	if err != nil {
		return "", err
	}

	return idEdit, nil
}

func DeleteById(id string) (string, error) {
	deleteId, err := DeleteByIdInf(id)
	if err != nil {
		return "", err
	}

	return deleteId, nil
}