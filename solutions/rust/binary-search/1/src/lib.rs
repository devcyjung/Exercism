use std::cmp::Ordering;

pub fn find(array: &[i32], key: i32) -> Option<usize> {
    let mut begin = 0;
    let mut end = array.len();
    let mut mid: usize;
    while begin < end {
        mid = (begin + end - 1) / 2;
        match array[mid].cmp(&key) {
            Ordering::Equal => {
                return Some(mid);
            },
            Ordering::Less => {
                begin = mid + 1;
            },
            Ordering::Greater => {
                end = mid;
            }
        }
    }
    None
}
