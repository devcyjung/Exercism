WITH byte_popcount(byte, popcount) AS (
    SELECT value, (
        (value >> 0 & 1)
        + (value >> 1 & 1)
        + (value >> 2 & 1)
        + (value >> 3 & 1)
        + (value >> 4 & 1)
        + (value >> 5 & 1)
        + (value >> 6 & 1)
        + (value >> 7 & 1)
    )
    FROM generate_series(0, 255)
)
UPDATE "eliuds-eggs"
SET result = (
    SELECT SUM(popcount)
    FROM byte_popcount
    WHERE byte IN (
        number >> 0 & 255,
        number >> 8 * 1 & 255,
        number >> 8 * 2 & 255,
        number >> 8 * 3 & 255,
        number >> 8 * 4 & 255,
        number >> 8 * 5 & 255,
        number >> 8 * 6 & 255,
        number >> 8 * 7 & 255
    )
);