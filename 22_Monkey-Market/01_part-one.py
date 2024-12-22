MOD = 16777216  # 2^24

def next_secret(s: int) -> int:
    # Step 1: multiply by 64, XOR, then prune
    val1 = s * 64
    s = (s ^ val1) % MOD
    
    # Step 2: integer divide by 32, XOR, then prune
    val2 = s // 32
    s = (s ^ val2) % MOD
    
    # Step 3: multiply by 2048, XOR, then prune
    val3 = s * 2048
    s = (s ^ val3) % MOD
    
    return s

def get_2000th_secret(initial_secret: int) -> int:
    s = initial_secret
    for _ in range(2000):
        s = next_secret(s)
    return s

def main():
    import sys
    total = 0
    for line in sys.stdin:
        line = line.strip()
        if not line:
            continue
        initial_secret = int(line)
        secret_2000 = get_2000th_secret(initial_secret)
        total += secret_2000

    print(total)

if __name__ == "__main__":
    main()
