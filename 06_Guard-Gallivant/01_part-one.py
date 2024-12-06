def solve():
    # Instead of reading from sys.stdin, we open input.txt directly:
    with open('input.txt', 'r') as f:
        grid = [line.rstrip('\n') for line in f]

    rows = len(grid)
    cols = len(grid[0]) if rows > 0 else 0

    # Find the guard's initial position and direction
    dir_map = {'^': 0, '>': 1, 'v': 2, '<': 3}
    start_r, start_c = None, None
    dir_idx = None

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] in dir_map:
                start_r, start_c = r, c
                dir_idx = dir_map[grid[r][c]]
                break
        if start_r is not None:
            break

    # Directions: up, right, down, left
    directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]

    visited = set()
    visited.add((start_r, start_c))

    current_r, current_c = start_r, start_c

    while True:
        dr, dc = directions[dir_idx]
        next_r = current_r + dr
        next_c = current_c + dc

        # Check if outside bounds
        if not (0 <= next_r < rows and 0 <= next_c < cols):
            # Guard leaves the map
            break

        # Check if obstacle ahead
        if grid[next_r][next_c] == '#':
            # Turn right
            dir_idx = (dir_idx + 1) % 4
            # Don't move this turn
        else:
            # Move forward
            current_r, current_c = next_r, next_c
            visited.add((current_r, current_c))

    print(len(visited))


if __name__ == "__main__":
    solve()
