import { BTNode, createTree } from "./binary_tree";

function walk(curr: BTNode<number> | null, path: number[]): number[] {
    if (!curr){
        return path;
    }

    walk(curr.left, path);
    path.push(curr.item);
    walk(curr.right, path);

    return path;
}

function BTInOrder() {
    const tree =createTree();
    const arr = walk(tree, []);
    console.log(arr);
}

BTInOrder();
