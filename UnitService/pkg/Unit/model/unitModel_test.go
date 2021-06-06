package UnitModel

import "testing"

func createCorrectUnitInputData() UnitInputData {
	unit := UnitInputData {
		"Test_Id",
		"Test Name",
		"Test Force Name",
		4,
		3,
		 2,
		 2,
		"Additional Rule",
	}

	return unit
}

func TestSuccessCreateUnit(t *testing.T) {
	unitIn := createCorrectUnitInputData()
	_, err := CreateUnit(unitIn)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateUnitInvalidId(t *testing.T) {
	unitIn := createCorrectUnitInputData()
	unitIn.Id = ""
	_, err := CreateUnit(unitIn)
	if err != InvalidUnitId {
		t.Fatalf("Error. Create invalide id")
	}
}

func TestCreateUnitInvalidName(t *testing.T) {
	unitIn := createCorrectUnitInputData()
	unitIn.Name = ""
	_, err := CreateUnit(unitIn)
	if err != InvalidUnitName {
		t.Fatalf("Error. Create invalide name")
	}
}

func TestCreateUnitInvalidForceName(t *testing.T) {
	unitIn := createCorrectUnitInputData()
	unitIn.ForceName = ""
	_, err := CreateUnit(unitIn)
	if err != InvalidUnitForceName {
		t.Fatalf("Error. Create invalide Force name")
	}
}

func TestCreateUnitInvalidBS(t *testing.T) {
	unitIn := createCorrectUnitInputData()
	unitIn.Bs = 0
	_, err := CreateUnit(unitIn)
	if err != InvalidUnitBs {
		t.Fatalf("Error. Create invalide BS")
	}
}

func TestCreateUnitInvalidFS(t *testing.T) {
	unitIn := createCorrectUnitInputData()
	unitIn.Fs = 0
	_, err := CreateUnit(unitIn)
	if err != InvalidUnitFs {
		t.Fatalf("Error. Create invalide FS")
	}
}

func TestCreateUnitInvalidHp(t *testing.T) {
	unitIn := createCorrectUnitInputData()
	unitIn.Hp = 0
	_, err := CreateUnit(unitIn)
	if err != InvalidUnitHp {
		t.Fatalf("Error. Create invalide Hp")
	}
}

func TestCreateUnitInvalidInitiative(t *testing.T) {
	unitIn := createCorrectUnitInputData()
	unitIn.Initiative = 0
	_, err := CreateUnit(unitIn)
	if err != InvalidUnitInitiative {
		t.Fatalf("Error. Create invalide Initiative")
	}
}
