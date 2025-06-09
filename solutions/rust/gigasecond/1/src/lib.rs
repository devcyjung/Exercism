use time::{PrimitiveDateTime, Duration};

pub fn after(start: PrimitiveDateTime) -> PrimitiveDateTime {
    start.checked_add(Duration::seconds(1_000_000_000)).unwrap()
}
