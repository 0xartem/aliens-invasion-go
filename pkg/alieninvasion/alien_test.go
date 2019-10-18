package alieninvasion

import "testing"

func TestMoveNotInCity(t *testing.T) {

	alien := NewAlien(0, nil)
	moved, err := alien.Move(&FakeFightObserver{})

	if moved {
		t.Error("TestMoveNotInCity failed. Alien 0 is not in a city and shouldn't have been moved")
	}

	if err == nil {
		t.Error("TestMoveNotInCity failed. An error must be returned")
	}
}

func TestMoveTrapped(t *testing.T) {

	alien := NewAlien(0, NewCity("Jakku"))
	moved, err := alien.Move(&FakeFightObserver{})

	if moved {
		t.Error("TestMoveTrapped failed. Alien 0 is trapped and shouldn't have been moved")
	}

	if err != nil {
		t.Errorf("TestMoveTrapped failed. Error returned %s", err.Error())
	}

	if !alien.IsTrapped() {
		t.Error("TestMoveTrapped failed. Alien 0 must be trapped")
	}
}

func TestMove(t *testing.T) {

	prevCity := NewCity("Naboo")
	newCity := NewCity("Tatooine")
	prevCity.SetDirection(North, newCity)

	alien := NewAlien(0, prevCity)
	prevCity.EnterCity(alien, &FakeFightObserver{})

	moved, err := alien.Move(&FakeFightObserver{})

	if !moved {
		t.Error("TestMove failed. Alien 0 must move to Tatooine")
	}

	if err != nil {
		t.Errorf("TestMove failed. Error returned %s", err.Error())
	}

	if alien.position.name != newCity.name {
		t.Errorf("TestMove failed. Alein 0 must be in Tatooine, but he is in %s", alien.position.name)
	}

	if newCity == nil || newCity.firstAlien.Id() != 0 {
		t.Error("TestMove failed. Tatooine must contain Aline 0")
	}
}
