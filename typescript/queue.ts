type QNode<T> = {
    item: T,
    next?: QNode<T>,
}

class Queue<T> {
    public length: number;
    private head?: QNode<T>;
    private tail?: QNode<T>;

    constructor() {
        this.length = 0;
        this.head = this.tail = undefined;
    }

    enqueue(item: T): void {
        const newNode: QNode<T> = { item: item }

        this.length++;
        if (!this.head) {
            this.head = this.tail = newNode;
        } else {
            this.tail.next = newNode;
            this.tail = newNode;
        }
    }

    deque(): T | undefined {
        if (!this.head) {
            return undefined;
        }

        this.length--;
        const item = this.head?.item
        this.head = this.head.next
        return item;
    }

    peek(): T | undefined {
        return this.head?.item;
    }
}

function test() {
    const newQueue = new Queue();
    console.log('length', newQueue.length);

    newQueue.enqueue('foo');
    console.log('\nenqueued foo');
    console.log('length', newQueue.length);
    console.log('peek', newQueue.peek());

    newQueue.enqueue('bar');
    console.log('\nenqueued bar');
    console.log('length', newQueue.length);
    console.log('peek', newQueue.peek());

    newQueue.enqueue(1);
    console.log('\nenqueued 1');
    console.log('length', newQueue.length);
    console.log('peek', newQueue.peek());

    newQueue.deque();
    console.log('\ndequed');
    console.log('length', newQueue.length);
    console.log('peek', newQueue.peek());

    newQueue.deque();
    console.log('\ndequed');
    console.log('length', newQueue.length);
    console.log('peek', newQueue.peek());

    newQueue.deque();
    console.log('\ndequed');
    console.log('length', newQueue.length);
    console.log('peek', newQueue.peek());
}
test();
