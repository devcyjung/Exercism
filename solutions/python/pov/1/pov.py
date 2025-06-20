from json import dumps

class Tree:
    def __init__(self, label, children=None):
        self.label = label
        self.children = children if children is not None else []

    def __dict__(self):
        return {self.label: [c.__dict__() for c in sorted(
            self.children, key = lambda ch: ch.label
        )]}

    def __hash__(self):
        return id(self)

    def __str__(self, indent=None):
        return dumps(self.__dict__(), indent=indent)

    def __lt__(self, other):
        return self.label < other.label

    def __eq__(self, other):
        return self.__dict__() == other.__dict__()

    def matches(self, name):
        return self.label == name
    
    def index(self, child):
        for i, ch in enumerate(self.children):
            if ch == child:
                return i
        return -1

    def add(self, child):
        self.children.append(child)

    def remove(self, child):
        cursor = self.index(child)
        if cursor >= 0:
            return self.children.pop(cursor)
        return None

    def swap(self, parent):
        if not parent.remove(self):
            return None
        self.add(parent)
        return parent

    def swap_to_root(self, trail):
        cursor = self
        while cursor in trail:
            parent = trail[cursor]
            cursor = cursor.swap(parent)
        return self

    def from_pov(self, from_node):
        trail = {}
        stack = [self]
        while stack:
            cursor = stack.pop()
            if cursor.matches(from_node):
                return cursor.swap_to_root(trail)
            for child in cursor.children:
                trail[child] = cursor
                stack.append(child)
        raise ValueError('Tree could not be reoriented')

    def ancestry(self, node_name):
        trail = {}
        stack = [self]
        while stack:
            cursor = stack.pop()
            if cursor.matches(node_name):
                ancestry_list = []
                while cursor is not None:
                    ancestry_list.append(cursor.label)
                    cursor = trail.get(cursor)
                return ancestry_list
            for child in cursor.children:
                trail[child] = cursor
                stack.append(child)
        return None

    def path_to(self, from_node, to_node):
        from_src = self.ancestry(from_node)
        from_dst = self.ancestry(to_node)
        if from_src is None:
            raise ValueError('Tree could not be reoriented')
        if from_dst is None:
            raise ValueError('No path found')
        last_removed = None
        while from_src and from_dst and from_src[-1] == from_dst[-1]:
            last_removed = from_src.pop()
            from_dst.pop()
        if last_removed is not None:
            from_src.append(last_removed)
        from_src.extend(from_dst[::-1])
        return from_src