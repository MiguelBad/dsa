class Node:
    def __init__(self, item):
        self.item = item
        self.priority = item[0]
        self.next = None


class PriorityQueueInsertion:
    def __init__(self):
        self.head = self.tail = None
        self._length = 0

    def add(self, item):
        new_node = Node(item)

        if self.is_empty():
            self.head = self.tail = new_node
        else:
            current_node = previous_node = self.head
            inserted = False

            while current_node:
                if current_node.priority > new_node.priority:
                    new_node.next = current_node

                    if current_node == self.head:
                        self.head = new_node
                    else:
                        previous_node.next = new_node

                    inserted = True
                    break

                previous_node = current_node
                current_node = current_node.next

            if not inserted:
                self.tail.next = new_node
                self.tail = new_node

        self._length += 1

    def remove_min(self):
        if self.is_empty():
            return "Nothing to remove"

        item = self.head.item
        self.head = self.head.next
        self._length -= 1
        return item

    def min(self):
        if self.is_empty():
            return "Nothing in queue"

        return self.head.item

    def is_empty(self):
        return self._length == 0

    def length(self):
        return self._length


# ============= TESTING
def print_test(priority_queue):

    if priority_queue.head:
        print('Head:', priority_queue.head.item)
    else:
        print('Head: None')

    if priority_queue.tail:
        print('Tail:', priority_queue.tail.item)
    else:
        print('Tail: None')

    print('Is empty:', priority_queue.is_empty())
    print('Min:', priority_queue.min())
    print('Length:', priority_queue.length())
    print('\n\n =======\n\n')


def test():
    priority_queue = PriorityQueueInsertion()
    print_test(priority_queue)
    to_add = [(4, "Noob"),
              (1, "Heh"),
              (2, "Foo"),
              (5, "Foo")]

    for i in to_add:
        priority_queue.add(i)
        print("Added: ", i)
        print_test(priority_queue)

    for i in range(4):
        remove = priority_queue.remove_min()
        print('Removed: ', remove)
        print_test(priority_queue)


test()
