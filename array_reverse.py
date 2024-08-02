import time
"""
def main(arr, start, stop):
    if start < stop - 1:
        arr[start], arr[stop - 1] = arr[stop - 1], arr[start]
        main(arr, start + 1, stop - 1)
"""


def main(arr):
    j = len(arr) - 1
    for i in range(len(arr) // 2):
        arr[i], arr[j] = arr[j], arr[i]
        j -= 1


numbers = [3, 7, 12, 18, 21, 26, 34, 45, 56, 62, 74, 88, 99]
start = 0
stop = len(numbers)
# main(numbers, start, stop)
a = time.time()
main(numbers)
b = time.time()

c = time.time()
numbers = numbers[::-1]
d = time.time()

print(numbers)
print('first', b - a)
print('second', d - c)
