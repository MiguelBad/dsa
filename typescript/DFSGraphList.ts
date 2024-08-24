type GraphEdge = {
    to: number;
    weight: number;
}
type WeightedAdjacencyList = GraphEdge[][];
// how weighted adjacency list would look like in this type:
// -> const graph = [[{ to: 1, weight: 4 }, { to: 2, weight: 3 }], [{ to: 2, weight: 1 }]]
// with no types, can also be structured with:
// -> const graph = { 0: [[1, 4], [2, 3]], 1: [[2, 1]] } or a tuple

function walk(
    graph: WeightedAdjacencyList,
    curr: number,
    needle: number,
    seen: boolean[],
    path: number[]): boolean {

    if (seen[curr]) {
        return false;
    }

    seen[curr] = true;
    path.push(curr); 
    if (curr === needle) {
        return true;
    }

    const list = graph[curr];
    for (let i = 0; i < list.length; i++) {
        const edge = list[i];
        if (walk(graph, edge.to, needle, seen, path)) {
            return true;
        }
    }

    path.pop();
    return false;

}

function DFSGraphList(
    graph: WeightedAdjacencyList,
    source: number,
    needle: number): number[] {

    const seen: boolean[] = new Array(graph.length).fill(false);
    const path: number[] = [];

    walk(graph, source, needle, seen, path);

    return path;
}

function main() {
    const graph: WeightedAdjacencyList = [
        [
            { to: 1, weight: 5 },
            { to: 2, weight: 3 },
        ],
        [
            { to: 2, weight: 2 },
            { to: 3, weight: 6 },
        ],
        [
            { to: 3, weight: 7 }
        ],
        [
            { to: 0, weight: 4 }
        ]
    ]

    console.log(DFSGraphList(graph, 0, 3));
}

main();
