WITH RECURSIVE compute(pos, next, array) AS (
    SELECT 0, substr(input, 1, 1), json_array()
    UNION
    SELECT pos + 1, substr(input, pos + 2, 1), (
        CASE
            WHEN next IN ("[", "(", "{") THEN
                json_insert(array, "$[#]", next)
            WHEN next = "]" THEN
                iif(
                    (array ->> "$[#-1]") = "[",
                    json_remove(array, "$[#-1]"),
                    json_insert(array, "$[#]", next)
                )
            WHEN next = ")" THEN
                iif(
                    (array ->> "$[#-1]") = "(",
                    json_remove(array, "$[#-1]"),
                    json_insert(array, "$[#]", next)
                )
            WHEN next = "}" THEN
                iif(
                    (array ->> "$[#-1]") = "{",
                    json_remove(array, "$[#-1]"),
                    json_insert(array, "$[#]", next)
                )
            ELSE
                array
        END
    )
    FROM compute
    WHERE pos < length(input)
)
UPDATE "matching-brackets"
SET result = (
    json_array_length((
        SELECT array
        FROM compute
        WHERE pos = length(input)
    )) = 0
);