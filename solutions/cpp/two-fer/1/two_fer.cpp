#include "two_fer.h"

std::string two_fer::two_fer(std::string_view name) noexcept {
    if (name.empty()) {
        name = "you";
    }
    return "One for " + std::string{name} + ", one for me.";
}