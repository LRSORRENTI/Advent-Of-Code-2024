<?php

// Read the input from 'input.txt'
$input = file_get_contents('input.txt');

// Regular expression to match valid mul(X,Y) instructions
$pattern = '/mul\(([0-9]{1,3}),([0-9]{1,3})\)/';

// Find all matches in the input string
preg_match_all($pattern, $input, $matches);

// Initialize total sum
$total = 0;

// Loop through all matches and calculate the sum of products
for ($i = 0; $i < count($matches[0]); $i++) {
    $x = intval($matches[1][$i]);
    $y = intval($matches[2][$i]);
    $total += $x * $y;
}

// Output the total sum
echo "$total\n";

?>