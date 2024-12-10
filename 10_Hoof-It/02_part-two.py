from collections import defaultdict
import sys
sys.setrecursionlimit(10**7)

def neighbors(r, c, R, C):
    for nr, nc in [(r-1,c),(r+1,c),(r,c-1),(r,c+1)]:
        if 0 <= nr < R and 0 <= nc < C:
            yield nr, nc

def dp(r, c, height_map, memo):
    if memo[r][c] is not None:
        return memo[r][c]

    R = len(height_map)
    C = len(height_map[0])
    h = height_map[r][c]

    # If already at height 9, one distinct path (itself)
    if h == 9:
        memo[r][c] = 1
        return 1

    total_paths = 0
    next_h = h + 1
    # Explore neighbors
    for nr, nc in neighbors(r, c, R, C):
        if height_map[nr][nc] == next_h:
            total_paths += dp(nr, nc, height_map, memo)

    memo[r][c] = total_paths
    return total_paths

if __name__ == "__main__":
    # Read from input.txt directly
    with open("input.txt", "r") as f:
        lines = [line.strip() for line in f if line.strip()]

    height_map = [list(map(int, list(line))) for line in lines]
    R = len(height_map)
    C = len(height_map[0])

    memo = [[None]*C for _ in range(R)]

    total_rating = 0
    # A trailhead is any cell with height=0
    for r in range(R):
        for c in range(C):
            if height_map[r][c] == 0:
                total_rating += dp(r, c, height_map, memo)

    print(total_rating)
