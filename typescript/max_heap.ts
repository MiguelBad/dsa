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
