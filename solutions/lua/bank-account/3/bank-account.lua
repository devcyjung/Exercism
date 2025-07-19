local BankAccount = {
  _balance = 0,
  _closed = false,
}

function BankAccount:new()
  local account = {}
  return setmetatable(account, {__index = BankAccount})
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
  self._closed = true
  self._balance = 0
end

return BankAccount