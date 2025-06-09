module LuciansLusciousLasagna exposing (elapsedTimeInMinutes, expectedMinutesInOven, preparationTimeInMinutes)
expectedMinutesInOven = 40
preparationTimeInMinutes layers = 2 * layers
elapsedTimeInMinutes layers time_in_oven = time_in_oven + preparationTimeInMinutes layers