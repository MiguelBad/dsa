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
            return;
        }

        this.insertNode(node, this.root);
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

        const deleteThis: (BTNode | string)[] = this.deleteNode(item, this.root);
        if (!deleteThis) {
            return undefined;
        }

        const parentNode = deleteThis[0] as BTNode;
        let childToDelete = deleteThis[1] === 'left' ? parentNode.left : parentNode.right
        const temp = childToDelete; 
        if (!childToDelete.left && !childToDelete.right) {
            childToDelete = null
            return temp.item;
        }

        if (!childToDelete.left || !childToDelete.right) {
            const childOfChild: BTNode = childToDelete.left || childToDelete.right
            childToDelete = childOfChild;
            return temp.item;
        }
        
        childToDelete = childToDelete.right
        childToDelete.left = temp.left
        return temp.item;
    }

    private deleteNode(item: number, curr: BTNode): (BTNode | string)[] | undefined {
        if (curr.left === null || curr.right === null) {
            return undefined
        }

        if (curr.left.item === item) {
            return [curr, 'left'];
        }

        if (curr.right.item === item) {
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
    const addLog = (item: number | null) => {
        console.log('root item', myBST.root?.item);
        console.log('root', myBST.root);

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
    console.log('Search for 7:', myBST.search(7));

    // delete test
    console.log('test test')
    console.log(myBST.delete(2));
}

BSTTest();
