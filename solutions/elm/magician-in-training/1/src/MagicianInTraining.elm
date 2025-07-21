module MagicianInTraining exposing (..)

import Array exposing (Array)
import Bitwise exposing (and)

getCard : Int -> Array Int -> Maybe Int
getCard = Array.get


setCard : Int -> Int -> Array Int -> Array Int
setCard = Array.set


addCard : Int -> Array Int -> Array Int
addCard = Array.push


removeCard : Int -> Array Int -> Array Int
removeCard index deck =
    let
        front = Array.slice 0 index deck
        back = Array.slice (index + 1) (Array.length deck) deck
    in
        Array.append front back

    
evenCardCount : Array Int -> Int
evenCardCount =
    Array.filter (\i -> and i 1 == 0) >> Array.length