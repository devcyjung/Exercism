#include "speedywagon.h"

namespace speedywagon {

int uv_light_heuristic(std::vector<int>* data_array) {
    double avg{};
    for (auto element : *data_array) {
        avg += element;
    }
    avg /= data_array->size();
    int uv_index{};
    for (auto element : *data_array) {
        if (element > avg) ++uv_index;
    }
    return uv_index;
}

bool connection_check(pillar_men_sensor* sensor) {
    return sensor;
}

int activity_counter(pillar_men_sensor* first, int capacity){
    int acc{0};
    for (int i = 0; i < capacity; ++i) {
        acc += (first + i)->activity;
    }
    return acc;
}

bool alarm_control(pillar_men_sensor* sensor) {
    if (!sensor) {
        return false;
    }
    return sensor->activity > 0;
}

bool uv_alarm(pillar_men_sensor* sensor) {
    if (!sensor) {
        return false;
    }
    return uv_light_heuristic(&(sensor->data)) > sensor->activity;
}
    
}  // namespace speedywagon
