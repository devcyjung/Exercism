#include <string>

namespace log_line {
std::string message(std::string line) {
    auto message_begin = line.find("]: ") + 3;
    return line.substr(message_begin);
}

std::string log_level(std::string line) {
    auto level_begin = line.find("[")+1;
    auto level_end = line.find("]");
    auto level_length = level_end - level_begin;
    return line.substr(level_begin, level_length);
}

std::string reformat(std::string line) {
    return message(line) + " (" + log_level(line) + ")";
}
}
