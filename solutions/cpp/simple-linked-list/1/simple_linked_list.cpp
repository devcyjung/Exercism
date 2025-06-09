#include "simple_linked_list.h"

#include <stdexcept>
#include <string>

namespace simple_linked_list {

std::size_t List::size() const {
    return current_size;
}

void List::push(int entry) {
    ++current_size;
    auto new_node = new Element(entry);
    if (!head) {
        head = new_node;
    } else {
        auto cur = head;
        while (cur->next) {
            cur = cur->next;
        }
        cur->next = new_node;
    }
}

int List::pop() {
    if (current_size == 0) {
        throw std::runtime_error("pop on empty list");
    }
    --current_size;
    if (!head) {
        throw std::runtime_error("head is nullptr when size is " + std::to_string(current_size));
    }
    if (!(head->next)) {
        auto r = head->data;
        delete head;
        head = nullptr;
        return r;
    }
    auto first = head;
    Element* second = nullptr;
    while (first->next) {
        second = first->next;
        if (!(second->next)) {
            auto r = second->data;
            delete second;
            second = nullptr;
            first->next = nullptr;
            return r;
        }
        first = second;
    }
    throw std::runtime_error("unreachable state");
}

void List::reverse() {
    if (current_size <= 1) {
        return;
    }
    if (!head || !(head->next)) {
        throw std::runtime_error("head or head->next is nullptr when size is " + std::to_string(current_size));
    }
    Element* prev = nullptr;
    auto first = head;
    Element* second = nullptr;
    while (first->next) {
        second = first->next;
        first->next = prev;
        prev = first;
        first = second;
    }
    first->next = prev;
    head = first;
}

List::~List() {
    current_size = 0;
    if (!head) {
        return;
    }
    if (!(head->next)) {
        delete head;
        head = nullptr;
        return;
    }
    auto first = head;
    head = nullptr;
    Element* second = nullptr;
    while (first->next) {
        second = first->next;
        first->next = nullptr;
        delete first;
        first = second;
    }
    delete first;
    first = nullptr;
}

}  // namespace simple_linked_list
