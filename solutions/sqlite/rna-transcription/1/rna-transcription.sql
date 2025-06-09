WITH RECURSIVE
    chars(pos, ch) AS (
        SELECT 1, (
            SELECT rna_ch
            FROM mapping
            WHERE substr(dna, 1, 1) = dna_ch
        )
        UNION
        SELECT pos + 1, (
            SELECT rna_ch
            FROM mapping
            WHERE substr(dna, pos + 1, 1) = dna_ch
        )
        FROM chars
        WHERE pos < length(dna)
    ),
    mapping(dna_ch, rna_ch) AS (
        VALUES
        ("G", "C"),
        ("C", "G"),
        ("T", "A"),
        ("A", "U")
    )
UPDATE "rna-transcription"
SET result = (
    SELECT coalesce(group_concat(ch, ""), "")
    FROM chars
    ORDER BY pos
);


select * from "rna-transcription";