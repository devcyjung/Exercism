module Robot
    ( Bearing(East,North,South,West)
    , bearing
    , coordinates
    , mkRobot
    , move
    ) where

import Data.Foldable (foldl')

data Bearing = North
             | East
             | South
             | West
             deriving (Eq, Show, Enum, Bounded)

data Robot = Robot { bearing :: Bearing
                   , coordinates :: (Integer, Integer)
                   }

mkRobot :: Bearing -> (Integer, Integer) -> Robot
mkRobot = Robot

move :: Robot -> String -> Robot
move = foldl' step
  where
    step robot c = case c of
      'A' -> advance robot
      'L' -> turnLeft robot
      'R' -> turnRight robot
      _ -> robot
      
    advance robot@Robot { bearing = b
                        ,coordinates = (x, y)
                        } =
      case b of
        East -> robot { coordinates = (x+1, y) }
        West -> robot { coordinates = (x-1, y) }
        North -> robot { coordinates = (x, y+1) }
        South -> robot { coordinates = (x, y-1) }
        
    turnLeft robot@Robot { bearing = b } =
      robot { bearing = toEnum $ (fromEnum b - 1) `mod` 4 }
      
    turnRight robot@Robot { bearing = b } =
      robot { bearing = toEnum $ (fromEnum b + 1) `mod` 4 }