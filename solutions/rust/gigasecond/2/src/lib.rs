use time::{PrimitiveDateTime as DateTime, Duration};

pub fn after(start: DateTime) -> DateTime {
    start.saturating_add(Duration::seconds_f64(1e9))
}
