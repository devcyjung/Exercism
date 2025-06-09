UPDATE "difference-of-squares"
    SET result =
        CASE property
            WHEN "squareOfSum" THEN
                number * number * (number + 1) * (number + 1) / 4
            WHEN "sumOfSquares" THEN
                number * (number + 1) * (2 * number + 1) / 6
            WHEN "differenceOfSquares" THEN
                number * number * (number + 1) * (number + 1) / 4 - number * (number + 1) * (2 * number + 1) / 6
        END;