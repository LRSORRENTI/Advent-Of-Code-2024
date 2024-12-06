def solve():
    # Read the input
    with open('input.txt', 'r') as f:
        grid = [line.rstrip('\n') for line in f]

    rows = len(grid)
    cols = len(grid[0]) if rows > 0 else 0

    # Identify the guard's start position and direction
    dir_map = {'^': 0, '>': 1, 'v': 2, '<': 3}
    start_r, start_c, dir_idx = None, None, None

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] in dir_map:
                start_r, start_c = r, c
                dir_idx = dir_map[grid[r][c]]
                break
        if start_r is not None:
            break

    directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]

    # We will need a helper function to simulate given a modified grid
    def will_loop_with_obstacle(obstacle_r, obstacle_c):
        # Create a modified grid as a list of lists for mutability
        mod_grid = [list(row) for row in grid]

        # Place the new obstacle
        mod_grid[obstacle_r][obstacle_c] = '#'

        # Convert back to something easy to index
        mod_grid = ["".join(row) for row in mod_grid]

        # Run simulation
        current_r, current_c = start_r, start_c
        current_dir = dir_idx

        visited_states = set()
        visited_states.add((current_r, current_c, current_dir))

        while True:
            dr, dc = directions[current_dir]
            next_r = current_r + dr
            next_c = current_c + dc

            # Check if leaving the map
            if not (0 <= next_r < rows and 0 <= next_c < cols):
                # No loop, escaped
                return False

            # Check obstacle
            if mod_grid[next_r][next_c] == '#':
                # Turn right
                current_dir = (current_dir + 1) % 4
                # After turning, check if we have seen this state again
                # Actually, we only mark movement states, but let's just continue
                # We'll mark states after a move or turn
                if (current_r, current_c, current_dir) in visited_states:
                    return True
                visited_states.add((current_r, current_c, current_dir))
            else:
                # Move forward
                current_r, current_c = next_r, next_c
                if (current_r, current_c, current_dir) in visited_states:
                    # Loop detected
                    return True
                visited_states.add((current_r, current_c, current_dir))

    # Try placing an obstacle in each empty cell (not obstacle, not start)
    loop_count = 0
    for r in range(rows):
        for c in range(cols):
            # Conditions:
            # - Not the start position
            # - Not currently an obstacle
            # - Not the initial direction character
            if (r == start_r and c == start_c):
                continue
            if grid[r][c] == '#' or grid[r][c] in dir_map:
                continue

            # Now test placing an obstacle here
            if will_loop_with_obstacle(r, c):
                loop_count += 1

    print(loop_count)

solve()