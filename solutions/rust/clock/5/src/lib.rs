use std::fmt;

#[derive(Debug, PartialEq, Eq, Clone)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

const DAY_IN_MINUTES: i32 = 24 * 60;

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let minutes_since_midnight = (hours * 60 + minutes).rem_euclid(DAY_IN_MINUTES);
        let (hours, minutes) = (minutes_since_midnight / 60, minutes_since_midnight % 60);
        Self {
            hours,
            minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Self::new(self.hours, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> Result<(), fmt::Error> {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}
