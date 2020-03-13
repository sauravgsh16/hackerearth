/*
Like Merge Sort, QuickSort is a Divide and Conquer algorithm.
It picks an element as pivot and partitions the given array around the picked pivot.
There are many different versions of quickSort that pick pivot in different ways.

    Always pick first element as pivot.
    Always pick last element as pivot (implemented below)
    Pick a random element as pivot.
    Pick median as pivot.

The key process in quickSort is partition(). Target of partitions is,
given an array and an element x of array as pivot, put x at its correct position in sorted array
and put all smaller elements (smaller than x) before x,
and put all greater elements (greater than x) after x.
All this should be done in linear time.

QuickSort psuedo code

quickSort(arr[], low, high)
{
    if (low < high)
    {
        // pi is partitioning index, arr[pi] is now at right place

	pi = partition(arr, low, high);

	quickSort(arr, low, pi - 1);  // Before pi
	quickSort(arr, pi + 1, high); // After pi
    }
}

This function takes last element as pivot, places the pivot element at its correct position in sorted
array, and places all smaller (smaller than pivot) to left of pivot and all greater elements to right of pivot

partition (arr[], low, high)
{
// pivot (Element to be placed at right position)

pivot = arr[high];

i = (low - 1)  // Index of smaller element

for (j = low; j <= high- 1; j++)
{
	// If current element is smaller than the pivot
	if (arr[j] < pivot) {
		i++;    		// increment index of smaller element
		swap arr[i] and arr[j]
	}
}
swap arr[i + 1] and arr[high])
return (i + 1)
}


ANALYSIS:

Worst case: O(n^2) - when partition process always picks smallest or greatest element as pivot.
		    Worst case would occur from above partition logic, if array is sorted in incresing/decreasing order.

Best Case : O(nlogn) - best case occurs when middle element is picked as pivot

Average Case : O(nlogn)

Stable : No

In-place: yes

*/
package main

func quickSort(arr []int, low, high int) {
	if low < high {
		// partition index
		pi := partition(arr, low, high)

		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	i := low - 1 // index of smaller element
	pivot := arr[high]

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++ // increment index of smaller element

			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
