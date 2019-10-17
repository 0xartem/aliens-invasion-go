package alieninvasion

import "errors"

// Alien struct contains information about each alien created
// id - unique identifier
// position - current city position on the map
// alive - dead or alive
type Alien struct {
	id       int
	position *City
	alive    bool
}

func NewAlien(id int, position *City) *Alien {
	return &Alien{id, position, true}
}

func (alien *Alien) Id() int {
	return alien.id
}

func (alien *Alien) IsTrapped() bool {
	return alien.position.IsIsolated()
}

func (alien *Alien) Kill() {
	alien.alive = false
}

// Move alien to the next city
// if 'moved' is true - the alien was moved, otherwise - it's trapped (not error)
// 'error' denotes the error occured during the moving process
func (alien *Alien) Move(observer FightObserver) (moved bool, err error) {
	moved = false
	err = nil

	if alien.position == nil {
		err = errors.New("The alien position is empty")
		return
	}

	city := alien.position.GetRandomDirection()
	if city == nil {
		return
	}

	err = alien.position.LeaveCity(alien)
	if err != nil {
		return
	}

	alien.position = city
	err = alien.position.EnterCity(alien, observer)
	if err != nil {
		return
	}

	moved = true
	return
}
