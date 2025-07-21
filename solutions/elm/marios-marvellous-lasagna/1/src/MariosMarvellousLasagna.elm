module MariosMarvellousLasagna exposing (remainingTimeInMinutes)

remainingTimeInMinutes : Int -> Int -> Int
remainingTimeInMinutes layers minutesInOven
    = let
        preparationTimeInMinutes : Int -> Int
        preparationTimeInMinutes numberOfLayers = 2 * numberOfLayers
        expectedMinutesInOven : Int
        expectedMinutesInOven = 40
    in
        preparationTimeInMinutes layers + expectedMinutesInOven - minutesInOven
    