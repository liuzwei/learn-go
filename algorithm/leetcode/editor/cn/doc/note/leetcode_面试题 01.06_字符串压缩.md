### 思路1：
遍历整个字符串，去当前字符和当前字符的数量为变量，声明一个压缩字符串变量；
每遍历完一个相同的字符就把该字符和数量拼接到压缩的字符串上，直到遍历结束；

判断压缩后的字符串，是否比压缩前的字符串短，是则返回压缩后的，不是则返回原字符串。
