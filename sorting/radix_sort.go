package sorting

import (
	"math"
)

func RadixSort(arr []int) {
	// 声明桶数组。
	var bucketArr [10][]int
	// 初始化arr中的最大元素。
	max := arr[0]
	// 初始化最大元素的位数（十进制）。
	figure := 1
	// 初始化循环次数。
	k := 0

	for k < figure {
		{
			// 在第一次遍历中，寻找数组中最大值。
			// 在所有遍历中，使用遍历次数来计算元素相应位的数字，根据该数字把元素添加到桶数组中。

			for _, ele := range arr {
				if figure == 1 {
					// 寻找arr中最大元素。
					if ele > max {
						max = ele
					}
				}

				// 根据k求指定位的数字。
				// 如求个位，计算公式为(ele % 10) / 1，十位为(ele % 100) / 10。
				lowFigure := int(math.Pow(float64(10), float64(k)))
				figureNumber := (ele % (lowFigure * 10)) / lowFigure

				// 根据figureNumber将ele添加到桶数组中。
				bucketArr[figureNumber] = append(bucketArr[figureNumber], ele)
			}
		}

		{
			// 将所有元素依次从桶数组中取出来，放入arr中。

			idx := 0
			for i := 0; i < 10; i++ {
				for j := 0; j < len(bucketArr[i]); j++ {
					arr[idx] = bucketArr[i][j]
					idx++
				}
			}

			// 清空桶数组中的所有元素。
			for i := 0; i < 10; i++ {
				bucketArr[i] = []int{}
			}
		}

		{
			// 计算最大元素的位数。

			if figure == 1 && max >= 10 {
				for {
					if max/int(math.Pow(float64(10), float64(figure))) == 0 {
						break
					}
					figure++
				}
			}
		}

		k++
	}
}
