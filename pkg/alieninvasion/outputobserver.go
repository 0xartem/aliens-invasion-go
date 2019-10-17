package alieninvasion

type OutputObserver interface {
	OnCityOutput(city *City)
	OnAliensDestroyedCity(fstAl *Alien, secAl *Alien, city *City)
}
