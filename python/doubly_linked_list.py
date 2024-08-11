class Node:
    def __init__(self, item):
        self.item = item
        self.prev = self.next = None


class DoubleLinkedList:
    def __init__(self):
        self.head = self.tail = None
        self._length = 0

    # add new node at the front of the list (head)
    def prepend(self, item):
        new_node = Node(item)

        if self.is_empty():
            self.head = self.tail = new_node
        else:
            new_node.next = self.head
            self.head.prev = new_node
            self.head = new_node

        self._length += 1

    # adds new node at the end of the list (tail)
    def append(self, item):
        new_node = Node(item)

        if self.is_empty():
            self.prepend(item)
        else:
            new_node.prev = self.tail
            self.tail.next = new_node
            self.tail = new_node

        self._length += 1

    # insert new node at given position
    def insert_at(self, item, index):
        self.index_error_check(index)

        new_node = Node(item)

        if index == 0:
            self.prepend(item)
        elif index == self._length:
            self.append(item)
        else:
            curr = self.head
            for i in range(index):
                curr = curr.next

            new_node.next = curr
            new_node.prev = curr.prev
            curr.prev = new_node
            new_node.prev.next = new_node

        self._length += 1

    # removes item on the list
    def remove(self, item):
        curr = self.head

        while curr:
            if curr.item == item:
                if curr.prev:
                    curr.prev.next = curr.next
                else:
                    self.head = curr.next

                if curr.next:
                    curr.next.prev = curr.prev
                else:
                    self.tail = curr.prev

                self._length -= 1
                return curr.item

            curr = curr.next

        return None

    # removes item with the given index
    def remove_at(self, index):
        self.index_error_check(index)

        curr = self.head

        if index == 0:
            if self.length() == 1:
                self.head = self.tail = None
            else:
                self.head = curr.next
                self.head.prev = None

        elif index == self._length - 1:
            self.tail = curr.prev
            self.tail.next = None

        else:
            for _ in range(index):
                curr = curr.next

            curr.next.prev = curr.prev
            curr.prev.next = curr.next

        self._length -= 1
        return curr

    # get item at specific index
    def get(self, index):
        self.index_error_check(index)

        curr = self.head
        for _ in range(index):
            curr = curr.next

        return curr

    def length(self):
        return self._length

    def is_empty(self):
        return self._length == 0

    def index_error_check(self, index):
        if index > self._length or index < 0:
            raise IndexError('Index out of range')
