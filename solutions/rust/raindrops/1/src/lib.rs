use lazy_static::lazy_static;
use std::collections::HashMap;

lazy_static! {
    static ref SOUND_MAP: HashMap<u32, String> = HashMap::from([
        (3, "Pling".into()), (5, "Plang".into()), (7, "Plong".into()),
    ]);
}

pub fn raindrops(n: u32) -> String {
    let mut sound = SOUND_MAP.iter().fold(String::new(), |mut acc, (divisor, sound)| {
        if n % divisor == 0 {
            acc.push_str(sound);
        }
        acc
    });
    if sound.is_empty() {
        sound.push_str(&n.to_string());
    }
    sound
}