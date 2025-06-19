def append(list1, list2):
    return [*list1, *list2]

def concat(lists):
    return [elem for list in lists for elem in list]

def filter(function, list):
    return [elem for elem in list if function(elem)]

def length(list):
    return sum(1 for _ in list)

def map(function, list):
    return [function(elem) for elem in list]

def foldl(function, list, initial):
    acc = initial
    [acc := function(acc, elem) for elem in list]
    return acc

def foldr(function, list, initial):
    acc = initial
    [acc := function(acc, elem) for elem in list[::-1]]
    return acc

def reverse(list):
    return list[::-1]