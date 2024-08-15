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
        const deleteThis: (BTNode | string)[] | string = this.deleteNode(item, this.root);
        const parentNode = deleteThis[0] as BTNode;
        // im avoiding ternary operation to be easier to understand
        // but i think this is clearer. 
        let childToDelete = deleteThis[1] === 'left' ? parentNode.left : parentNode.right

        if (!deleteThis) {
            return undefined;
        }

        // both child notes is null
        if (!childToDelete.left && !childToDelete.right) {
            if (deleteThis[1] === 'left') {
                parentNode.left = null;
            } else if (deleteThis[1] === 'right') {
                parentNode.right = null;
            } else {
                this.root = null
            }
            return childToDelete.item;
        }

        // one of the child node is null
        if (!childToDelete.left || !childToDelete.right) {
            const childOfChild: BTNode = childToDelete.left || childToDelete.right
            if (deleteThis[1] === 'left') {
                parentNode.left = childOfChild
            } else if (deleteThis[1] === 'right') {
                parentNode.right = childOfChild
            } else {
                this.root = childOfChild;
            }
            return childToDelete.item;
        }

        // item to delete is equal to root val
        if (deleteThis[1] === 'root') {
            const successorParent = this.findSucessorParent(this.root.right);
            const successor = successorParent.left;

            successor.left = this.root.left;
            successor.right = this.root.right;
            this.root = successor;
            successorParent.left = null;
            return childToDelete.item;
        }

        // item to delete have both left and right children
        let successorParent = this.findSucessorParent(childToDelete);
        const successor = successorParent.left;

console.log(parentNode)
console.log(successorParent)
console.log(successor)

        if (deleteThis[1] === 'left') {
        } else  {
        }
        successorParent.left = successor.left;
        successorParent.right = successor.right;
        successorParent = successor;
        return childToDelete.item;
    }

    private deleteNode(item: number, curr: BTNode): (BTNode | string)[] | undefined {
        if (item === curr.item) {
            return [curr, 'root'];
        }
        if (curr.left === null && curr.right === null) {
            return undefined;
        }

        if (curr.left?.item === item) {
            return [curr, 'left'];
        }

        if (curr.right?.item === item) {
            return [curr, 'right'];
        }

        if (curr.item > item) {
            return this.deleteNode(item, curr.left);
        };
        return this.deleteNode(item, curr.right);
    }

    // is the smallest node after the root node
    // e.g. if  root is 7, successor is should be something like 8
    findSucessor(): number | null {
        if (!this.root || !this.root.right) {
            return null;
        }

        return this.findSucessorParent(this.root.right)?.left.item || null;
    }

    private findSucessorParent(curr: BTNode): BTNode {
        let parent: BTNode = null;
        while (curr.left) {
            parent = curr
            curr = curr.left;
        }

        return parent;
    }

    // similar to  successor, predecessor is the largest node before root node
    // e.g. if  root is 7, successor is should be something like 6
    findPredecessor(): number | null {
        if (!this.root || !this.root.left) {
            return null;
        }

        return this.findPredecessorParent(this.root.left)?.right.item || null;
    }

    private findPredecessorParent(curr: BTNode): BTNode {
        let parent: BTNode = null;
        while (curr.right) {
            parent = curr
            curr = curr.right;
        }

        return parent;
    }

    // preOrderTraversal(): number[] {
    // }

    // inOrderTraversal(): number[] {
    // }

    // postOrderTraversal(): number[] {
    // }

    // findMin(): number {
    // }

    // findMax(): number {
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

    myBST.insert(8);
    log(8, 'Insert');

    myBST.insert(10);
    log(10, 'Insert');

    // search test
    console.log('\n================\nSearch Test:');
    console.log('Search for 8:', myBST.search(8));
    console.log('Search for 7:', myBST.search(7));

    // delete test
    console.log('\n================\nDelete Test:')
    // console.log('Deleted:', myBST.delete(9));
    myBST.delete(9)
    log(9, 'Delete')

    // find predecessor
    console.log('\n================\nPredecessor Test:')
    console.log(myBST.findPredecessor());

    // find successor
    console.log('\n================\nSuccessor Test:')
    console.log(myBST.findSucessor());
}

BSTTest();
