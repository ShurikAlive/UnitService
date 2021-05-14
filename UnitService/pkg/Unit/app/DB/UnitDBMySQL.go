package DB

type UnitMySQL struct {
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

type UnitInputMySQL struct {
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

func GetAllUnitsInMySQL() ([]UnitMySQL, error) {
	// TODO SELECT * FROM
	units := []UnitMySQL {}
	testUnit := UnitMySQL{
		"c01d7cf6-ec3f-47f0-9556-a5d6e9009a43",
		"Gregiory W. Morris",
		"SPECIAL FORCES AIRBORNE",
		4,
		8,
		1,
		1,
		"When Morris move into hand to hand combat his roll is at +2.",
	}

	units = append(units, testUnit)
	units = append(units, testUnit)
	units = append(units, testUnit)

	return units, nil
}

func GetUnitInMySQLById(id string) (UnitMySQL, error) {
	// TODO SELECT * WHERE unitId = id ...
	return UnitMySQL{
		"c01d7cf6-ec3f-47f0-9556-a5d6e9009a43",
		"Gregiory W. Morris",
		"SPECIAL FORCES AIRBORNE",
		4,
		8,
		1,
		1,
		"When Morris move into hand to hand combat his roll is at +2.",
	}, nil
}

func InsertNewUnitInMySQL(unit UnitInputMySQL) (string, error) {
	// TODO INSERT INTO ...

	return "c01d7cf6-ec3f-47f0-9556-a5d6e9009a43", nil
}

func UpdateUnitInMySQL(unit UnitInputMySQL) (string, error) {
	// TODO UPDATE ...

	return "c01d7cf6-ec3f-47f0-9556-a5d6e9009a43", nil
}

func DeleteUnitInMySQL(id string) (string, error) {
	// TODO DELETE ...

	return "c01d7cf6-ec3f-47f0-9556-a5d6e9009a43", nil
}

func UnitExistMySQL(id string) (bool) {
	// TODO SELECT WHERE unitId = id

	return true
}