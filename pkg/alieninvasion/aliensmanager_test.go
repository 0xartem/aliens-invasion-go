package alieninvasion

import "testing"

type FakeCityProvider struct {
	city *City
}

func (f *FakeCityProvider) GetRandomCity() *City {
	return f.city
}

func (f *FakeCityProvider) DestroyCity(name string) error {
	return nil
}

type FakeOutputObserver struct{}

func (f *FakeOutputObserver) OnCityOutput(city *City)                                      {}
func (f *FakeOutputObserver) OnAliensDestroyedCity(fstAl *Alien, secAl *Alien, city *City) {}

func TestAddNewAlien(t *testing.T) {
	fakeCityProvider := FakeCityProvider{NewCity("Jakku")}
	aliensManager := NewAliensManager(&fakeCityProvider, &FakeOutputObserver{})

	added, err := aliensManager.AddNewAlien(0)

	if !added {
		t.Error("TestAddNewAlien failed. Alien 0 wasn't added")
	}
	if err != nil {
		t.Errorf("TestAddNewAlien failed. An error occured: %s", err.Error())
	}

	if fakeCityProvider.city.firstAlien.Id() != 0 {
		t.Error("TestAddNewAlien failed. Alien didn't enter the city")
	}
}
