object BinarySearch {
    fun search(list: List<Int>, item: Int): Int {
        var left = 0
        var right = list.size
        while (left < right) {
            val mid = (left + right - 1) / 2
            if (list[mid] == item) {
                return mid
            } else if (list[mid] > item) {
                right = mid
            } else {
                left = mid + 1
            }
        }
        throw NoSuchElementException("$list does not have $item")
    }
}