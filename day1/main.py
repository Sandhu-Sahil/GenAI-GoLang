from time import time

def sumation(n):
    total = 0
    for i in range(n+1):
        total += i
    return total

def main():
    n=1000000000
    st = time()
    print(sumation(n))
    et = time()
    print("Time taken: ", et-st, "sec")

if __name__ == "__main__":
    main()