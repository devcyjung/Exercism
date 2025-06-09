pub fn is_armstrong_number(num: u32) -> bool {
    let power = num_digits(num);
    let mut rem = num;
    let mut sum = 0;
    while rem != 0 {
        sum += pow(rem % 10, power);
        rem /= 10;
    }
    sum == num
}

fn num_digits(num: u32) -> u32 {
    let mut rem = num;
    let mut ret = 0;
    while rem != 0 {
        rem /= 10;
        ret += 1;
    }
    ret
}

fn pow(num: u32, power: u32) -> u32 {
    let mut ret = 1;
    for i in 0..power {
        ret *= num;
    }
    ret
}