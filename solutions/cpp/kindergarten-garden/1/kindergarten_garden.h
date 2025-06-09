#pragma once

#include <array>
#include <map>
#include <stdexcept>
#include <string>

namespace kindergarten_garden {

    enum class Plants {
        clover, grass, violets, radishes,
    };

    constexpr Plants decode(char s) {
        switch (s) {
            case 'C':
                return Plants::clover;
            case 'G':
                return Plants::grass;
            case 'V':
                return Plants::violets;
            case 'R':
                return Plants::radishes;
        }
        throw new std::runtime_error("plant not found");
    }

    const std::map<std::string, unsigned int> students = {
        {"Alice", 0}, {"Bob", 1}, {"Charlie", 2}, {"David", 3},
        {"Eve", 4}, {"Fred", 5}, {"Ginny", 6}, {"Harriet", 7},
        {"Ileana", 8}, {"Joseph", 9}, {"Kincaid", 10}, {"Larry", 11},
    };

    const std::array<Plants, 4> plants(const std::string&, const std::string&);

}  // namespace kindergarten_garden
