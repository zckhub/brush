package main

// 堆排序 - 使用大顶堆实现升序排序
// 时间复杂度: O(nlogn), 空间复杂度: O(1)
import "fmt"

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	sortedArr := HeapSort(arr)
	fmt.Println("Sorted array is:", sortedArr)
	fmt.Println("Original array is:", arr)
}

// HeapSort 对数组进行堆排序
func HeapSort(arr []int) []int {
	n := len(arr)

	// 从最后一个非叶子节点开始，自底向上构建大顶堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 逐个提取堆顶元素（最大值）放到数组末尾
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 堆顶与末尾元素交换
		heapify(arr, i, 0)              // 对剩余元素重新堆化
	}

	return arr
}

// heapify 对以i为根的子树进行堆化，n为堆大小
func heapify(arr []int, n, i int) {
	largest := i     // 初始化最大值为当前节点
	left := 2*i + 1  // 左子节点索引
	right := 2*i + 2 // 右子节点索引

	// 找出当前节点、左子节点、右子节点中的最大值
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	// 如果最大值不是当前节点，交换并递归堆化受影响的子树
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}
