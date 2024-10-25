function sort(arr: number[]): void {
    let low: number = 0;
    let mid: number = 0;
    let high: number = arr.length - 1;

    while (mid <= high) {
        if (arr[mid] === 0) {
            const temp = arr[low];
            arr[low] = arr[mid];
            arr[mid] = temp;
            low++;
            mid++;
        } else if (arr[mid] === 1) {
            mid++;
        } else if (arr[mid] === 2) {
            const temp = arr[mid];
            arr[mid] = arr[high];
            arr[high] = temp;
            high--;
        }
    }

}

function DutchNationalFlagDefence(): void {
    const arr: number[] = [0, 0, 1, 1, 2, 2, 2, 1, 2, 2, 1, 0, 0, 0, 2, 0, 0, 2, 1, 0, 1, 1, 1, 2, 2, 0, 1, 1, 0, 2, 1, 1, 1, 2, 2, 2, 0, 2, 1, 0, 2, 2, 2, 1, 1, 1, 2, 0, 1, 0, 0, 0, 2, 1, 1, 2, 2, 0, 1, 2, 0, 0, 0, 2, 1, 2, 0, 1, 2, 0, 1, 0, 0, 1, 2, 0, 0, 2, 2, 2, 0, 2, 0, 2, 2, 0, 0, 1, 0, 0, 1, 2, 0, 0, 1, 1, 2, 0, 2, 2];
    sort(arr);
    console.log(arr);
}

DutchNationalFlagDefence();
