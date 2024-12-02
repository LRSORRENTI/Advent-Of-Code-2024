<?php
// Read the input file
$inputFile = 'input.txt';
$lines = file($inputFile, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);

// Initialize the safe report counter
$safeReportCount = 0;

// Process each report
foreach ($lines as $lineNumber => $line) {
    // Parse the report into an array of integers
    $levels = preg_split('/\s+/', trim($line));
    $levels = array_map('intval', $levels);

    // Skip empty lines or reports with less than 2 levels
    if (count($levels) < 2) {
        continue;
    }

    $isSafe = true;
    $differences = [];
    $direction = null; // 'increasing' or 'decreasing'

    for ($i = 0; $i < count($levels) - 1; $i++) {
        $current = $levels[$i];
        $next = $levels[$i + 1];
        $diff = $next - $current;

        // Check for zero difference
        if ($diff == 0) {
            // Levels must differ by at least 1
            $isSafe = false;
            break;
        }

        // Check if difference is between -3 and -1 or 1 and 3
        if (abs($diff) < 1 || abs($diff) > 3) {
            $isSafe = false;
            break;
        }

        // Determine direction on the first comparison
        if ($i == 0) {
            $direction = $diff > 0 ? 'increasing' : 'decreasing';
        } else {
            // Check for consistent direction
            if (($direction == 'increasing' && $diff < 0) || ($direction == 'decreasing' && $diff > 0)) {
                $isSafe = false;
                break;
            }
        }
    }

    // Increment the safe report count if the report is safe
    if ($isSafe) {
        $safeReportCount++;
    }

    // Debugging Output (Optional)
    /*
    echo "Report #" . ($lineNumber + 1) . ": " . implode(' ', $levels) . "\n";
    echo $isSafe ? "Safe\n\n" : "Unsafe\n\n";
    */
}

// Output the total number of safe reports
echo "Total Safe Reports: $safeReportCount\n";
