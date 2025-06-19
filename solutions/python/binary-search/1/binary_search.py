def find(search_list, value):
    left, right = 0, len(search_list)
    while left < right:
        mid = (left + right - 1) >> 1
        match search_list[mid] - value:
            case cmp if cmp > 0:
                right = mid
            case cmp if cmp < 0:
                left = mid + 1
            case cmp if cmp == 0:
                return mid
    raise ValueError("value not in array")