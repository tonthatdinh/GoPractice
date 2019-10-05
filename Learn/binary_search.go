package main

import "fmt"

func binarySearch(array []int, target int, low int, high int) int{
    if high < low {
        return -1;
    }
    mid := int ((low + high) / 2)
    if array[mid] > target {
        return binarySearch(array, target, low, mid)
    } else if array[mid] < target{
        return binarySearch(array, target, mid + 1, high)
    } else {
        return mid
    }
}

func iterBinarySearch(array []int, target int, low int, high int) int{
    start := low
    end := high
    var mid int
    for start < end {
        mid = int((start + end) / 2)
        if array[mid] > target{
            end = mid
        } else if array[mid] < target{
            start = mid
        } else{
            return mid
        }
    }
    return -1
}

func main()  {
    arr1 := []int{1, 2, 3, 4, 5}
    fmt.Println(arr1[1])
    fmt.Println(binarySearch(arr1, 3, 0, 4))
    fmt.Println(iterBinarySearch(arr1, 4, 0, 4));
}


