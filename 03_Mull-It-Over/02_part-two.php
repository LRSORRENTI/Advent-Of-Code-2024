<?php

// Read the input from 'input.txt'
$input = file_get_contents('input.txt');

// Regular expressions to match the instructions
$mul_pattern = '/mul\(([0-9]{1,3}),([0-9]{1,3})\)/';
$do_pattern = '/do\(\)/';
$dont_pattern = '/don\'t\(\)/';

// Initialize total sum and the state (enabled by default)
$total = 0;
$enabled = true;

// Position tracker for regex matching
$offset = 0;

while ($offset < strlen($input)) {
    // Find the next instruction (mul, do, or don't)
    $patterns = [
        'mul' => '/mul\(([0-9]{1,3}),([0-9]{1,3})\)/',
        'do' => '/do\(\)/',
        'dont' => '/don\'t\(\)/',
    ];

    $nextMatch = null;
    $nextType = null;
    $nextPos = strlen($input);

    // Find the earliest next instruction
    foreach ($patterns as $type => $pattern) {
        if (preg_match($pattern, $input, $match, PREG_OFFSET_CAPTURE, $offset)) {
            $pos = $match[0][1];
            if ($pos < $nextPos) {
                $nextPos = $pos;
                $nextMatch = $match;
                $nextType = $type;
            }
        }
    }

    // If no more instructions are found, break the loop
    if ($nextMatch === null) {
        break;
    }

    // Move the offset to the end of the matched instruction
    $offset = $nextPos + strlen($nextMatch[0][0]);

    // Handle the instruction based on its type
    switch ($nextType) {
        case 'mul':
            if ($enabled) {
                $x = intval($nextMatch[1][0]);
                $y = intval($nextMatch[2][0]);
                $total += $x * $y;
            }
            break;
        case 'do':
            $enabled = true;
            break;
        case 'dont':
            $enabled = false;
            break;
    }
}

// Output the total sum
echo "$total\n";

?>