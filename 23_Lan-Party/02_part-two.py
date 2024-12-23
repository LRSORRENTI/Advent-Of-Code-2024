def build_graph(connections):
    """
    Given a list of lines describing connections like 'kh-tc',
    return an adjacency list (dict[node, set[node]]).
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

def bron_kerbosch(R, P, X, graph, cliques):
    """
    Bron–Kerbosch helper function to find all maximal cliques.
    R: set of nodes in the current clique
    P: set of candidate nodes that can still be added to R
    X: set of nodes that should not be added to R
    graph: adjacency list
    cliques: list to collect all maximal cliques found
    """
    if not P and not X:
        # R is a maximal clique
        cliques.append(R)
        return
    
    # Copy P because we're going to modify it
    for v in list(P):
        # Explore adding v to the clique
        neighbors_v = graph[v]
        bron_kerbosch(
            R.union({v}), 
            P.intersection(neighbors_v), 
            X.intersection(neighbors_v), 
            graph, 
            cliques
        )
        # Move v from P to X
        P.remove(v)
        X.add(v)

def find_maximum_clique(graph):
    """
    Use the Bron–Kerbosch algorithm to find the largest clique in the graph.
    """
    # Initialize
    R = set()
    P = set(graph.keys())
    X = set()
    all_cliques = []
    
    # Run Bron–Kerbosch to find all maximal cliques
    bron_kerbosch(R, P, X, graph, all_cliques)
    
    # Pick the largest clique by size
    max_clique = max(all_cliques, key=len)
    return max_clique

def solve_advent_day23_part2():
    # Read input from input.txt
    with open("input.txt", "r") as file:
        puzzle_input = file.read().strip().splitlines()
    
    # 1) Build the graph
    graph = build_graph(puzzle_input)
    
    # 2) Find the largest clique (the LAN party)
    largest_clique = find_maximum_clique(graph)
    
    # 3) Form the "password": sorted, then joined by commas
    password = ",".join(sorted(largest_clique))
    
    print(f"The LAN party computers (largest clique) are: {sorted(largest_clique)}")
    print(f"The password to get in is: {password}")

if __name__ == "__main__":
    solve_advent_day23_part2()
