import sys
from collections import deque

def neighbors(r, c, R, C):
    for nr, nc in [(r-1,c),(r+1,c),(r,c-1),(r,c+1)]:
        if 0 <= nr < R and 0 <= nc < C:
            yield nr, nc

def find_score_for_trailhead(height_map, start_r, start_c):
    R = len(height_map)
    C = len(height_map[0])
    start_height = height_map[start_r][start_c]
    if start_height != 0:
        return 0

    visited = set()
    queue = deque()
    queue.append((start_r, start_c))
    visited.add((start_r, start_c))
    found_nines = set()

    while queue:
        r, c = queue.popleft()
        h = height_map[r][c]
        if h == 9:
            found_nines.add((r, c))
        # Move to neighbors with height h+1
        next_h = h + 1
        if next_h <= 9:
            for nr, nc in neighbors(r, c, R, C):
                if height_map[nr][nc] == next_h and (nr, nc) not in visited:
                    visited.add((nr, nc))
                    queue.append((nr, nc))

    return len(found_nines)

if __name__ == "__main__":
    # Read the entire input from input.txt
    with open("input.txt", "r") as f:
        lines = [line.strip() for line in f if line.strip()]

    height_map = [list(map(int, list(line))) for line in lines]
    R = len(height_map)
    C = len(height_map[0])

    # Find all trailheads (height=0)
    trailheads = [(r, c) for r in range(R) for c in range(C) if height_map[r][c] == 0]

    total_score = 0
    for (r,c) in trailheads:
        score = find_score_for_trailhead(height_map, r, c)
        total_score += score

    print(total_score)
