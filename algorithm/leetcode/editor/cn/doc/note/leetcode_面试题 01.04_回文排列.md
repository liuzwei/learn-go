## 判断是否是回文字符串

### 思路1：

将字符串转成字符数组，双指针进行循环数组，匹配到成对的就将数组的value置成空字符串；
遍历完之后将数组合并成字符串；
如果新的字符串长度为0或1即说明是回文字符串；

```go
func canPermutePalindrome(s string) bool {
	na := strings.Split(s, "")
	for i := 0; i < len(na); i++ {
		for j := i + 1; j < len(na); j++ {
			if na[i] != "" && na[i] == na[j] {
				na[i] = ""
				na[j] = ""
			}
		}
	}
	rs := strings.Join(na, "")
	return len(rs) <= 1
}
```

### 思路2：

利用Map结构来计数，遍历字符串，如果Map结构中没有待配对的字符，则计数+1，如果匹配到成对的数据，则计数-1；
最后看计数是否小于2，小于2则说明是回文字符串。

```go
func canPermutePalindrome(s string) bool {
	n := 0
	runeMap := make(map[rune]bool)
	for _, r := range s {
		if runeMap[r] {
			runeMap[r] = false
			n--
		}else {
			runeMap[r] = true
			n++
		}
	}

	return n < 2
}
```