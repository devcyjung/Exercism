UPDATE "difference-of-squares"
SET result = number * number * (number + 1) * (number + 1) / 4
WHERE property = "squareOfSum";

UPDATE "difference-of-squares"
SET result = number * (number + 1) * (2 * number + 1) / 6
WHERE property = "sumOfSquares";

UPDATE "difference-of-squares" AS d
SET result = (
    (SELECT result FROM "difference-of-squares" WHERE number = d.number AND property = "squareOfSum")
    -
    (SELECT result FROM "difference-of-squares" WHERE number = d.number AND property = "sumOfSquares")
)
WHERE property = "differenceOfSquares";