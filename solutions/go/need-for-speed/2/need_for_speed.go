package speed

type Car struct {
    battery, batteryDrain, speed, distance int
}

func NewCar(speed, batteryDrain int) Car {
	return Car {
        100,
        batteryDrain,
        speed,
        0,
    }
}

type Track struct {
    distance int
}

func NewTrack(distance int) Track {
	return Track {distance}
}

func Drive(car Car) Car {
    if car.battery < car.batteryDrain {
        return car
    }
	car.battery -= car.batteryDrain
    car.distance += car.speed
    return car
}

func CanFinish(car Car, track Track) bool {
	return track.distance * car.batteryDrain <= car.speed * car.battery
}
