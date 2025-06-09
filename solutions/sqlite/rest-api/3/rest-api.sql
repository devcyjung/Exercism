UPDATE "rest-api" AS rest
SET result = 
CASE
    WHEN rest.payload = "{}" THEN
        rest.database
    ELSE (
        SELECT json_object("users", json_array(user.value))
        FROM json_each(json_extract(rest.database, "$.users")) AS user
        WHERE json_extract(user.value, "$.name") IN (
            SELECT value FROM json_each(json_extract(rest.payload, "$.users"))       
        )
    )
END
WHERE url = "/users";

UPDATE "rest-api" AS rest
SET result = json_object(
    "name", json_extract(rest.payload, "$.user"),
    "owes", json_object(),
    "owed_by", json_object(),
    "balance", 0
)
WHERE url = "/add";

WITH
    lender AS (
        SELECT
            user.value AS lender,
            json_extract(user.value, "$.name") AS lender_name
        FROM json_each(json_extract(rest.database, "$.users")) as user
        WHERE json_extract(user.value, "$.name") = json_extract(rest.payload, "$.lender")
    ),
    borrower AS (
        SELECT
            user.value AS borrower,
            json_extract(user.value, "$.name") AS borrower_name
        FROM json_each(json_extract(rest.database, "$.users")) as user
        WHERE json_extract(user.value, "$.name") = json_extract(rest.payload, "$.borrower")
    ),
    json_selectors AS (
        SELECT
            "$.owes." || borrower_name AS owes_borrower,
            "$.owed_by." || borrower_name AS owed_by_borrower,
            "$.owes." || lender_name AS owes_lender,
            "$.owed_by." || lender_name AS owed_by_lender
        FROM lender, borrower
    ),
    amount AS (
        SELECT
            json_extract(rest.payload, "$.amount") AS amount,
            IFNULL(json_extract(lender, owes_borrower), 0) AS lender_owes,
            IFNULL(json_extract(lender, owed_by_borrower), 0) AS lender_owed_by,
            IFNULL(json_extract(borrower, owes_lender), 0) AS borrower_owes,
            IFNULL(json_extract(borrower, owed_by_lender), 0) AS borrower_owed_by,
            json_extract(lender, "$.balance") AS lender_balance,
            json_extract(borrower, "$.balance") AS borrower_balance
        FROM lender, borrower, json_selectors
    ),
    net_owe_amount AS (
        SELECT
            lender_owes - lender_owed_by - amount AS lender_net_owes,
            borrower_owes - borrower_owed_by + amount AS borrower_net_owes,
            lender_name,
            borrower_name
        FROM amount, lender, borrower
    ),
    amount_updated_lender AS (
        SELECT json_set(
            lender,
            "$.balance",
            lender_balance + amount.amount
        ) AS lender
        FROM lender, amount
    ),
    amount_updated_borrower AS (
        SELECT json_set(
            borrower,
            "$.balance",
            borrower_balance - amount.amount
        ) AS borrower
        FROM borrower, amount
    ),
    
    owe_updated_lender AS (
        SELECT
            CASE
                WHEN lender_net_owes = 0 THEN
                    json_remove(
                        json_remove(lender, owes_borrower),
                        owed_by_borrower
                    )
                WHEN lender_net_owes > 0 THEN
                    json_set(
                        json_remove(lender, owed_by_borrower),
                        owes_borrower,
                        lender_net_owes
                    )
                WHEN lender_net_owes < 0 THEN
                    json_set(
                        json_remove(lender, owes_borrower),
                        owed_by_borrower,
                        -lender_net_owes
                    )
            END AS lender
        FROM net_owe_amount, amount_updated_lender, json_selectors
    ),
    owe_updated_borrower AS (
        SELECT
            CASE
                WHEN borrower_net_owes = 0 THEN
                    json_remove(
                        json_remove(borrower, owes_lender),
                        owed_by_lender
                    )
                WHEN borrower_net_owes > 0 THEN
                    json_set(
                        json_remove(borrower, owed_by_lender),
                        owes_lender,
                        borrower_net_owes
                    )
                WHEN borrower_net_owes < 0 THEN
                    json_set(
                        json_remove(borrower, owes_lender),
                        owed_by_lender,
                        -borrower_net_owes
                    )
            END AS borrower
        FROM net_owe_amount, amount_updated_borrower, json_selectors
    )
UPDATE "rest-api" AS rest
SET result = (
    SELECT json_object("users", json_group_array(json(user)))
    FROM (
        SELECT
            lender AS user,
            lender_name as name
        FROM net_owe_amount, owe_updated_lender
        UNION
        SELECT
            borrower AS user,
            borrower_name as name
        FROM net_owe_amount, owe_updated_borrower
        ORDER BY name
    )
)
WHERE url = "/iou";