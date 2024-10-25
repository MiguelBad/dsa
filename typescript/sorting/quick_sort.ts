import data from './sample_arraylist.json'

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


function sortList(): number[] {
    const givenList: number[] = data;
    quickSort(givenList, 0, givenList.length - 1);
    console.log(givenList)
    return givenList;
}

export default sortList;
