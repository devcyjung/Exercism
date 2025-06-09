#include "bob.h"

#include <algorithm>
#include <locale>

namespace bob {
    std::string hey(std::string input) {
        const auto is_question{
            input[input.find_last_not_of("\n\r\t ")] == '?'
        };
        const auto is_yelling{
            std::all_of(
                input.begin(),
                input.end(),
                [](char c){ return std::isupper(c) || !std::isalpha(c); }
            ) && std::any_of(
                input.begin(),
                input.end(),
                [](char c){ return std::isupper(c); }
            )
        };
        const auto is_silent{
            std::all_of(
                input.begin(),
                input.end(),
                [](char c){ return std::isspace(c); }
            )
        };
        if (!is_yelling && is_question) {
            return "Sure.";
        }
        if (is_yelling && !is_question) {
            return "Whoa, chill out!";
        }
        if (is_yelling && is_question) {
            return "Calm down, I know what I'm doing!";
        }
        if (is_silent) {
            return "Fine. Be that way!";
        }
        return "Whatever.";
    }
}
