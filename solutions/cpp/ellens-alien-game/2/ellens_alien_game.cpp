#include <algorithm>

namespace targets {
    class Alien {
        public:
            Alien(int x, int y): x_coordinate{x}, y_coordinate{y} {
            }
            int get_health() const {
                return health;
            }
            bool hit() {
                health = std::max(0, health-1);
                return true;
            }
            bool is_alive() const {
                return health > 0;
            }
            bool teleport(int x_new, int y_new) {
                x_coordinate = x_new;
                y_coordinate = y_new;
                return true;
            }
            bool collision_detection(Alien other) const {
                return x_coordinate == other.x_coordinate && y_coordinate == other.y_coordinate;
            }
            int x_coordinate{};
            int y_coordinate{};
        private:
            int health{3};
    };
}