module SqueakyClean exposing (clean, clean1, clean2, clean3, clean4)


clean1 : String -> String
clean1 = String.replace " " "_"


clean2 : String -> String
clean2 = clean1 >> String.replace "\n" "[CTRL]" >> String.replace "\r" "[CTRL]" >> String.replace "\t" "[CTRL]"


clean3 : String -> String
clean3 =
    let
        capitalize : Int -> String -> String
        capitalize index str =
            if index == 0 then
                str
            else
                String.uncons str
                |> Maybe.map (\(first, rest) -> String.cons (Char.toUpper first) rest)
                |> Maybe.withDefault ""
    in
        clean2
        >> String.split "-"
        >> List.indexedMap capitalize
        >> String.concat


clean4 : String -> String
clean4 = clean3 >> String.filter (Char.isDigit >> not)


clean : String -> String
clean =
    let
        alphaCode = Char.toCode 'α'
        omegaCode = Char.toCode 'ω'
        isNotLowerGreek : Char -> Bool
        isNotLowerGreek ch =
            let code = Char.toCode ch
            in alphaCode > code || code > omegaCode              
    in
        clean4 >> String.filter isNotLowerGreek
