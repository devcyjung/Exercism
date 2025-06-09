#ifndef LINKED_LIST_H
#define LINKED_LIST_H

#include <stddef.h>

typedef int ll_data_t;
struct list;

struct list *list_create(void);

size_t list_count(const struct list *const list);

void list_push(struct list *const list, const ll_data_t item_data);

ll_data_t list_pop(struct list *const list);

void list_unshift(struct list *const list, const ll_data_t item_data);

ll_data_t list_shift(struct list *const list);

void list_delete(struct list *const list, const ll_data_t data);

void list_destroy(struct list *const list);

#endif
