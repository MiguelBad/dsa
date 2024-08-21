class MaxHeap:
    # initialise empty python array
    def __init__(self) -> None:
        self.array: list[int] = []

    # since using python's array, we can use builtin length function
    def length(self) -> int:
        return len(self.array)

    # insert method
    def insert(self, item: int) -> None:
        self.array.append(item)
        self.up_heap(self.length() - 1)  # perform up heap

    # delete method
    def delete(self) -> int | None:
        # return none if array is empty
        if self.length() == 0:
            return None

        # if only 1 item in array return element and empty the array
        if self.length() - 1 == 0:
            return self.array.pop()

        # return the first item in array and replace it with the last item
        item = self.array[0]
        self.array[0] = self.array.pop()
        self.down_heap(0)  # perform down heap
        return item

    # up heap method
    def up_heap(self, idx: int) -> None:
        # base case
        if idx < 1:
            return

        parent_idx = self.parent_idx(idx)
        parent_val = self.array[parent_idx]

        # swap if parent value is lower than the current value
        if parent_val < self.array[idx]:
            self.array[parent_idx] = self.array[idx]
            self.array[idx] = parent_val
            self.up_heap(parent_idx)  # recurse

    # down heap method
    def down_heap(self, idx: int) -> None:
        left_idx = self.left_idx(idx)
        right_idx = self.right_idx(idx)
        largest = idx

        # check if left index exist and if its value is larger than current value
        if left_idx < self.length() and self.array[left_idx] > self.array[largest]:
            largest = left_idx

        # check if right index exist and if its value is larger than current largest
        if right_idx < self.length() and self.array[right_idx] > self.array[largest]:
            largest = right_idx

        # swap the current value to the largest value found
        if largest != idx:
            self.array[largest], self.array[idx] = self.array[idx], self.array[largest]
            self.down_heap(largest)  # recurse

    # formula to get parent index
    def parent_idx(self, idx: int) -> int:
        return (idx - 1) // 2

    # formula to get the left child index
    def left_idx(self, idx: int) -> int:
        return idx * 2 + 1

    def right_idx(self, idx: int) -> int:
        return idx * 2 + 2


# test
def test():
    min_heap = MaxHeap()
    items = [2, 7, 4, 13, 9, 10, 5, 6, 13]

    print(min_heap.array)
    for i in items:
        min_heap.insert(i)
        print(min_heap.array)

    for _ in items:
        min_heap.delete()
        print(min_heap.array)


test()
