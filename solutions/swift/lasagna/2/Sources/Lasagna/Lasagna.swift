let expectedMinutesInOven = 40

func remainingMinutesInOven(elapsedMinutes: Int) -> Int {
  expectedMinutesInOven - elapsedMinutes
}

func preparationTimeInMinutes(layers: Int) -> Int {
  2 * layers
}

func totalTimeInMinutes(layers: Int, elapsedMinutes: Int) -> Int {
  preparationTimeInMinutes(layers: layers) + elapsedMinutes
}