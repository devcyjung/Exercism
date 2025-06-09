package binarysearch

func SearchInts(list []int, key int) int {
	var begin, mid, end int
    end = len(list)
    for begin < end {
        mid = (begin + end - 1) / 2
        if list[mid] == key {
            return mid
        }
        if list[mid] < key {
            begin = mid + 1
            continue
        }
        end = mid
    }
    return -1
}
