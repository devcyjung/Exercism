use std::cmp::Ordering;

#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist(first_list: &[i32], second_list: &[i32]) -> Comparison {
    let len_1 = first_list.len();
    let len_2 = second_list.len();
    match len_1.cmp(&len_2) {
        Ordering::Equal => {
            let first_in_second = match_against(first_list, second_list);
            let second_in_first = match_against(second_list, first_list);
            if first_in_second && second_in_first {
                return Comparison::Equal;
            } else if first_in_second {
                return Comparison::Sublist;
            } else if second_in_first {
                return Comparison::Superlist;
            }
            Comparison::Unequal
        },
        Ordering::Less => {
            if match_against(first_list, second_list) {
                return Comparison::Sublist;
            }
            Comparison::Unequal
        },
        Ordering::Greater => {
            if match_against(second_list, first_list) {
                return Comparison::Superlist;
            }
            Comparison::Unequal
        }
    }
}

fn match_against(pattern: &[i32], against: &[i32]) -> bool {
    if pattern.is_empty() {
        return true;
    }
    if against.is_empty() {
        return false;
    }
    let p_len = pattern.len();
    let a_len = against.len();
    let p = process_pattern(pattern);
    let mut match_len = 0;
    let mut i = 0;
    while i < a_len {
        if against[i] == pattern[match_len] {
            match_len += 1;
            i += 1;
            if match_len == p_len {
                return true;
            }
        } else if match_len == 0 {
            i += 1;
        } else {
            match_len = p[match_len - 1];
        }
    }
    false
}

/// KMP algorithm
fn process_pattern(pattern: &[i32]) -> Vec<usize> {
    let p_len = pattern.len();
    let mut p = vec![0; p_len];
    p[0] = 0;
    let mut match_len = 0;
    let mut i = 1;
    while i < p_len {
        if pattern[i] == pattern[match_len] {
            match_len += 1;
            p[i] = match_len;
            i += 1;
        } else if match_len == 0 {
            p[i] = 0;
            i += 1;
        } else {
            match_len = p[match_len - 1];
        }
    }
    p
}