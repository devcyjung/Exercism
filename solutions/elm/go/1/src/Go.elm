module Go exposing (..)

import GoSupport exposing (..)


applyRules : Game -> Rule -> NonValidatingRule -> Rule -> Rule -> Game
applyRules game oneStonePerPointRule captureRule libertyRule koRule =
    let
        aggregate : Result String Game -> Game
        aggregate result = case result of
            Ok newGame -> newGame
            Err newError -> {game | error = newError}
    in
        game
        |> oneStonePerPointRule
        |> Result.map captureRule
        |> Result.andThen libertyRule
        |> Result.andThen koRule
        |> Result.map changePlayer
        |> aggregate