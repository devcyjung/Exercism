defmodule Lasagna do
  def expected_minutes_in_oven() do
    40
  end
  
  def remaining_minutes_in_oven(elapsed_time) do
    expected_minutes_in_oven() - elapsed_time
  end
  
  def preparation_time_in_minutes(layers) do
    layers * 2
  end
  
  def total_time_in_minutes(layers, elapsed_time) do
    preparation_time_in_minutes(layers) + elapsed_time
  end
  
  def alarm do
    "Ding!"
  end
end
