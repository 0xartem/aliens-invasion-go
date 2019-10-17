package alieninvasion

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

func (d Direction) String() string {
	names := [...]string{"north", "south", "west", "east"}

	if d < North || d > East {
		return "Unknown"
	}

	return names[d]
}

func (d Direction) Opposite() Direction {
	if d%2 == 0 {
		return d + 1
	} else {
		return d - 1
	}
}

var nameToDirection = map[string]Direction{"north": North, "south": South, "west": West, "east": East}

func directionFromString(rawDirection string) Direction {
	return nameToDirection[rawDirection]
}
