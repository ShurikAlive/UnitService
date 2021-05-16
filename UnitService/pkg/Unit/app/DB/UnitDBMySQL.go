package DB

import (
	DB "UnitService/pkg/DB/MySQLDB"
	"errors"
)

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

type UnitDBMySQL struct {
	// ID unit
	Id string `db:"id"`
	// FULL NAME unit
	Name string `db:"Name"`
	// force name unit
	ForceName string `db:"ForceName"`
	// count heals point unit
	Hp int32 `db:"Hp"`
	// initiative unit
	Initiative int32 `db:"Initiative"`
	// ability to shoot unit
	Bs int32 `db:"Bs"`
	// ability to fight unit
	Fs int32 `db:"Fs"`
	// Additionat ability soldes
	AdditionalRule string `db:"AdditionalRule"`
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

type RequiredParametersMySQL struct {
	// FULL NAME unit
	Name string
	// force name unit
	ForceName string
}

func SerialitheUnitDBMySQLByUnitMySQL(unitDB UnitDBMySQL) (UnitMySQL) {
	unit := UnitMySQL {
		unitDB.Id,
		unitDB.Name,
		unitDB.ForceName,
		unitDB.Hp,
		unitDB.Initiative,
		unitDB.Bs,
		unitDB.Fs,
		unitDB.AdditionalRule,
	}

	return unit
}

func GetUnitInMySQLByRequiredParameters(requiredParams RequiredParametersMySQL) (UnitMySQL, error) {
	db := DB.GetDBInstance()
	if db == nil {
		return UnitMySQL {}, errors.New("Error conection DB")
	}

	query := "query := SELECT * FROM unit_db.units where name = '" + requiredParams.Name + "' AND ForceName = '" + requiredParams.ForceName + "';";
	rows, err := db.Query(query)
	if err != nil {
		return UnitMySQL {}, err
	}
	defer rows.Close()

	unitDBMySQL := UnitDBMySQL {}
	for rows.Next() {
		err = rows.Scan(&unitDBMySQL.Id,
			&unitDBMySQL.Name,
			&unitDBMySQL.ForceName,
			&unitDBMySQL.Hp,
			&unitDBMySQL.Initiative,
			&unitDBMySQL.Bs,
			&unitDBMySQL.Fs,
			&unitDBMySQL.AdditionalRule)
		if err != nil {
			return UnitMySQL {}, err
		}
	}

	unitDB := UnitDBMySQL {}
	if unitDBMySQL.Id == unitDB.Id {
		return UnitMySQL {}, errors.New("Record Not Found")
	}

	return SerialitheUnitDBMySQLByUnitMySQL(unitDBMySQL), nil
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
	db := DB.GetDBInstance()
	if db == nil {
		return UnitMySQL {}, errors.New("Error conection DB")
	}

	query := "query := SELECT * FROM unit_db.units where id = '" + id + "';";
	rows, err := db.Query(query)
	if err != nil {
		return UnitMySQL {}, err
	}
	defer rows.Close()

	unitDBMySQL := UnitDBMySQL {}
	for rows.Next() {
		err = rows.Scan(&unitDBMySQL.Id,
			&unitDBMySQL.Name,
			&unitDBMySQL.ForceName,
			&unitDBMySQL.Hp,
			&unitDBMySQL.Initiative,
			&unitDBMySQL.Bs,
			&unitDBMySQL.Fs,
			&unitDBMySQL.AdditionalRule)
		if err != nil {
			return UnitMySQL {}, err
		}
	}

	unitDB := UnitDBMySQL {}
	if unitDBMySQL.Id == unitDB.Id {
		return UnitMySQL {}, errors.New("Record Not Found")
	}

	return SerialitheUnitDBMySQLByUnitMySQL(unitDBMySQL), nil
/*
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
 */
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

