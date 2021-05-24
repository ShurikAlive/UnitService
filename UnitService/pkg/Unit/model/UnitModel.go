package UnitModel

import "errors"

var InvalidUnitName = errors.New("incorrect name unit")
var InvalidUnitForceName = errors.New("Incorrect force name unit")
var InvalidUnitId = errors.New("Incorrect id unit")
var InvalidUnitHp = errors.New("Incorrect hp value unit")
var InvalidUnitFs = errors.New("Incorrect Fs value unit")
var InvalidUnitBs = errors.New("Incorrect Bs value unit")
var InvalidUnitInitiative = errors.New("Incorrect Initiative value unit")

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

func IsEmpty(param string) bool {
	return param == ""
}

func IsNotNaturalNumber(param int32) bool {
	return param <= 0
}

func SerializationUnit(unitInData UnitInputData) (Unit) {
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
	if IsEmpty(unitInData.Id) {
		return Unit{}, InvalidUnitId
	}

	if IsEmpty(unitInData.Name) {
		return Unit{}, InvalidUnitName
	}

	if IsEmpty(unitInData.ForceName)  {
		return Unit{}, InvalidUnitForceName
	}

	if IsNotNaturalNumber(unitInData.Hp) {
		return Unit{}, InvalidUnitHp
	}

	if IsNotNaturalNumber(unitInData.Fs) {
		return Unit{}, InvalidUnitFs
	}

	if IsNotNaturalNumber(unitInData.Bs) {
		return Unit{}, InvalidUnitBs
	}

	if IsNotNaturalNumber(unitInData.Initiative) {
		return Unit{}, InvalidUnitInitiative
	}

	return SerializationUnit(unitInData), nil
}
