#[derive(PartialEq, Eq, Debug)]
pub enum Direction {
    North,
    East,
    South,
    West,
}

const ALL_DIR: [Direction; 4] = [Direction::North, Direction::East, Direction::South, Direction::West];

pub struct Robot {
    x: i32,
    y: i32,
    dirIdx: usize,
}

impl Robot {
    pub fn new(x: i32, y: i32, d: Direction) -> Self {
        Self {
            x, y, dirIdx: ALL_DIR.iter().position(|e| *e == d).expect("Direction not found")
        }
    }

    #[must_use]
    pub fn turn_right(mut self) -> Self {
        self.dirIdx += 1;
        self.dirIdx &= 3;
        self
    }

    #[must_use]
    pub fn turn_left(mut self) -> Self {
        self.dirIdx += 3;
        self.dirIdx &= 3;
        self
    }

    #[must_use]
    pub fn advance(mut self) -> Self {
        let delta = if (self.dirIdx >> 1) == 0 { 1 } else { -1 };
        if (self.dirIdx & 1) == 0 {
            self.y += delta;
        } else {
            self.x += delta;
        }
        self
    }

    #[must_use]
    pub fn instructions(mut self, instructions: &str) -> Self {
        instructions.chars().fold(self, |mut robot, ch| {
            match ch {
                'A' => robot.advance(),
                'L' => robot.turn_left(),
                'R' => robot.turn_right(),
                _ => unreachable!("Invalid command"),
            }
        })
    }

    pub fn position(&self) -> (i32, i32) {
        (self.x, self.y)
    }

    pub fn direction(&self) -> &Direction {
        &ALL_DIR[self.dirIdx]
    }
}
