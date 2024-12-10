def parse_input(file_path):
    """Read the input file and parse the antenna map."""
    with open(file_path, "r") as file:
        lines = file.readlines()
    return [line.strip() for line in lines]

def find_antennas(map_data):
    """Find all antennas and their positions in the map."""
    antennas = {}
    for y, row in enumerate(map_data):
        for x, char in enumerate(row):
            if char.isalnum():  # Check for lowercase, uppercase, or digit
                if char not in antennas:
                    antennas[char] = []
                antennas[char].append((x, y))
    return antennas

def calculate_antinodes_with_harmonics(antennas, map_width, map_height):
    """Calculate all antinodes including resonant harmonics."""
    antinodes = set()

    for freq, positions in antennas.items():
        n = len(positions)
        if n > 1:  # Only consider frequencies with more than one antenna
            # Add each antenna's position as an antinode
            for position in positions:
                antinodes.add(position)

            # Calculate all possible antinodes along the lines formed by pairs of antennas
            for i in range(n):
                for j in range(i + 1, n):
                    x1, y1 = positions[i]
                    x2, y2 = positions[j]

                    # Calculate the direction vector
                    dx, dy = x2 - x1, y2 - y1

                    # Generate antinodes along the line in both directions
                    k = 1
                    while True:
                        # Forward direction
                        xf, yf = x2 + k * dx, y2 + k * dy
                        # Backward direction
                        xb, yb = x1 - k * dx, y1 - k * dy

                        # Check if positions are within map bounds
                        if 0 <= xf < map_width and 0 <= yf < map_height:
                            antinodes.add((xf, yf))
                        if 0 <= xb < map_width and 0 <= yb < map_height:
                            antinodes.add((xb, yb))

                        # Stop if out of bounds in both directions
                        if (xf < 0 or xf >= map_width or yf < 0 or yf >= map_height) and \
                           (xb < 0 or xb >= map_width or yb < 0 or yb >= map_height):
                            break

                        k += 1

    return antinodes

def main_part_two(file_path):
    map_data = parse_input(file_path)
    map_width = len(map_data[0])
    map_height = len(map_data)
    antennas = find_antennas(map_data)
    antinodes = calculate_antinodes_with_harmonics(antennas, map_width, map_height)
    return len(antinodes)

# Run the solution for Part Two
if __name__ == "__main__":
    input_file = "input.txt"  # Replace with the actual input file path
    result = main_part_two(input_file)
    print(f"Number of unique antinode locations (Part Two): {result}")
