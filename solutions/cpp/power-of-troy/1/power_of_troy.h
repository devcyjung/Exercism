#pragma once

#include <memory>
#include <string>

namespace troy {

struct artifact {
    // constructors needed (until C++20)
    artifact(std::string name) : name(name) {}
    std::string name;
};

struct power {
    // constructors needed (until C++20)
    power(std::string effect) : effect(effect) {}
    std::string effect;
};

struct human {
    std::unique_ptr<artifact> possession;
    std::shared_ptr<power> own_power;
    std::shared_ptr<power> influenced_by;
};

void give_new_artifact(human&, std::string);
void exchange_artifacts(std::unique_ptr<artifact>&, std::unique_ptr<artifact>&);
void manifest_power(human&, std::string);
void use_power(human&, human&);
int power_intensity(human&);

}  // namespace troy
