"""
Problem Statement:

Write a function called bubbleSort that takes an array of numbers as an input and returns a new array that contains the same numbers sorted in ascending order using the bubble sort algorithm.

Input:

    An array of numbers, for example: [64, 34, 25, 12, 22, 11, 90]

Output:

    A new array with the numbers sorted in ascending order, for example: [11, 12, 22, 25, 34, 64, 90]


PSUEDOCODE

n = array length

for i in n do
    for j in range i +
"""


def main(arr):
    n = len(arr)

    for i in range(n):
        for j in range(n - i - 1):
            if i > j:
                temp = arr[i]
                arr[i] = arr[j]
                arr[j + 1] = temp
    print(arr)


arr = [64, 34, 25, 12, 22, 11, 90]
main(arr)
