package alieninvasion

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

// MapManager contains methods and data to manage the cities map.
// cittiesMap - container for the cities storage.
type MapManager struct {
	citiesMap map[string]*City
}

func NewMapManager() *MapManager {
	return &MapManager{make(map[string]*City)}
}

// BuildCity parses a raw city data and builds a new city with its connections.
func (m *MapManager) BuildCity(rawCity string) *City {

	if len(rawCity) == 0 {
		return nil
	}

	rawCity = strings.Trim(rawCity, " ")
	cityParams := strings.FieldsFunc(rawCity, func(r rune) bool {
		return r == ' ' || r == '='
	})

	// Parse the city name
	newCity, ok := m.citiesMap[cityParams[0]]
	if !ok {
		newCity = NewCity(cityParams[0])
		m.citiesMap[cityParams[0]] = newCity
	}

	// Process pairs (direction -> city) from the cityParams slice.
	for i, j := 1, 2; j < len(cityParams); i, j = i+2, j+2 {
		connectedCityName := cityParams[j]
		connectedCity, ok := m.citiesMap[connectedCityName]

		if !ok {
			connectedCity = NewCity(connectedCityName)
			m.citiesMap[connectedCityName] = connectedCity
		}

		direction := directionFromString(strings.ToLower(cityParams[i]))
		newCity.SetDirection(direction, connectedCity)

		/// Fix links in case they are not in the input file
		if !connectedCity.DirectionExists(direction.Opposite()) {
			connectedCity.SetDirection(direction.Opposite(), newCity)
		}
	}

	return newCity
}

func (m *MapManager) IsMapEmpty() bool {
	return len(m.citiesMap) == 0
}

// GetRandomCity returns a random city on the map.
// If the map is empty - return nil.
// The method should be used during the aliens initialization stage.
func (m *MapManager) GetRandomCity() *City {

	mapLen := len(m.citiesMap)
	if mapLen == 0 {
		return nil
	}

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	rndPos := rnd.Intn(mapLen)

	//TODO: optimize
	counter := 0
	for _, city := range m.citiesMap {
		if counter == rndPos {
			return city
		}
		counter++
	}

	return nil
}

// DestroyCity ruins the city by its name and deletes it from the storage.
func (m *MapManager) DestroyCity(name string) error {
	city, ok := m.citiesMap[name]
	if !ok {
		return errors.New("The city " + name + " is not on the map")
	}

	err := city.RuinSelf()
	if err != nil {
		return err
	}

	delete(m.citiesMap, name)
	return nil
}

// OutputAll - outputs all existing cities on the map.
func (m *MapManager) OutputAll(observer OutputObserver) {
	for _, city := range m.citiesMap {
		observer.OnCityOutput(city)
	}
}
