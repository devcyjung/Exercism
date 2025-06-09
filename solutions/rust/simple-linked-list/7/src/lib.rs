use std::mem;

type Link<T> = Option<Box<Node<T>>>;

pub struct Node<T> {
    elem: T,
    next: Link<T>,
}

pub struct SimpleLinkedList<T> {
    len: usize,
    head: Link<T>,
}

impl<T> SimpleLinkedList<T> {
    #[must_use]
    pub const fn new() -> Self {
        Self { len: 0, head: None }
    }

    #[must_use]
    pub const fn is_empty(&self) -> bool {
        self.len == 0
    }

    #[must_use]
    pub const fn len(&self) -> usize {
        self.len
    }

    pub fn push(&mut self, elem: T) {
        self.len += 1;
        self.head = Some(Box::new(Node {
            elem,
            next: self.head.take(),
        }));
    }

    #[must_use]
    pub fn pop(&mut self) -> Option<T> {
        self.head.take().map(|node| {
            let Node { elem, next } = *node;
            self.len -= 1;
            self.head = next;
            elem
        })
    }

    #[must_use]
    pub fn peek(&self) -> Option<&T> {
        self.head.as_deref().map(|node| &node.elem)
    }

    #[must_use]
    pub fn rev(&self) -> Self
    where
        T: Clone,
    {
        let mut list = Self::new();
        let mut current = self.head.as_deref();
        while let Some(node) = current {
            list.push(node.elem.clone());
            current = node.next.as_deref();
        }
        list
    }

    #[must_use]
    pub fn iter(&self) -> SimpleLinkedListIter<'_, T> {
        SimpleLinkedListIter(self.head.as_deref())
    }
}

pub struct SimpleLinkedListIter<'a, T>(Option<&'a Node<T>>);

impl<'a, T> Iterator for SimpleLinkedListIter<'a, T> {
    type Item = &'a T;

    #[must_use]
    fn next(&mut self) -> Option<Self::Item> {
        self.0.map(|node| {
            self.0 = node.next.as_deref();
            &node.elem
        })
    }
}

pub struct SimpleLinkedListIntoIter<T>(Link<T>);

impl<T> Iterator for SimpleLinkedListIntoIter<T> {
    type Item = T;

    #[must_use]
    fn next(&mut self) -> Option<Self::Item> {
        self.0.take().map(|node| {
            let Node { elem, next } = *node;
            self.0 = next;
            elem
        })
    }
}

impl<T> IntoIterator for SimpleLinkedList<T> {
    type Item = T;
    type IntoIter = SimpleLinkedListIntoIter<T>;

    #[must_use]
    fn into_iter(mut self) -> Self::IntoIter {
        SimpleLinkedListIntoIter(mem::take(&mut self.head))
    }
}

impl<T> FromIterator<T> for SimpleLinkedList<T> {
    #[must_use]
    fn from_iter<I>(iter: I) -> Self
    where
        I: IntoIterator<Item = T>,
    {
        iter.into_iter().fold(Self::new(), |mut list, item| {
            list.push(item);
            list
        })
    }
}

impl<T> From<SimpleLinkedList<T>> for Vec<T> {
    #[must_use]
    fn from(list: SimpleLinkedList<T>) -> Vec<T> {
        let mut vec: Vec<_> = list.into_iter().collect();
        vec.reverse();
        vec
    }
}

impl<T> Drop for SimpleLinkedList<T> {
    fn drop(&mut self) {
        let mut current = self.head.take();
        while let Some(mut node) = current {
            current = node.next.take();
        }
    }
}