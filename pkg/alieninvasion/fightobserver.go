package alieninvasion

// FightObserver - interface that process when the aliens meet, fight and destroy the city
type FightObserver interface {
	OnAliensFight(fstAl *Alien, secAl *Alien, city *City)
}
