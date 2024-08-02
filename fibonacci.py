"""
list = []
current = 0

for i:=0 to n do

    if current < 1 then
        current += 1
    else:
        current += list[-1]

    append current to list

fib = [0]
current = 0

for i in range(20):
    if current < 1:
        current += 1
    else:
        current += fib[-2]

    fib.append(current)

print(fib)

===============================

0, 1, 1, 2, 3, 5, 8, 13, 21
"""


def main(n):
    if n < 2:
        return [0, 1]
    else:
        fib = main(n - 1)
        fib.append(fib[-1] + fib[-2])
        return fib


print(main(20))
