package alieninvasion

import "testing"

func TestOppositeDirectionNorth(t *testing.T) {
	d := North
	if d.Opposite() != South {
		t.Errorf("TestOppositeDirectionNorth failed. The opposite must be South but it's %d instead", d.Opposite())
	}
}

func TestOppositeDirectionSouth(t *testing.T) {
	d := South
	if d.Opposite() != North {
		t.Errorf("TestOppositeDirectionNorth failed. The opposite must be North but it's %d instead", d.Opposite())
	}
}

func TestOppositeDirectionWest(t *testing.T) {
	d := West
	if d.Opposite() != East {
		t.Errorf("TestOppositeDirectionNorth failed. The opposite must be East but it's %d instead", d.Opposite())
	}
}

func TestOppositeDirectionEast(t *testing.T) {
	d := East
	if d.Opposite() != West {
		t.Errorf("TestOppositeDirectionNorth failed. The opposite must be West but it's %d instead", d.Opposite())
	}
}
