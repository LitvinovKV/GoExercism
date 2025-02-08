package annalyn

import "testing"

func TestCanFastAttack_WhenKnightIsAwake(t *testing.T) {
	knightIsAwake := true
	expectedCanFastAttack := false

	testCanFastAttack(t, knightIsAwake, expectedCanFastAttack)
}

func TestCanFastAttack_WhenKnightIsSleeping(t *testing.T) {
	knightIsAwake := false
	expectedCanFastAttack := true

	testCanFastAttack(t, knightIsAwake, expectedCanFastAttack)
}

func TestCanSpy_WhenEveryoneIsSleeping(t *testing.T) {
	knightIsAwake := false
	archerIsAwake := false
	prisonerIsAwake := false
	expectedCanSpy := false

	testCanSpy(t, knightIsAwake, archerIsAwake, prisonerIsAwake, expectedCanSpy)
}

func TestCanSpy_WhenOnlyKnightIsAwake(t *testing.T) {
	knightIsAwake := true
	archerIsAwake := false
	prisonerIsAwake := false
	expectedCanSpy := true

	testCanSpy(t, knightIsAwake, archerIsAwake, prisonerIsAwake, expectedCanSpy)
}

func TestCanSpy_WhenOnlyArcherIsAwake(t *testing.T) {
	knightIsAwake := false
	archerIsAwake := true
	prisonerIsAwake := false
	expectedCanSpy := true

	testCanSpy(t, knightIsAwake, archerIsAwake, prisonerIsAwake, expectedCanSpy)
}

func TestCanSpy_WhenOnlyPrisonerIsAwake(t *testing.T) {
	knightIsAwake := false
	archerIsAwake := false
	prisonerIsAwake := true
	expectedCanSpy := true

	testCanSpy(t, knightIsAwake, archerIsAwake, prisonerIsAwake, expectedCanSpy)
}

func TestCanSpy_WhenEveryoneIsAwake(t *testing.T) {
	knightIsAwake := true
	archerIsAwake := true
	prisonerIsAwake := true
	expectedCanSpy := true

	testCanSpy(t, knightIsAwake, archerIsAwake, prisonerIsAwake, expectedCanSpy)
}

func TestCanSpy_WhenSeveralCharactersAreAwake(t *testing.T) {
	knightIsAwake := true
	archerIsAwake := false
	prisonerIsAwake := true
	expectedCanSpy := true

	testCanSpy(t, knightIsAwake, archerIsAwake, prisonerIsAwake, expectedCanSpy)
}

func TestCanSignalPrisoner_WhenArcherIsAwakeAndPrisonerIsSleeping(t *testing.T) {
	archerIsAwake := true
	prisonerIsAwake := false
	expectedCanSignalPrisoner := false

	testCanSignalPrisoner(t, archerIsAwake, prisonerIsAwake, expectedCanSignalPrisoner)
}

func TestCanSignalPrisoner_WhenArcherIsAwakeAndPrisonerIsAwake(t *testing.T) {
	archerIsAwake := true
	prisonerIsAwake := true
	expectedCanSignalPrisoner := false

	testCanSignalPrisoner(t, archerIsAwake, prisonerIsAwake, expectedCanSignalPrisoner)
}

func TestCanSignalPrisoner_WhenArcherIsSleepingAndPrisonerIsSleeping(t *testing.T) {
	archerIsAwake := false
	prisonerIsAwake := false
	expectedCanSignalPrisoner := false

	testCanSignalPrisoner(t, archerIsAwake, prisonerIsAwake, expectedCanSignalPrisoner)
}

func TestCanSignalPrisoner_WhenArcherIsSleepingAndPrisonerIsAwake(t *testing.T) {
	archerIsAwake := false
	prisonerIsAwake := true
	expectedCanSignalPrisoner := true

	testCanSignalPrisoner(t, archerIsAwake, prisonerIsAwake, expectedCanSignalPrisoner)
}

func TestCanFreePrisoner_WhenPetDogIsPresentAndArcherIsAwake(t *testing.T) {
	knightsIsAwake := false
	archerIsAwake := true
	prisonerIsAwake := false
	petDogIsPresent := true
	expectedCanFreePrisoner := false

	testCanFreePrisoner(t, knightsIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent, expectedCanFreePrisoner)
}

func TestCanFreePrisoner_WhenPetDogIsAbsentAndEveryoneIsSleeping(t *testing.T) {
	knightsIsAwake := false
	archerIsAwake := false
	prisonerIsAwake := false
	petDogIsPresent := false
	expectedCanFreePrisoner := false

	testCanFreePrisoner(t, knightsIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent, expectedCanFreePrisoner)
}

func TestCanFreePrisoner_WhenPetDogIsAbsentAndOnlyPrisonerIsAwake(t *testing.T) {
	knightsIsAwake := false
	archerIsAwake := false
	prisonerIsAwake := true
	petDogIsPresent := false
	expectedCanFreePrisoner := true

	testCanFreePrisoner(t, knightsIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent, expectedCanFreePrisoner)
}

func TestCanFreePrisoner_WhenPetDogIsAbsentAndEveryoneIsAwake(t *testing.T) {
	knightsIsAwake := true
	archerIsAwake := true
	prisonerIsAwake := true
	petDogIsPresent := false
	expectedCanFreePrisoner := false

	testCanFreePrisoner(t, knightsIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent, expectedCanFreePrisoner)
}

func TestCanFreePrisoner_WhenPetDogIsAbsentAndPrisonerIsAwakeAndOneGuardIsAwake(t *testing.T) {
	knightsIsAwake := true
	archerIsAwake := false
	prisonerIsAwake := true
	petDogIsPresent := false
	expectedCanFreePrisoner := false

	testCanFreePrisoner(t, knightsIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent, expectedCanFreePrisoner)
}

func testCanFastAttack(
	t *testing.T,
	knightIsAwake, expectedCanFastAttack bool) {

	actualCanFastAttack := CanFastAttack(knightIsAwake)

	if expectedCanFastAttack != actualCanFastAttack {
		t.Errorf(
			"CanFastAttack(%v) = %v. Expected %v",
			knightIsAwake,
			actualCanFastAttack,
			expectedCanFastAttack)
	}
}

func testCanSpy(
	t *testing.T,
	knightIsAwake, archerIsAwake, prisonerIsAwake, expectedCanSpy bool) {

	actualCanSpy := CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake)

	if expectedCanSpy != actualCanSpy {
		t.Errorf(
			"CanSpy(%v, %v, %v) = %v. Expected %v",
			knightIsAwake,
			archerIsAwake,
			prisonerIsAwake,
			actualCanSpy,
			expectedCanSpy)
	}
}

func testCanSignalPrisoner(
	t *testing.T,
	archerIsAwake, prisonerIsAwake, expectedCanSignalPrisoner bool) {

	actualCanSignalPrisoner := CanSignalPrisoner(archerIsAwake, prisonerIsAwake)

	if expectedCanSignalPrisoner != actualCanSignalPrisoner {
		t.Errorf(
			"CanSignalPrisoner(%v, %v) = %v. Expected %v",
			archerIsAwake,
			prisonerIsAwake,
			actualCanSignalPrisoner,
			expectedCanSignalPrisoner)
	}
}

func testCanFreePrisoner(
	t *testing.T,
	knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent, expectedCanFreePrisoner bool) {

	actualCanFreePrisoner := CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent)

	if expectedCanFreePrisoner != actualCanFreePrisoner {
		t.Errorf(
			"CanFreePrisoner(%v, %v, %v, %v) = %v. Expected %v",
			knightIsAwake,
			archerIsAwake,
			prisonerIsAwake,
			petDogIsPresent,
			actualCanFreePrisoner,
			expectedCanFreePrisoner)
	}
}
