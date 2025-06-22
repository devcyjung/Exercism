#ifndef REACT_H
#define REACT_H

typedef struct reactor reactor_t;
typedef struct cell cell_t;

typedef int (*compute1)(int);
typedef int (*compute2)(int, int);

reactor_t *create_reactor(void);
void destroy_reactor(reactor_t *);

cell_t *create_input_cell(reactor_t *, int);
cell_t *create_compute1_cell(reactor_t *, const cell_t *, const compute1);
cell_t *create_compute2_cell(reactor_t *, const cell_t *, const cell_t *,
                             const compute2);

int get_cell_value(const cell_t *);
void set_cell_value(cell_t *, int);

typedef void (*callback)(void *, int);
typedef unsigned long long callback_id;

callback_id add_callback(cell_t *, void *, callback);
void remove_callback(cell_t *, callback_id);

#endif