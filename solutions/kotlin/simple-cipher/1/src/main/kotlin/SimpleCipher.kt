data class Cipher(val key: String = (1..100).map{ ('a'..'z').random() }.joinToString("")) {
    init {
        require(!key.isEmpty() && key.all { it in 'a'..'z' })
    }
    
    fun encode(s: String): String = s.mapIndexed { sIdx, sChar ->
        val keyCode = key[sIdx % key.length].code - 'a'.code
        ((sChar.code - 'a'.code + keyCode) % 26 + 'a'.code).toChar()
    }.joinToString("")

    fun decode(s: String): String = s.mapIndexed { sIdx, sChar ->
        val keyCode = key[sIdx % key.length].code - 'a'.code
        ((sChar.code - 'a'.code - keyCode).mod(26) + 'a'.code).toChar()
    }.joinToString("")
}