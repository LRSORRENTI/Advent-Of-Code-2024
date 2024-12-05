<?php
// Read the input file without skipping empty lines
$lines = file('input.txt', FILE_IGNORE_NEW_LINES);

// Initialize variables
$ordering_rules = array();
$updates = array();
$parsing_rules = true;

// Parse the input into ordering rules and updates
foreach ($lines as $line) {
    $line = trim($line);
    if ($line === '') {
        $parsing_rules = false;
        continue;
    }

    if ($parsing_rules) {
        // Parse ordering rules
        list($x, $y) = array_map('intval', explode('|', $line));
        $ordering_rules[] = array($x, $y);
    } else {
        // Parse updates
        $update_pages = array_map('intval', explode(',', $line));
        $updates[] = $update_pages;
    }
}

// Function to check if an update is in correct order
function is_correct_order($update_pages, $ordering_rules)
{
    // Build a map from page number to its index in the update
    $position = array();
    foreach ($update_pages as $index => $page) {
        $position[$page] = $index;
    }
    // For each ordering rule
    foreach ($ordering_rules as $rule) {
        $x = $rule[0];
        $y = $rule[1];
        // If both x and y are in the update
        if (isset($position[$x]) && isset($position[$y])) {
            // Check if position[x] < position[y]
            if ($position[$x] >= $position[$y]) {
                // Rule violated
                return false;
            }
        }
        // else, rule does not apply for this update
    }
    // All rules satisfied
    return true;
}

// Process each update
$total = 0;
foreach ($updates as $update_pages) {
    if (is_correct_order($update_pages, $ordering_rules)) {
        // Update is in correct order
        // Find the middle page number
        $length = count($update_pages);
        $middle_index = floor($length / 2);
        $middle_page = $update_pages[$middle_index];
        $total += $middle_page;
    }
}

// Output the total sum of middle page numbers
echo $total . PHP_EOL;
?>