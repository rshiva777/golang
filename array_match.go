// this program compare 2 arrays and give the match elements between arrays
package main
import "fmt"
func main (){
	ab := []int{1,4,4,6,7,1}
	ac := []int{5,6,2,2,1,4}
	ab = uniq(ab)
	ac = uniq(ac)
	common(ab,ac)
	}

func common(ab []int, ac []int){
	var union []int
	for _, i := range ab{
	for _, j := range ac{
		if i == j{
			union = append(union, i)
		}
	}
}
fmt.Println(union)
}
func uniq(arr []int)[]int{
	mapping := map[int]bool{}
	var result []int
	for e := range arr {
		if mapping[arr[e]] != true {
			mapping[arr[e]] = true
			result = append(result, arr[e])
		}
	}
	return result
}
