import { BTNode, createTree, createTree2 } from './binary_tree'

function compare(leaf1: BTNode<number> | null, leaf2: BTNode<number> | null,): boolean {
    if (leaf1 === null && leaf2 === null) {
        return true;
    }

    if (leaf1 === null || leaf2 === null) {
        return false;
    }

    if (leaf1.item !== leaf2.item) {
        return false;
    }

    return compare(leaf1.left, leaf2.left) && compare(leaf1.right, leaf2.right)
}

function compareTree(): void {
    // similar compare
    const same1: BTNode<number> = createTree();
    const same2: BTNode<number> = createTree();
    console.log(compare(same1, same2));

    // similar compare
    const dif1: BTNode<number> = createTree();
    const dif2: BTNode<number> = createTree2();
    console.log(compare(dif1, dif2));
}

compareTree();
