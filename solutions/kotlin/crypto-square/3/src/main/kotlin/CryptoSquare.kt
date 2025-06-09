import kotlin.math.sqrt

object CryptoSquare {
    fun ciphertext(plaintext: String): String {
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
        return (0..<c).flatMap{ ci ->
            (0..<r).map{ ri ->
                plaintext.getOrElse(ri * c + ci) { ' ' }
            }.plus(if (ci == c - 1) "" else " ")
        }.joinToString(separator = "")
    }
}