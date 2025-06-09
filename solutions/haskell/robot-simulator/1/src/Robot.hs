module Robot
    ( Bearing(East,North,South,West)
    , bearing
    , coordinates
    , mkRobot
    , move
    ) where

data Bearing = North
             | East
             | South
             | West
             deriving (Eq, Show)

data Robot = Robot { bearing :: Bearing
                   , coordinates :: (Integer, Integer)
                   }

mkRobot :: Bearing -> (Integer, Integer) -> Robot
mkRobot = Robot

move :: Robot -> String -> Robot
move robot "" = robot
move robot ('R':xs) = case bearing robot of
  North -> move robot{ bearing = East } xs
  East -> move robot{ bearing = South } xs
  South -> move robot{ bearing = West } xs
  West -> move robot{ bearing = North } xs
  
move robot ('L':xs) = case bearing robot of
  North -> move robot{ bearing = West } xs
  East -> move robot{ bearing = North } xs
  South -> move robot{ bearing = East } xs
  West -> move robot{ bearing = South } xs
  
move robot ('A':xs) = case bearing robot of
  North -> let (x, y) = coordinates robot
           in move robot{ coordinates = (x, y + 1) } xs
  East -> let (x, y) = coordinates robot
           in move robot{ coordinates = (x + 1, y) } xs
  South -> let (x, y) = coordinates robot
           in move robot{ coordinates = (x, y - 1) } xs
  West -> let (x, y) = coordinates robot
           in move robot{ coordinates = (x - 1, y) } xs