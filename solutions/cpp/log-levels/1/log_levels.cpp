#include <string>

namespace log_line {
std::string message(std::string line) {
    int message_begin = line.find("]: ") + 3;
    return line.substr(message_begin);
}

std::string log_level(std::string line) {
    int level_begin = line.find("[")+1;
    int level_end = line.find("]");
    int level_length = level_end - level_begin;
    return line.substr(level_begin, level_length);
}

std::string reformat(std::string line) {
    return message(line) + " (" + log_level(line) + ")";
}
}
