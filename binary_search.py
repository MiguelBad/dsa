def main(arr, n):
    left = 0
    right = len(arr) - 1

    while left > right:
        mid = left + ((right - left) // 2)
        if n == arr[mid]:
            return mid
        elif n > arr[mid]:
            left = mid + 1
        else:
            right = mid - 1

    return -1


# def main(arr, n, left, right):
#     mid = left + ((right - left) // 2)
#
#     if left > right:
#         return -1
#
#     if n == arr[mid]:
#         return mid
#     elif n > arr[mid]:
#         return main(arr, n, mid + 1, right)
#     else:
#         return main(arr, n, left, mid - 1)


numbers = [3, 7, 12, 18, 21, 26, 34, 45, 56, 62, 74, 88, 99]
target = 10
left = 0
right = len(numbers) - 1

# result = main(numbers, target, left, right)
result = main(numbers, target)
print(result)
print(numbers.index(target))
