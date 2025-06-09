WITH RECURSIVE byte_popcount(byte, popcount) AS (
    SELECT 0, 0
    UNION
    SELECT byte + 1, (
        (((byte + 1) >> 0) & 1)
        + (((byte + 1) >> 1) & 1)
        + (((byte + 1) >> 2) & 1)
        + (((byte + 1) >> 3) & 1)
        + (((byte + 1) >> 4) & 1)
        + (((byte + 1) >> 5) & 1)
        + (((byte + 1) >> 6) & 1)
        + (((byte + 1) >> 7) & 1)
    )
    FROM byte_popcount
    LIMIT 256
)
UPDATE "eliuds-eggs"
SET result = (
    SELECT SUM(popcount)
    FROM byte_popcount
    WHERE byte IN (
        (number >> 0) & 255,
        (number >> 8) & 255,
        (number >> 16) & 255,
        (number >> 24) & 255,
        (number >> 32) & 255,
        (number >> 40) & 255,
        (number >> 48) & 255,
        (number >> 56) & 255
    )
);