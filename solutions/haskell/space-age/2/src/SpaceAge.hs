module SpaceAge (Planet(..), ageOn) where

ageOn :: Planet -> Float -> Float
ageOn p s  = s / secondsInPlanetYear p

secondsInPlanetYear :: Planet -> Float
secondsInPlanetYear p  = 365.25 * 24 * 60 * 60 * orbitalPeriod p

orbitalPeriod :: Planet -> Float
orbitalPeriod Mercury  = 0.2408467
orbitalPeriod Venus    = 0.61519726
orbitalPeriod Earth    = 1
orbitalPeriod Mars     = 1.8808158
orbitalPeriod Jupiter  = 11.862615
orbitalPeriod Saturn   = 29.447498
orbitalPeriod Uranus   = 84.016846
orbitalPeriod Neptune  = 164.79132

data Planet = Mercury
            | Venus
            | Earth
            | Mars
            | Jupiter
            | Saturn
            | Uranus
            | Neptune