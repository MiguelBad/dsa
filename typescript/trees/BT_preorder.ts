import { BTNode, createTree } from './binary_tree'

function walk(curr: BTNode<number> | null, path: number[]): number[] {
    if (!curr) {
        return path;
    }
    
    path.push(curr.item);
    walk(curr.left, path);
    walk(curr.right, path);

    return path;
}

function BTPreOrder(): void {
    const tree = createTree();
    const arr = walk(tree, []);
    console.log(arr);
}
BTPreOrder();
