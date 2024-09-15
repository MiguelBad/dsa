from collections import deque


def is_clique(subset, graph):
    # Check if every pair of vertices in the subset is connected
    for i in range(len(subset)):
        for j in range(i + 1, len(subset)):
            if subset[j] not in graph[subset[i]]:
                return False
    return True

def bfs_find_k_clique(k, graph):
    # Initialize the queue for BFS
    queue = deque()

    # Start BFS from each vertex
    for vertex in graph:
        queue.append((vertex, [vertex]))

    while queue:
        current_vertex, path = queue.popleft()

        # If we've found a clique of size k
        if len(path) == k:
            if is_clique(path, graph):
                return True

        # Continue BFS, add neighbors to the path
        for neighbor in graph[current_vertex]:
            if neighbor not in path:
                new_path = path + [neighbor]
                if is_clique(new_path, graph):
                    queue.append((neighbor, new_path))

    return False  # No k-clique found

