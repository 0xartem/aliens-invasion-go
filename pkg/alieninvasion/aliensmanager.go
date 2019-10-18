package alieninvasion

// AliensManager interface provides a way to manage aliens from outside the packages.
type AliensManager interface {
	AddNewAlien(id int) (bool, error)
	MoveAll() error
	AreAllDestroyed() bool
}

func NewAliensManager(cp CityProvider, ob OutputObserver) AliensManager {
	return &aliensManager{cp, ob, make(map[int]*Alien)}
}

// aliensManager struct contains data for aliens management.
// cityProvider - interface for getting a random city/destroying a city.
// outputObserver - interface must be implemented by the package user to provide information to output.
// aliensMap - container for aliens storage.
// aliensManager struct is encapsulated.
type aliensManager struct {
	cityProvider   CityProvider
	outputObserver OutputObserver
	aliensMap      map[int]*Alien
}

// AddNewAlien spawns a new alien and puts it on the map.
// If 2 aliens meet during the initialization process - they destroy the city immediately.
// If the CityProvider can't generate a city (map is empty) - 'created' returns false (not error).
func (a *aliensManager) AddNewAlien(id int) (created bool, err error) {

	cityPosition := a.cityProvider.GetRandomCity()
	if cityPosition == nil {
		return false, nil
	}

	alien := NewAlien(id, cityPosition)
	a.aliensMap[id] = alien

	err = cityPosition.EnterCity(alien, a)
	if err != nil {
		return false, err
	}

	return true, nil
}

// MoveAll processes all the aliens and moves those who are alive and not trapped.
func (a *aliensManager) MoveAll() error {
	for id, alien := range a.aliensMap {

		if !alien.alive {
			delete(a.aliensMap, id)
			continue
		}

		if !alien.IsTrapped() {
			_, err := alien.Move(a)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// OnAliensFight - implements interface method to reac on the event when 2 aliens meet.
func (a *aliensManager) OnAliensFight(fstAl *Alien, secAl *Alien, city *City) {

	fstAl.Kill()
	secAl.Kill()
	a.cityProvider.DestroyCity(city.name)

	a.outputObserver.OnAliensDestroyedCity(fstAl, secAl, city)
}

func (a *aliensManager) AreAllDestroyed() bool {
	return len(a.aliensMap) == 0
}
