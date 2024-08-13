type BTNode = {
    item: number;
    left: BTNode | null;
    right: BTNode | null;
}

class BTS {
    height: number;
    head: BTNode;
    pivot: number;

    constructor() {
        this.head = undefined;
        this.height = 0;
    }

    insert(item: number): void {
        const node: BTNode = {
            item: item,
            left: null,
            right: null,
        };

        if (!this.head) {
            this.head = node;
            this.height++;
            return; 
        } 
        if (self.head.item > node.item) {
        }

    }

    delete(item: number): number {
    }

    search(n: number): number {
    }

    preOrderTraversal(): number[] {
    }

    inOrderTraversal(): number[] {
    }

    postOrderTraversal(): number[] {
    }

    // findMin(): number {
    // }
    //
    // findMax(): number {
    // }
    //
    // findPredecessor() {
    // }
    //
    // findSucessor() {
    // }
}
