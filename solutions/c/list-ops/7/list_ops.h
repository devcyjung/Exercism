#ifndef LIST_OPS_H
#define LIST_OPS_H

#include <stddef.h>
#include <stdbool.h>
#include <stdlib.h>

typedef int list_element_t;

typedef struct {
   size_t length;
   list_element_t elements[];
} list_t;

list_t *new_list(const size_t length, const list_element_t *const elements);

list_t *append_list(const list_t *const list1, const list_t *const list2);

list_t *filter_list(const list_t *const list, bool (*const filter)(const list_element_t));

size_t length_list(const list_t *const list);

list_t *map_list(const list_t *const list, list_element_t (*const map)(const list_element_t));

list_element_t foldl_list(const list_t *const list, const list_element_t initial,
                          list_element_t (*const foldl)(const list_element_t, const list_element_t));

list_element_t foldr_list(const list_t *const list, const list_element_t initial,
                          list_element_t (*const foldr)(const list_element_t, const list_element_t));

list_t *reverse_list(const list_t *const list);

void delete_list(list_t *const list);

#endif
