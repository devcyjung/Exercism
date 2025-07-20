module SavingsAccount
  def self.interest_rate(balance)
    case balance
    in ...0
      3.213
    in 0...1000
      0.5
    in 1000...5000
      1.621
    else
      2.475
    end
  end

  def self.annual_balance_update(balance)
    (1 + self.interest_rate(balance) / 100) * balance
  end

  def self.years_before_desired_balance(current_balance, desired_balance)
    years = 0
    until current_balance >= desired_balance
      current_balance = self.annual_balance_update(current_balance)
      years += 1
    end
    years
  end
end
