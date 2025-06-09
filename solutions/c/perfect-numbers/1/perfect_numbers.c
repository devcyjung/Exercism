#include "perfect_numbers.h"
#include <stdio.h>
static long long simple_aliquot_sum(const long long number) {
    long long quo = 0;
    long long divisor = 1;
    long long sum = 0;
    while ((quo = number / divisor) >= divisor) {
        if (quo * divisor == number) {
            sum += divisor;
            if (quo != divisor) {
                sum += quo;
            }
        }
        ++divisor;
    }
    return sum - number;
}

kind classify_number(const long long number) {
    if (number <= 0) {
        return ERROR;
    }
    long long sum = simple_aliquot_sum(number);
    printf("%llu %llu", number, sum);
    if (number == sum) {
        return PERFECT_NUMBER;
    }
    if (number < sum) {
        return ABUNDANT_NUMBER;
    }
    return DEFICIENT_NUMBER;
}