class MaxHeap {
    arrayList: number[];
    length: number;

    constructor() {
        this.arrayList = [];
        this.length = 0;
    }

    insert(item: number): void {
        // move item to end of array
        this.arrayList[this.length] = item;
        this.length++;

        // perform upHeap
        this.upHeap(this.length - 1);
    }

    delete(): number | undefined {
        // return undefined if empty heap
        if (this.length === 0) {
            return undefined;
        }

        const item = this.arrayList[0];
        this.length--;

        // clear array when length is 1
        if (this.length === 0) {
            this.arrayList = []
            return item;
        }

        // move the last element to the first 
        this.arrayList[0] = this.arrayList[this.length];
        this.arrayList.splice(this.length, 1);
        this.downHeap(0); // perform down heap
        return item;
    }

    private downHeap(idx: number): void {
        const leftIdx = this.leftChild(idx);
        const rightIdx = this.rightChild(idx);

        // basecase
        if (idx >= this.length || leftIdx >= this.length) {
            return;
        }

        const leftVal = this.arrayList[leftIdx];
        const rightVal = this.arrayList[rightIdx];
        const currVal = this.arrayList[idx];

        // check if left or right child value is lower to switch to 
        if (leftVal > rightVal && leftVal > currVal) {
            // swap
            this.arrayList[idx] = leftVal;
            this.arrayList[leftIdx] = currVal;

            // recurse
            this.downHeap(leftIdx);
        } else if (leftVal < rightVal && rightVal > currVal) {
            // swap
            this.arrayList[idx] = rightVal;
            this.arrayList[rightIdx] = currVal;

            // recurse
            this.downHeap(rightIdx);
        }

    }

    private upHeap(idx: number): void {
        // base case
        if (idx === 0) {
            return;
        }

        const parentIdx = this.parent(idx);
        const parentVal = this.arrayList[parentIdx];
        const currVal = this.arrayList[idx];

        // switch positions if child > parent
        if (currVal > parentVal) {
            // swap
            this.arrayList[parentIdx] = currVal;
            this.arrayList[idx] = parentVal;

            // recurse
            this.upHeap(parentIdx);
        }
    }

    // formula to get the parent index
    private parent(idx: number): number {
        return Math.floor((idx - 1) / 2);
    }

    // formula to get the left child index
    private leftChild(idx: number): number {
        return idx * 2 + 1;
    }

    // formula to get the right child index
    private rightChild(idx: number): number {
        return idx * 2 + 2;
    }
}

// test
const heap = new MaxHeap();

// add
console.log(heap.arrayList);
heap.insert(1);
console.log(heap.arrayList);
heap.insert(4);
console.log(heap.arrayList);
heap.insert(3);
console.log(heap.arrayList);
heap.insert(7);
console.log(heap.arrayList);
heap.insert(7);
console.log(heap.arrayList);

// delete
heap.delete();
console.log(heap.arrayList);
heap.delete();
console.log(heap.arrayList);
heap.delete();
console.log(heap.arrayList);
heap.delete();
console.log(heap.arrayList);
heap.delete();
console.log(heap.arrayList);
