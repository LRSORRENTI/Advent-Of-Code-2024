import * as fs from 'fs';
import * as path from 'path';

interface State {
  x: number;
  y: number;
  direction: number; // 0: East, 1: South, 2: West, 3: North
  score: number;
}

const directions = [
  { dx: 1, dy: 0 },  // East
  { dx: 0, dy: 1 },  // South
  { dx: -1, dy: 0 }, // West
  { dx: 0, dy: -1 }, // North
];

const reindeerMaze = (input: string[]): number => {
  const rows = input.length;
  const cols = input[0].length;
  let start: { x: number; y: number } | null = null;
  let end: { x: number; y: number } | null = null;

  // Find the start (S) and end (E) positions
  for (let y = 0; y < rows; y++) {
    for (let x = 0; x < cols; x++) {
      if (input[y][x] === "S") start = { x, y };
      if (input[y][x] === "E") end = { x, y };
    }
  }

  if (!start || !end) throw new Error("Start or End not found");

  const visited = new Set<string>();
  const queue: State[] = [{ x: start.x, y: start.y, direction: 0, score: 0 }];

  const encodeState = (state: State) =>
    `${state.x},${state.y},${state.direction}`;

  while (queue.length > 0) {
    // Get the state with the smallest score
    queue.sort((a, b) => a.score - b.score); // Priority queue
    const current = queue.shift()!;
    const key = encodeState(current);

    // Skip if we've already visited this state
    if (visited.has(key)) continue;
    visited.add(key);

    // Check if we reached the end
    if (current.x === end.x && current.y === end.y) return current.score;

    // Try all possible moves
    for (let i = 0; i < 4; i++) {
      const nextDirection = i;
      const rotationCost = Math.min(
        Math.abs(nextDirection - current.direction),
        4 - Math.abs(nextDirection - current.direction)
      ) * 1000;

      const nextScore = current.score + rotationCost;

      // Move forward
      const { dx, dy } = directions[nextDirection];
      const nx = current.x + dx;
      const ny = current.y + dy;

      if (nx >= 0 && nx < cols && ny >= 0 && ny < rows && input[ny][nx] !== "#") {
        queue.push({ x: nx, y: ny, direction: nextDirection, score: nextScore + 1 });
      }
    }
  }

  return -1; // If no path is found
};

// Main Function to Read File Input
const main = () => {
  const filePath = process.argv[2];
  if (!filePath) {
    console.error("Please provide the path to the input file.");
    process.exit(1);
  }

  const absolutePath = path.resolve(filePath);
  try {
    const input = fs.readFileSync(absolutePath, 'utf8').split('\n').map(line => line.trim());
    const result = reindeerMaze(input);
    console.log(`Minimum Score: ${result}`);
  } catch (error: any) {
    console.error(`Error reading file: ${error.message}`);
    process.exit(1);
  }
};

// Run the Main Function
main();
