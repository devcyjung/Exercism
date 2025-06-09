use lazy_static::lazy_static;
use std::collections::BTreeMap;

lazy_static! {
    static ref SOUND_MAP: BTreeMap<u32, String> = BTreeMap::from([
        (3, "Pling".into()), (5, "Plang".into()), (7, "Plong".into()),
    ]);
}

pub fn raindrops(n: u32) -> String {
    let sound: String = SOUND_MAP.iter().filter(|&(divisor, _)| n % divisor == 0)
        .map(|(_, sound)| sound.to_owned()).collect();
    if sound.is_empty() {
        return n.to_string();
    }
    sound
}