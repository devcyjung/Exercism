import kotlin.math.sqrt

object CryptoSquare {
    fun ciphertext(plaintext: String): String = run {
        val plaintext = plaintext.filter{ it.isLetterOrDigit() }.lowercase()
        val length = plaintext.length
        val root = sqrt(length.toDouble()).toInt()
        val (r, c) = if (root * root == length) {
            root to root
        } else if (root * (root + 1) >= length) {
            root to (root + 1)
        } else {
            (root + 1) to (root + 1)
        }
        (0..<c).map{ ci ->
            (0..<r).map{ ri ->
                plaintext.getOrElse(ri * c + ci) { ' ' }
            }.joinToString(separator = "")
        }.joinToString(separator = " ")
    }
}