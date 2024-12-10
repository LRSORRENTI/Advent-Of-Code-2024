def parse_disk_map(disk_map_str):
    """
    Parse the disk map into a list of characters representing file blocks and free blocks.
    Also return the number of files (the highest file ID + 1).
    """
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
    return blocks, file_id

def find_files(blocks):
    """
    Identify all files in the blocks array.
    Returns a list of (file_id, start_index, length).
    The files are guaranteed to appear in ascending order of their IDs.
    """
    # Since files are assigned sequential IDs and each block of a file has the same digit,
    # we can find contiguous runs of the same digit to determine file boundaries.
    files = []
    current_id = None
    current_start = None
    current_length = 0

    for i, b in enumerate(blocks):
        if b == '.':
            # If we were tracking a file, close it out
            if current_id is not None:
                files.append((int(current_id), current_start, current_length))
                current_id = None
                current_start = None
                current_length = 0
            # Else just continue
        else:
            # b is a file block
            if b != current_id:
                # We've encountered a new file or a break in continuity
                # If we were tracking another file, close it
                if current_id is not None:
                    files.append((int(current_id), current_start, current_length))
                # Start tracking a new file
                current_id = b
                current_start = i
                current_length = 1
            else:
                # Same file continues
                current_length += 1
    # If ended on a file
    if current_id is not None:
        files.append((int(current_id), current_start, current_length))
    
    return files

def move_files(blocks):
    """
    Attempt to move each file once in order of decreasing file ID.
    """
    files = find_files(blocks)
    # Sort by file ID descending
    files.sort(key=lambda x: x[0], reverse=True)

    for file_id, start_idx, length in files:
        # Check if this file still exists at the same location (it might have moved if a previous move affected it)
        # We must re-check because the disk has changed after each move.
        # Identify the file again by scanning from start_idx for 'length' blocks of the same digit
        if start_idx >= len(blocks):
            # Something went off, or the file was possibly moved, skip
            continue
        # Validate the file is still intact and find its actual current location
        current_file_id = blocks[start_idx]
        if current_file_id == '.' or int(current_file_id) != file_id:
            # The file might have moved or changed; we must re-locate it
            # Let's find this file again by scanning all blocks
            new_files = find_files(blocks)
            # Extract the current position of the file with file_id
            new_files_dict = {f[0]: (f[1], f[2]) for f in new_files}
            if file_id not in new_files_dict:
                # File not found (?). Possibly size 0 or something unexpected. Just skip.
                continue
            start_idx, length = new_files_dict[file_id]
            current_file_id = blocks[start_idx]

        # Now we know the file's current location and length
        # Find a free space segment to the left of start_idx that can hold the entire file
        # We'll scan for contiguous '.' runs from the start of the disk up to start_idx-1
        suitable_start = None
        count = 0
        run_start = None
        
        for i in range(start_idx):
            if blocks[i] == '.':
                if run_start is None:
                    run_start = i
                count += 1
                if count >= length:
                    # Found a suitable segment
                    suitable_start = run_start
                    break
            else:
                # Reset run
                run_start = None
                count = 0
        
        # If no suitable segment was found, do not move the file
        if suitable_start is None:
            continue
        
        # Move the file there
        # We know the file currently occupies [start_idx, start_idx + length - 1]
        # Copy the file blocks to [suitable_start, suitable_start + length - 1]
        # Then set old location to '.'
        file_blocks = blocks[start_idx:start_idx+length]
        # Place them in the new free space
        for i in range(length):
            blocks[suitable_start + i] = file_blocks[i]
        # Free old space
        for i in range(length):
            blocks[start_idx + i] = '.'

def compute_checksum(blocks):
    """
    Compute the checksum as described:
    For each block that is not '.', sum i * file_id.
    """
    checksum = 0
    for i, b in enumerate(blocks):
        if b != '.':
            checksum += i * int(b)
    return checksum

if __name__ == "__main__":
    # Read input
    with open("input.txt") as f:
        disk_map_line = f.read().strip()

    # Parse disk map
    blocks, _ = parse_disk_map(disk_map_line)

    # Move files according to the part two rules
    move_files(blocks)

    # Compute checksum
    result = compute_checksum(blocks)
    print(result)
