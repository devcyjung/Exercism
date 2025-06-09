module Bandwagoner exposing (..)

type alias Coach =
    { name: String
    , formerPlayer: Bool
    }

type alias Stats =
    { wins: Int
    , losses: Int
    }

type alias Team =
    { name: String
    , coach: Coach
    , stats: Stats
    }

replaceCoach : Coach -> { a | coach : Coach } -> { a | coach : Coach }
replaceCoach newCoach team =
    { team | coach = newCoach }

createTeam : String -> Stats -> Coach -> Team
createTeam name stats coach =
    Team name coach stats

rootForTeam : { a | stats: Stats } -> Bool
rootForTeam team =
    team.stats.wins > team.stats.losses
