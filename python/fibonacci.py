def fib_iterative(n, low, high, fib):
    if n == 0:
        return 0
    elif n == 1:
        return 1

    for _ in range(1, n):
        fib = low + high
        temp = high
        high = low + high
        low = temp

    return fib


def fib_list_recursion(n):
    if n == 0:
        return [0]
    elif n == 1:
        return [0, 1]
    else:
        fib = fib_list_recursion(n - 1)
        fib.append(fib[-1] + fib[-2])
        return fib


print(fib_iterative(20, 0, 1, 0))
print(fib_list_recursion(20))
