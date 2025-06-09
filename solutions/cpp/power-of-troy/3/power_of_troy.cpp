#include "power_of_troy.h"

namespace troy {
    
void give_new_artifact(human& receiver, std::string artifact_name) {
    receiver.possession = std::make_unique<artifact>(artifact(artifact_name));
}

void exchange_artifacts(std::unique_ptr<artifact>& artifact_1, std::unique_ptr<artifact>& artifact_2) {
    std::swap(artifact_1, artifact_2);
}
    
void manifest_power(human& manifester, std::string power_name) {
    manifester.own_power = std::make_shared<power>(power(power_name));
}
    
void use_power(human& caster, human& target) {
    target.influenced_by = caster.own_power;
}
    
int power_intensity(human& power_owner) {
    if (!power_owner.own_power) {
        return 0;
    }
    return power_owner.own_power.use_count();
}
    
}  // namespace troy
