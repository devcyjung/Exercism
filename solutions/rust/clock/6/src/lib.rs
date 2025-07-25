use std::fmt;

#[derive(Debug, PartialEq, Eq, Clone)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let total_minutes = hours * 60 + minutes;
        let (hours, minutes) = (
            total_minutes
                .div_euclid(60)
                .rem_euclid(24),
            total_minutes.rem_euclid(60),
        );
        Self{ hours, minutes }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Self::new(self.hours, self.minutes + minutes)
    }

    pub fn checked_new(hours: i32, minutes: i32) -> Option<Self> {
        let total_minutes = hours * 60 + minutes;
        let (hours, minutes) = (
            total_minutes
                .checked_div_euclid(60)
                ?.checked_rem_euclid(24)?,
            total_minutes.checked_rem_euclid(60)?,
        );
        Some(Self{ hours, minutes })
    }

    pub fn checked_add_minutes(&self, minutes: i32) -> Option<Self> {
        Self::checked_new(self.hours, self.minutes + minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> Result<(), fmt::Error> {
        write!(f, "{:02}:{:02}", self.hours, self.minutes)
    }
}
