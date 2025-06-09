object EliudsEggs {
    private val byteTable = Array<Int>(0x100){ 0 }

    init {
        for (i in 0x1..0xFF) {
            byteTable[i] = byteTable[i shr 1] + (i and 0x1)
        }
    }
    
    fun eggCount(number: Int): Int {
        return byteTable[number shr 8 * 0 and 0xFF] +
               byteTable[number shr 8 * 1 and 0xFF] +
               byteTable[number shr 8 * 2 and 0xFF] +
               byteTable[number shr 8 * 3 and 0xFF]
    }
}