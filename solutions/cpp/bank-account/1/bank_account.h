#if !defined(BANK_ACCOUNT_H)
#define BANK_ACCOUNT_H

#include <mutex>

namespace Bankaccount {
class Bankaccount {
public:
    Bankaccount() : is_open_{false}, account_balance_{0}, account_mutex_{} {}
    void open();
    void close();
    void deposit(const int&);
    void withdraw(const int&);
    [[nodiscard]] int balance();
private:
    bool is_open_;
    int account_balance_;
    std::mutex account_mutex_;
};  // class Bankaccount
}  // namespace Bankaccount

#endif  // BANK_ACCOUNT_H