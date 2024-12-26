import sys
import re
from collections import defaultdict

def parse_input(filename):
    initial_values = {}
    gates = []

    with open(filename, "r") as file:
        lines = file.read().splitlines()

    for line in lines:
        if ":" in line:  # Initial values
            wire, value = line.split(": ")
            initial_values[wire] = int(value)
        elif "->" in line:  # Gates
            gates.append(line)

    return initial_values, gates

def evaluate_gates(initial_values, gates):
    values = initial_values.copy()
    dependencies = defaultdict(list)
    ready_wires = set(initial_values.keys())

    # Parse gates and track dependencies
    parsed_gates = []
    for gate in gates:
        match = re.match(r"(\w+)\s+(AND|OR|XOR)\s+(\w+)\s+->\s+(\w+)", gate)
        if match:
            left, op, right, out = match.groups()
            parsed_gates.append((left, op, right, out))
            if left not in values:
                dependencies[left].append(out)
            if right not in values:
                dependencies[right].append(out)
        else:
            raise ValueError(f"Invalid gate format: {gate}")

    while parsed_gates:
        for gate in parsed_gates[:]:
            left, op, right, out = gate

            if left in values and right in values:
                if op == "AND":
                    values[out] = values[left] & values[right]
                elif op == "OR":
                    values[out] = values[left] | values[right]
                elif op == "XOR":
                    values[out] = values[left] ^ values[right]

                ready_wires.add(out)
                parsed_gates.remove(gate)

                # Propagate values to dependent wires
                for dependent in dependencies[out]:
                    if dependent not in ready_wires:
                        ready_wires.add(dependent)

    return values

def calculate_output(values):
    # Combine the bits from all wires starting with 'z'
    z_wires = {k: v for k, v in values.items() if k.startswith('z')}
    binary_representation = "".join(
        str(z_wires[f"z{i:02}"]) for i in range(len(z_wires))
    )[::-1]  # Reverse to match significance order

    return int(binary_representation, 2)

def main():
    if len(sys.argv) != 2:
        print("Usage: python script.py <input.txt>")
        sys.exit(1)

    filename = sys.argv[1]
    initial_values, gates = parse_input(filename)
    values = evaluate_gates(initial_values, gates)
    result = calculate_output(values)
    print(result)

if __name__ == "__main__":
    main()
