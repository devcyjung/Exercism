#include "list_ops.h"
#include <stdlib.h>
#include <string.h>

__attribute__ ((warn_unused_result))
static list_t *list_alloc(const size_t length) {
    list_t *list = malloc(sizeof(list_t) + sizeof(list_element_t) * length);
    list->length = length;
    return list;
}

__attribute__ ((warn_unused_result))
static list_t *list_realloc(list_t *list, const size_t length) {
    list = realloc(list, sizeof(list_t) + sizeof(list_element_t) * length);
    list->length = length;
    return list;
}

static void list_cpy(
    list_t *const dst,
    const list_element_t src[],
    const size_t begin_dst_index,
    const size_t length
) {
    memcpy(dst->elements + begin_dst_index, src, sizeof(list_element_t) * length);
}

list_t *new_list(const size_t length, const list_element_t *const elements) {
    list_t *list = list_alloc(length);
    if (!list) {
        return NULL;
    }
    list_cpy(list, elements, 0, length);
    return list;
}

list_t *append_list(const list_t *const list1, const list_t *const list2) {
    size_t new_length = list1->length + list2->length;
    list_t *list = list_alloc(new_length);
    if (!list) {
        return NULL;
    }
    list_cpy(list, list1->elements, 0, list1->length);
    list_cpy(list, list2->elements, list1->length, list2->length);
    return list;
}

list_t *filter_list(const list_t *const list, bool (*const filter)(const list_element_t)) {
    size_t result_length = 0;
    list_t *result = list_alloc(list->length);
    if (!result) {
        return NULL;
    }
    for (size_t i = 0; i < list->length; ++i) {
        if (filter(list->elements[i])) {
            result->elements[result_length] = list->elements[i];
            ++result_length;
        }
    }
    return list_realloc(result, result_length);
}

size_t length_list(const list_t *const list) {
    return list->length;
}

list_t *map_list(const list_t *const list, list_element_t (*const map)(const list_element_t)) {
    list_t *result = list_alloc(list->length);
    if (!result) {
        return NULL;
    }
    for (size_t i = 0; i < list->length; ++i) {
        result->elements[i] = map(list->elements[i]);
    }
    return result;
}

list_element_t foldl_list(
    const list_t *const list,
    const list_element_t initial,
    list_element_t (*const foldl)(const list_element_t element, const list_element_t accumulator)
) {
    list_element_t result = initial;
    for (size_t i = 0; i < list->length; ++i) {
        result = foldl(list->elements[i], result);
    }
    return result;
}

list_element_t foldr_list(
    const list_t *const list,
    const list_element_t initial,
    list_element_t (*const foldr)(const list_element_t element, const list_element_t accumulator)
) {
    list_element_t result = initial;
    for (size_t i = list->length; i != 0; --i) {
        result = foldr(list->elements[i - 1], result);
    }
    return result;
}

list_t *reverse_list(const list_t *const list) {
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

void delete_list(list_t *const list) {
    free(list);
}