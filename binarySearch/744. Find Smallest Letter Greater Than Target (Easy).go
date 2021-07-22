package binarysearch

//给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。
//
//在比较时，字母是依序循环出现的。举个例子：
//
//如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'

// NextGreatestLetter 因为letters队列是按照从小到大的顺序排列的, 所以遍历一旦得到一个bigger than target的元素, 那就是大于target的最小的一个元素 :)
func NextGreatestLetter(letters []byte, target byte) byte {
	for _, v := range letters {
		if v > target {
			return v
		}
	}
	return 0
}
