package speed

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	car := Car{battery: 100, distance: 0, speed: speed, batteryDrain: batteryDrain}
	return car
}

type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	track := Track{distance: distance}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	carUpdated := Car{}
	carUpdated.batteryDrain = car.batteryDrain
	carUpdated.speed = car.speed

	if car.battery > car.batteryDrain {
		carUpdated.distance = car.distance + car.speed
		carUpdated.battery = car.battery - car.batteryDrain
	} else {
		carUpdated = car
	}

	return carUpdated

}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	if car.batteryDrain > car.battery {
		return false
	}

	carMaximumDistance := ((car.battery / car.batteryDrain) * car.speed) + car.distance
	if carMaximumDistance >= track.distance {
		return true
	}
	return false
}
