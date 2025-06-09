object RunLengthEncoding {
    fun encode(input: String): String =
        input.fold(Triple(StringBuilder(), ' ', 0)) { (builder, run, length), current ->
            when {
                length == 0 || current == run -> Triple(builder, current, length + 1)
                length == 1 -> {
                    builder.append(run)
                    Triple(builder, current, 1)
                }
                else -> {
                    builder.append(length.toString() + run)
                    Triple(builder, current, 1)
                }
            }
        }.also { (builder, run, length) ->
            if (length > 1)
                builder.append(length.toString() + run)
            else if (length == 1)
                builder.append(run)
        }.first.toString()

    fun decode(input: String): String =
        input.fold(StringBuilder() to 0) { (builder, length), current ->
            when {
                current in '0'..'9' -> builder to (current - '0' + length * 10)
                length == 0 -> {
                    builder.append(current.toString())
                    builder to 0
                }
                else -> {
                    builder.append(current.toString().repeat(length))
                    builder to 0
                }
            }
        }.first.toString()
}