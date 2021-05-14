package DB

type UnitDB struct {
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

type UnitInputDB struct {
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

func SerializationUnitDB(unitMySQL UnitMySQL) (UnitDB) {
	unit := UnitDB {
		unitMySQL.Id,
		unitMySQL.Name,
		unitMySQL.ForceName,
		unitMySQL.Hp,
		unitMySQL.Initiative,
		unitMySQL.Bs,
		unitMySQL.Fs,
		unitMySQL.AdditionalRule,
	}

	return unit
}

func SerialitheUnitInputMySQL(unit UnitInputDB) (UnitInputMySQL) {
	unitMySQL := UnitInputMySQL {
		unit.Id,
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule,
	}

	return unitMySQL
}

func GetUnitInDBById(id string) (UnitDB, error) {
	unitMySQL, err := GetUnitInMySQLById(id)
	if err != nil {
		return UnitDB{}, err
	}
	return SerializationUnitDB(unitMySQL), nil
}

func GetAllUnits() ([]UnitDB, error) {
	unitsMySQL, err := GetAllUnitsInMySQL()
	if err != nil {
		return []UnitDB {}, err
	}

	units := []UnitDB {}

	for i := 0; i < len(unitsMySQL); i++ {
		unitMySQL := unitsMySQL[i]
		unitDB := SerializationUnitDB(unitMySQL)
		units = append(units, unitDB)
	}

	return units, nil
}

func InsertNewUnit(unit UnitInputDB) (string, error) {
	unitInputMySQL := SerialitheUnitInputMySQL(unit)
	id, err := InsertNewUnitInMySQL(unitInputMySQL)
	if err != nil {
		return "", err
	}
	return id, nil
}

func UpdateUnit(unit UnitInputDB) (string, error) {
	unitUpdateMySQL := SerialitheUnitInputMySQL(unit)
	id, err := UpdateUnitInMySQL(unitUpdateMySQL)
	if err != nil {
		return "", err
	}
	return id, nil
}

func DeleteUnit(id string) (string, error) {
	deleteId, err := DeleteUnitInMySQL(id)
	if err != nil {
		return "", err
	}
	return deleteId, nil
}

func UnitExist(id string) (bool) {
	isExist := UnitExistMySQL(id)
	return isExist
}