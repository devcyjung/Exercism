#include "robot_simulator.h"
#include <stddef.h>

const int DIRECTION_MAX = 4;

robot_status_t robot_create(robot_direction_t direction, int x, int y) {
    return (robot_status_t) {
        .direction = direction,
        .position = {
            .x = x,
            .y = y,
        },
    };
}

void robot_move(robot_status_t *robot, const char *const commands) {
    if (!commands) {
        return;
    }
    size_t i = 0;
    char cur;
    while ((cur = commands[i++])) {
        switch (cur) {
        case 'R':
            ++(robot->direction);
            robot->direction %= DIRECTION_MAX;
            break;
        case 'L':
            --(robot->direction);
            robot->direction += DIRECTION_MAX;
            robot->direction %= DIRECTION_MAX;
            break;
        case 'A':
            switch (robot->direction) {
            case DIRECTION_NORTH:
                ++(robot->position.y);
                break;
            case DIRECTION_SOUTH:
                --(robot->position.y);
                break;
            case DIRECTION_EAST:
                ++(robot->position.x);
                break;
            case DIRECTION_WEST:
                --(robot->position.x);
                break;
            }
        }
    }
}