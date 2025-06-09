#include "linked_list.h"
#include <stdlib.h>
#include <stddef.h>

struct list_node {
    struct list_node *prev, *next;
    ll_data_t data;
};

struct list {
    struct list_node *first, *last;
    size_t length;
};

__attribute__((warn_unused_result))
static struct list_node *node_create(
    const ll_data_t item_data,
    struct list_node *const prev,
    struct list_node *const next
) {
    struct list_node *node = malloc(sizeof(struct list_node));
    if (!node) {
        return NULL;
    }
    node->data = item_data;
    node->prev = prev;
    node->next = next;
    if (prev) {
        prev->next = node;
    }
    if (next) {
        next->prev = node;
    }
    return node;
}

static void node_delete(struct list_node *const node) {
    if (!node) {
        return;
    }
    if (node->prev) {
        node->prev->next = node->next;
    }
    if (node->next) {
        node->next->prev = node->prev;
    }
    free(node);
}

__attribute__((warn_unused_result))
struct list *list_create() {
    struct list *list = malloc(sizeof(struct list));
    list->length = 0;
    list->first = NULL;
    list->last = NULL;
    return list;
}

size_t list_count(const struct list *const list) {
    return list->length;
}

void list_push(struct list *const list, const ll_data_t item_data) {
    list->last = node_create(item_data, list->last, NULL);
    if (list->length == 0) {
        list->first = list->last;
    }
    ++(list->length);
}

ll_data_t list_pop(struct list *const list) {
    if (list->length != 0) {
        --(list->length);
    }
    ll_data_t result = list->last->data;
    struct list_node *new_last = list->last->prev;
    node_delete(list->last);
    list->last = new_last;
    if (list->length == 0) {
        list->first = NULL;
    }
    return result;
}

void list_unshift(struct list *const list, const ll_data_t item_data) {
    list->first = node_create(item_data, NULL, list->first);
    if (list->length == 0) {
        list->last = list->first;
    }
    ++(list->length);
}

ll_data_t list_shift(struct list *const list) {
    if (list->length != 0) {
        --(list->length);
    }
    ll_data_t result = list->first->data;
    struct list_node *new_first = list->first->next;
    node_delete(list->first);
    list->first = new_first;
    if (list->length == 0) {
        list->last = NULL;
    }
    return result;
}

void list_delete(struct list *const list, const ll_data_t data) {
    struct list_node *current = list->first;
    while (current) {
        if (current->data == data) {
            if (current == list->first) {
                list->first = list->first->next;
            }
            if (current == list->last) {
                list->last = list->last->prev;
            }
            node_delete(current);
            --(list->length);
            break;
        }
        current = current->next;
    }
}

void list_destroy(struct list *const list) {
    struct list_node *current = list->first;
    struct list_node *next = NULL;
    while (current) {
        next = current->next;
        node_delete(current);
        current = next;
    }
    free(list);
}