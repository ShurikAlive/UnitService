package UnitModel

import "errors"

var InvalidUnitName = errors.New("incorrect name unit")
var InvalidUnitForceName = errors.New("incorrect force name unit")
var InvalidUnitId = errors.New("incorrect id unit")
var InvalidUnitHp = errors.New("incorrect hp value unit")
var InvalidUnitFs = errors.New("incorrect Fs value unit")
var InvalidUnitBs = errors.New("incorrect Bs value unit")
var InvalidUnitInitiative = errors.New("incorrect Initiative value unit")

type Unit struct {
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

type UnitInputData struct {
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

func isEmpty(param string) bool {
	return param == ""
}

func isNotNaturalNumber(param int32) bool {
	return param <= 0
}

func assertIdEmptiness(id string) error {
	if isEmpty(id) {
		return InvalidUnitId
	}
	return nil
}

func assertNameEmptiness(name string) error {
	if isEmpty(name) {
		return InvalidUnitName
	}
	return nil
}

func assertForceNameEmptiness(forceName string) error {
	if isEmpty(forceName) {
		return InvalidUnitForceName
	}
	return nil
}

func assertHpIncorrect(hp int32) error {
	if isNotNaturalNumber(hp) {
		return InvalidUnitHp
	}
	return nil
}

func assertFsIncorrect(fs int32) error {
	if isNotNaturalNumber(fs) {
		return InvalidUnitFs
	}
	return nil
}

func assertBsIncorrect(bs int32) error {
	if isNotNaturalNumber(bs) {
		return InvalidUnitBs
	}
	return nil
}

func assertInitiativeIncorrect(initiative int32) error {
	if isNotNaturalNumber(initiative) {
		return InvalidUnitInitiative
	}
	return nil
}

func ConvertUnitInputDataToUnit(unitInData UnitInputData) Unit {
	unit := Unit{
		unitInData.Id,
		unitInData.Name,
		unitInData.ForceName,
		unitInData.Hp,
		unitInData.Initiative,
		unitInData.Bs,
		unitInData.Fs,
		unitInData.AdditionalRule,
	}

	return unit
}

func CreateUnit(unitInData UnitInputData) (Unit, error) {
	err := assertIdEmptiness(unitInData.Id)
	if err != nil {
		return Unit{}, err
	}

	err = assertNameEmptiness(unitInData.Name)
	if err != nil {
		return Unit{}, err
	}

	err = assertForceNameEmptiness(unitInData.ForceName)
	if err != nil {
		return Unit{}, err
	}

	err = assertHpIncorrect(unitInData.Hp)
	if err != nil {
		return Unit{}, err
	}

	err = assertBsIncorrect(unitInData.Bs)
	if err != nil {
		return Unit{}, err
	}

	err = assertFsIncorrect(unitInData.Fs)
	if err != nil {
		return Unit{}, err
	}

	err =  assertInitiativeIncorrect(unitInData.Initiative)
	if err != nil {
		return Unit{}, err
	}

	return ConvertUnitInputDataToUnit(unitInData), nil
}
