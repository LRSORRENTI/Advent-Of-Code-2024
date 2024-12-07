from itertools import product

def evaluate_expression(nums, ops):
    """Evaluate the expression defined by nums and ops strictly left-to-right.
       ops can include '+', '*', '||'."""
    val = nums[0]
    for i, op in enumerate(ops, start=1):
        if op == '+':
            val = val + nums[i]
        elif op == '*':
            val = val * nums[i]
        else:  # op == '||'
            # Concatenate digits of val and nums[i]
            val = int(str(val) + str(nums[i]))
    return val

def can_form_target(target, nums):
    """Check if we can form the target by placing '+', '*', or '||' 
       between the given nums."""
    n = len(nums)
    if n == 1:
        return nums[0] == target
    
    # For each position between numbers, we can have '+', '*', or '||'
    for ops in product(['+', '*', '||'], repeat=n-1):
        val = evaluate_expression(nums, ops)
        if val == target:
            return True
    return False

if __name__ == "__main__":
    total = 0
    # Read from input.txt in the same directory
    with open("input.txt", "r") as f:
        for line in f:
            line = line.strip()
            if not line:
                continue
            # Format: "<target>: <num1> <num2> ..."
            target_part, nums_part = line.split(':')
            target = int(target_part.strip())
            nums = list(map(int, nums_part.strip().split()))
            
            if can_form_target(target, nums):
                total += target
    
    print(total)
