// This program sort the elements in an array using quicksort method

package main
import "fmt"
func main (){
	arr := []int{6,4,5,3,9,7,2,10,3}
fmt.Println(arr)
	low := 0
	high:= len(arr)-1
	quickSort(arr,low,high)
	fmt.Println(arr)
}
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	fmt.Println(i,high)
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			fmt.Println(i,j)
			fmt.Println(arr)
		}

	}
	fmt.Println(i,high)
	arr[i], arr[high] = arr[high], arr[i]
	fmt.Println(arr)
	fmt.Println("done")
	return arr, i

}
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		//fmt.Println(low,high)
		arr, p = partition(arr, low, high)
		fmt.Println(arr,p,low,high)
		fmt.Println("done1")
		//fmt.Println(arr,p,low,high)
		arr = quickSort(arr, low, p-1)
		fmt.Println("done2")
		arr = quickSort(arr, p+1, high)
		fmt.Println("done3")
	}
	return arr
}
