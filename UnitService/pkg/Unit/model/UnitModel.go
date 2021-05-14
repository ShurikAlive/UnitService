package Unit

import "errors"

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
		return Unit{}, errors.New("Incorrect id unit")
	}

	if IsEmpty(unitInData.Name) {
		return Unit{}, errors.New("Incorrect name unit")
	}

	if IsEmpty(unitInData.ForceName)  {
		return Unit{}, errors.New("Incorrect force name unit")
	}

	if IsNotNaturalNumber(unitInData.Hp) {
		return Unit{}, errors.New("Incorrect hp value unit")
	}

	if IsNotNaturalNumber(unitInData.Fs) {
		return Unit{}, errors.New("Incorrect Fs value unit")
	}

	if IsNotNaturalNumber(unitInData.Bs) {
		return Unit{}, errors.New("Incorrect Bs value unit")
	}

	if IsNotNaturalNumber(unitInData.Initiative) {
		return Unit{}, errors.New("Incorrect Initiative value unit")
	}

	return SerializationUnit(unitInData), nil
}
