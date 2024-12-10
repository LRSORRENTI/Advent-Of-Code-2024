def compact_disk(disk_map_str):
    # Parse the disk map into blocks
    blocks = []
    file_id = 0
    for i, ch in enumerate(disk_map_str):
        length = int(ch)
        if i % 2 == 0:
            # file segment
            for _ in range(length):
                blocks.append(str(file_id))
            file_id += 1
        else:
            # free segment
            for _ in range(length):
                blocks.append('.')

    # Perform the compaction process
    while True:
        # Find the leftmost free block '.'
        try:
            leftmost_free = blocks.index('.')
        except ValueError:
            # No free block at all means we are done
            break

        # Find the rightmost file block (non '.')
        # Starting from the end
        rightmost_file = -1
        for idx in range(len(blocks)-1, -1, -1):
            if blocks[idx] != '.':
                rightmost_file = idx
                break

        # If no file found or the rightmost file is not to the right of the leftmost free, no more moves
        if rightmost_file == -1 or rightmost_file <= leftmost_free:
            break

        # Move the file block from rightmost_file to leftmost_free
        blocks[leftmost_free] = blocks[rightmost_file]
        blocks[rightmost_file] = '.'

    # Compute checksum
    checksum = 0
    for i, b in enumerate(blocks):
        if b != '.':
            # b is a file ID digit; convert to int
            checksum += i * int(b)

    return checksum

# Example usage with a provided snippet:
# If the input is large, it will still follow the same pattern.

if __name__ == "__main__":
    # Read the single line input from input.txt
    with open("input.txt") as f:
        disk_map_line = f.read().strip()
    result = compact_disk(disk_map_line)
    print(result)
