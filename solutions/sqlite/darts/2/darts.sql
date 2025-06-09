UPDATE darts SET score = CASE
    WHEN sqrt(x*x + y*y) <= 1 THEN 10
    WHEN sqrt(x*x + y*y) <= 5 THEN 5
    WHEN sqrt(x*x + y*y) <= 10 THEN 1
    ELSE 0
END;