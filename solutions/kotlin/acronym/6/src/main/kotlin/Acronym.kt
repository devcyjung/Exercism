object Acronym

private fun Char.isDelimiter() = this.isWhitespace() || this == '-'

fun Acronym.generate(phrase: String) = phrase
    .filter { it.isLetter() || it.isDelimiter() }
    .let {
        it.filterIndexed { index, char ->
            char.isLetter() && (index == 0 || it[index - 1].isDelimiter())
        }
    }
    .uppercase()
