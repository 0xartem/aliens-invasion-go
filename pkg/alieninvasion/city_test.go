package alieninvasion

import "testing"

type FakeFightObserver struct {
	called int
}

func (f *FakeFightObserver) OnAliensFight(fstAl *Alien, secAl *Alien, city *City) {
	f.called++
}

func TestIsIsolated(t *testing.T) {
	city := NewCity("Naboo")
	if !city.IsIsolated() {
		t.Error("TestIsIsolated failed. City Naboo must be isolated")
	}
}

func TestName(t *testing.T) {
	cityName := "Tatooine"
	city := NewCity(cityName)
	if city.Name() != cityName {
		t.Errorf("TestName failed. City name %s is incorrect", cityName)
	}
}

func TestSetDirection(t *testing.T) {
	city := NewCity("Naboo")
	connCity := NewCity("Tatooine")
	city.SetDirection(North, connCity)

	_, ok := city.directions[North]
	if !ok {
		t.Fatal("TestSetDirection failed. North direction doesn't exist")
	}

	if city.directions[North].name != connCity.name {
		t.Errorf("TestSetDirection failed. The connected city is %s. Must be Tatooine", city.directions[North].name)
	}
}

func TestDirectionExists(t *testing.T) {
	city := NewCity("Naboo")
	connCity := NewCity("Tatooine")
	city.SetDirection(South, connCity)

	if !city.DirectionExists(South) {
		t.Error("TestDirectionExists failed. South direction doesn't exist")
	}
}

func TestGetRandomDirection(t *testing.T) {
	city := NewCity("Naboo")
	city.SetDirection(North, NewCity("Tatooine"))
	city.SetDirection(South, NewCity("Jupiter"))
	city.SetDirection(West, NewCity("Jakku"))
	city.SetDirection(East, NewCity("SanJuperdino"))

	testMap := map[string]bool{"Tatooine": false, "Jupiter": false, "Jakku": false, "SanJuperdino": false}
	for i := 0; i < 50; i++ {
		randCity := city.GetRandomDirection()
		_, ok := testMap[randCity.name]

		if !ok {
			t.Errorf("TestGetRandomDirection failed. Direction %s doesn't exist", randCity.name)
		}

		testMap[randCity.name] = true
	}

	for name, ok := range testMap {
		if !ok {
			t.Errorf("TestGetRandomDirection failed. The city %s was not generated at random after 50 tries", name)
		}
	}
}

func TestEnterCityError(t *testing.T) {
	city := NewCity("Jakku")
	city.secondAlien = NewAlien(0, city)
	err := city.EnterCity(NewAlien(1, city), &FakeFightObserver{})
	if err == nil {
		t.Errorf("TestEnterCityError failed. If the second alien is in the city - an error must be returned")
	}
}

func TestEnterCityFirst(t *testing.T) {
	city := NewCity("Jakku")
	firstAlien := NewAlien(0, city)

	fakeObserver := FakeFightObserver{}
	err := city.EnterCity(firstAlien, &fakeObserver)

	if err != nil {
		t.Errorf("TestEnterCityFirst failed. Error returned %s", err.Error())
	}

	if firstAlien != city.firstAlien {
		t.Error("TestEnterCityFirst failed. The first alien is different")
	}

	if city.secondAlien != nil {
		t.Error("TestEnterCityFirst failed. The second alien is not nil")
	}

	if fakeObserver.called != 0 {
		t.Error("TestEnterCityFirst failed. FightObserver was called")
	}
}

func TestEnterCitySecond(t *testing.T) {
	city := NewCity("Jakku")
	firstAlien := NewAlien(0, city)
	secondAlien := NewAlien(1, city)

	fakeObserver := FakeFightObserver{}
	err := city.EnterCity(firstAlien, &fakeObserver)
	if err != nil {
		t.Errorf("TestEnterCitySecond failed. Error returned %s", err.Error())
	}

	err = city.EnterCity(secondAlien, &fakeObserver)
	if err != nil {
		t.Errorf("TestEnterCitySecond failed. Error returned %s", err.Error())
	}

	if firstAlien != city.firstAlien {
		t.Error("TestEnterCitySecond failed. The first alien is different")
	}

	if secondAlien != city.secondAlien {
		t.Error("TestEnterCitySecond failed. The second alien is different")
	}

	if fakeObserver.called != 1 {
		t.Error("TestEnterCityFirst failed. FightObserver was not called")
	}
}

func TestLeaveCity(t *testing.T) {
	city := NewCity("Jakku")
	alien := NewAlien(0, city)

	err := city.EnterCity(alien, &FakeFightObserver{})
	if err != nil {
		t.Errorf("TestLeaveCity failed. Error returned %s", err.Error())
	}

	err = city.LeaveCity(alien)
	if err != nil {
		t.Errorf("TestLeaveCity failed. Error returned %s", err.Error())
	}

	if city.firstAlien != nil || city.secondAlien != nil {
		t.Error("TestLeaveCity failed. An alien is still in the city")
	}
}

func TestRemoveConnection(t *testing.T) {
	city := NewCity("Naboo")
	city.SetDirection(North, NewCity("Tatooine"))
	city.SetDirection(South, NewCity("Jupiter"))

	ok := city.RemoveConnection(North, "Tatooine")
	if !ok {
		t.Error("TestRemoveConnection failed. The connection Tatooine was not removed")
	}

	ok = city.RemoveConnection(South, "Jupiter")
	if !ok {
		t.Error("TestRemoveConnection failed. The connection Jupiter was not removed")
	}

	ok = city.RemoveConnection(West, "Jakku")
	if ok {
		t.Error("TestRemoveConnection failed. The connection Jupiter shouldn't have been removed")
	}

	if len(city.directions) != 0 {
		t.Error("TestRemoveConnection failed. Some connections still persist")
	}
}

func TestRuinSelf(t *testing.T) {
	city := NewCity("Naboo")
	city.SetDirection(North, NewCity("Tatooine"))
	city.SetDirection(South, NewCity("Jupiter"))
	city.SetDirection(West, NewCity("Jakku"))
	city.SetDirection(East, NewCity("SanJuperdino"))

	city.firstAlien = NewAlien(0, city)
	city.secondAlien = NewAlien(0, city)

	err := city.RuinSelf()
	if err != nil {
		t.Errorf("TetRuinSelf failed. Error returned %s", err.Error())
	}

	if city.firstAlien != nil || city.secondAlien != nil {
		t.Error("TetRuinSelf failed. On of the aliens is still in the city")
	}

	if len(city.directions) != 0 {
		t.Error("TetRuinSelf failed. Some connections still persist")
	}
}
