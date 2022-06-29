package algorithm

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/valid-parentheses
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func ValidString(s string) bool {
	mp := map[string]string{
		")": "(", "]": "[", "}": "{",
	}
	stack := []string{}
	for _, item := range s {
		if item == '(' || item == '[' || item == '{' {
			stack = append(stack, string(item))
		} else {
			// 如果栈中没有元素，说明不匹配
			if len(stack) == 0 {
				return false
			}
			// 不是左括号就取出一个，然后和当前比较是否配对
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if mp[string(item)] != pop {
				return false
			}
		}
	}
	return true
}
