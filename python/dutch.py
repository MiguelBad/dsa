def sortByColours(arr: list[int]) -> None:
    """
    Sorts a list of integers representing a colour.
        - 0 == black
        - 1 == red
        - 2 == yellow

    This sort only itterates over the array once, moving number 0 to the front part
    of the array and move the number 2 to the end part of the array. This gives
    the sort a time complexity of O(n).

    Assumptions:
    - The input list contains only integers 0, 1, and 2.

    """

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
