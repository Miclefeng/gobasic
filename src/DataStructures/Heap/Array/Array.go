package Array

/**
 * Author : MicleFengzss@gmail.com
 * Time : 2019/4/8 上午9:35
 */

type Array struct {
	Data []int
	Size int
}

var (
	Instance *Array
)

func init()  {
	Instance = new(Array)
}

func (arr *Array) GetSize() int {
	return arr.Size
}

func (arr *Array) IsEmpty() bool {
	return 0 == arr.Size
}

func (arr *Array) Add(idx int, e int) {

}