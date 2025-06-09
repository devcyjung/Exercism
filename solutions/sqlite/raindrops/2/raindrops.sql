UPDATE raindrops
SET sound = iif(mod(number, 3) = 0, "Pling", "") ||
            iif(mod(number, 5) = 0, "Plang", "") ||
            iif(mod(number, 7) = 0, "Plong", "");

UPDATE raindrops
SET sound = number
WHERE sound = "";