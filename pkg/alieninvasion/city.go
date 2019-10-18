package alieninvasion

import (
    "errors"
    "fmt"
    "math/rand"
    "strings"
    "time"
)

// City - data structure that contains information about each city on the map
// name - text name of the city
// directions - connections/roads available in the city that lead to other cities
// firstAlien, secondAlien - placeholders for aliens to arrive to the city
type City struct {
    name                    string
    directions              map[Direction]*City
    firstAlien, secondAlien *Alien
}

func NewCity(name string) *City {
    return &City{name: name, directions: make(map[Direction]*City)}
}

func (city *City) Name() string {
    return city.name
}

func (city *City) IsIsolated() bool {
    return len(city.directions) == 0
}

func (city *City) DirectionExists(d Direction) bool {
    _, ok := city.directions[d]
    return ok
}

// GetRandomDirection returns a random connected city
// If the city is isolated - return nil
func (city *City) GetRandomDirection() *City {

    mapLen := len(city.directions)
    if mapLen == 0 {
        return nil
    }

    src := rand.NewSource(time.Now().UnixNano())
    rnd := rand.New(src)
    rndDir := rnd.Intn(mapLen)

    return city.directions[Direction(rndDir)]
}

// SetDirections sets/updates a connected city
func (city *City) SetDirection(d Direction, connectedCity *City) {
    city.directions[d] = connectedCity
}

// EnterCity lets an alien into the city
// If 2 aliens arrived to the city - call the implementation of OnAliensFight method
// If both aliens are in the city and it have not been destroyed yet - return an 'error', it should never happen
func (city *City) EnterCity(alien *Alien, observer FightObserver) error {
    if city.secondAlien != nil {
        return errors.New("There are 2 aliens in the city already. The city must have been destroyed.")
    }

    if city.firstAlien == nil {
        city.firstAlien = alien
    } else {
        city.secondAlien = alien
        observer.OnAliensFight(city.firstAlien, city.secondAlien, city)
    }

    return nil
}

// LeaveCity lets the first alien out of the city
// It must always be the first, otherwise the city must have been destroyed
// If the second alien is in the city - return error, it should never happen
func (city *City) LeaveCity(alien *Alien) error {
    if city.secondAlien != nil {
        return errors.New("There are 2 aliens in the city already. The city must have been destroyed.")
    } else if city.firstAlien != alien {
        return errors.New("The alien is not in city so he cannot leave it.")
    }

    city.firstAlien = nil
    return nil
}

// RemoveConnection removes a connection from an opposite city
// It means the opposite city is being destroyed
func (city *City) RemoveConnection(d Direction, cityName string) bool {
    val, ok := city.directions[d]
    if ok && val.name == cityName {
        delete(city.directions, d)
        return true
    }
    return false
}

// RuinSelf destroys the city, all its connections and aliens
func (city *City) RuinSelf() error {
    if city.firstAlien == nil || city.secondAlien == nil {
        return errors.New("2 aliens must be in the city to ruin it")
    }

    for d, connectedCity := range city.directions {
        connectedCity.RemoveConnection(d.Opposite(), city.name)
        delete(city.directions, d)
    }
    city.firstAlien = nil
    city.secondAlien = nil
    return nil
}

func (city *City) String() string {
    var b strings.Builder
    b.WriteString(city.name)
    b.WriteByte(' ')

    for d, connectedCity := range city.directions {
        b.WriteString(d.String())
        b.WriteByte('=')
        b.WriteString(connectedCity.name)
        b.WriteByte(' ')
    }
    return b.String()
}

func (city *City) StringWithAlien() string {
    var b strings.Builder
    b.WriteString(city.String())

    if city.firstAlien != nil {
        if city.firstAlien.IsTrapped() {
            b.WriteString(fmt.Sprintf("(Alien %d is trapped)", city.firstAlien.id))
        } else {
            b.WriteString(fmt.Sprintf("(Alien %d is wandering here)", city.firstAlien.id))
        }
    }

    return b.String()
}
