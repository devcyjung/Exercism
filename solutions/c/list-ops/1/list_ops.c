#include "list_ops.h"
#include <stdlib.h>

list_t *new_list(size_t length, list_element_t elements[]) {
    list_t *list = malloc(sizeof(list_t) + sizeof(list_element_t) * length);
    if (!list) {
        return NULL;
    }
    list->length = length;
    for (size_t i = 0; i < length; ++i) {
        list->elements[i] = elements[i];
    }
    return list;
}

list_t *append_list(list_t *list1, list_t *list2) {
    size_t new_length = list1->length + list2->length;
    list_t *list = malloc(sizeof(list_t) + sizeof(list_element_t) * new_length);
    if (!list) {
        return NULL;
    }
    list->length = new_length;
    for (size_t i = 0; i < list1->length; ++i) {
        list->elements[i] = list1->elements[i];
    }
    for (size_t i = 0; i < list2->length; ++i) {
        list->elements[i + list1->length] = list2->elements[i];
    }
    return list;
}

list_t *filter_list(list_t *list, bool (*filter)(list_element_t)) {
    size_t result_length = 0;
    size_t indices[list->length];
    for (size_t i = 0; i < list->length; ++i) {
        if (filter(list->elements[i])) {
            indices[result_length++] = i;
        }
    }
    list_t *result = malloc(sizeof(list_t) + sizeof(list_element_t) * result_length);
    if (!result) {
        return NULL;
    }
    result->length = result_length;
    for (size_t i = 0; i < result->length; ++i) {
        result->elements[i] = list->elements[indices[i]];
    }
    return result;
}

size_t length_list(list_t *list) {
    return list->length;
}

list_t *map_list(list_t *list, list_element_t (*map)(list_element_t)) {
    list_t *result = malloc(sizeof(list_t) + sizeof(list_element_t) * list->length);
    if (!result) {
        return NULL;
    }
    result->length = list->length;
    for (size_t i = 0; i < result->length; ++i) {
        result->elements[i] = map(list->elements[i]);
    }
    return result;
}

list_element_t foldl_list(
    list_t *list,
    list_element_t initial,
    list_element_t (*foldl)(list_element_t element, list_element_t accumulator)
) {
    list_element_t result = initial;
    for (size_t i = 0; i < list->length; ++i) {
        result = foldl(list->elements[i], result);
    }
    return result;
}

list_element_t foldr_list(
    list_t *list,
    list_element_t initial,
    list_element_t (*foldr)(list_element_t element, list_element_t accumulator)
) {
    list_element_t result = initial;
    for (size_t i = list->length; i != 0; --i) {
        result = foldr(list->elements[i - 1], result);
    }
    return result;
}

list_t *reverse_list(list_t *list) {
    list_t *result = malloc(sizeof(list_t) + sizeof(list_element_t) * list->length);
    if (!result) {
        return NULL;
    }
    result->length = list->length;
    for (size_t i = 0; i < result->length; ++i) {
        result->elements[i] = list->elements[result->length - 1 - i];
    }
    return result;
}

void delete_list(list_t *list) {
    free(list);
}