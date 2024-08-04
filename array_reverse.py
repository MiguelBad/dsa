def reverse_recursive(arr, start, stop):
    if start < stop - 1:
        arr[start], arr[stop - 1] = arr[stop - 1], arr[start]
        reverse_recursive(arr, start + 1, stop - 1)


numbers = [3, 7, 12, 18, 21, 26, 34, 45, 56, 62, 74, 88, 99]
start = 0
stop = len(numbers)
print(reverse_recursive(numbers, start, stop))


def reverse_iterative(arr):
    j = len(arr) - 1
    for i in range(len(arr) // 2):
        arr[i], arr[j] = arr[j], arr[i]
        j -= 1


print(reverse_iterative(numbers))
