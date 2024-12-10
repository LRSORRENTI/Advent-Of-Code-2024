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

def calculate_antinodes(antennas, map_width, map_height):
    """Calculate all antinodes for given antennas."""
    antinodes = set()
    for freq, positions in antennas.items():
        n = len(positions)
        for i in range(n):
            for j in range(i + 1, n):
                x1, y1 = positions[i]
                x2, y2 = positions[j]
                
                # Ensure the second antenna is twice as far as the first
                dx, dy = x2 - x1, y2 - y1
                x3, y3 = x1 - dx, y1 - dy
                x4, y4 = x2 + dx, y2 + dy
                
                # Check bounds and add antinodes
                if 0 <= x3 < map_width and 0 <= y3 < map_height:
                    antinodes.add((x3, y3))
                if 0 <= x4 < map_width and 0 <= y4 < map_height:
                    antinodes.add((x4, y4))
    return antinodes

def main(file_path):
    map_data = parse_input(file_path)
    map_width = len(map_data[0])
    map_height = len(map_data)
    antennas = find_antennas(map_data)
    antinodes = calculate_antinodes(antennas, map_width, map_height)
    return len(antinodes)

# Run the solution
if __name__ == "__main__":
    input_file = "input.txt"  # Replace with the actual input file path
    result = main(input_file)
    print(f"Number of unique antinode locations: {result}")
