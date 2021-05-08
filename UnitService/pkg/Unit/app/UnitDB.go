package Unit

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

func GetUnitInDBById(id string) (UnitDB) {
	unitMySQL := GetUnitInMySQLById(id)
	return SerializationUnitDB(unitMySQL)
}
