#include "difference_of_squares.h"

unsigned int sum_of_squares(unsigned int number) {
    unsigned int sq = 0;
    for (unsigned int i = 1; i <= number; ++i) {
        sq += i * i;
    }
    return sq;
}

unsigned int square_of_sum(unsigned int number) {
    unsigned int s = 0;
    for (unsigned int i = 1; i <= number; ++i) {
        s += i;
    }
    return s * s;
}

unsigned int difference_of_squares(unsigned int number) {
    unsigned int ss1 = sum_of_squares(number);
    unsigned int ss2 = square_of_sum(number);
    if (ss1 > ss2) {
        return ss1 - ss2;
    }
    return ss2 - ss1;
}