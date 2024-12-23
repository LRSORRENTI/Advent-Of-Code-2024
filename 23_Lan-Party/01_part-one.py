def build_graph(connections):
    """
    Given a list of lines describing connections like 'kh-tc',
    return an adjacency list (dict[str, set[str]]).
    """
    graph = {}
    for line in connections:
        # Each line is of the form "node1-node2".
        node1, node2 = line.strip().split('-')
        
        if node1 not in graph:
            graph[node1] = set()
        if node2 not in graph:
            graph[node2] = set()
        
        graph[node1].add(node2)
        graph[node2].add(node1)
    
    return graph


def find_triangles(graph):
    """
    Return a set of tuples (a, b, c) representing
    all triangles (3-cliques) in the undirected graph.
    """
    triangles = set()
    
    # Iterate through each node in the graph
    for a in graph:
        neighbors_a = graph[a]
        for b in neighbors_a:
            if b <= a:  # Avoid redundant checks
                continue
            neighbors_b = graph[b]
            # Find common neighbors of a and b
            common_neighbors = neighbors_a.intersection(neighbors_b)
            for c in common_neighbors:
                if c > b:  # Ensure consistent ordering a < b < c
                    triangles.add(tuple(sorted([a, b, c])))
    
    return triangles


def triangles_with_t(triangles):
    """
    Given a set of 3-node tuples, return those that have
    at least one node starting with 't'.
    """
    result = []
    for tri in triangles:
        # tri is a tuple like (a, b, c)
        if any(node.startswith('t') for node in tri):
            result.append(tri)
    return result


def solve_advent_day23():
    # Read input from input.txt
    with open("input.txt", "r") as file:
        puzzle_input = file.read().strip().splitlines()
    
    # 1) Build the graph
    graph = build_graph(puzzle_input)
    
    # 2) Find all triangles
    all_triangles = find_triangles(graph)
    
    # 3) Filter for triangles that contain a node starting with 't'
    triangles_with_t_ = triangles_with_t(all_triangles)
    
    # Print result
    print(f"Number of triangles (3-computer sets) that contain at least one 't': {len(triangles_with_t_)}")
    
    # (Optional) Print them to check (sorting each triangle for readability)
    # for tri in sorted(triangles_with_t_):
    #     print(",".join(tri))


if __name__ == "__main__":
    solve_advent_day23()
