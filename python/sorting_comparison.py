from time import monotonic
from random import randint


def selection_sort(arr: list[int]) -> None:
    for i in range(len(arr) - 1):
        min_idx = i

        for j in range(i + 1, len(arr)):
            if arr[min_idx] > arr[j]:
                min_idx = j

        arr[min_idx], arr[i] = arr[i], arr[min_idx]


def insertion_sort(arr: list[int]) -> None:
    for i in range(1, len(arr)):
        key = arr[i]
        j = i - 1

        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1

        arr[j + 1] = key


def bubble_sort(arr: list[int]) -> None:
    for i in range(len(arr)):
        swapped = False

        for j in range(len(arr) - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                swapped = True

        if not swapped:
            break


def quick_sort(arr: list[int], low: int, high: int) -> None:

    # base case
    if low >= high:
        return

    # selecting the last element as pivot
    pivot_idx = low - 1

    # move items in the front of the array if item is less than pivot
    for i in range(low, high):
        if arr[i] < arr[high]:
            pivot_idx += 1
            arr[i], arr[pivot_idx] = arr[pivot_idx], arr[i]

    # swap pivot with the first element that is greater than pivot
    pivot_idx += 1
    arr[pivot_idx], arr[high] = arr[high], arr[pivot_idx]

    # divide the array based on pivot idx then recurse
    quick_sort(arr, low, pivot_idx - 1)
    quick_sort(arr, pivot_idx + 1, high)


def merge_sort(arr: list[int]) -> None:

    # base case
    if len(arr) < 2:
        return

    # divide array into two
    mid = len(arr) // 2
    left = arr[:mid]
    right = arr[mid:]

    # recurse to every element in array into their own single array
    merge_sort(left)
    merge_sort(right)

    # counter
    i = j = k = 0

    # perform loop until either left or array is "empty"
    while i < len(left) and j < len(right):

        # push the lower element in the array
        if left[i] < right[j]:
            arr[k] = left[i]
            i += 1
        else:
            arr[k] = right[j]
            j += 1

        k += 1

    # push remaining items in the array
    while len(left) > i:
        arr[k] = left[i]
        i += 1
        k += 1

    # push remaining items in the array
    while len(right) > j:
        arr[k] = right[j]
        j += 1
        k += 1


def do_sorting(
    m_time: list[float],
    q_time: list[float],
    t_time: list[float],
    b_time: list[float],
    i_time: list[float],
    s_time: list[float],
    item_to_sort: int,
) -> None:

    arr_to_sort = [randint(1, 1000) for _ in range(item_to_sort)]

    arr = {
        "bubble": arr_to_sort.copy(),
        "insertion": arr_to_sort.copy(),
        "selection": arr_to_sort.copy(),
        "merge": arr_to_sort.copy(),
        "quick": arr_to_sort.copy(),
        "tim": arr_to_sort.copy(),
    }

    m_start = monotonic()
    merge_sort(arr["merge"])
    m_end = monotonic()
    m_time.append(m_end - m_start)

    q_start = monotonic()
    quick_sort(arr["quick"], 0, len(arr["quick"]) - 1)
    q_end = monotonic()
    q_time.append(q_end - q_start)

    t_start = monotonic()
    sorted(arr["tim"])
    t_end = monotonic()
    t_time.append(t_end - t_start)

    # b_start = monotonic()
    # bubble_sort(arr["bubble"])
    # b_end = monotonic()
    # b_time.append(b_end - b_start)

    # i_start = monotonic()
    # insertion_sort(arr["insertion"])
    # i_end = monotonic()
    # i_time.append(i_end - i_start)

    # s_start = monotonic()
    # selection_sort(arr["selection"])
    # s_end = monotonic()
    # s_time.append(s_end - s_start)


def main() -> None:
    m_time, q_time, t_sort, b_time, i_time, s_time = [], [], [], [], [], []
    loop_for = 500
    item_to_sort = 500_000

    for _ in range(loop_for):
        do_sorting(m_time, q_time, t_sort, b_time, i_time, s_time, item_to_sort)

    print(f"Number of items to sort: {item_to_sort}\nRepeat for {loop_for} times\n")
    print("Average time:")
    # print(f"\t* Bubble sort: {sum(x for x in b_time)/loop_for:.6f} seconds")
    # print(f"\t* Insertion sort: {sum(x for x in i_time)/loop_for:.6f} seconds")
    # print(f"\t* Selection sort: {sum(x for x in s_time)/loop_for:.6f} seconds")
    print(f"\t* Merge sort: {sum(x for x in m_time)/loop_for:.6f} seconds")
    print(f"\t* Quick sort: {sum(x for x in q_time)/loop_for:.6f} seconds")
    print(f"\t* Python's Tim sort: {sum(x for x in t_sort)/loop_for:.6f} seconds")


if __name__ == "__main__":
    main()
