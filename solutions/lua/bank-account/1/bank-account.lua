return {
  new = function(self)
    local closed_ = false
    local balance_ = 0
    return {
      balance = function(self)
        return balance_
      end,
      deposit = function(self, amount)
        if type(amount) ~= "number" or amount <= 0 then
          error(string.format([[            
            deposit() expects positive number
            received type: %s
            received value: %s
          ]], type(amount), amount))
        end
        if closed_ then
          error([[            
            deposit() cannot be called on a closed account
          ]])
        end
        balance_ = balance_ + amount
      end,
      withdraw = function(self, amount)
        if type(amount) ~= "number" or amount <= 0 then
          error(string.format([[            
            withdraw() expects positive number
            received type: %s
            received value: %s
          ]], type(amount), amount))
        end
        if balance_ - amount < 0 then
          error(string.format([[
            withdraw() expects current balance >= withdraw amount
            current balance: %s
            withdraw amount: %s
          ]], tostring(balance_), tostring(amount)))
        end
        if closed_ then
          error([[            
            withdraw() cannot be called on a closed account
          ]])
        end
        balance_ = balance_ - amount
      end,
      close = function(self)
        closed_ = true
      end,
    }
  end
}