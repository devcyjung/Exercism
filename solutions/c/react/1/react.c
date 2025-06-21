#include "react.h"
#include <assert.h>
#include <string.h>

#define INIT_CAP 1

typedef struct reactor reactor_t;
typedef struct cell cell_t;

struct reactor {
  size_t len, cap;
  cell_t **cells;
};

typedef struct callback_info {
  callback cb;
  void *arg;
} cb_info_t;

typedef struct compute_info {
  compute1 comp1;
  compute2 comp2;
  cell_t *cell1, *cell2;
} comp_info_t;

struct cell {
  int value;
  size_t cell_idx, len, cap;
  reactor_t *reactor;
  comp_info_t compute;
  cb_info_t *cb_infos;
};

reactor_t *create_reactor(void) {
  reactor_t *new = calloc(1, sizeof(reactor_t));
  assert(new != NULL);
  new->cells = calloc(INIT_CAP, sizeof(cell_t *));
  assert(new->cells != NULL);
  new->cap = INIT_CAP;
  return new;
}

static reactor_t *resize_reactor(reactor_t *reactor) {
  if (reactor->cap >= reactor->len + 1)
    return reactor;
  size_t old_cap = reactor->cap;
  reactor->cap <<= 1;
  reactor->cells = realloc(reactor->cells, sizeof(cell_t *) * reactor->cap);
  memset(&reactor->cells[old_cap], 0,
         sizeof(cell_t *) * (reactor->cap - old_cap));
  assert(reactor->cells != NULL);
  return reactor;
}

void destroy_reactor(reactor_t *reactor) {
  assert(reactor != NULL);
  for (size_t cell_id = 0; cell_id < reactor->len; ++cell_id) {
    cell_t *cell = reactor->cells[cell_id];
    if (cell == NULL)
      continue;
    free(cell->cb_infos);
    free(cell);
  }
  free(reactor->cells);
  free(reactor);
}

cell_t *create_input_cell(reactor_t *reactor, int initial_value) {
  assert(reactor != NULL);
  resize_reactor(reactor);
  cell_t *new = calloc(1, sizeof(cell_t));
  assert(new != NULL);
  reactor->cells[reactor->len++] = new;
  new->value = initial_value;
  new->cell_idx = reactor->len - 1;
  new->reactor = reactor;
  return new;
}

static cell_t *create_compute_cell_proto(reactor_t *reactor) {
  assert(reactor != NULL);
  resize_reactor(reactor);
  cell_t *new = calloc(1, sizeof(cell_t));
  assert(new != NULL);
  reactor->cells[reactor->len++] = new;
  new->cell_idx = reactor->len - 1;
  new->cap = INIT_CAP;
  new->reactor = reactor;
  new->cb_infos = calloc(INIT_CAP, sizeof(cb_info_t));
  assert(new->cb_infos != NULL);
  return new;
}

cell_t *create_compute1_cell(reactor_t *reactor, cell_t *cell1,
                             compute1 comp1) {
  cell_t *new = create_compute_cell_proto(reactor);
  new->compute.comp1 = comp1;
  new->compute.cell1 = cell1;
  new->value = comp1(get_cell_value(cell1));
  return new;
}

cell_t *create_compute2_cell(reactor_t *reactor, cell_t *cell1, cell_t *cell2,
                             compute2 comp2) {
  cell_t *new = create_compute_cell_proto(reactor);
  new->compute.comp2 = comp2;
  new->compute.cell1 = cell1;
  new->compute.cell2 = cell2;
  new->value = comp2(get_cell_value(cell1), get_cell_value(cell2));
  return new;
}

int get_cell_value(cell_t *cell) { return cell->value; }

void set_cell_value(cell_t *cell, int new_value) {
  if (cell->value == new_value)
    return;
  cell->value = new_value;
  assert(cell->reactor != NULL);
  for (size_t cell_id = cell->cell_idx + 1; cell_id < cell->reactor->len;
       ++cell_id) {
    cell_t *next_cell = cell->reactor->cells[cell_id];
    if (next_cell->compute.comp1 == NULL && next_cell->compute.comp2 == NULL)
      continue;
    int prev_value = next_cell->value;
    comp_info_t comp = next_cell->compute;
    int new_value;
    if (comp.comp1 != NULL) {
      assert(comp.cell1 != NULL);
      compute1 comp1 = comp.comp1;
      new_value = comp1(comp.cell1->value);
    } else {
      assert(comp.comp2 != NULL && comp.cell1 != NULL && comp.cell2 != NULL);
      compute2 comp2 = comp.comp2;
      new_value = comp2(comp.cell1->value, comp.cell2->value);
    }
    next_cell->value = new_value;
    if (prev_value != new_value && next_cell->len > 0) {
      assert(next_cell->cb_infos != NULL);
      size_t done = 0;
      for (size_t cb_idx = 0; cb_idx < next_cell->cap; ++cb_idx) {
        if (next_cell->cb_infos[cb_idx].cb == NULL)
          continue;
        cb_info_t cb_info = next_cell->cb_infos[cb_idx];
        callback cb = cb_info.cb;
        cb(cb_info.arg, new_value);
        if (++done == next_cell->len)
          break;
      }
    }
  }
}

callback_id add_callback(cell_t *cell, void *arg, callback cb) {
  assert(cell->cb_infos != NULL);
  if (cell->cap <= cell->len) {
    size_t old_cap = cell->cap;
    cell->cap <<= 1;
    cell->cb_infos = realloc(cell->cb_infos, sizeof(cb_info_t) * cell->cap);
    memset(&cell->cb_infos[old_cap], 0,
           sizeof(cb_info_t) * (cell->cap - old_cap));
    assert(cell->cb_infos != NULL);
  }
  size_t cb_idx;
  for (cb_idx = 0; cb_idx < cell->cap; ++cb_idx) {
    if (cell->cb_infos[cb_idx].cb == NULL) {
      break;
    }
  }
  cell->cb_infos[cb_idx].cb = cb;
  cell->cb_infos[cb_idx].arg = arg;
  ++cell->len;
  return cb_idx;
}

void remove_callback(cell_t *cell, callback_id cb_idx) {
  cell->cb_infos[cb_idx].cb = NULL;
  cell->cb_infos[cb_idx].arg = NULL;
  --cell->len;
}