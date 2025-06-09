#include "kindergarten_garden.h"

namespace kindergarten_garden {

    const std::array<Plants, 4> plants(const std::string& layout, const std::string& name) {
        auto it = students.find(name);
        if (it == students.end()) {
            throw std::runtime_error("student name " + name + " not found" );
        }
        auto id = it->second;
        auto nl = layout.find("\n");
        if (nl == std::string::npos) {
            throw std::runtime_error("invalid layout");
        }
        if (nl <= 2 * id + 1) {
            throw std::runtime_error("student " + name + " doesn't have a spot in " + layout);
        }
        std::array<Plants, 4> result{};
        result[0] = decode(layout[2*id]);
        result[1] = decode(layout[2*id+1]);
        result[2] = decode(layout[nl+1+2*id]);
        result[3] = decode(layout[nl+2+2*id]);
        return result;
    }

}  // namespace kindergarten_garden
