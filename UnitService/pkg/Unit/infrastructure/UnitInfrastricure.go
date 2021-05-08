package Unit

import . "UnitService/pkg/Unit/model"

type UnitInf struct {
	// ID unit
	Id string
	// FULL NAME unit
	Name string
	// force name unit
	ForceName string
	// count heals point unit
	Hp int32
	// initiative unit
	Initiative int32
	// ability to shoot unit
	Bs int32
	// ability to fight unit
	Fs int32
	// Additionat ability soldes
	AdditionalRule string
}

func GetUnitById(id string) (UnitInf) {
	unitFromDB := GetUnitInDBById(id)
	// Сериализация unitFromDB в данные для входа CreateUnit
	CreateUnit()
	// Сереализация из возвращаемой модели в UnitInf
	// return UnitInf
}
