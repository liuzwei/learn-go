<p>一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>

<pre><strong>输入:</strong> [0,1,3]
<strong>输出:</strong> 2
</pre>

<p><strong>示例&nbsp;2:</strong></p>

<pre><strong>输入:</strong> [0,1,2,3,4,5,6,7,9]
<strong>输出:</strong> 8</pre>

<p>&nbsp;</p>

<p><strong>限制：</strong></p>

<p><code>1 &lt;= 数组长度 &lt;= 10000</code></p>
<details><summary><strong>Related Topics</strong></summary>位运算 | 数组 | 哈希表 | 数学 | 二分查找</details><br>

<div>👍 307, 👎 0</div>

<div id="labuladong"><hr>

**通知：[数据结构精品课 V1.8](https://aep.h5.xeknow.com/s/1XJHEO) 持续更新中，[第 11 期刷题打卡挑战](https://mp.weixin.qq.com/s/eUG2OOzY3k_ZTz-CFvtv5Q) 开始报名。**

<details><summary><strong>labuladong 思路</strong></summary>

## 基本思路

这道题考察 [二分查找算法](https://labuladong.github.io/article/fname.html?fname=二分查找详解)。

常规的二分搜索让你在 `nums` 中搜索目标值 `target`，体现在代码中就是判断 `nums[mid]` 和 `target` 的关系。

这道题寻找缺失的那个数字，体现在代码中就是判断 `nums[mid]` 和 `mid` 的关系。

[二分查找算法](https://labuladong.github.io/article/fname.html?fname=二分查找详解) 中说到了二分搜索的几种形式，我就用搜索左侧边界的二分搜索定位缺失的元素位置。

**标签：[二分搜索](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)，[数组](https://mp.weixin.qq.com/mp/appmsgalbum?__biz=MzAxODQxMDM0Mw==&action=getalbum&album_id=2120601117519675393)**

## 解法代码

```java
class Solution {
    public int missingNumber(int[] nums) {
        // 搜索左侧的二分搜索
        int left = 0, right = nums.length - 1;
        while (left <= right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] > mid) {
                // mid 和 nums[mid] 不对应，说明左边有元素缺失
                right = mid - 1;
            } else {
                // mid 和 nums[mid] 对应，说明元素缺失在右边
                left = mid + 1;
            }
        }
        return left;
    }
}
```

</details>
</div>



