//Program to sort the elements incrementally in array

package main
import "fmt"
func main (){
	arr := []int{2,5,1,6,3,7,4}
	fmt.Println("before sort", arr)
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-1;j++{
			if arr[j] > arr[j+1]{
				a := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = a
			}
		}
	}
	fmt.Println("after  sort",arr)
}
