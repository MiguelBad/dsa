type VerticesPath = [number, number[]][];

function isClique(subset: number[], graph: Graph) {
    for (let i = 0; i < subset.length; i++) {
        for (let j = i + 1; j < subset.length; i++) {
            if (!graph[subset[i]].includes(subset[j])) {
                return false;
            }
        }
    }

    return true;
}

function findKClique(k: number, graph: Graph): boolean {
    // using js array for queue (which is not good, O(n) on deque/shift)
    const vertices: number[][] = [];

    while (vertices) {
        const vertex: number = vertices.shift();


    }

    return false;
}

function bfsClique(): void {
    const k = 3
    const graph: Graph = {
        0: [1, 2],
        1: [0, 2],
        2: [0, 1]
    }

    findKClique(k, graph);
}

bfsClique();
