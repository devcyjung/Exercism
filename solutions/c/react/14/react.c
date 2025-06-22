#include "react.h"
#include <assert.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

#define INIT_CAP 1

struct reactor {
  size_t len, cap;
  cell_t **cells;
};

typedef struct callback_info {
  callback cb;
  void *arg;
  callback_id cb_id;
} cb_info_t;

typedef struct compute_info {
  compute1 comp1;
  compute2 comp2;
  const cell_t *cell1, *cell2;
} comp_info_t;

struct cell {
  int value;
  size_t cell_idx, len, cap;
  reactor_t *reactor;
  comp_info_t compute;
  cb_info_t *cb_infos;
};

reactor_t *create_reactor(void) {
  reactor_t *new = malloc(sizeof(*new));
  assert(new);
  *new =
      (reactor_t){.cells = calloc(INIT_CAP, sizeof(*new->cells)), .cap = INIT_CAP};
  assert(new->cells);
  return new;
}

static reactor_t *resize_reactor(reactor_t *reactor) {
  if (reactor->cap >= reactor->len + 1)
    return reactor;
  size_t old_cap = reactor->cap;
  reactor->cap <<= 1;
  reactor->cells =
      realloc(reactor->cells, sizeof(*reactor->cells) * reactor->cap);
  memset(&reactor->cells[old_cap], 0,
         sizeof(*reactor->cells) * (reactor->cap - old_cap));
  assert(reactor->cells);
  return reactor;
}

void destroy_reactor(reactor_t *reactor) {
  assert(reactor);
  for (size_t cell_id = 0; cell_id < reactor->len; ++cell_id) {
    cell_t *cell = reactor->cells[cell_id];
    if (!cell)
      continue;
    free(cell->cb_infos);
    free(cell);
  }
  free(reactor->cells);
  free(reactor);
}

static cell_t *create_cell_proto(reactor_t *reactor, bool is_input) {
  assert(reactor);
  resize_reactor(reactor);
  cell_t *new = malloc(sizeof(*new));
  assert(new);
  reactor->cells[reactor->len++] = new;
  *new = (cell_t){.cell_idx = reactor->len - 1,
                  .cap = is_input ? 0 : INIT_CAP,
                  .reactor = reactor,
                  .cb_infos = is_input ? NULL
                      : calloc(INIT_CAP, sizeof(*new->cb_infos))};
  if (!is_input)
    assert(new->cb_infos);
  return new;
}

cell_t *create_input_cell(reactor_t *reactor, int initial_value) {
  cell_t *new = create_cell_proto(reactor, true);
  new->value = initial_value;
  return new;
}

cell_t *create_compute1_cell(reactor_t *reactor, const cell_t *cell1,
                             const compute1 comp1) {
  cell_t *new = create_cell_proto(reactor, false);
  new->compute.comp1 = comp1;
  new->compute.cell1 = cell1;
  new->value = comp1(get_cell_value(cell1));
  return new;
}

cell_t *create_compute2_cell(reactor_t *reactor, const cell_t *cell1,
                             const cell_t *cell2, const compute2 comp2) {
  cell_t *new = create_cell_proto(reactor, false);
  new->compute.comp2 = comp2;
  new->compute.cell1 = cell1;
  new->compute.cell2 = cell2;
  new->value = comp2(get_cell_value(cell1), get_cell_value(cell2));
  return new;
}

int get_cell_value(const cell_t *cell) { return cell->value; }

void set_cell_value(cell_t *cell, int new_value) {
  if (cell->value == new_value)
    return;
  cell->value = new_value;
  assert(cell->reactor);
  for (size_t cell_idx = cell->cell_idx + 1; cell_idx < cell->reactor->len;
       ++cell_idx) {
    cell_t *next_cell = cell->reactor->cells[cell_idx];
    if (!next_cell->compute.comp1 && !next_cell->compute.comp2)
      continue;
    int prev_value = next_cell->value;
    comp_info_t comp = next_cell->compute;
    int new_value;
    if (comp.comp1) {
      assert(comp.cell1);
      compute1 comp1 = comp.comp1;
      new_value = comp1(comp.cell1->value);
    } else {
      assert(comp.comp2 && comp.cell1 && comp.cell2);
      compute2 comp2 = comp.comp2;
      new_value = comp2(comp.cell1->value, comp.cell2->value);
    }
    next_cell->value = new_value;
    if (prev_value != new_value && next_cell->len > 0) {
      assert(next_cell->cb_infos);
      size_t done = 0;
      for (size_t cb_idx = 0; cb_idx < next_cell->cap; ++cb_idx) {
        if (!next_cell->cb_infos[cb_idx].cb)
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
  assert(cell->cb_infos);
  if (cell->cap <= cell->len) {
    size_t old_cap = cell->cap;
    cell->cap <<= 1;
    cell->cb_infos = realloc(cell->cb_infos, sizeof(*cell->cb_infos) * cell->cap);
    memset(&cell->cb_infos[old_cap], 0,
           sizeof(*cell->cb_infos) * (cell->cap - old_cap));
    assert(cell->cb_infos);
  }
  size_t cb_idx;
  for (cb_idx = 0; cb_idx < cell->cap; ++cb_idx) {
    if (!cell->cb_infos[cb_idx].cb) {
      break;
    }
  }
  callback_id cb_id = (callback_id) cb ^ (callback_id) arg;
  cell->cb_infos[cb_idx] =
      (cb_info_t){.cb = cb, .arg = arg, .cb_id = cb_id};
  ++cell->len;
  return cb_id;
}

void remove_callback(cell_t *cell, callback_id cb_id) {
  for (size_t cb_idx = 0; cb_idx < cell->cap; ++cb_idx) {
    cb_info_t cb_info = cell->cb_infos[cb_idx];
    if (cb_info.cb_id == cb_id) {
      memset(&cell->cb_infos[cb_idx], 0, sizeof(*cell->cb_infos));
      --cell->len;
      return;
    }
  }
}