import data from './sample_arraylist.json'

function average(arr: number[]): number | undefined {
    if (arr.length === 0) {
        return undefined;
    }

    let sum: number = 0;

    for (let i = 0; i < arr.length; i++) {
        sum += arr[i]
    }

    return sum/arr.length;
}

console.log(`average is ${average(data)}`)
