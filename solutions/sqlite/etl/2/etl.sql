WITH flattened AS (
    SELECT
        input,
        json_group_object(output_key, output_value) AS output
    FROM (
        SELECT
            input,
            lower(sublevel.value) AS output_key,
            CAST(toplevel.key AS INTEGER) AS output_value
        FROM
            etl,
            json_each(input) AS toplevel,
            json_each(toplevel.value) AS sublevel
        ORDER BY
            input, output_key
    )
    GROUP BY input
)
UPDATE etl
SET result = (
    SELECT output
    FROM flattened
    WHERE flattened.input = etl.input
);