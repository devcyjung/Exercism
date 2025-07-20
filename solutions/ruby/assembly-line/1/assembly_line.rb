class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    221 * @speed * case @speed
    in 1..4
      1
    in 5..8
      0.9
    in 9
      0.8
    in 10
      0.77
    else
      raise "invalid speed"
    end
  end

  def working_items_per_minute
    (production_rate_per_hour / 60).floor
  end
end
