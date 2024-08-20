class MinHeap:
    def __init__(self) -> None:
        self.array: list[int] = []

    def length(self) -> int:
        return len(self.array)

    def insert(self, item: int) -> None:
        self.array.append(item)
        self.up_heap(self.length() - 1)

    def delete(self) -> int | None:
        if self.length() == 0:
            return None

        if self.length() - 1 == 0:
            return self.array.pop()

        item = self.array[0]
        self.array[0] = self.array.pop()
        self.down_heap(0)
        return item

    def up_heap(self, idx: int) -> None:
        if idx < 1:
            return

        parent_idx = self.parent_idx(idx)
        parent_val = self.array[parent_idx]

        if parent_val > self.array[idx]:
            self.array[parent_idx] = self.array[idx]
            self.array[idx] = parent_val
            self.up_heap(parent_idx)

    def down_heap(self, idx: int) -> None:
        left_idx = self.left_idx(idx)
        right_idx = self.right_idx(idx)

        if left_idx >= self.length() - 1 or not self.array[idx]:
            return

        left_val = self.array[left_idx]
        right_val = self.array[right_idx]

        if left_val <= right_val and left_val < self.array[idx]:
            self.array[left_idx] = self.array[idx]
            self.array[idx] = left_val
            self.down_heap(left_idx)
        elif left_val > right_val and right_val < self.array[idx]:
            self.array[right_idx] = self.array[idx]
            self.array[idx] = right_val
            self.down_heap(right_idx)

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
    min_heap = MinHeap()
    items = [9, 7, 4, 1, 5, 4, 2]

    print(min_heap.array)
    for i in items:
        min_heap.insert(i)
        print(min_heap.array)

    for _ in items:
        min_heap.delete()
        print(min_heap.array)


test()
