package speed

import "testing"

func TestDrive_WhenBatteryIsFull(t *testing.T) {
	car := Car{
		battery:      100,
		batteryDrain: 10,
		speed:        10,
		distance:     0,
	}

	expectedCar := Car{
		battery:      90,
		batteryDrain: 10,
		speed:        10,
		distance:     10,
	}

	testDrive(t, car, expectedCar)
}

func TestDrive_WhenBatteryIsExactlyEnough(t *testing.T) {
	car := Car{
		battery:      10,
		batteryDrain: 10,
		speed:        10,
		distance:     0,
	}

	expectedCar := Car{
		battery:      0,
		batteryDrain: 10,
		speed:        10,
		distance:     10,
	}

	testDrive(t, car, expectedCar)
}

func TestDrive_WhenBatteryIsBelow(t *testing.T) {
	car := Car{
		battery:      5,
		batteryDrain: 10,
		speed:        10,
		distance:     0,
	}

	expectedCar := Car{
		battery:      5,
		batteryDrain: 10,
		speed:        10,
		distance:     0,
	}

	testDrive(t, car, expectedCar)
}

func TestCanFinish_WhenBatteryIsLow(t *testing.T) {
	car := Car{
		battery:      20,
		batteryDrain: 10,
		speed:        2,
		distance:     0,
	}

	track := Track{
		distance: 6,
	}

	expectedCanFinish := false

	testCanFinish(t, car, track, expectedCanFinish)
}

func TestCanFinish_WhenBatteryIsEnough(t *testing.T) {
	car := Car{
		battery:      20,
		batteryDrain: 10,
		speed:        2,
		distance:     0,
	}

	track := Track{
		distance: 4,
	}

	expectedCanFinish := true

	testCanFinish(t, car, track, expectedCanFinish)
}

func TestCanFinish_WhenBatteryIsOut(t *testing.T) {
	car := Car{
		battery:      0,
		batteryDrain: 10,
		speed:        2,
		distance:     0,
	}

	track := Track{
		distance: 4,
	}

	expectedCanFinish := false

	testCanFinish(t, car, track, expectedCanFinish)
}

func testDrive(t *testing.T, car, expectedCar Car) {
	actualCar := Drive(car)

	if expectedCar != actualCar {
		t.Errorf(
			"Drive(%+v) = %+v. Expected %+v",
			car,
			actualCar,
			expectedCar)
	}
}

func testCanFinish(t *testing.T, car Car, track Track, expectedCanFinish bool) {
	actualCanFinish := CanFinish(car, track)

	if expectedCanFinish != actualCanFinish {
		t.Errorf(
			"CanFinish(%+v, %+v) = %t. Expected %t",
			car,
			track,
			actualCanFinish,
			expectedCanFinish)
	}
}
