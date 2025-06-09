pub struct SimpleLinkedList<T> {
    len: usize,
    head: Link<T>,
}

type Link<T> = Option<Box<Node<T>>>;

pub struct Node<T> {
    elem: T,
    next: Link<T>,
}

impl<T> SimpleLinkedList<T> {
    pub fn new() -> Self {
        Self {
            len: 0,
            head: None,
        }
    }

    pub fn is_empty(&self) -> bool {
        self.len == 0
    }

    pub fn len(&self) -> usize {
        self.len
    }

    pub fn push(&mut self, elem: T) {
        self.len += 1;
        self.head = Some(Box::new(Node {
            elem,
            next: self.head.take(),
        }));
    }

    pub fn pop(&mut self) -> Option<T> {
        self.len = self.len.saturating_sub(1);
        self.head.take().map(|link| {
            self.head = link.next;
            link.elem
        })
    }

    pub fn peek(&self) -> Option<&T> {
        self.head.as_deref().map(|node| {
            &node.elem
        })
    }

    #[must_use]
    pub fn rev(mut self) -> SimpleLinkedList<T> {
        let mut current = self.head.take();
        let mut prev = None;
        let mut next;
        while let Some(ref mut node) = current {
            next = node.next.take();
            node.next = prev;
            prev = current;
            current = next;
        }
        self.head = prev;
        self
    }
}

impl<T> FromIterator<T> for SimpleLinkedList<T> {
    fn from_iter<I>(iter: I) -> Self
    where
        I: IntoIterator<Item = T>
    {
        iter.into_iter().fold(Self::new(), |mut list, item| {
            list.push(item);
            list
        })
    }
}

impl<T> IntoIterator for SimpleLinkedList<T> {
    type Item = T;
    type IntoIter = SimpleLinkedListIntoIter<T>;
    fn into_iter(self) -> Self::IntoIter {
        SimpleLinkedListIntoIter(self.rev())
    }
}

pub struct SimpleLinkedListIntoIter<T>(SimpleLinkedList<T>);

impl<T> Iterator for SimpleLinkedListIntoIter<T> {
    type Item = T;
    fn next(&mut self) -> Option<Self::Item> {
        self.0.pop()
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

impl<T> From<SimpleLinkedList<T>> for Vec<T> {
    fn from(mut linked_list: SimpleLinkedList<T>) -> Vec<T> {
        linked_list.into_iter().collect()
    }
}
