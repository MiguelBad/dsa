type BTNode = {
    item: number;
    left: BTNode | null;
    right: BTNode | null;
}

class BST {
    height: number;
    head?: BTNode;

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
            return;
        }

        this.insertNode(node, this.head);
    }

    private insertNode(node: BTNode, curr: BTNode): void {
        if (curr.item > node.item) {
            if (curr.left === null) {
                curr.left = node;
                return;
            }
            this.insertNode(node, curr.left);
        } else {
            if (curr.right === null) {
                curr.right = node;
                return;
            }
            this.insertNode(node, curr.right);
        }
    }

    // delete(item: number): number {
    // }
    //
    // search(n: number): number {
    // }
    //
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

function addTestBST() {
    const log = (item: number | null) => {
        console.log('head item', myBST.head?.item);
        console.log('head', myBST.head);

        if (item) {
            console.log(`\ninsert ${item}`);
        }
    }

    const myBST = new BST();

    log(7);
    myBST.insert(7);

    log(1);
    myBST.insert(1);

    log(3);
    myBST.insert(3);

    log(9);
    myBST.insert(9);

    log(7);
    myBST.insert(7);

    log(null);
}

addTestBST();
