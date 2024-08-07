function fibonacci(i: number): number {
    if (i === 0) return 0;
    if (i === 1) return 1;

    let prev: number = 0;
    let current: number = 1;

    for (let j = 1; j < i; j++) {
        const temp: number = current;
        current = current + prev;
        prev = temp;
    }

    return current;
}

const fib: number = fibonacci(6);
console.log(fib);

// Give the list of fibonacci
function fibonacci_list(i: number): number[] {
    if (i === 0) return [0];
    if (i === 1) return [0, 1];

    let arr: number[] = [0, 1];
    for (let j = 2; j < i; j++) {
        arr.push(arr[j - 1] + arr[j - 2]);
    }

    return arr;
}

const fib_list: number[] = fibonacci_list(6);
console.log(fib_list);
