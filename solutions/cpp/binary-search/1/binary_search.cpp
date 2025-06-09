#include "binary_search.h"

#include <stdexcept>

namespace binary_search {

    size_t find(const std::vector<int>& data, const int& value) {
        if (data.size() == 0) {
            throw std::domain_error("empty vector");
        }
        size_t slice_begin {0U};
        size_t slice_end {data.size()};
        size_t slice_mid {};
        while (slice_begin < slice_end) {
            slice_mid = (slice_begin + slice_end - 1) / 2;
            if (value < data[slice_mid]) {
                slice_end = slice_mid;
                continue;
            }
            if (data[slice_mid] < value) {
                slice_begin = slice_mid + 1;
                continue;
            }
            if (data[slice_mid] == value) {
                return slice_mid;
            }
        }
        if (data[slice_begin] == value) {
            return slice_begin;
        }
        throw std::domain_error("value not found");
    }    
    
}  // namespace binary_search
