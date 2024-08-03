# Queue with linked list
class Node:
    def __init__(self, item):
        self.item = item
        self.next = None


class Queue:
    def __init__(self):
        self.head = None
        self.tail = None
        self._length = 0

    def enqueue(self, item):
        new_node = Node(item)

        if self.is_empty():
            self.head = self.tail = new_node

        else:
            self.tail.next = new_node
            self.tail = new_node

        self._length += 1

    def deque(self):
        if self.is_empty():
            return None

        temp = self.head

        if self._length == 1:
            self.head = self.tail = None
        else:
            self.head = self.head.next

        self._length -= 1

        return temp.item

    def is_empty(self):
        return self.head is None

    def peek(self):
        return self.head.item or None

    def length(self):
        return self._length


# ===========================
# Queue using Python's List

class Queue_PyList:
    def __init__(self):
        self.queue = []

    def enqueue(self, item):
        self.queue.append(item)

    def deque(self, item):
        if self.is_empty():
            return None

        return self.queue.pop(0)

    def peek(self):
        return self.queue[0]

    def is_empty(self):
        return len(self.queue) == 0

    def length(self):
        return len(self.queue)
