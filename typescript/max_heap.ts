class MaxHeap {
    arrayList: number[];
    length: number;

    constructor() {
        this.arrayList = [];
        this.length = 0;
    }

    insert(item: number): void {
        this.arrayList[this.length] = item;
        this.length++;
        this.heapifyUp(this.length - 1);
    }

    delete(): number {
        const itemIdx: number = this.length - 1;
        const item: number = this.arrayList[itemIdx];
        this.heapifyDown(itemIdx, item);

        this.arrayList.splice(itemIdx, 1);
        return this.arrayList[itemIdx];
    }

    private heapifyDown(idx: number, item: number): void {
        const parentIdx = this.parent(idx);
        const parentVal = this.arrayList[parentIdx];
        const leftChildIdx = this.leftChild(idx);
        const leftChildVal = this.arrayList[leftChildIdx];
        const rightChildIdx = this.rightChild(idx);
        const rightChildVal = this.arrayList[rightChildIdx];

        if (!leftChildVal) {
            this.arrayList[idx] = leftChildVal;
        }

        if (!rightChildVal) {
            this.arrayList[idx] = rightChildVal;
        }

        if (leftChildVal && rightChildVal) {
            if (rightChildVal > leftChildVal) {
                this.arrayList[idx] = rightChildVal;
            } else {
                this.arrayList[idx] = leftChildVal;
            }
        }
    }

    private heapifyUp(idx: number): void {
        if (idx === 0) {
            return;
        }

        const parentIdx = this.parent(idx);
        const parentVal = this.arrayList[parentIdx];
        const currentVal = this.arrayList[idx];

        if (currentVal > parentVal) {
            this.arrayList[parentIdx] = currentVal;
            this.arrayList[idx] = parentVal;
            this.heapifyUp(parentIdx);
        }
    }

    private parent(idx: number): number {
        return Math.floor((idx - 1) / 2);
    }

    private leftChild(idx: number): number {
        return idx * 2 + 1;
    }

    private rightChild(idx: number): number {
        return idx * 2 + 2;
    }
}

const heap = new MaxHeap();

console.log(heap.arrayList)
console.log(heap.insert(1))
console.log(heap.arrayList)
console.log(heap.insert(4))
console.log(heap.arrayList)
console.log(heap.insert(3))
console.log(heap.arrayList)
console.log(heap.insert(7))
console.log(heap.arrayList)
console.log(heap.insert(7))
console.log(heap.arrayList)
