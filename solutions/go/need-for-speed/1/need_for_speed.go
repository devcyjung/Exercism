package speed

// TODO: define the 'Car' type struct
type Car struct {
    battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car {speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
}

// TODO: define the 'Track' type struct
type Track struct {
    distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track {distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
    if car.battery < car.batteryDrain {
        return car
    }
	car.battery -= car.batteryDrain
    car.distance += car.speed
    return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	return track.distance <= car.speed*car.battery/car.batteryDrain
}
