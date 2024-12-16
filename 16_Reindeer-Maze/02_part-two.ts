import * as fs from 'fs';
import * as path from 'path';

type Point = any;

const RIGHT: Point = [0, 1];
const DOWN: Point = [1, 0];
const LEFT: Point = [0, -1];
const UP: Point = [-1, 0];

const DIRECTIONS: Point[] = [RIGHT, DOWN, LEFT, UP];

function parse(input: string): [number, number] {
    // Create a 2D grid from input string
    const grid = input.trim().split('\n').map(row => row.split(''));
    
    // Find start and end positions
    const start = findPosition(grid, 'S');
    const end = findPosition(grid, 'E');
    
    if (!start || !end) {
        throw new Error('Start or end position not found');
    }

    // Initialize data structures
    const buckets: [Point, number][][] = Array.from({ length: 1001 }, () => []);
    const seen: number[][][] = Array.from({ length: grid.length }, () => 
        Array.from({ length: grid[0].length }, () => new Array(4).fill(Number.MAX_SAFE_INTEGER))
    );
    
    let cost = 0;
    let remaining = 1;
    let lowest = Number.MAX_SAFE_INTEGER;
    
    // Initial state
    buckets[0].push([start, 0]);
    seen[start[0]][start[1]][0] = 0;
    
    while (remaining > 0) {
        const currentBucket = buckets[cost % 1001];
        
        while (currentBucket.length > 0) {
            const [position, direction] = currentBucket.pop()!;
            remaining--;
            
            // Check if reached end
            if (position[0] === end[0] && position[1] === end[1]) {
                lowest = Math.min(lowest, cost);
                continue;
            }
            
            const left = (direction + 3) % 4;
            const right = (direction + 1) % 4;
            
            const nextMoves = [
                [addPoints(position, DIRECTIONS[direction]), direction, cost + 1],
                [position, left, cost + 1000],
                [position, right, cost + 1000]
            ];
            
            for (const [nextPosition, nextDirection, nextCost] of nextMoves) {
                if (
                    isValidMove(grid, nextPosition) && 
                    nextCost < seen[nextPosition[0]][nextPosition[1]][nextDirection]
                ) {
                    buckets[(nextCost % 1001)].push([nextPosition, nextDirection]);
                    seen[nextPosition[0]][nextPosition[1]][nextDirection] = nextCost;
                    remaining++;
                }
            }
        }
        
        cost++;
    }
    
    // Backward path reconstruction
    const todo: [Point, number, number][] = [];
    const path = Array.from({ length: grid.length }, () => 
        new Array(grid[0].length).fill(false)
    );
    
    for (let direction = 0; direction < 4; direction++) {
        if (seen[end[0]][end[1]][direction] === lowest) {
            todo.push([end, direction, lowest]);
        }
    }
    
    while (todo.length > 0) {
        const [position, direction, currentCost] = todo.pop()!;
        path[position[0]][position[1]] = true;
        
        if (position[0] === start[0] && position[1] === start[1]) {
            continue;
        }
        
        const left = (direction + 3) % 4;
        const right = (direction + 1) % 4;
        
        const nextMoves = [
            [subtractPoints(position, DIRECTIONS[direction]), direction, currentCost - 1],
            [position, left, currentCost - 1000],
            [position, right, currentCost - 1000]
        ];
        
        for (const [nextPosition, nextDirection, nextCost] of nextMoves) {
            if (nextCost === seen[nextPosition[0]][nextPosition[1]][nextDirection]) {
                todo.push([nextPosition, nextDirection, nextCost]);
                seen[nextPosition[0]][nextPosition[1]][nextDirection] = Number.MAX_SAFE_INTEGER;
            }
        }
    }
    
    return [
        lowest, 
        path.flat().filter(Boolean).length
    ];
}

// Utility functions
function findPosition(grid: string[][], target: string): Point | null {
    for (let i = 0; i < grid.length; i++) {
        for (let j = 0; j < grid[i].length; j++) {
            if (grid[i][j] === target) return [i, j];
        }
    }
    return null;
}

function addPoints(a: Point, b: Point): Point {
    return [a[0] + b[0], a[1] + b[1]];
}

function subtractPoints(a: Point, b: Point): Point {
    return [a[0] - b[0], a[1] - b[1]];
}

function isValidMove(grid: string[][], point: Point): boolean {
    return (
        point[0] >= 0 && 
        point[0] < grid.length && 
        point[1] >= 0 && 
        point[1] < grid[0].length && 
        grid[point[0]][point[1]] !== '#'
    );
}


function part2(input: ReturnType<typeof parse>): number {
    return input[1];
}

// Main function to handle file input and run the solution
function main() {
    // Get the input file path from command line arguments
    const inputFilePath = process.argv[2];

    if (!inputFilePath) {
        console.error('Please provide an input file path');
        process.exit(1);
    }

    try {
        // Read the input file
        const inputContent = fs.readFileSync(path.resolve(inputFilePath), 'utf-8');

        // Parse the input
        const result = parse(inputContent);

        // Output results

        console.log('Part 2:', part2(result));
    } catch (error) {
        console.error('Error reading or processing the input file:', error);
        process.exit(1);
    }
}

// Run the main function
main();