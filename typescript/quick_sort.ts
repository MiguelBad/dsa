function partition(arr: number[], low: number, high: number): number {
    const pivot: number = arr[high];
    let index: number = low - 1;

    for (let i = low; i < high; i++) {
        if (arr[i] <= pivot) {
            index++;
            const temp = arr[i];
            arr[i] = arr[index];
            arr[index] = temp;
        }
    }

    index++;
    arr[high] = arr[index];
    arr[index] = pivot;
    return index;
}

function quickSort(arr: number[], low: number, high: number): void {
    if (low >= high) {
        return;
    }

    const pivotIndex = partition(arr, low, high)

    quickSort(arr, low, pivotIndex - 1)
    quickSort(arr, pivotIndex + 1, high)
}

function sortList(): void {
    const givenList: number[] = [
        42, 17, 23, 56, 9, 33, 89, 5, 78, 61, 102, 47, 3, 88, 24, 75, 11, 94,
        28, 36, 66, 83, 12, 57, 91, 64, 19, 85, 48, 14, 73, 27, 30, 54, 68, 7,
        99, 38, 41, 15, 81, 93, 21, 60, 6, 70, 18, 95, 31, 53, 76, 8, 32, 50,
        44, 22, 87, 13, 62, 39, 80, 20, 46, 4, 98, 74, 35, 58, 100, 67, 25, 16,
        77, 34, 92, 49, 29, 84, 10, 40, 82, 55, 63, 2, 45, 86, 26, 97, 37, 72,
        59, 1, 65, 90, 79, 52, 43, 71, 51, 96
    ];

    quickSort(givenList, 0, givenList.length - 1);
    console.log(givenList);
}

sortList();
