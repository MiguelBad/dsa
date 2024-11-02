def fib_iterative(n):
    # where n is the nth fib
    if n < 0:
        print("n must be greater than 0")
    elif n == 0:
        return 0
    elif n == 1:
        return 1

    low = 0
    high = 1
    result = 0
    for _ in range(n - 1):
        result = low + high
        low = high
        high = result

    return result


def fib_recursive(n):
    temp = {}

    def fib_helper(n, temp) -> int:
        if n in temp:
            return temp[n]

        if n == 0:
            return 0
        if n == 1:
            return 1

        result = fib_helper(n - 1, temp) + fib_helper(n - 2, temp)
        temp[n] = result
        return result

    return fib_helper(n, temp)


print(fib_recursive(10))
