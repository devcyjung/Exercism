module MazeMaker exposing (..)


import Random exposing (Generator)


type Maze
    = DeadEnd
    | Room Treasure
    | Branch (List Maze)


type Treasure
    = Gold
    | Diamond
    | Friendship


deadend : Generator Maze
deadend = Random.constant DeadEnd


treasure : Generator Treasure
treasure = Random.uniform Gold [Diamond, Friendship]


room : Generator Maze
room = Random.map Room treasure


branch : Generator Maze -> Generator Maze
branch generator = Random.map Branch <|
    Random.andThen (\len -> Random.list len generator) (Random.int 2 4)


maze : Generator Maze
maze =
    Random.andThen identity <|
        Random.weighted (60, deadend) [(15, room), (25, branch (Random.lazy (\_ -> maze)))]


mazeOfDepth : Int -> Generator Maze
mazeOfDepth depth =
        if depth <= 0 then
            Random.andThen identity <| Random.uniform deadend [room]
        else
            branch (Random.lazy (\_ -> mazeOfDepth (depth - 1)))

