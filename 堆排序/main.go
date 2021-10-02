// 思路堆排序实际是数组排序，使用数组下标构造一个二叉树，
// 加入有数组a，那可以将a[0]视为根节点，它的左子节点为a[2*i + 1]，右子节点为a[2*i +2]
// 构建一个大顶堆，构建成功后a[0]便是最大的
// 之后交换a[0]和a[len(a) - 1]
// 然后在调整a[:len(a)-2]为大顶堆
// 重复上面的步骤，得到有序数组
// 时间复杂度为nlog(n)

package main
import (
    "fmt"
)
func main(){
	num := []int64{1,2,3,4,5,6}
	heapSort(num)
	
	fmt.Println("heap sort over:", num)

}

func heapSort(arr []int64) {
	//构建
	for i := int64(len(arr))/2;i>=0;i-- {
		adjectHeap(arr, i, int64(len(arr)))
	}
	// 调整
	for i := int64(len(arr)) -1 ; i > 0 ; i-- {
		if arr[i] > arr[0] {
			continue
		}
		arr[0],arr[i] = arr[i],arr[0]
		adjectHeap(arr, 0, i)
	}
}

func adjectHeap(arr []int64, pos int64, length int64) {
	for {
		child := pos * 2 + 1

		if child >= length-1 {
			break
		}
			
		if arr[child] < arr[child+1] {
			child++
		}

		if arr[pos] < arr[child] {
			arr[pos], arr[child] = arr[child], arr[pos]
			pos = child
		} else {
			break
		}
	}
}