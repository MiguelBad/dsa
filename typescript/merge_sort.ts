import arr from './sample_arraylist.json'

function mergeSort(arr: number[]): void {
    if (arr.length < 2) {
        return;
    }

    const mid: number = Math.floor(arr.length / 2);
    const left: number[] = arr.slice(0, mid);
    const right: number[] = arr.slice(mid, arr.length);

    mergeSort(left);
    mergeSort(right);


    let i: number, j: number, k: number;
    i = j = k = 0;

    while (left.length > i && right.length > j) {
        if (left[i] <= right[j]) {
            arr[k] = left[i];
            i++;
        } else {
            arr[k] = right[j];
            j++;
        }

        k++;
    }

    while (left.length > i) {
        arr[k] = left[i];
        i++;
        k++;
    }

    while (right.length > j) {
        arr[k] = right[j];
        j++;
        k++;
    }
}

function test(): void {
    mergeSort(arr);
    console.log(arr);
}

test()
