defmodule FreelancerRates do
  def daily_rate(hourly_rate) do
    hourly_rate
    |> Kernel.*(8.0)
  end

  def apply_discount(before_discount, discount) do
    100
    |> Kernel.-(discount)
    |> Kernel./(100)
    |> Kernel.*(before_discount)
  end

  def monthly_rate(hourly_rate, discount) do
    hourly_rate
    |> daily_rate
    |> Kernel.*(22)
    |> apply_discount(discount)
    |> ceil
  end

  def days_in_budget(budget, hourly_rate, discount) do
    daily_discounted_rate =
      hourly_rate
      |> daily_rate
      |> apply_discount(discount)
      
    budget
    |> Kernel./(daily_discounted_rate)
    |> Float.floor(1)
  end
end
