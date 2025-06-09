#include "bank_account.h"

#include <mutex>
#include <stdexcept>
#include <string>
#include <system_error>

namespace Bankaccount {
    void Bankaccount::open() {
        try {
            std::unique_lock lock(account_mutex_);
            if (is_open_) {
                throw std::runtime_error("account already open");
            }
            is_open_ = true;
        } catch (const std::system_error& e) {
            throw std::runtime_error("lock error in open() " + std::string(e.what()));
        }
    }
    
    void Bankaccount::close() {
        try {
            std::unique_lock lock(account_mutex_);
            if (!is_open_) {
                throw std::runtime_error("account unopen or already closed");
            }
            is_open_ = false;
            account_balance_ = 0;
        } catch (const std::system_error& e) {
            throw std::runtime_error("lock error in close() " + std::string(e.what()));
        }
    }
    void Bankaccount::deposit(const int& deposit_amount) {
        try {
            std::unique_lock lock(account_mutex_);
            if (!is_open_) {
                throw std::runtime_error("account is closed");
            }
            if (deposit_amount < 0) {
                throw std::runtime_error("cannot deposit negative amount");
            }
            account_balance_ += deposit_amount;
        } catch (const std::system_error& e) {
            throw std::runtime_error("lock error in deposit() " + std::string(e.what()));
        }
    }
    
    void Bankaccount::withdraw(const int& withdraw_amount) {
        try {
            std::unique_lock lock(account_mutex_);
            if (!is_open_) {
                throw std::runtime_error("account is closed");
            }
            if (withdraw_amount < 0) {
                throw std::runtime_error("cannot withdraw negative amount");
            }
            if (account_balance_ < withdraw_amount) {
                throw std::runtime_error("account balance less than withdrawal amount");
            }
            account_balance_ -= withdraw_amount;
        } catch (const std::system_error& e) {
            throw std::runtime_error("lock error in withdraw() " + std::string(e.what()));
        }
    }
    
    int Bankaccount::balance() const {
        try {
            std::shared_lock lock(account_mutex_);
            if (!is_open_) {
                throw std::runtime_error("account is closed");
            }
            return account_balance_;
        } catch (const std::system_error& e) {
            throw std::runtime_error("lock error in balance() " + std::string(e.what()));
        }
    }
}