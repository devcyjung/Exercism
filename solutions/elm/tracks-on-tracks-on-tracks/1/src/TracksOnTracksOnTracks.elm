module TracksOnTracksOnTracks exposing (..)

newList : List String
newList = []

existingList : List String
existingList = [ "Elm", "Clojure", "Haskell" ]

addLanguage : String -> List String -> List String
addLanguage language languages = language :: languages

countLanguages : List String -> Int
countLanguages languages =
    case languages of
        [] -> 0
        _ :: xs -> 1 + countLanguages xs
        
reverseList : List String -> List String
reverseList languages =
    let
        reverser input output = 
            case input of
                [] -> output
                x :: xs -> reverser xs (x :: output)
    in
    reverser languages []

excitingList : List String -> Bool
excitingList languages =
    case languages of
        "Elm" :: _ -> True
        _ :: "Elm" :: [] -> True
        _ :: "Elm" :: _ :: [] -> True
        _ -> False
