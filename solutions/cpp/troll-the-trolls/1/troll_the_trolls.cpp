namespace hellmath {

    enum class AccountStatus {
        troll, guest, user, mod,
    };

    enum class Action {
        read, write, remove,
    };

    bool display_post(AccountStatus poster, AccountStatus viewer) {
        if (poster == AccountStatus::troll) {
            return viewer == AccountStatus::troll;
        }
        return true;
    }

    bool permission_check(Action action, AccountStatus account) {
        switch (account) {
            case AccountStatus::mod:
                return true;
            case AccountStatus::user:
                return action == Action::read || action == Action::write;
            case AccountStatus::troll:
                return action == Action::read || action == Action::write;
            case AccountStatus::guest:
                return action == Action::read;
        }
        return false;
    }

    bool valid_player_combination(AccountStatus p1, AccountStatus p2) {
        if (p1 == AccountStatus::troll || p2 == AccountStatus::troll) {
            return p1 == p2;
        }
        return p1 != AccountStatus::guest && p2 != AccountStatus::guest;
    }

    bool has_priority(AccountStatus a1, AccountStatus a2) {
        return a1 > a2;
    }

}  // namespace hellmath