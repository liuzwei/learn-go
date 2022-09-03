//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。请完成一个高效的函数，输入这样的一个二维数组和一个
//整数，判断数组中是否含有该整数。 
//
// 
//
// 示例: 
//
// 现有矩阵 matrix 如下： 
//
// 
//[
//  [1,   4,  7, 11, 15],
//  [2,   5,  8, 12, 19],
//  [3,   6,  9, 16, 22],
//  [10, 13, 14, 17, 24],
//  [18, 21, 23, 26, 30]
//]
// 
//
// 给定 target = 5，返回 true。 
//
// 给定 target = 20，返回 false。 
//
// 
//
// 限制： 
//
// 0 <= n <= 1000 
//
// 0 <= m <= 1000 
//
// 
//
// 注意：本题与主站 240 题相同：https://leetcode-cn.com/problems/search-a-2d-matrix-ii/ 
// Related Topics数组 | 二分查找 | 分治 | 矩阵 
//
// 👍 773, 👎 0 
//
//
//
//


<div id="labuladong"><hr>

**通知：[数据结构精品课 V1.8](https://aep.h5.xeknow.com/s/1XJHEO) 持续更新中，[第 11 期刷题打卡挑战](https://mp.weixin.qq.com/s/eUG2OOzY3k_ZTz-CFvtv5Q) 开始报名。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题和 [240. 搜索二维矩阵 II](/problems/search-a-2d-matrix-ii) 相同。

作为 [74. 搜索二维矩阵](/problems/search-a-2d-matrix) 的进阶题目，这道题不算 [二分搜索算法](https://labuladong.github.io/article/fname.html?fname=二分查找详解)，更像 [一个方法秒杀所有 N 数之和问题](https://labuladong.github.io/article/fname.html?fname=nSum)，因为它们的思想上有些类似。

这道题说 `matrix` 从上到下递增，从左到右递增，显然左上角是最小元素，右下角是最大元素。**我们如果想高效在 `matrix` 中搜索一个元素，肯定需要从某个角开始**，比如说从左上角开始，然后每次只能向右或向下移动，不要走回头路。

如果真从左上角开始的话，就会发现无论向右还是向下走，元素大小都会增加，那么到底向右还是向下？不确定，那只好用类似 [动态规划算法](https://labuladong.github.io/article/fname.html?fname=动态规划详解进阶) 的思路穷举了。

但实际上不用这么麻烦，我们不要从左上角开始，而是从右上角开始，规定只能向左或向下移动。

你注意，如果向左移动，元素在减小，如果向下移动，元素在增大，这样的话我们就可以根据当前位置的元素和 `target` 的相对大小来判断应该往哪移动，不断接近从而找到 `target` 的位置。

当然，如果你想从左下角开始，规定只能向右或向上移动也可以，具体看代码吧。

**标签：[数组](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)，[数组双指针](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)**

## 解法代码

```java
class Solution {
    public boolean searchMatrix(int[][] matrix, int target) {
        int m = matrix.length, n = matrix[0].length;
        // 初始化在右上角
        int i = 0, j = n - 1;
        while (i < m && j >= 0) {
            if (matrix[i][j] == target) {
                return true;
            }
            if (matrix[i][j] < target) {
                // 需要大一点，往下移动
                i++;
            } else {
                // 需要小一点，往左移动
                j--;
            }
        }
        // while 循环中没有找到，则 target 不存在
        return false;
    }
}
```

**类似题目**：
  - [剑指 Offer 04. 二维数组中的查找 🟠](/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof)

</details>
</div>

