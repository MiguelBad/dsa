class MinHeap {
    length: number;
    array: number[];

    constructor() {
        this.length = 0;
        this.array = [];
    }

    insert(item: number): void {
        this.array[this.length] = item;
        this.length += 1;
        this.upHeap(this.length - 1);
    }

    delete(): number | undefined {
        if (this.length === 0) {
            return undefined;
        }

        const out = this.array[0];
        this.length--;
        this.array[0] = this.array[this.length];
        this.array.splice(this.length, 1);
        this.downHeap(0);
        return out;
    }

    private downHeap(idx: number): void {
        const leftIdx = this.leftIdx(idx);
        const rightIdx = this.rightIdx(idx);

        if (!this.array[leftIdx]) {
            return;
        }

        const leftVal = this.array[leftIdx];
        const rightVal = this.array[rightIdx];

        if (leftVal > rightVal && rightVal < this.array[idx]) {
            this.array[rightIdx] = this.array[idx];
            this.array[idx] = rightVal;
            this.downHeap(rightIdx);
        } else if (leftVal < rightVal && leftVal < this.array[idx]) {
            this.array[leftIdx] = this.array[idx];
            this.array[idx] = leftVal;
            this.downHeap(leftIdx);
        }
    }

    private upHeap(idx: number): void {
        if (this.length < 2) {
            return;
        }

        const parentIdx = this.parentIdx(idx);
        const parentVal = this.array[parentIdx];
        const currentVal = this.array[idx];

        if (currentVal < parentVal) {
            this.array[idx] = parentVal;
            this.array[parentIdx] = currentVal;
            this.upHeap(parentIdx);
        }
    }

    private parentIdx(idx: number): number {
        return Math.floor((idx - 1) / 2);
    }

    private leftIdx(idx: number): number {
        return idx * 2 + 1;
    }

    private rightIdx(idx: number): number {
        return idx * 2 + 2;
    }
}

// test
const heap = new MaxHeap();

// add
console.log(heap.array);
heap.insert(1);
console.log(heap.array);
heap.insert(4);
console.log(heap.array);
heap.insert(3);
console.log(heap.array);
heap.insert(7);
console.log(heap.array);
heap.insert(7);
console.log(heap.array);

// delete
heap.delete();
console.log(heap.array);
heap.delete();
console.log(heap.array);
heap.delete();
console.log(heap.array);
heap.delete();
console.log(heap.array);
heap.delete();
console.log(heap.array);
