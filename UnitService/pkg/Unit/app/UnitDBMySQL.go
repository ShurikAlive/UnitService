package Unit

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

func GetUnitInMySQLById(id string) (UnitMySQL) {
	// TODO SELECT * FROM ...
	return UnitMySQL{}
}
