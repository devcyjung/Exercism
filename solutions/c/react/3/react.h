#ifndef REACT_H
#define REACT_H

#include <stdlib.h>

struct reactor;
struct cell;

typedef int (*compute1)(int);
typedef int (*compute2)(int, int);

struct reactor *create_reactor(void);
void destroy_reactor(struct reactor *);

struct cell *create_input_cell(struct reactor *, int);
struct cell *create_compute1_cell(struct reactor *, struct cell *, compute1);
struct cell *create_compute2_cell(struct reactor *, struct cell *,
                                  struct cell *, compute2);

int get_cell_value(struct cell *);
void set_cell_value(struct cell *, int);

typedef void (*callback)(void *, int);
typedef size_t callback_id;

callback_id add_callback(struct cell *, void *, callback);
void remove_callback(struct cell *, callback_id);

#endif