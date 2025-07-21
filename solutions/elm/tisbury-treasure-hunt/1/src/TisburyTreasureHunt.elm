module TisburyTreasureHunt exposing (..)


type alias TreasureLocation = (Int, Char)
type alias Treasure = (String, TreasureLocation)
type alias PlaceLocation = (Char, Int)
type alias Place = (String, PlaceLocation)


placeLocationToTreasureLocation : PlaceLocation -> TreasureLocation
placeLocationToTreasureLocation (ch, i) = (i, ch)


treasureLocationMatchesPlaceLocation : PlaceLocation -> TreasureLocation -> Bool
treasureLocationMatchesPlaceLocation placeLocation treasureLocation =
    treasureLocation == placeLocationToTreasureLocation placeLocation


countPlaceTreasures : Place -> List Treasure -> Int
countPlaceTreasures (_, placeLocation) =
    let
        locationToSearch = placeLocationToTreasureLocation placeLocation
    in
        List.filter (\(_, treasureLocation) -> treasureLocation == locationToSearch) >> List.length
    

specialCaseSwapPossible : Treasure -> Place -> Treasure -> Bool
specialCaseSwapPossible (foundTreasure, _) (place, _) (desiredTreasure, _) =
    case (foundTreasure, place) of
        ("Brass Spyglass", "Abandoned Lighthouse") -> True
        ("Amethyst Octopus", "Stormy Breakwater") -> case desiredTreasure of
            "Crystal Crab" -> True
            "Glass Starfish" -> True
            _ -> False
        ("Vintage Pirate Hat", "Harbor Managers Office") -> case desiredTreasure of
            "Model Ship in Large Bottle" -> True
            "Antique Glass Fishnet Float" -> True
            _ -> False
        _ -> False
