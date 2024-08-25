type Graph = {
    [key: number]: number[];
}

function bronKerbosch(
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

        if (bronKerbosch(k, graph, newCurr, newPotential, newExcluded)) { 
            return true
        }

        potential.delete(v);
        excluded.add(v);
    }

    return false;
}

function main(k: number, graph: Graph): Boolean {
    const potential: Set<number> = new Set(Object.keys(graph).map(x => parseInt(x)));
    return bronKerbosch(k, graph, [], potential, new Set());
}

const k: number = 3;
const graph: Graph = {
    0: [1, 2],
    1: [0, 2],
    2: [0, 1],
}

console.log(main(k, graph));
