package main

//problem: create one bounded ring buffer (FIFO)
//which provide Get and Put functions, and write test cases for it.
import "fmt"

func main() {
	var nums FIFOList
	nums.PutFIFO(1)
	fmt.Println("nums", nums)
	nums.GetFIFO()
	fmt.Println("nums", nums)

	fmt.Println("wwww")
}

//type FIFO struct {
//	Values int
//	Next *FIFO
//}
type FIFOList []int

func (nums *FIFOList) GetFIFO() (res int) {
	if len(*nums) < 1 {
		return -1
	}
	res = (*nums)[0]
	*nums = (*nums)[1:]
	return res
}
func (nums *FIFOList) PutFIFO(value int) {
	*nums = append(*nums, value)
	return
}
