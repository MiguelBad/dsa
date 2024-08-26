type Graph = {
    [key: number]: number[];
}

function findClique(
    k: number,
    graph: Graph,
    curr: number[],
    potential: Set<number>,
    excluded: Set<number>): Boolean {

    if (curr.length === k) {
        return true;
    }

    if (!potential && !excluded) {
        return false;
    }

    const potentialCopy = potential;
    for (const v of potentialCopy) {
        const newCurr = curr.concat(v);
        const edges = new Set(graph[v]);

        // intersection operation in a set
        const newPotential = new Set([...potential].filter(item => edges.has(item)));
        const newExcluded = new Set([...excluded].filter(item => edges.has(item)));

        if (findClique(k, graph, newCurr, newPotential, newExcluded)) { 
            return true
        }

        potential.delete(v);
        excluded.add(v);
    }

    return false;
}

function bronKerbosch(): Boolean {
    const k: number = 3;
    const graph: Graph = {
        0: [1, 2],
        1: [0, 2],
        2: [0, 1],
    };

    const potential: Set<number> = new Set(Object.keys(graph).map(x => parseInt(x)));
    return findClique(k, graph, [], potential, new Set());
}

bronKerbosch();
