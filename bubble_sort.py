"""
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
