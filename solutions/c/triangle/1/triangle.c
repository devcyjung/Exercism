#include "triangle.h"
#include <math.h>

static bool fequal(double a, double b) {
    return fabs(a-b) < 0.0000000001;
} 

static double max_side(const triangle_t t) {
    double max = t.a;
    if (t.b > max) {
        max = t.b;
    }
    if (t.c > max) {
        max = t.c;
    }
    return max;
}

static double min_side(const triangle_t t) {
    double min = t.a;
    if (t.b < min) {
        min = t.b;
    }
    if (t.c < min) {
        min = t.c;
    }
    return min;
}

static bool is_triangle(const triangle_t t) {
    return min_side(t) > 0 && 2 * max_side(t) < t.a + t.b + t.c;
}

bool is_equilateral(const triangle_t t) {
    return is_triangle(t) && fequal(min_side(t), max_side(t));
}

bool is_isosceles(const triangle_t t) {
    double max = max_side(t);
    double min = min_side(t);
    double mid = t.a + t.b + t.c - min - max;
    
    return is_triangle(t) && (fequal(mid, max) || fequal(mid, min));
}

bool is_scalene(const triangle_t t) {
    return is_triangle(t) && !fequal(t.a, t.b) && !fequal(t.b, t.c) && !fequal(t.c, t.a);
}