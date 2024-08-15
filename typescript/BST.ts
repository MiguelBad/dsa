type BTNode = {
    item: number;
    left: BTNode | null;
    right: BTNode | null;
}

class BST {
    height: number;
    root: BTNode | null;

    constructor() {
        this.root = null;
        this.height = 0;
    }

    insert(item: number): void {
        const node: BTNode = {
            item: item,
            left: null,
            right: null,
        };

        if (!this.root) {
            this.root = node;
            this.height++;
            return;
        }

        const height: number = this.insertNode(node, this.root);
        if (height > this.height) {
            this.height = height
        }
    }

    private insertNode(node: BTNode, curr: BTNode): number {
        if (curr.item > node.item) {
            if (curr.left === null) {
                curr.left = node;
                return 1;
            } else {
                return 1 + this.insertNode(node, curr.left);
            }
        } else {
            if (curr.right === null) {
                curr.right = node;
                return 1;
            } else {
                return 1 + this.insertNode(node, curr.right);
            }
        }
    }

    search(n: number): BTNode | undefined {
        if (!this.root) {
            return undefined;
        }

        return this.searchNode(n, this.root);
    }

    private searchNode(n: number, curr: BTNode): BTNode | undefined {
        if (!curr) {
            return undefined;
        }

        if (curr.item === n) {
            return curr;
        }

        if (curr.item > n) {
            return this.searchNode(n, curr.left);
        }
        return this.searchNode(n, curr.right);
    }

    delete(item: number): number | undefined {
        if (this.root.item === item) {
            // TODO: CHANGE ROOT
        }

        const deleteThis: (BTNode | string)[] | string= this.deleteNode(item, this.root);
        if (!deleteThis) {
            return undefined;
        }

        const parentNode = deleteThis[0] as BTNode;
        // im avoiding ternary operation to be easier to understand
        // but i think this is clearer. 
        let childToDelete = deleteThis[1] === 'left' ? parentNode.left : parentNode.right

        // both child notes is null
        if (!childToDelete.left && !childToDelete.right) {
            if (deleteThis[1] === 'left') {
                parentNode.left = null;
            } else {
                parentNode.right = null;
            }
            return childToDelete.item;
        }

        // one of the child node is null
        if (!childToDelete.left || !childToDelete.right) {
            console.log('here')
            const childOfChild: BTNode = childToDelete.left || childToDelete.right
            if (deleteThis[1] === 'left') {
                parentNode.left = childOfChild
            } else {
                parentNode.right = childOfChild
            }
            return childToDelete.item;
        }

        // no child node is null
        // parentNode.left = childToDelete.left
        // childToDelete = childToDelete.right
        // childToDelete.left = childToDelete.left
        return childToDelete.item;
    }

    private deleteNode(item: number, curr: BTNode): (BTNode | string)[] | undefined {
        if (curr.left === null && curr.right === null) {
            return undefined
        }

        if (curr.left?.item === item) {
            return [curr, 'left'];
        }

        if (curr.right?.item === item) {
            return [curr, 'right'];
        }

        if (curr.item > item) {
            return this.deleteNode(item, curr.left);
        }
        return this.deleteNode(item, curr.right);
    }

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
    const log = (item: number | null, action: string | undefined) => {
        if (item) {
            console.log(`\n${action} ${item}`);
        }

        console.log('root item', myBST.root?.item);
        console.log('root', myBST.root);
        console.log('height', myBST.height);
    }

    const myBST = new BST();

    // add test
    myBST.insert(7);
    log(7, 'Insert');

    myBST.insert(1);
    log(1, 'Insert');

    myBST.insert(3);
    log(3, 'Insert');

    myBST.insert(9);
    log(9, 'Insert');

    myBST.insert(7);
    log(7, 'Insert');

    // search test
    console.log('\n================\nSearch Test:');
    console.log('Search for 8:', myBST.search(8));
    console.log('Search for 7:', myBST.search(7));

    // delete test
    console.log('\n================\nDelete Test:')
    console.log('Deleted:', myBST.delete(3));
    log(3, 'Delete')
}

BSTTest();
