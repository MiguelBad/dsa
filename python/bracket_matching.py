''' Implementation of  Stack '''


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
            raise IndexError('Array is Empty')

        item = self.head.item
        self.head = self.head.next
        self._length -= 1

        return item

    def is_empty(self):
        return self._length == 0

    def peek(self):
        if self.is_empty():
            raise IndexError('Array is Empty')

        return self.head.item


def main(input):
    '''
        Checks if the input string of brackets is balanced.

        Parameters:
        - input (str): A string containing brackets. ('()', '[]', '{}')
        - works even if input not a bracket

        Returns:
        - bool: True if brackets are balanced

        note to self:
        - using sets is more performant than lists on membership test
    '''
    pairs = {'(': ')', '[': ']', '{': '}'}

    left = set(pairs.keys())
    right = set(pairs.values())
    left_so_far = Stack()

    for char in input:
        if char in left:
            left_so_far.push(char)

        elif char in right:
            if left_so_far.is_empty():
                return False

            if pairs[left_so_far.peek()] == char:
                left_so_far.pop()
            else:
                return False

        else:
            continue

    return left_so_far.is_empty()


test = ['( )(( )){([( )])}',
        '((( )(( )){([( )])}',
        ')(( )){([( )])}',
        '({[ ])}',
        '(',]

for i in test:
    print(main(i))
