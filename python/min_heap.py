class MinHeap:
    def __init__(self) -> None:
        self.length: int = 0
        self.array: list[int] = []

    def insert(self, item: int) -> None:
        self.array.append(item)
        self.length += 1
        self.upHeap(self.length - 1)

    def delete(self) -> int | None:
        if self.length == 0:
            return None

        out = self.array[self.length]
        self.length -= 1
        self.array[0] = self.array[self.length]
        self.array.pop()
        self.downHeap(self.length)
        return out

    def upHeap(self, idx: int) -> None:
        if self.length < 2:
            return

        parentIdx = self.parentIdx(idx)
        parentVal = self.array[parentIdx]
        currentVal = self.array[idx]

        if parentVal > currentVal:
            self.array[parentIdx] = currentVal
            self.array[idx] = parentVal
            self.upHeap(parentIdx)

    def downHeap(self, idx: int) -> None:
        leftIdx = self.leftIdx(idx)
        rightIdx = self.rightIdx(idx)

        if self.length < leftIdx:
            return

        leftVal = self.array[leftIdx]
        rightVal = self.array[rightIdx]
        currVal = self.array[idx]

        if leftIdx > rightIdx and rightIdx < currVal:
            self.array[rightIdx] = currVal
            self.array[currVal] = rightVal
            self.downHeap(rightIdx)
        elif leftIdx < rightIdx and leftIdx < currVal:
            self.array[leftIdx] = currVal
            self.array[currVal] = leftVal
            self.downHeap(leftIdx)

    def parentIdx(self, idx: int) -> int:
        return (idx - 1) // 2

    def leftIdx(self, idx: int) -> int:
        return idx * 2 + 1

    def rightIdx(self, idx: int) -> int:
        return idx * 2 + 2


test = MinHeap()

test.insert(9)
print(test.array)
test.insert(7)
print(test.array)
test.insert(4)
print(test.array)
test.insert(1)
print(test.array)
test.insert(6)
print(test.array)
test.insert(4)
print(test.array)
test.insert(3)
print(test.array)
