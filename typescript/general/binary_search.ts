import sortList from './quick_sort';

function binary_search(arr: number[], n: number): number {
    let high: number = arr.length - 1;
    let low: number = 0;

    while (high > low) {
        let idx: number = Math.floor((high - low) / 2) + low;
        if (n === arr[idx]) {
            return idx;
        }

        if (n > arr[idx]) {
            low = idx + 1;
        } else {
            high = idx - 1;
        }

    }

    return low;
}

const arr: number[] = sortList();
console.log(binary_search(arr, 100));
