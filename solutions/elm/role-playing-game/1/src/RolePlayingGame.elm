module RolePlayingGame exposing (Player, castSpell, introduce, revive)


type alias Player =
    { name : Maybe String
    , level : Int
    , health : Int
    , mana : Maybe Int
    }


introduce : Player -> String
introduce { name } =
    Maybe.withDefault "Mighty Magician" name


revive : Player -> Maybe Player
revive player =
    if player.health > 0 then
        Nothing
    else
        Just {player | health = 100, mana = if player.level < 10 then Nothing else Just 100}


castSpell : Int -> Player -> ( Player, Int )
castSpell manaCost player =
    case player.mana of
        Nothing -> ({player | health = max (player.health - manaCost) 0}, 0)
        Just mana -> if mana >= manaCost then ({player | mana = Maybe.map ((+) -manaCost) player.mana}, 2 * manaCost) else (player, 0)
