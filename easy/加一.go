package main

//给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。
//
//最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。
//
//你可以假设除了整数 0 之外，这个整数不会以零开头。
//
//示例 1:
//
//输入: [1,2,3]
//输出: [1,2,4]
//解释: 输入数组表示数字 123。
//
//示例 2:
//
//输入: [4,3,2,1]
//输出: [4,3,2,2]
//解释: 输入数组表示数字 4321。
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/plus-one
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func plusOne(digits []int) []int {
	//// 首先, 按照数组的次序, 得到一个数字, 然后把这个数字加一, 然后把数字转化为数组
	//sum := 0
	//for i:=0; i<len(digits); i++ {
	//	v := digits[i]*int(math.Pow10(len(digits)-1-i)) // 10的多少次方?
	//	sum += v
	//}
	//fmt.Println(sum)
	//sum += 1
	//want_slice := make([]int, 0, 100)
	//for ;sum!=0;{
	//	v := sum%10
	//	want_slice = append(want_slice, v)
	//	sum = sum/10
	//}
	//// 再来一个反转就阔以了吧
	//want_slice2 := make([]int, 0, 100)
	//for i:=len(want_slice)-1; i>-1; i-- {
	//	want_slice2 = append(want_slice2, want_slice[i])
	//}
	//return want_slice2

	// 上面的, 没有考虑一个非常大的数, 溢出的问题
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		digits[i] = digits[i] % 10 // 等于0说明有需要进一位的, 则需要执行循环, 如果小于10, 则%10之后还是之前的那个数
		if digits[i] != 0 {
			return digits
		}
	}
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
