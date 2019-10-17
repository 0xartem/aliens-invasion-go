package alieninvasion

type FightObserver interface {
	OnAliensFight(fstAl *Alien, secAl *Alien, city *City)
}
