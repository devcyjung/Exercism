#include "doctor_data.h"

heaven::Vessel::Vessel(std::string name, int generation)
    : heaven::Vessel::Vessel(name, generation, star_map::System::Sol) {
}

heaven::Vessel::Vessel(std::string name, int generation, star_map::System system)
    : name{name}, generation{generation}, current_system{system} {
}

heaven::Vessel heaven::Vessel::replicate(std::string new_name) {
    return heaven::Vessel(new_name, generation+1, current_system);
}

void heaven::Vessel::make_buster() {
    ++busters;
}

bool heaven::Vessel::shoot_buster() {
    if (busters <= 0) {
        return false;
    }
    --busters;
    return true;
}

std::string heaven::get_older_bob(heaven::Vessel v1, heaven::Vessel v2) {
    return (v1.generation < v2.generation) ? v1.name : v2.name;
}

bool heaven::in_the_same_system(heaven::Vessel v1, heaven::Vessel v2) {
    return v1.current_system == v2.current_system;
}
