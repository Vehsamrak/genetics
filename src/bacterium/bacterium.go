package bacterium

import "github.com/vehsamrak/genetics/src/applicationError"

const defaultLifePoints = 10

type bacterium struct {
	gameField  gameField
	lifePoints int
	x          int
	y          int
}

// New bacterium constructor
func New() bacterium {
	return bacterium{lifePoints: defaultLifePoints}
}

func (bacterium *bacterium) GetLifePoints() int {
	return bacterium.lifePoints
}

func (bacterium *bacterium) AddLifePoints(points int) {
	if points <= 0 {
		return
	}

	bacterium.lifePoints += points
}

func (bacterium *bacterium) DeductLifePoints(points int) {
	if points <= 0 {
		return
	}

	if bacterium.lifePoints-points < 0 {
		bacterium.lifePoints = 0
	} else {
		bacterium.lifePoints -= points
	}
}

func (bacterium *bacterium) IsAlive() bool {
	return bacterium.lifePoints > 0
}

func (bacterium *bacterium) IsDead() bool {
	return bacterium.lifePoints <= 0
}

func (bacterium *bacterium) Move(direction Direction) (error error) {
	if bacterium.IsDead() {
		error = new(applicationError.CanNotMove)
	}

	destinationX := bacterium.x
	destinationY := bacterium.y

	switch direction {
	case directionNorth:
		destinationY++
	case directionEast:
		destinationX++
	case directionSouth:
		destinationY--
	case directionWest:
		destinationX--
	}

	for _, fieldBacterium := range bacterium.gameField.allBacterias() {
		if destinationX == fieldBacterium.X() && destinationY == fieldBacterium.Y() {
			return new(applicationError.CanNotMove)
		}
	}

	bacterium.x = destinationX
	bacterium.y = destinationY
	bacterium.lifePoints--

	return error
}

func (bacterium *bacterium) X() int {
	return bacterium.x
}

func (bacterium *bacterium) Y() int {
	return bacterium.y
}
