package UnitMySQLDB


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

type RequiredParameters struct {
	// FULL NAME unit
	Name string
	// force name unit
	ForceName string
}

type IUnitDB interface {
	GetUnitInDBById(id string) (UnitDB, error)
	GetUnitInDBByRequiredParameters(unitParams RequiredParameters) (UnitDB, error)
	GetAllUnits() ([]UnitDB, error)
	InsertNewUnit(unit UnitInputDB) (string, error)
	UpdateUnit(unit UnitInputDB) (string, error)
	DeleteUnit(id string) (string, error)
}

