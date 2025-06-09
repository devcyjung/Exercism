UPDATE leap
SET is_leap = year % 4 = 0 AND year % 100 <> 0 OR year % 400 = 0;