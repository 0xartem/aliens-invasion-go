package alieninvasion

// OutputObserver - is an interface to implement by a package user.
// Its responsibility to define display behavior.
type OutputObserver interface {
	OnCityOutput(city *City)
	OnAliensDestroyedCity(fstAl *Alien, secAl *Alien, city *City)
}
