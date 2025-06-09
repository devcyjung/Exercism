#include "matching_brackets.h"

#include <vector>

namespace matching_brackets {
    enum class brackets {
        square_open, square_close, curly_open, curly_close, round_open, round_close, none,
    };

    brackets get_type(char c) {
        switch (c) {
            case '[':
                return brackets::square_open;
            case ']':
                return brackets::square_close;
            case '{':
                return brackets::curly_open;
            case '}':
                return brackets::curly_close;
            case '(': 
                return brackets::round_open;
            case ')': 
                return brackets::round_close;
        }
        return brackets::none;
    }

    bool check(std::string input) {
        std::vector<brackets> stack(0);
        brackets cur {};

        for (auto it = input.begin(); it != input.end(); ++it) {
            cur = get_type(*it);
            switch (cur) {
                case brackets::none:
                    break;
                case brackets::curly_open:
                    stack.push_back(cur);
                    break;
                case brackets::square_open:
                    stack.push_back(cur);
                    break;
                case brackets::round_open:
                    stack.push_back(cur);
                    break;
                case brackets::curly_close:
                    if (!stack.empty() && stack.back() == brackets::curly_open) {
                        stack.pop_back();
                        break;
                    } else {
                        return false;
                    }
                case brackets::square_close:
                    if (!stack.empty() && stack.back() == brackets::square_open) {
                        stack.pop_back();
                        break;
                    } else {
                        return false;
                    }
                case brackets::round_close:
                    if (!stack.empty() && stack.back() == brackets::round_open) {
                        stack.pop_back();
                        break;
                    } else {
                        return false;
                    }
            }
        }
        return stack.empty();
    }

}  // namespace matching_brackets
