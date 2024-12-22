MOD = 16777216  # 2^24

def next_secret(s: int) -> int:
    """
    Part One's rule to evolve the current secret into the next secret:
      1. Multiply by 64, XOR with s, then mod 16,777,216
      2. Integer-divide s by 32, XOR with s, then mod 16,777,216
      3. Multiply by 2048, XOR with s, then mod 16,777,216
    """
    # Step 1
    val1 = s * 64
    s = (s ^ val1) % MOD
    
    # Step 2
    val2 = s // 32
    s = (s ^ val2) % MOD
    
    # Step 3
    val3 = s * 2048
    s = (s ^ val3) % MOD

    return s

def generate_prices(initial_secret: int, count: int = 2000) -> list[int]:
    """
    Return a list of `count+1` prices:
      - The first price is the ones digit of `initial_secret`
      - Then generate `count` more secrets (and prices),
        each time applying next_secret(...).
    """
    prices = []
    s = initial_secret
    # The "initial" price (before any evolutions)
    prices.append(s % 10)
    # Generate 'count' more
    for _ in range(count):
        s = next_secret(s)
        prices.append(s % 10)
    return prices  # length = count + 1

def find_first_occurrences(prices: list[int]) -> dict[tuple[int,int,int,int], int]:
    """
    Given a buyer's sequence of prices (length 2001),
    compute the 2000 consecutive changes and find the
    earliest occurrence of each 4-change pattern.

    Returns a dictionary:
      pattern_of_4_changes -> earliest index i
    meaning the earliest index where that pattern starts.
    """
    earliest = {}
    # Generate the changes (length = len(prices) - 1)
    changes = []
    for i in range(len(prices) - 1):
        changes.append(prices[i+1] - prices[i])  # difference in ones digits

    # For each sliding window of size 4 in changes
    for i in range(len(changes) - 3):
        pattern = (changes[i], changes[i+1], changes[i+2], changes[i+3])
        if pattern not in earliest:
            earliest[pattern] = i

    return earliest

def main():
    import sys
    # Read all initial secrets from stdin
    initial_secrets = []
    for line in sys.stdin:
        line = line.strip()
        if line:
            initial_secrets.append(int(line))

    # Global dictionary: pattern -> total bananas
    from collections import defaultdict
    sum_for_pattern = defaultdict(int)

    # Process each buyer
    for secret in initial_secrets:
        # Generate the buyer's first 2001 prices
        prices = generate_prices(secret, count=2000)
        # Find earliest occurrence of each 4-change pattern
        earliest_map = find_first_occurrences(prices)
        # For each pattern, add the sale price
        for pattern, idx in earliest_map.items():
            # The sell price is prices[idx+4] (since the pattern covers changes at idx..idx+3)
            sale_price = prices[idx+4]
            sum_for_pattern[pattern] += sale_price

    # Find the pattern with the maximum total bananas
    best_total = max(sum_for_pattern.values()) if sum_for_pattern else 0

    print(best_total)

if __name__ == "__main__":
    main()
