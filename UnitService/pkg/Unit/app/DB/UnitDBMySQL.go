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

	query := "SELECT * FROM unit_db.units where name = '" + requiredParams.Name + "' AND ForceName = '" + requiredParams.ForceName + "';";
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
	db := DB.GetDBInstance()
	if db == nil {
		return []UnitMySQL {}, errors.New("Error conection DB")
	}

	rows, err := db.Query("SELECT * FROM unit_db.units;")
	if err != nil {
		return []UnitMySQL {}, err
	}
	defer rows.Close()

	units := []UnitMySQL {}

	for rows.Next() {
		unitDBMySQL := UnitDBMySQL {}
		err = rows.Scan(&unitDBMySQL.Id,
			&unitDBMySQL.Name,
			&unitDBMySQL.ForceName,
			&unitDBMySQL.Hp,
			&unitDBMySQL.Initiative,
			&unitDBMySQL.Bs,
			&unitDBMySQL.Fs,
			&unitDBMySQL.AdditionalRule)
		if err != nil {
			return []UnitMySQL {}, err
		}
		units = append(units, SerialitheUnitDBMySQLByUnitMySQL(unitDBMySQL))
	}

	return units, nil
}

func GetUnitInMySQLById(id string) (UnitMySQL, error) {
	db := DB.GetDBInstance()
	if db == nil {
		return UnitMySQL {}, errors.New("Error conection DB")
	}

	query := "SELECT * FROM unit_db.units where id = '" + id + "';";
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

func InsertNewUnitInMySQL(unit UnitInputMySQL) (string, error) {
	db := DB.GetDBInstance()
	if db == nil {
		return "", errors.New("Error conection DB")
	}

	_, err := db.Exec("INSERT INTO `unit_db`.`units` 	(`id`, `Name`, `ForceName`,	`Hp`, `Initiative`, `Bs`, `Fs`, `AdditionalRule`) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		unit.Id,
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule)

	if err != nil {
		return "", err
	}

	return unit.Id, nil
}

func UpdateUnitInMySQL(unit UnitInputMySQL) (string, error) {
	db := DB.GetDBInstance()
	if db == nil {
		return "", errors.New("Error conection DB")
	}

	_, err := db.Exec("UPDATE `unit_db`.`units` SET `Name` = ?, `ForceName` = ?, `Hp` = ?, `Initiative` = ?, `Bs` = ?, `Fs` = ?, `AdditionalRule` = ? WHERE id = ?;",
		unit.Name,
		unit.ForceName,
		unit.Hp,
		unit.Initiative,
		unit.Bs,
		unit.Fs,
		unit.AdditionalRule,
		unit.Id)

	if err != nil {
		return "", err
	}

	return unit.Id, nil
}

func DeleteUnitInMySQL(id string) (string, error) {
	db := DB.GetDBInstance()
	if db == nil {
		return "", errors.New("Error conection DB")
	}

	_, err := db.Exec("DELETE FROM `unit_db`.`units` WHERE id = ?;", id)

	if err != nil {
		return "", err
	}

	return id, nil
}

