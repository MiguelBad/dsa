from time import monotonic
from random import randint


def selection_sort(arr):
    for i in range(len(arr) - 1):
        min_idx = i

        for j in range(i + 1, len(arr)):
            if arr[min_idx] > arr[i]:
                min_idx = j

        arr[min_idx], arr[i] = arr[i], arr[min_idx]


def insertion_sort(arr):
    for i in range(1, len(arr)):
        key = arr[i]
        j = i - 1

        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
            arr[j + 1] = key


def bubble_sort(arr):
    for i in range(len(arr)):
        for j in range(len(arr) - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]


def quick_sort(arr, low, high):
    if low >= high:
        return

    def partition(arr, low, high):
        idx = low - 1

        for i in range(low, high):
            if arr[i] < arr[high]:
                idx += 1
                arr[idx], arr[i] = arr[i], arr[idx]

        idx += 1
        arr[idx], arr[high] = arr[high], arr[idx]
        return idx

    partition_idx = partition(arr, low, high)
    quick_sort(arr, low, partition_idx - 1)
    quick_sort(arr, partition_idx + 1, high)


def merge_sort(arr):
    if len(arr) < 2:
        return

    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]

    merge_sort(left)
    merge_sort(right)

    i = j = k = 0

    while len(left) > i and len(right) > j:
        if left[i] <= right[j]:
            arr[k] = left[i]
            i += 1
        else:
            arr[k] = right[j]
            j += 1

        k += 1

    while len(left) > i:
        arr[k] = left[i]
        i += 1
        k += 1

    while len(right) > j:
        arr[k] = right[j]
        j += 1
        k += 1

    if j == 250:
        meow = 0
        for i in arr:
            meow += i
        print(meow)

    if j == 500:
        meow = 0
        for i in arr:
            meow += i

        print(meow)


def main():
    arr_1000 = [randint(1, 1000) for i in range(1000)]

    arr = {
        "bubble": arr_1000.copy(),
        "insertion": arr_1000.copy(),
        "selection": arr_1000.copy(),
        "merge": arr_1000.copy(),
        "quick": arr_1000.copy(),
    }

    m_start = monotonic()
    merge_sort(arr["merge"])
    m_end = monotonic()
    print(f"Merge sort time: {m_end - m_start} seconds")

    q_start = monotonic()
    quick_sort(arr["quick"], 0, len(arr["quick"]) - 1)
    q_end = monotonic()
    print(f"Quick sort time: {q_end - q_start} seconds")

    b_start = monotonic()
    bubble_sort(arr["bubble"])
    b_end = monotonic()
    print(f"Bubble sort time: {b_end - b_start} seconds")

    i_start = monotonic()
    insertion_sort(arr["insertion"])
    i_end = monotonic()
    print(f"Insertion sort time: {i_end - i_start} seconds")

    s_start = monotonic()
    selection_sort(arr["selection"])
    s_end = monotonic()
    print(f"Selection sort time: {s_end - s_start} seconds")


main()
