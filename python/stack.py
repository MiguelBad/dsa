class Node:
    def __init__(self, item):
        self.item = item
        self.next = None


class Stack:
    def __init__(self):
        self.head = None
        self._length = 0

    def push(self, item):
        new_node = Node(item)
        new_node.next = self.head
        self.head = new_node
        self._length += 1

    def pop(self):
        if self.is_empty():
            return None

        head = self.head.item
        self.head = self.head.next
        self._length -= 1

        return head

    def length(self):
        return self._length

    def is_empty(self):
        return self._length == 0

    def peek(self):
        if self.is_empty():
            return None

        return self.head.item


def push_test():
    print('\nPush Test\n')
    new_stack = Stack()

    print('Length: ', new_stack.length())
    print('Peek: ', new_stack.peek())
    print('Empty?: ', new_stack.is_empty())

    new_stack.push('bar')

    print()
    print('Length: ', new_stack.length())
    print('Peek: ', new_stack.peek())
    print('Empty?: ', new_stack.is_empty())

    new_stack.push('foo')

    print()
    print('Length: ', new_stack.length())
    print('Peek: ', new_stack.peek())
    print('Empty?: ', new_stack.is_empty())


def pop_test():
    print('\nPop Test\n')
    new_stack = Stack()

    print('Length: ', new_stack.length())
    print('Peek: ', new_stack.peek())
    print('Empty?: ', new_stack.is_empty())

    new_stack.push('bar')
    new_stack.push('foo')
    new_stack.push('baz')

    print()
    print('Length: ', new_stack.length())
    print('Peek: ', new_stack.peek())
    print('Empty?: ', new_stack.is_empty())

    popped = new_stack.pop()

    print()
    print(popped)
    print('Length: ', new_stack.length())
    print('Peek: ', new_stack.peek())
    print('Empty?: ', new_stack.is_empty())


if __name__ == "__main__":
    push_test()
    pop_test()
