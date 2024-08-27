type Graph = { [key: number]: number[] };
type VerticesQueue = [number, number[]][];

function isClique(subset: number[], graph: Graph) {
    for (let i = 0; i < subset.length; i++) {
        for (let j = i + 1; j < subset.length; j++) {
            if (!graph[subset[i]].includes(subset[j])) {
                return false;
            }
        }
    }

    return true;
}

function findKClique(k: number, graph: Graph): boolean {
    // using js array for queue (which is not good, O(n) on deque/shift)
    const verticesQueue: VerticesQueue = [];

    for (let v of Object.keys(graph).map(item => parseInt(item))) {
        verticesQueue.push([v, [v]])
    }


    while (verticesQueue.length > 0) {
        const queueItem: [number, number[]] = verticesQueue.shift();
        const vertex: number = queueItem[0];
        const path: number[] = queueItem[1];

        if (path.length === k) {
            return true;
        }

        for (let neighbor of graph[vertex]) {
            if (!path.includes(neighbor)) {
                const newPath = path.concat(neighbor);
                if (isClique(newPath, graph)) {
                    verticesQueue.push([neighbor, newPath]);
                }
            }
        }

    }

    return false;
}

function bfsClique(): void {
    const k = 3
    const graph: Graph = {
        0: [1, 2],
        1: [0],
        2: [0],
    }

    console.log(findKClique(k, graph));
}

bfsClique();
