#include "power_of_troy.h"

namespace troy {
    
void give_new_artifact(human& h, std::string s) {
    h.possession = std::make_unique<artifact>(artifact(s));
}

void exchange_artifacts(std::unique_ptr<artifact>& a1, std::unique_ptr<artifact>& a2) {
    auto temp = std::move(a1);
    a1 = std::move(a2);
    a2 = std::move(temp);
}
    
void manifest_power(human& h, std::string s) {
    h.own_power = std::make_shared<power>(power(s));
}
    
void use_power(human& caster, human& target) {
    target.influenced_by = caster.own_power;
}
    
int power_intensity(human& h) {
    if (!h.own_power) {
        return 0;
    }
    return h.own_power.use_count();
}
    
}  // namespace troy
