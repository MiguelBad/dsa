export type BTNode<T> = {
    item: T;
    left: BTNode<T> | null;
    right: BTNode<T> | null;
}

export function createTree(): BTNode<number> {
    const createNode = (item: number): BTNode<number> => ({
        item,
        left: null,
        right: null
    })

    const node1 = createNode(1);
    const node2 = createNode(2);
    const node3 = createNode(3);
    const node4 = createNode(4);
    const node5 = createNode(5);
    const node6 = createNode(6);
    const node7 = createNode(7);
    const node8 = createNode(8);
    const node9 = createNode(9);
    const node10 = createNode(10);
    const node11 = createNode(11);
    const node12 = createNode(12);
    const node13 = createNode(13);

    node1.left = node2;
    node1.right = node3;

    node2.left = node4;
    node2.right = node5;

    node3.left = node6;
    node3.right = node7;

    node4.left = node8;
    node4.right = node9;

    node5.left = node10;
    node5.right = node11;

    node6.right = node12;
    node6.right = node13;

    //          __  1  __
    //         /         \
    //        2           3
    //      /   \       /   \
    //     4     5     6     7 
    //    / \   / \   / \   / \
    //   8   9 10 11 12 13 -   -

    return node1;
}

export function createTree2(): BTNode<number> {
    const createNode = (item: number): BTNode<number> => ({
        item,
        left: null,
        right: null
    })

    const node1 = createNode(1);
    const node2 = createNode(2);
    const node3 = createNode(3);
    const node4 = createNode(4);
    const node5 = createNode(5);
    const node6 = createNode(6);
    const node7 = createNode(7);
    const node8 = createNode(8);
    const node9 = createNode(9);
    const node10 = createNode(20);
    const node11 = createNode(21);
    const node12 = createNode(22);
    const node13 = createNode(23);

    node1.left = node2;
    node1.right = node3;

    node2.left = node4;
    node2.right = node5;

    node3.left = node6;
    node3.right = node7;

    node4.left = node8;
    node4.right = node9;

    node5.left = node10;
    node5.right = node11;

    node6.right = node12;
    node6.right = node13;

    //          __  1  __
    //         /         \
    //        2           3
    //      /   \       /   \
    //     4     5     6     7 
    //    / \   / \   / \   / \
    //   8   9 20 21 22 23 -   -

    return node1;
}
