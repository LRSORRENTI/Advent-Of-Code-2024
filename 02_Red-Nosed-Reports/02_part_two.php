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

    // First, check if the report is already safe
    if (isReportSafe($levels)) {
        $safeReportCount++;
        continue;
    }

    // If not safe, try removing each level one at a time
    $isSafeAfterRemoval = false;
    for ($i = 0; $i < count($levels); $i++) {
        // Create a copy of the levels without the ith level
        $modifiedLevels = $levels;
        array_splice($modifiedLevels, $i, 1);

        // Check if the modified report is safe
        if (count($modifiedLevels) >= 2 && isReportSafe($modifiedLevels)) {
            $isSafeAfterRemoval = true;
            break; // No need to check further removals
        }
    }

    if ($isSafeAfterRemoval) {
        $safeReportCount++;
    }

    // Optional Debugging Output
    /*
    echo "Report #" . ($lineNumber + 1) . ": " . implode(' ', $levels) . "\n";
    echo $isSafeAfterRemoval ? "Safe after removing a level\n\n" : "Unsafe\n\n";
    */
}

// Output the total number of safe reports
echo "Total Safe Reports: $safeReportCount\n";

// Function to check if a report is safe
function isReportSafe($levels)
{
    $direction = null; // 'increasing' or 'decreasing'

    for ($i = 0; $i < count($levels) - 1; $i++) {
        $current = $levels[$i];
        $next = $levels[$i + 1];
        $diff = $next - $current;

        // Check for zero difference
        if ($diff == 0) {
            // Levels must differ by at least 1
            return false;
        }

        // Check if difference is between -3 and -1 or 1 and 3
        if (abs($diff) < 1 || abs($diff) > 3) {
            return false;
        }

        // Determine direction on the first comparison
        if ($i == 0) {
            $direction = $diff > 0 ? 'increasing' : 'decreasing';
        } else {
            // Check for consistent direction
            if (($direction == 'increasing' && $diff < 0) || ($direction == 'decreasing' && $diff > 0)) {
                return false;
            }
        }
    }

    // If all checks pass, the report is safe
    return true;
}
