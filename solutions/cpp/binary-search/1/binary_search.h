#if !defined(BINARY_SEARCH_H)
#define BINARY_SEARCH_H

#include <cstddef>
#include <vector>

namespace binary_search {

    [[nodiscard]] size_t find(const std::vector<int>&, const int&);    

}  // namespace binary_search

#endif  // BINARY_SEARCH_H