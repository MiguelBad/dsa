def partition(arr, low, high):
    pivot = arr[high]
    idx = low - 1

    for i in range(low, high):
        if arr[i] <= pivot:
            idx += 1
            temp = arr[i]
            arr[i] = arr[idx]
            arr[idx] = temp

    idx += 1
    arr[high] = arr[idx]
    arr[idx] = pivot

    return idx


def quick_sort(arr, low, high):
    if low >= high:
        return

    pivot_idx = partition(arr, low, high)

    quick_sort(arr, low, pivot_idx - 1)
    quick_sort(arr, pivot_idx + 1, high)


def main():
    unsorted_list = [
        42, 17, 23, 56, 9, 33, 89, 5, 78, 61, 102, 47, 3, 88, 24, 75, 11, 94,
        28, 36, 66, 83, 12, 57, 91, 64, 19, 85, 48, 14, 73, 27, 30, 54, 68, 7,
        99, 38, 41, 15, 81, 93, 21, 60, 6, 70, 18, 95, 31, 53, 76, 8, 32, 50,
        44, 22, 87, 13, 62, 39, 80, 20, 46, 4, 98, 74, 35, 58, 100, 67, 25, 16,
        77, 34, 92, 49, 29, 84, 10, 40, 82, 55, 63, 2, 45, 86, 26, 97, 37, 72,
        59, 1, 65, 90, 79, 52, 43, 71, 51, 96
    ]
    quick_sort(unsorted_list, 0, len(unsorted_list) - 1)
    print(unsorted_list)


if __name__ == '__main__':
    main()
