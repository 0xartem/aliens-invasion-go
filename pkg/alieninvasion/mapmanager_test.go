package alieninvasion

import "testing"

func TestBuildCityAllDirections(t *testing.T) {
	mapManager := NewMapManager()
	city := mapManager.BuildCity("Jupiter north=Tatooine south=Earth west=Mars east=SanJuperdino")

	if city.name != "Jupiter" {
		t.Errorf("TestBuildCity failed. City's name must be Jupiter, it's %s instead", city.name)
	}

	dirCity, ok := city.directions[North]
	if !ok {
		t.Error("TestBuildCity failed. City's North direction doesn't exist")
	}
	if dirCity.name != "Tatooine" {
		t.Errorf("TestBuildCity failed. City's North direction must be Tatooine, it's %s instead", dirCity.name)
	}

	dirCity, ok = city.directions[South]
	if !ok {
		t.Error("TestBuildCity failed. City's South direction doesn't exist")
	}
	if dirCity.name != "Earth" {
		t.Errorf("TestBuildCity failed. City's South direction must be Earth, it's %s instead", dirCity.name)
	}

	dirCity, ok = city.directions[West]
	if !ok {
		t.Error("TestBuildCity failed. City's West direction doesn't exist")
	}
	if dirCity.name != "Mars" {
		t.Errorf("TestBuildCity failed. City's West direction must be Mars, it's %s instead", dirCity.name)
	}

	dirCity, ok = city.directions[East]
	if !ok {
		t.Error("TestBuildCity failed. City's East direction doesn't exist")
	}
	if dirCity.name != "SanJuperdino" {
		t.Errorf("TestBuildCity failed. City's East direction must be SanJuperdino, it's %s instead", dirCity.name)
	}
}

func TestBuildCityTwoDirections(t *testing.T) {
	mapManager := NewMapManager()
	city := mapManager.BuildCity("Tatooine south=Jupiter east=Naboo")

	if city.name != "Tatooine" {
		t.Errorf("TestBuildCity failed. City's name must be Tatooine, it's %s instead", city.name)
	}

	dirCity, ok := city.directions[South]
	if !ok {
		t.Error("TestBuildCity failed. City's South direction doesn't exist")
	}
	if dirCity.name != "Jupiter" {
		t.Errorf("TestBuildCity failed. City's South direction must be Jupiter, it's %s instead", dirCity.name)
	}

	dirCity, ok = city.directions[East]
	if !ok {
		t.Error("TestBuildCity failed. City's East direction doesn't exist")
	}
	if dirCity.name != "Naboo" {
		t.Errorf("TestBuildCity failed. City's East direction must be Naboo, it's %s instead", dirCity.name)
	}
}

func TestBuildCityEmpty(t *testing.T) {
	mapManager := NewMapManager()
	city := mapManager.BuildCity("")
	if city != nil {
		t.Errorf("TestBuildCity failed. City must be nil")
	}
}
