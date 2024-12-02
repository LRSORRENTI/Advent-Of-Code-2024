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

// Create a frequency map for the right list
$rightFreqMap = array_count_values($rightList);

// Calculate the similarity score
$similarityScore = 0;
foreach ($leftList as $number) {
    $occurrences = isset($rightFreqMap[$number]) ? $rightFreqMap[$number] : 0;
    $score = $number * $occurrences;
    $similarityScore += $score;

    // Debug: Show calculation for the first few numbers
    static $debugCount = 0;
    if ($debugCount < 5) {
        echo "Number: $number, Occurrences in rightList: $occurrences, Score: $score\n";
        $debugCount++;
    }
}

// Output the total similarity score
echo "Similarity Score: $similarityScore\n";
