type BTNode = {
    item: number;
    left: BTNode | null;
    right: BTNode | null;
}

class BST {
    height: number;
    head: BTNode | null;

    constructor() {
        this.head = null;
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
            return;
        }

        this.insertNode(node, this.head);
    }

    private insertNode(node: BTNode, curr: BTNode): void {
        if (curr.item > node.item) {
            if (curr.left === null) {
                curr.left = node;
            } else {
                this.insertNode(node, curr.left);
            }
        } else {
            if (curr.right === null) {
                curr.right = node;
            } else {
                this.insertNode(node, curr.right);
            }
        }
    }

    search(n: number): BTNode | undefined {
        if (!this.head) {
            return undefined;
        }

        return this.searchNode(n, this.head);
    }

    private searchNode(n: number, curr: BTNode): BTNode | undefined {
        if (!curr) {
            return undefined;
        }

        if (curr.item === n) {
            return curr;
        } else {
            if (curr.item > n) {
                return this.searchNode(n, curr.left);
            } else {
                return this.searchNode(n, curr.right);
            }
        }
    }

    // delete(item: number): number {
    //     
    // }

    // preOrderTraversal(): number[] {
    // }
    //
    // inOrderTraversal(): number[] {
    // }
    //
    // postOrderTraversal(): number[] {
    // }

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

function BSTTest() {
    const addLog = (item: number | null) => {
        console.log('head item', myBST.head?.item);
        console.log('head', myBST.head);

        if (item) {
            console.log(`\ninsert ${item}`);
        }
    }

    const myBST = new BST();

    // add test
    addLog(7);
    myBST.insert(7);

    addLog(1);
    myBST.insert(1);

    addLog(3);
    myBST.insert(3);

    addLog(9);
    myBST.insert(9);

    addLog(7);
    myBST.insert(7);

    addLog(null);


    // search test
    console.log('\nSearch Test:');
    console.log('Search for 8:', myBST.search(8));
    console.log('Search for 1:', myBST.search(1));
}

BSTTest();
