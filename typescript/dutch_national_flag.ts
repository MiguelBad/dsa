function sortNumber(arr: number[]): void {
    let low: number = 0;
    let mid: number = 0;
    let high: number = arr.length - 1;

    while (mid < high) {
        if (arr[mid] === 0) {
            const temp = arr[mid];
            arr[mid] = arr[low];
            arr[low] = temp;
            low++;
            mid++;
        }

        if (arr[mid] === 1) {
            mid++;
        }

        if (arr[mid] === 2) {
            const temp = arr[mid];
            arr[mid] = arr[high];
            arr[high] = temp;
            high--;
        }
    }
}

function main(): void {
    const arr: number[] = [0, 0, 1, 1, 2, 2, 2, 1, 2, 2, 1, 0, 0, 0, 2, 0, 0, 2, 1, 0, 1, 1, 1, 2, 2, 0, 1, 1, 0, 2, 1, 1, 1, 2, 2, 2, 0, 2, 1, 0, 2, 2, 2, 1, 1, 1, 2, 0, 1, 0, 0, 0, 2, 1, 1, 2, 2, 0, 1, 2, 0, 0, 0, 2, 1, 2, 0, 1, 2, 0, 1, 0, 0, 1, 2, 0, 0, 2, 2, 2, 0, 2, 0, 2, 2, 0, 0, 1, 0, 0, 1, 2, 0, 0, 1, 1, 2, 0, 2, 2];
    sortNumber(arr);
    console.log(arr);
}

main();
