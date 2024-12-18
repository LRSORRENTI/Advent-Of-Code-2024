import * as fs from 'fs';

// Read the input from a file
const input = fs.readFileSync('input.txt', 'utf-8').trim().split('\n');

// Initialize Registers
let A = 0;
let B = 0;
let C = 0;

// Parse input for register values and program instructions
for (const line of input) {
  if (line.startsWith("Register A:")) A = parseInt(line.split(":")[1].trim());
  if (line.startsWith("Register B:")) B = parseInt(line.split(":")[1].trim());
  if (line.startsWith("Register C:")) C = parseInt(line.split(":")[1].trim());
}

// Extract the program as an array of numbers
const programLine = input.find(line => line.startsWith("Program:"));
const program = programLine
  ? programLine.split(":")[1].trim().split(",").map(Number)
  : [];

// Execute program
function runProgram(): string {
  const output: number[] = [];
  let instructionPointer = 0;

  while (instructionPointer < program.length) {
    const opcode = program[instructionPointer];
    const operand = program[instructionPointer + 1];
    const powerOf2 = (exp: number) => Math.pow(2, exp);

    switch (opcode) {
      case 0: // adv: A = A / 2^operand
        A = Math.trunc(A / powerOf2(operand < 4 ? operand : [A, B, C][operand - 4]));
        break;
      case 1: // bxl: B = B ^ operand (literal)
        B = B ^ operand;
        break;
      case 2: // bst: B = operand % 8
        B = (operand < 4 ? operand : [A, B, C][operand - 4]) % 8;
        break;
      case 3: // jnz: Jump if A != 0
        if (A !== 0) {
          instructionPointer = operand;
          continue;
        }
        break;
      case 4: // bxc: B = B ^ C
        B = B ^ C;
        break;
      case 5: // out: Output operand % 8
        output.push((operand < 4 ? operand : [A, B, C][operand - 4]) % 8);
        break;
      case 6: // bdv: B = A / 2^operand
        B = Math.trunc(A / powerOf2(operand < 4 ? operand : [A, B, C][operand - 4]));
        break;
      case 7: // cdv: C = A / 2^operand
        C = Math.trunc(A / powerOf2(operand < 4 ? operand : [A, B, C][operand - 4]));
        break;
      default:
        throw new Error(`Unknown opcode: ${opcode}`);
    }
    instructionPointer += 2; // Move to the next instruction
  }

  return output.join(",");
}

// Run the program and print the output
const finalOutput = runProgram();
console.log(`Output: ${finalOutput}`);
