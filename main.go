// You can edit this code!
// Click here and start typing.
package main

import "fmt"


func insertionSort(arr []int) {
    n := len(arr)
    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
}

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func merge(arr []int, left int, mid int, right int) {
    n1 := mid - left + 1
    n2 := right - mid

    leftArr := make([]int, n1)
    rightArr := make([]int, n2)

    for i := 0; i < n1; i++ {
        leftArr[i] = arr[left+i]
    }

    for j := 0; j < n2; j++ {
        rightArr[j] = arr[mid+1+j]
    }

    i := 0
    j := 0
    k := left

    for i < n1 && j < n2 {
        if leftArr[i] <= rightArr[j] {
            arr[k] = leftArr[i]
            i++
        } else {
            arr[k] = rightArr[j]
            j++
        }
        k++
    }

    for i < n1 {
        arr[k] = leftArr[i]
        i++
        k++
    }

    for j < n2 {
        arr[k] = rightArr[j]
        j++
        k++
    }
}

func mergeSort(arr []int, left int, right int) {
    if left < right {
        mid := (left + right) / 2

        mergeSort(arr, left, mid)
        mergeSort(arr, mid+1, right)

        merge(arr, left, mid, right)
    }
}

func partition(arr []int, low int, high int) int {
    pivot := arr[high]
    i := low - 1

    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }

    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}

func quickSort(arr []int, low int, high int) {
    if low < high {
        pivot := partition(arr, low, high)

        quickSort(arr, low, pivot-1)
        quickSort(arr, pivot+1, high)
    }
}

func binarySearch(arr []int, target int) int {
    low := 0
    high := len(arr) - 1
    
    for low <= high {
        mid := (low + high) / 2
        if arr[mid] == target {
            return mid
        } else if arr[mid] < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    
    return -1
}

 
func main() {
	arr := []int{34,45,56,67,78,43,45,22,33,76,12,43,78,89,43,99,67,98}

	bubbleSort(arr)
	// mergeSort(arr, 0, len(arr)-1)
	// quickSort(arr, 0, len(arr)-1)
	// insertionSort(arr)
	// selectionSort(arr)
	fmt.Println("This is bubbleSort result", arr)
	target := 22
	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Println("Indice do elemento (target): ", index)
	}
    
}

