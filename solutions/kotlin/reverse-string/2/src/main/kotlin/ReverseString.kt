fun reverse(input: String): String =
    buildString {
        for (index in input.lastIndex downTo 0) {
            append(input[index])
        }
    }