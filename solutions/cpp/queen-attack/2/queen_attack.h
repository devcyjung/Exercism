#ifndef QUEEN_ATTACK_H_
#define QUEEN_ATTACK_H_

#include <stdexcept>
#include <utility>

namespace queen_attack {

class chess_board {
public:
    using position = std::pair<int, int>;
    
private:
    position white_, black_;

public:
    static inline constexpr int MASK = ~7;
    
    [[nodiscard]]
    inline constexpr chess_board(const position& white, const position& black):
        white_(white), black_(black)
    {
        if (white == black) {
            throw std::domain_error("Duplicate positions");
        }
        const auto off_bounds = (white_.first & MASK) | (black_.first & MASK)
            | (white_.second & MASK) | (black_.second & MASK);
        if (off_bounds != 0) {
            throw std::domain_error("Position is off the board");
        }
    }

    [[nodiscard]]    
    inline constexpr position white() const noexcept {
        return white_;    
    }

    [[nodiscard]]
    inline constexpr position black() const noexcept {
        return black_;
    }
    
    [[nodiscard]]
    inline constexpr bool can_attack() const noexcept {
        const auto diff_first = white_.first - black_.first;
        const auto diff_second = white_.second - black_.second;
        return diff_first * diff_second
            * (diff_first + diff_second)
            * (diff_first - diff_second) == 0;
    }
};
    
}  // namespace queen_attack

#endif