WITH RECURSIVE product(product, rem) AS (
        SELECT pow(CAST(number AS INTEGER) % 10, len), CAST(number AS INTEGER) / 10
        FROM len
        UNION
        SELECT pow(rem % 10, len), rem / 10
        FROM product, len
        WHERE rem > 0
    ),
    len(len) AS (
        SELECT length(cast(number as text))
    ),
    result(result) AS (
        SELECT sum(product)
        FROM product
    )
    
UPDATE "armstrong-numbers"
SET result = (
    SELECT result = number
    FROM result
);