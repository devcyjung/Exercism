class AssemblyLine
  PER_HOUR_UNIT = 221
  
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    PER_HOUR_UNIT * @speed * case @speed
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
