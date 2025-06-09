return {
  new = function(self)
    local function bank_account_actor()
      local balance = 0
      local closed = false
    
      while true do
        local req = coroutine.yield()
    
        if req.cmd == "balance" then
          req.resp(true, balance)
    
        elseif req.cmd == "deposit" then
          if closed then
            req.resp(false, "Account closed")
          elseif type(req.amount) ~= "number" or req.amount <= 0 then
            req.resp(false, "Invalid deposit amount")
          else
            balance = balance + req.amount
            req.resp(true, balance)
          end
    
        elseif req.cmd == "withdraw" then
          if closed then
            req.resp(false, "Account closed")
          elseif type(req.amount) ~= "number" or req.amount <= 0 then
            req.resp(false, "Invalid withdraw amount")
          elseif req.amount > balance then
            req.resp(false, "Insufficient funds")
          else
            balance = balance - req.amount
            req.resp(true, balance)
          end
    
        elseif req.cmd == "close" then
          closed = true
          req.resp(true, "Account closed")
        end
      end
    end

    local co = coroutine.create(bank_account_actor)
    coroutine.resume(co)
  
    local function send_request(cmd, amount)
      local result, response
      local function callback(ok, res)
        result = ok
        response = res
      end
      coroutine.resume(co, { cmd = cmd, amount = amount, resp = callback })
      return result, response
    end
  
    return {
      balance = function(self)
        local ok, res = send_request("balance")
        if not ok then
          error(string.format([[              
              Error occurred in balance(): %s
          ]], tostring(res)))
        end
        return res
      end,
      deposit = function(self, amount)
        local ok, res = send_request("deposit", amount)
        if not ok then
          error(string.format([[              
              Error occurred in deposit(%s): %s
          ]], tostring(amount), tostring(res)))
        end
      end,
      withdraw = function(self, amount)
        local ok, res = send_request("withdraw", amount)
        if not ok then
          error(string.format([[              
              Error occurred in withdraw(%s): %s
          ]], tostring(amount), tostring(res)))
        end
      end,
      close = function(self)
        local ok, res = send_request("close")
        if not ok then
          error(string.format([[              
              Error occurred in close(): %s
          ]], tostring(res)))
        end
      end,
    }
  end
}