<?php
// Read input from a file
$input = file_get_contents('input.txt');
$lines = explode("\n", trim($input));
$grid = array_map('str_split', $lines); // Convert input into a 2D array
$rows = count($grid);
$cols = count($grid[0]);
$count = 0;

// Helper function to check bounds
function inBounds($x, $y, $rows, $cols)
{
    return $x >= 0 && $x < $rows && $y >= 0 && $y < $cols;
}

// Function to check for an X-MAS pattern at a given center
function isXMAS($grid, $x, $y, $rows, $cols)
{
    $patterns = [
        'MAS',
        'SAM',
    ];

    foreach ($patterns as $pattern1) {
        foreach ($patterns as $pattern2) {
            // First diagonal (top-left to bottom-right)
            $x1 = $x - 1;
            $y1 = $y - 1; // Should be pattern1[0]
            $x2 = $x;
            $y2 = $y;     // Should be 'A'
            $x3 = $x + 1;
            $y3 = $y + 1; // Should be pattern1[2]

            // Second diagonal (top-right to bottom-left)
            $x4 = $x - 1;
            $y4 = $y + 1; // Should be pattern2[0]
            // Center is at ($x2, $y2)
            $x5 = $x + 1;
            $y5 = $y - 1; // Should be pattern2[2]

            if (
                inBounds($x1, $y1, $rows, $cols) &&
                inBounds($x3, $y3, $rows, $cols) &&
                inBounds($x4, $y4, $rows, $cols) &&
                inBounds($x5, $y5, $rows, $cols) &&
                $grid[$x1][$y1] == $pattern1[0] &&
                $grid[$x2][$y2] == 'A' && // Center must be 'A'
                $grid[$x3][$y3] == $pattern1[2] &&
                $grid[$x4][$y4] == $pattern2[0] &&
                $grid[$x5][$y5] == $pattern2[2]
            ) {
                return 1; // Found a valid X-MAS pattern
            }
        }
    }
    return 0;
}

// Search for X-MAS patterns
for ($x = 0; $x < $rows; $x++) {
    for ($y = 0; $y < $cols; $y++) {
        if ($grid[$x][$y] == 'A') {
            $count += isXMAS($grid, $x, $y, $rows, $cols);
        }
    }
}

echo "Total occurrences of X-MAS: $count\n";
?>