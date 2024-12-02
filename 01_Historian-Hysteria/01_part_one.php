<?php
// Read the input file
$inputFile = 'input.txt';
$lines = file($inputFile, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);

// Initialize arrays for left and right lists
$leftList = [];
$rightList = [];

// Parse the input into two separate lists
foreach ($lines as $line) {
    // Split the line on any whitespace
    $parts = preg_split('/\s+/', trim($line));

    // Ensure that we have exactly two parts
    if (count($parts) == 2) {
        $leftList[] = (int) $parts[0];
        $rightList[] = (int) $parts[1];
    } else {
        // Handle lines that don't have exactly two numbers
        echo "Warning: Line does not contain exactly two numbers: $line\n";
    }
}

// Debug: Output the size of the lists
echo "Number of entries: " . count($leftList) . PHP_EOL;

// Sort both lists
sort($leftList);
sort($rightList);

// Debug: Print the first few elements after sorting
echo "First 5 sorted leftList: " . implode(', ', array_slice($leftList, 0, 5)) . PHP_EOL;
echo "First 5 sorted rightList: " . implode(', ', array_slice($rightList, 0, 5)) . PHP_EOL;

// Calculate the total distance
$totalDistance = 0;
for ($i = 0; $i < count($leftList); $i++) {
    $difference = abs($leftList[$i] - $rightList[$i]);
    $totalDistance += $difference;

    // Debug: Show the pair and difference calculation
    if ($i < 5) { // Limit debugging output to the first 5 pairs
        echo "Pair: {$leftList[$i]} - {$rightList[$i]}, Difference: $difference" . PHP_EOL;
    }
}

// Output the total distance
echo "Total Distance: $totalDistance\n";
