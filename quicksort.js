function quickSort(arr) {
    if (arr.length <=1) {
        return arr
    }

    let p = arr[arr.length -1]

    let left = []
    let right = []

    for(let i =0; i < arr.length-1; i++) {
        if (arr[i] < p) {
            left.push(arr[i])
        } else {
            right.push(arr[i])
        }
    }

    if (left.length > 0 && right.length > 0) {
        return [...quickSort(left), p, ...quickSort(right)]
    } else if (left.length > 0) {
        return [...quickSort(left), p]
    } else {
        return [p, ...quickSort(right)]
    }
}