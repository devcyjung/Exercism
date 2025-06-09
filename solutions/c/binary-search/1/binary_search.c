#include "binary_search.h"

const int *binary_search(int value, const int *arr, size_t length) {
    size_t begin = 0;
    size_t end = length;
    size_t mid;
    int cur;
    while (begin < end) {
        mid = (begin + end - 1) / 2;
        cur = arr[mid];
        if (cur == value) {
            return arr + mid;
        }
        if (cur < value) {
            begin = mid + 1;
        } else {
            end = mid;
        }
    }
    return NULL;
}