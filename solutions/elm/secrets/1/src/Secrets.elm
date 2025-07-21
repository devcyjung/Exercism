module Secrets exposing (clearBits, decrypt, flipBits, setBits, shiftBack)

import Bitwise


shiftBack : Int -> Int -> Int
shiftBack = Bitwise.shiftRightZfBy


setBits : Int -> Int -> Int
setBits = Bitwise.or


flipBits : Int -> Int -> Int
flipBits = Bitwise.xor


clearBits : Int -> Int -> Int
clearBits mask = Bitwise.and <| Bitwise.complement mask


decrypt : Int -> Int
decrypt = setBits 1996 >> flipBits 2009 >> shiftBack 5 >> clearBits 17
