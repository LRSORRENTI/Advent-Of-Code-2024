<?php
// Read input from a file
$input = file_get_contents('input.txt');
$lines = explode("\n", trim($input));
$grid = array_map('str_split', $lines); // Convert input into a 2D array
$word = "XMAS";
$wordLength = strlen($word);
$rows = count($grid);
$cols = count($grid[0]);
$count = 0;

// Helper function to check bounds
function inBounds($x, $y, $rows, $cols)
{
    return $x >= 0 && $x < $rows && $y >= 0 && $y < $cols;
}

// Directions to search: (dx, dy)
$directions = [
    [0, 1],  // Horizontal right
    [0, -1], // Horizontal left
    [1, 0],  // Vertical down
    [-1, 0], // Vertical up
    [1, 1],  // Diagonal down-right
    [-1, -1],// Diagonal up-left
    [1, -1], // Diagonal down-left
    [-1, 1]  // Diagonal up-right
];

// Search for the word in all directions
foreach ($grid as $row => $line) {
    foreach ($line as $col => $char) {
        foreach ($directions as [$dx, $dy]) {
            $found = true;
            for ($i = 0; $i < $wordLength; $i++) {
                $nx = $row + $i * $dx;
                $ny = $col + $i * $dy;
                if (!inBounds($nx, $ny, $rows, $cols) || $grid[$nx][$ny] !== $word[$i]) {
                    $found = false;
                    break;
                }
            }
            if ($found) {
                $count++;
            }
        }
    }
}

echo "Total occurrences of '$word': $count\n";
?>