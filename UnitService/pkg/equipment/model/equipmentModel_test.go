package EquipmentModel

import "testing"

func createCorrectEquipmentInputData() EquipmentInputData {
	equipment := EquipmentInputData {
		"test Id",
		"test Name",
		2,
		4,
		"",
		"test Rule",
		-1,
		2,
	}

	return equipment
}

func TestSuccessCreateEquipment(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	_, err := CreateEquipment(equipmentIn)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateEquipmentInvalidId(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.Id = ""
	_, err := CreateEquipment(equipmentIn)
	if err != InvalidEquipmentId {
		t.Fatalf("Error. Create invalide id")
	}
}

func TestCreateEquipmentInvalidName(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.Name = ""
	_, err := CreateEquipment(equipmentIn)
	if err != InvalidEquipmentName {
		t.Fatalf("Error. Create invalide name")
	}
}

func TestCreateEquipmentInvalidLimitOnUnit(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.LimitOnUnit = 0
	_, err := CreateEquipment(equipmentIn)
	if err != InvalidEquipmentLimitOnUnit {
		t.Fatalf("Error. Create invalide Limit On Unit")
	}
}

func TestSuccessCreateEquipmentWithUnlimitOnUnit(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.LimitOnUnit = -1
	_, err := CreateEquipment(equipmentIn)
	if err == InvalidEquipmentLimitOnUnit {
		t.Fatal(err)
	}
}

func TestCreateEquipmentInvalidLimitOnRoster(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.LimitOnTeam = 0
	_, err := CreateEquipment(equipmentIn)
	if err != InvalidEquipmentLimitOnTeam {
		t.Fatalf("Error. Create invalide Limit On Team")
	}
}

func TestSuccessCreateEquipmentWithUnlimitOnRoster(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.LimitOnTeam = -1
	_, err := CreateEquipment(equipmentIn)
	if err == InvalidEquipmentLimitOnTeam {
		t.Fatal(err)
	}
}

func TestCreateEquipmentInvalidAmmo(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.Ammo = 0
	_, err := CreateEquipment(equipmentIn)
	if err != InvalidEquipmentAmmo {
		t.Fatalf("Error. Create invalide Ammo")
	}
}

func TestSuccessCreateEquipmentWithUnlimitAmmo(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.Ammo = -1
	_, err := CreateEquipment(equipmentIn)
	if err == InvalidEquipmentAmmo {
		t.Fatal(err)
	}
}

func TestCreateEquipmentInvalidCost(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.Cost = -3
	_, err := CreateEquipment(equipmentIn)
	if err != InvalidEquipmentCost {
		t.Fatalf("Error. Create invalide Cost")
	}
}

func TestCreateEquipmentInvalidRule(t *testing.T) {
	equipmentIn := createCorrectEquipmentInputData()
	equipmentIn.Rule = ""
	_, err := CreateEquipment(equipmentIn)
	if err != InvalidEquipmentRule {
		t.Fatalf("Error. Create invalide Rule")
	}
}