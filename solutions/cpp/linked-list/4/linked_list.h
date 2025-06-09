#ifndef LINKED_LIST_H_
#define LINKED_LIST_H_

#include <memory>
#include <stdexcept>

namespace linked_list {

template<typename T>
struct Node: std::enable_shared_from_this<Node<T>> {
    T value;
    std::weak_ptr<Node<T>> prev;
    std::shared_ptr<Node<T>> next;
    explicit Node(T initial): value(initial) {}
    Node(const Node&) = delete;
    Node& operator=(const Node&) = delete;
    Node(Node&&) = default;
    Node& operator=(Node&&) = default;
    virtual ~Node() = default;
};

template<typename T>
class List {
private:
    size_t length;
    std::shared_ptr<Node<T>> head;
    std::shared_ptr<Node<T>> tail;
public:
    explicit List(): length(0) {}
    List(const List&) = delete;
    List& operator=(const List&) = delete;
    List(List&&) = delete;
    List& operator=(List&&) = delete;
    virtual ~List() = default;
    
    [[nodiscard]]
    inline constexpr
    size_t count() const noexcept {
        return length;
    }
    
    inline constexpr
    void push(T new_value) noexcept {      
        auto old_tail = tail;
        tail = std::make_shared<Node<T>>(new_value);
        tail->prev = old_tail;
        if (old_tail) {
            old_tail->next = tail;
        } else {
            head = tail;
        }
        ++length;
    }
    
    inline constexpr
    void unshift(T new_value) noexcept {
        auto old_head = head;
        head = std::make_shared<Node<T>>(new_value);
        head->next = old_head;
        if (old_head) {
            old_head->prev = head;   
        } else {
            tail = head;
        }
        ++length;
    }
    
    inline constexpr
    T pop() {
        if (!tail) {
            throw std::range_error("List is empty");
        }
        auto popped = tail;
        tail = tail->prev.lock();
        if (tail) {
            tail->next = nullptr;
        } else {
            head = nullptr;
        }
        popped->prev.reset();
        --length;
        return popped->value;
    }
    
    inline constexpr
    T shift() {
        if (!head) {
            throw std::range_error("List is empty");
        }
        auto popped = head;
        head = head->next;
        if (head) {
            head->prev.reset();
        } else {
            tail = nullptr;
        }
        popped->next = nullptr;
        --length;
        return popped->value;
    }
    
    inline constexpr
    bool erase(T value) noexcept {
        auto node = head;
        while (node) {
            if (node->value != value) {
                node = node->next;
                continue;
            }
            auto prev = node->prev.lock();
            auto next = node->next;
            node->prev.reset();
            if (prev) {
                prev->next = next;
            } else {
                head = next;
            }
            if (next) {
                next->prev = prev;
            } else {
                tail = prev;
            }
            --length;
            return true;
        }
        return false;
    }
};

}  // namespace linked_list

#endif  // LINKED_LIST_H_