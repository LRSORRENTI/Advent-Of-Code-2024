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

// Function to perform topological sort
function topological_sort($nodes, $edges)
{
    $visited = array();
    $temp_mark = array();
    $sorted = array();

    foreach ($nodes as $node) {
        if (!isset($visited[$node])) {
            $result = visit($node, $edges, $visited, $temp_mark, $sorted);
            if (!$result) {
                // Cycle detected
                return false;
            }
        }
    }
    return array_reverse($sorted);
}

function visit($node, &$edges, &$visited, &$temp_mark, &$sorted)
{
    if (isset($temp_mark[$node])) {
        // Not a DAG
        return false;
    }
    if (!isset($visited[$node])) {
        $temp_mark[$node] = true;
        if (isset($edges[$node])) {
            foreach ($edges[$node] as $m) {
                $result = visit($m, $edges, $visited, $temp_mark, $sorted);
                if (!$result) {
                    return false;
                }
            }
        }
        $visited[$node] = true;
        unset($temp_mark[$node]);
        $sorted[] = $node;
    }
    return true;
}

// Process each update
$total = 0;
$incorrect_updates = array();
foreach ($updates as $update_pages) {
    if (is_correct_order($update_pages, $ordering_rules)) {
        // Update is in correct order
        // Do nothing
    } else {
        // Update is not in correct order
        $incorrect_updates[] = $update_pages;
    }
}

// For each incorrect update, reorder it
foreach ($incorrect_updates as $update_pages) {
    // Build the dependency graph
    $nodes = $update_pages;
    $edges = array();
    // For each ordering rule
    foreach ($ordering_rules as $rule) {
        $x = $rule[0];
        $y = $rule[1];
        // If both x and y are in the update
        if (in_array($x, $update_pages) && in_array($y, $update_pages)) {
            // Add edge from x to y
            if (!isset($edges[$x])) {
                $edges[$x] = array();
            }
            $edges[$x][] = $y;
        }
    }
    // Perform topological sort
    $sorted_pages = topological_sort($nodes, $edges);
    if ($sorted_pages === false) {
        // Cycle detected, cannot sort
        // Should not happen in this problem
        continue;
    }
    // Get the middle page number
    $length = count($sorted_pages);
    $middle_index = floor($length / 2);
    $middle_page = $sorted_pages[$middle_index];
    $total += $middle_page;
}

echo $total . PHP_EOL;
?>