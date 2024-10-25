def sortByColours(arr: list[int]) -> None:
    low, mid, high = 0, 0, len(arr) - 1

    while mid <= high:

        # swaps current low to current mid
        if arr[mid] == 0:
            arr[low], arr[mid] = arr[mid], arr[low]
            low += 1
            mid += 1

        elif arr[mid] == 1:
            mid += 1

        # swaps current high to current mid
        elif arr[mid] == 2:
            arr[mid], arr[high] = arr[high], arr[mid]
            high -= 1


numColour: list[int] = [0, 0, 1, 1, 2, 2, 2, 1, 2, 2, 1, 0, 0, 0, 2, 0, 0, 2, 1, 0, 1, 1, 1, 2, 2, 0, 1, 1, 0, 2, 1, 1, 1, 2, 2, 2, 0, 2, 1, 0, 2, 2, 2, 1, 1, 1, 2, 0, 1, 0, 0, 0, 2, 1, 1, 2, 2, 0, 1, 2, 0, 0, 0, 2, 1, 2, 0, 1, 2, 0, 1, 0, 0, 1, 2, 0, 0, 2, 2, 2, 0, 2, 0, 2, 2, 0, 0, 1, 0, 0, 1, 2, 0, 0, 1, 1, 2, 0, 2, 2]
sortByColours(numColour)
print(numColour)
