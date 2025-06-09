package robot

const (
    N Dir = iota
    E
    S
    W
)

var directions = [][]int {
    {0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

func Right() {
    Step1Robot.Dir += 1
    Step1Robot.Dir %= 4
}

func Left() {
    Step1Robot.Dir += 3
    Step1Robot.Dir %= 4
}

func Advance() {
	dir := directions[Step1Robot.Dir]
    Step1Robot.X += dir[0]
    Step1Robot.Y += dir[1]
}

func (d Dir) String() string {
	switch d {
    case N:
        return "North"
    case E:
        return "East"
    case S:
        return "South"
    case W:
        return "West"
    }
    return "Invalid"
}

type Action struct {
    dirDelta, posDelta int
}

func StartRobot(command <-chan Command, action chan<- Action) {
	for cmd := range command {
        switch cmd {
        case 'A':
            action <- Action{0, 1}
        case 'L':
            action <- Action{-1, 0}
        case 'R':
            action <- Action{1, 0}
        }
    }
    close(action)
}

func Room(extent Rect, robot Step2Robot, action <-chan Action, report chan<- Step2Robot) {
    for act := range action {
        robot.Dir += Dir(4 + act.dirDelta)
        robot.Dir %= 4
        dir := directions[robot.Dir]
        robot.Easting += RU(dir[0] * act.posDelta)
        robot.Northing += RU(dir[1] * act.posDelta)
        if extent.Min.Easting <= robot.Easting &&
        	robot.Easting <= extent.Max.Easting &&
        	extent.Min.Northing <= robot.Northing &&
        	robot.Northing <= extent.Max.Northing {
            continue
        }
        robot.Easting -= RU(dir[0])
        robot.Northing -= RU(dir[1])
    }
    report <- robot
    close(report)
}

type Action3 struct {
    action	Action
    name	string
    done	bool
}

const (
    logNoName = "No name"
    logDuplicateName = "Duplicate names"
    logSamePlace = "Same place"
    logOutOfRoom = "Out of room"
    logUnknownCommand = "Unknown command"
    logWallCollision = "Wall collision"
    logRobotCollision = "Robot collision"
)

func StartRobot3(name, script string, action chan<- Action3, log chan<- string) {
    defer func() {
        action <- Action3{Action{0, 0}, name, true}
    }()
	if name == "" {
        log <- logNoName
    }
    for _, cmd := range script {
        switch cmd {
        case 'A':
            action <- Action3{Action{0, 1}, name, false}
        case 'L':
            action <- Action3{Action{-1, 0}, name, false}
        case 'R':
            action <- Action3{Action{1, 0}, name, false}
        default:
            log <- logUnknownCommand
            return
        }
    }
}

func Room3(extent Rect, robots []Step3Robot, action chan Action3, rep chan<- []Step3Robot, log chan<- string) {
    remainingRobots := len(robots) 
    defer func() {
        for remainingRobots > 0 {
            act := <- action
            if act.done {
                remainingRobots--
            }
        }
        close(action)
        rep <- robots
    	close(rep)
    }()
    inBounds := func(pos Pos) bool {
        return extent.Min.Easting <= pos.Easting && pos.Easting <= extent.Max.Easting &&
        	extent.Min.Northing <= pos.Northing && pos.Northing <= extent.Max.Northing
    }
    robotMap := make(map[string]int)
    locationSet := make(map[Pos]struct{})
    for i, robot := range robots {
        if _, ok := robotMap[robot.Name]; ok {
            log <- logDuplicateName
            return
        }
        if _, ok := locationSet[robot.Pos]; ok {
            log <- logSamePlace
            return
        }
        robotMap[robot.Name] = i
        locationSet[robot.Pos] = struct{}{}
        if !inBounds(robot.Pos) {
            log <- logOutOfRoom
            return
        }
    }
    for remainingRobots > 0 {
        act := <- action
        if act.done {
            remainingRobots--
            continue
        }
        robotIdx := robotMap[act.name]
        delete(locationSet, robots[robotIdx].Pos)
        robots[robotIdx].Dir += Dir(4 + act.action.dirDelta)
        robots[robotIdx].Dir %= 4
        dir := directions[robots[robotIdx].Dir]
        robots[robotIdx].Easting += RU(dir[0] * act.action.posDelta)
        robots[robotIdx].Northing += RU(dir[1] * act.action.posDelta)
        _, collision := locationSet[robots[robotIdx].Pos]
        if inBounds(robots[robotIdx].Pos) && !collision {
            locationSet[robots[robotIdx].Pos] = struct{}{}
            continue
        }
        if !inBounds(robots[robotIdx].Pos) {
            log <- logWallCollision
        }
        if collision {
            log <- logRobotCollision
        }
        robots[robotIdx].Easting -= RU(dir[0])
        robots[robotIdx].Northing -= RU(dir[1])
        locationSet[robots[robotIdx].Pos] = struct{}{}
    }
}