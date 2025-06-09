UPDATE color_code
SET result =
    CASE
        WHEN color = "black" THEN 0
        WHEN color = "brown" THEN 1
        WHEN color = "red" THEN 2
        WHEN color = "orange" THEN 3
        WHEN color = "yellow" THEN 4
        WHEN color = "green" THEN 5
        WHEN color = "blue" THEN 6
        WHEN color = "violet" THEN 7
        WHEN color = "grey" THEN 8
        WHEN color = "white" THEN 9
    END;