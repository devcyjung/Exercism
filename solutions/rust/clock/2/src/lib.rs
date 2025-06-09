#[derive(Debug, PartialEq, Eq, Clone)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        let unfolded = (hours * 60 + minutes) % (24 * 60);
        let neg_handled = if unfolded < 0 { 24 * 60 + unfolded } else { unfolded };
        let (hours, minutes) = (neg_handled / 60, neg_handled % 60);
        Self {
            hours,
            minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let unfolded = (self.hours * 60 + self.minutes + minutes) % (24 * 60);
        let neg_handled = if unfolded < 0 { 24 * 60 + unfolded } else { unfolded };
        let (hours, minutes) = (neg_handled / 60, neg_handled % 60);
        Self {
            hours,
            minutes,
        }
    }

    pub fn to_string(&self) -> String {
        format!("{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}
