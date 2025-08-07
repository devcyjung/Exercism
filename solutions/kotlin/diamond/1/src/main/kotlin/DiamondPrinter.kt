class DiamondPrinter

internal fun DiamondPrinter.printToList(char: Char): List<String> {
    require(char in 'A'..'Z') { "Invalid character: $char" }
    val sideLength = ((char - 'A') shl 1) + 1
    val buffer = CharArray(sideLength) { ' ' }
    val diamond = Array(sideLength) { "" }
    var top = 0
    var bottom = sideLength - 1
    var left = sideLength shr 1
    var right = left
    for (ch in 'A'..char) {
        buffer[left] = ch
        buffer[right] = ch
        diamond[top] = String(buffer)
        diamond[bottom--] = diamond[top++]
        buffer[left--] = ' '
        buffer[right++] = ' '
    }
    return diamond.toList()
}
