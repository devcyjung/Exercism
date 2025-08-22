local BankAccount = {}

function BankAccount:new()
  local account = {
    _balance = 0,
    _closed = true,
  }
  return setmetatable(account, {__index = BankAccount})
end

function BankAccount:open()
  if not self._closed then
    error("account is open")
  end
  self._closed = false
end

function BankAccount:deposit(amount)
  if amount <= 0 then
    error("deposit only takes positive amount")
  end
  if self._closed then
    error("account is closed")
  end
  self._balance = self._balance + amount
end

function BankAccount:balance()
  if self._closed then
    error("account is closed")
  end
  return self._balance
end

function BankAccount:withdraw(amount)
  if amount <= 0 then
    error("withdraw only takes positive amount")
  end
  if self._balance < amount then
    error("withdraw amount exceeds current balance")
  end
  if self._closed then
    error("account is closed")
  end
  self._balance = self._balance - amount
end

function BankAccount:close()
  if self._closed then
    error("account is closed")
  end
  self._closed = true
  self._balance = 0
end

return BankAccount