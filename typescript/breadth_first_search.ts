import { BTNode, createTree } from "./binary_tree"

type Node<T> = {
    item: T,
    next?: Node<T>,
}

class Queue<T> {
    public length: number;
    private head?: Node<T>;
    private tail?: Node<T>;

    constructor() {
        this.length = 0;
        this.head = this.tail = undefined;
    }

    enqueue(item: T): void {
        const newNode: Node<T> = { item: item }

        this.length++;
        if (!this.head) {
            this.head = this.tail = newNode;
        } else {
            this.tail.next = newNode;
            this.tail = newNode
        }
    }

    deque(): T | undefined {
        if (!this.head) {
            return undefined;
        }

        this.length--;
        const item = this.head.item;
        this.head = this.head.next;
        return item;
    }
}

function breadthFirstSearch(head: BTNode<number>, needle: number): boolean {
    const BTQueue = new Queue();
    BTQueue.enqueue(head)

    while (BTQueue.length) {
        const curr = BTQueue.deque() as BTNode<number> | undefined | null;
        if (!curr) {
            continue;
        }

        if (curr?.item === needle) {
            return true;
        }

        BTQueue.enqueue(curr.left)
        BTQueue.enqueue(curr.right)
    }
}

function main() {
    const tree = createTree();
    const arr = breadthFirstSearch(tree, 100);
    console.log(arr)
}

main();
