import kotlin.math.floor
import kotlin.math.sqrt

object CryptoSquare {

    fun ciphertext(plaintext: String): String {
        val plaintext = plaintext.filter{ it.isLetterOrDigit() }.lowercase()
        val size = plaintext.length
        var r = floor(sqrt(size.toDouble())).toInt()
        var c = r
        var change_r = false
        while (r * c < size) {
            if (change_r) {
                r++
            } else {
                c++
            }
            change_r = !change_r
        }
        return (0..<c)
            .fold(StringBuilder()) { builder, ci ->
                (0..<r)
                    .fold(builder) { builder, ri ->
                        builder.append(plaintext.getOrElse(ri*c+ci) { ' ' })
                        builder
                    }
                    .also {
                        if (ci + 1 != c) {
                            builder.append(' ')   
                        }
                    }
            }
            .toString()
    }

}
