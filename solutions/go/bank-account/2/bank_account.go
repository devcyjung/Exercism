package account

import (
    "sync"
    "sync/atomic"
)

type Account struct {
    balance	atomic.Int64
    closed	atomic.Bool
    lock	sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
        return nil
    }
    acc := Account{}
    acc.balance.Store(amount)
    return &acc
}

func (a *Account) Balance() (bal int64, ok bool) {
	a.lock.RLock()
    defer a.lock.RUnlock()
    if a.closed.Load() {
        return
    }
    bal, ok = a.balance.Load(), true
    return
}

func (a *Account) Deposit(amount int64) (bal int64, ok bool) {
	a.lock.Lock()
    defer a.lock.Unlock()
    if a.closed.Load() {
        return
    }
    if amount < 0 {
        current := a.balance.Load()
        if current + amount < 0 {
            return
        }
    }
    bal, ok = a.balance.Add(amount), true
    return
}

func (a *Account) Close() (bal int64, ok bool) {
	a.lock.Lock()
    defer a.lock.Unlock()
    if a.closed.Load() {
        return
    }
    a.closed.Store(true)
    bal, ok = a.balance.Swap(0), true
    return
}
