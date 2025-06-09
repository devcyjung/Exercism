#pragma once

#include <string>

namespace star_map {
    enum class System {
        BetaHydri, EpsilonEridani, Sol, AlphaCentauri, DeltaEridani, Omicron2Eridani,
    };
}

namespace heaven {
    class Vessel {
        public:
            Vessel(std::string, int);
            Vessel(std::string, int, star_map::System);
            Vessel replicate(std::string);
            
            void make_buster();
            bool shoot_buster();
            
            std::string name{};
            int generation{};
            star_map::System current_system{};
            int busters{0};
    };

    std::string get_older_bob(Vessel, Vessel);
    bool in_the_same_system(Vessel, Vessel);
}