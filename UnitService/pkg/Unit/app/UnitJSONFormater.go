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

func GetJSONUnitById(unitId string) ([]byte, error) {
	unitInf := GetUnitById(unitId)
	unit := SerializationUnit(unitInf)
	b, err := json.Marshal(unit)

	if err != nil {
		return nil, err
	}

	return b, nil
}
