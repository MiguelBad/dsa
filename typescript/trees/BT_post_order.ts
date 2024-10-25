import { BTNode, createTree } from "./binary_tree";

function walk(curr: BTNode<number> | null, path: number[]): number[] {
    if (!curr) {
        return path;
    }

    walk(curr.left, path);
    walk(curr.right, path);
    path.push(curr.item);

    return path;
}

function BTPostOrder() {
    const tree = createTree();
    const arr = walk(tree, []);
    console.log(arr);
}

BTPostOrder();
