module Go exposing (..)

import GoSupport exposing (..)


aggregate : Game -> Result String Game -> Game
aggregate game result = case result of
    Ok newGame -> newGame
    Err newError -> {game | error = newError}
    

applyRules : Game -> Rule -> NonValidatingRule -> Rule -> Rule -> Game
applyRules game rule1 nonValRule rule2 rule3 = game |> rule1 |> Result.map nonValRule
    |> Result.andThen rule2 |> Result.andThen rule3 |> Result.map changePlayer |> aggregate game