# Median of Two Sorted Arrays - 寻找两个正序数组的中位数

地址：https://leetcode.com/problems/median-of-two-sorted-arrays/

难度：困难

知识点：



## 题目：

There are two sorted arrays nums1 and nums2 of size m and n respectively.
Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

**Example 1:**

​	nums1 = [1, 3]
​	nums2 = [2]

​	The median is 2.0

**Example 2:**
	nums1 = [1, 2]
	nums2 = [3, 4]

​	The median is (2 + 3)/2 = 2.5



给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。

请你找出这两个正序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

 

**示例 1:**

​	nums1 = [1, 3]
​	nums2 = [2]

​	则中位数是 2.0
**示例 2:**

​	nums1 = [1, 2]
​	nums2 = [3, 4]

​	则中位数是 (2 + 3)/2 = 2.5


## 题目分析

题目要求求两个有序数组的中位数。

**中位数：**将一个集合划分为两个长度相等的子集，其中一个子集中的元素总是大于另一个子集中的元素。

**计算中位数：**假设 A 集合有 n 个数值，A 下标从 0 开始

**奇数：**A[(n+1)/2]，取最中间的那个数作为中位数

**偶数：**(A[n/2] + A[n/2+1])/2，取最中间的那两个数的和除以2得到的值作为中位数



例如：

​	[1 2 3] 的中位数是 2

​	[1 2 3 4] 的中位数是 (2+3)/2 = 2.5



**题目要求为：两个有序数组取中位数，并且空间复杂度为 O(log(m + n)) **



## 方法一：

将两个合并数组合并为一个，然后取按奇数或者偶数计算出中位数的位置 `i`，直接 A[i] 取就可以。

但这种方式，合并数组，需要重新排序数组，排序需要遍历两个数组，时间复杂度为 $O(m+n)$，大于 $O(log (m+n))$ ，**不符合要求。**



> 时间复杂度与 log 有关的，一般会用到二分法。



## 方法二：求第 K 小的数

> 求中位数是求第 K 小的数是一个特殊的情况）

总个数为奇数的时候，$k = (n+1)/2 = parseInt(n/2) + 1$，第 `k` 小的数就是中位数。

总个数为偶数的时候，$k = n/2+1 = parseInt(n/2) + 1$，第 `k` 小的数和第 `k-1` 小的数的和除以2得到的值就是中位数



例如：

​	[1 2 3] 的长度 n 是 3，$k = parseInt(3/2) + 1 = 2$，中位数就是 A[1] = 2。（数组下标从0开始，所以第2个数的下标是 1）

​	[1 2 3 4] 的长度 n 是 4，$k = parseInt(4/2) + 1 = 3$，中位数就是 (A[1] + A[2])/2 = (2+3)/2 = 2.5



### 思路：

题目要求是两个有数组，假设为 A 数组 和 B 数组，在两个有序数组中找第 `k` 小的值。

因为是有序的数组，所以只需要循环从数组的头部剔除掉 `k-1` 个数，就可以得到 k 的值。

但两个数组中比第 `k` 小的值还小的那些值，可能在 A 数组中，也可能在 B 数组中。

所以，在剔除的时候，需要判断 A 数组的第一个数值和 B 数组的第一个数值哪个小，然后就剔除小的那个。



例如：

```
A = [1 2 3]

B = [1 3 5 7]

A + B = [1 1 2 3 3 5 7]，中位数是 3。~一眼就看出来，真厉害。~
```



接下来用找第 `k` 小的值的方式，来找出中位数。假设数组长度分别是 `n` 跟 `m`。

$k = parseInt((m + n)/2) + 1 = parseInt((3 + 4)/2) + 1 = 4$

得到 `k` 值后，开始从数组中剔除 `k - 1 = 3` 个数



第一个：A[0] = B[0] = 1，两个数相等，则随便剔除哪一个都可以，可以剔除 A 的先。

剔除后 A = [2 3]，B = [1 3 5 7]

第二个：(A[0] = 2) > (B[0] = 1)，剔除小的那个，B[0]

剔除后：A = [2 3]，B = [3 5 7]

第三个：(A[0] = 2) > (B[0] = 3)，剔除小的那个，A[0]

剔除后：A = [3]，B = [3 5 7]



第四个：就是我们要的中位数， A[3] = B[3]，取小的那个，这里相等，所以取哪个都可以。

这种情况是总个数` m + n` 为奇数。如果是偶数，还需要保留 `k - 1` 位的数，然后再和 `k` 位的数相加除以 2



这种方式需要计算 `(m + n)/2` 时间复杂度为 $O(m + n)$ 还是不满足题干要求。



### 优化：

题目要求时间复杂度为 $O(log(m + n))$，一般看到 log 的，都需要用**二分法**，所以可以将上面都过程再优化一下，不要一次“剔除”一个值，而是一次剔除一部分值。



这个一部分应该是多少，才能安全都剔除呢？既然是二分法，那应该是一半一半的剔除，也就是 k/2（向下取整）。

那有没有可能 k 也被剔除，下面分析一下看看



下面例子，A 跟 B 的长度 m 跟 n 分别都是 8，k = (8+8)/2 + 1 = 9，k/2（向下取整）= 4

```
A = [0 2 4 6  |  8 10 12 14]

B = [1 3 5 7  |  9 11 13 15]
```

A 的第四位是 7，B 的第 4 位是 6，优先剔除小的部分，B 的第 4 位跟前面的几位都可以安全剔除。

调整一下 A B

```
A = [0 1 2 4 6  |  8 10 12 14]

B = [3 5 7      |  9 11 13 15]
```

A 的第四位是 9，B 的第 4 位是 4。B 的小，

无论怎么调整，只要取数值小的那一组，就可以保证被剔除的会是第 k 小的数。



但这是已知的简单例子，并不能完全证明。

更一般的情况（这里下标表示在数组中的第 k 个，不是数组下标，从 1 开始）

```
A[1]，A[2]，A[3]，... A[k/2]
B[1]，B[2]，B[3]，... B[k/2]
```

假设 A[k/2] < B[k/2]，并且 B[k/2] 前面的值小于 A[k/2] 的个数是 `x` 个， `0 <= x < k/2`。

因为 A[k/2] 之前的值也比 A[k/2] 小，个数是 `k/2 - 1` 个。

所以，此时比 A[k/2] 小的个数是 `k/2 - 1 + x`，可知 x 越大时，A[k/2] 越可能超过第 `k` 小的数。

因为 `0 <= x < k/2`，x 是整数，所以 `x` 最大是 `k/2 - 1`。

代入得比 A[k/2] 小的值最多有 `k/2 - 1 + x = (k/2 - 1) + (k/2 - 1) = k - 2` 个

所以 A[k/2] 最大是第 `k-1` 小的数，不会超过第 `k` 小的数，所以 A[k/2] 和它前面比它小的值都可以安全剔除。



每次剔除 `k/2`，并且剔除后 `k = k - k/2`（向下取整），直到 `k = 1`



特殊情况处理：

当数组个数小于 `k/2` 的时候，取最大的那个值比较。

当其中一个数组个数都剔除完了，则可以按剩下的个数直接计算中位数的位置。



举个例子：

```
A = [0 2 4 6 8 10 12 14]
B = [1 3 5 7 9 11 13 15]
```

A 跟 B 的长度都是 8，我们要在这两个数组中找第 `k` 小的数，`k = (8+8)/2 + 1 = 9`，`k/2= 4`（向下取整）

（数组下标从 0 开始）

分别取A B 的第 4 个数，(A[3] = 6) < (B[3] = 7)，剔除 A 的前 4 个

```
A = [8 10 12 14]
B = [1 3 5 7 9 11 13 15]
k = k - k/2 = 5      // 剔除前面 4 个后，k 重新赋值
k/2 = 2
```

分别取A B 的第 2 个数，(A[1] = 10) > (B[1] = 3)，剔除 B 的前 2 个

```
A = [8 10 12 14]
B = [5 7 9 11 13 15]
k = k - k/2 = 3
k/2 = 1
```

分别取A B 的第 1 个数，(A[0] = 8) > (B[0] = 5)，剔除 B 的前 1 个

```
A = [8 10 12 14]
B = [7 9 11 13 15]
k = k - k/2 = 2
k/2 = 1
```

分别取A B 的第 1 个数，(A[0] = 8) > (B[0] = 7)，剔除 B 的前 1 个

```
A = [8 10 12 14]
B = [9 11 13 15]
k = k - k/2 = 1
```

`k = 1`，分别取A B 的第 1 个数，(A[0] = 8) < (B[0] = 9)，可得 A[0] = 8 为第 `k` 小的数。

因为 A 和 B 的长度 `m + n` 是偶数，所以，还需要取 `k - 1` 小的值，两个相加除以 2 才是中位数。



### 代码实现：

```go
package main

import (
	"fmt"
	"math"
)
/**
 * findMedianSortedArrays: 方法二
 *
 * Runtime: 16ms
 * Memory: 5.6MB
 *
 */
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    k := int((m+n)/2) + 1
    // 偶数返回第 k 和 k - 1 小的值除以2，奇数直接返回第 k 小的值
    if (m+n)%2 == 0 {
        return (findKth(nums1, nums2, m, n, 0, 0, k-1) + findKth(nums1, nums2, m, n, 0, 0, k)) / 2
    } else {
        return findKth(nums1, nums2, m, n, 0, 0, k)
    }
}

func maxInt(a, b int) int {
    return int(math.Max(float64(a), float64(b)))
}
func minInt(a, b int) int {
    return int(math.Min(float64(a), float64(b)))
}

/**
 * findKth: 寻找第 k 小的数
 */
func findKth(nums1, nums2 []int, m, n, start1, start2, k int) float64 {
    // 根据传入的开始位置，计算数组中剩余长度
    len1 := m - start1
    len2 := n - start2
    halfK := int(k / 2)
    if len1 == 0 {
        return float64(nums2[start2 + k - 1])
    }
    if len2 == 0 {
        return float64(nums1[start1 + k - 1])
    }
    if k == 1 {
        return math.Min(float64(nums1[start1]), float64(nums2[start2]))
    }

    // 取剩余长度中第 k/2 位的数，如果 k/2 大于剩余长度，则取最后一位
    i := start1 + minInt(len1, halfK) - 1
    j := start2 + minInt(len2, halfK) - 1
    if nums1[i] > nums2[j] {
        // 如果 nums2 中的数值小，则截取 nums2，把 nums2 的开始位置往右移动，
        // 因为已经取 j 位的值判断，所以 nums2 开始位置移动到 j + 1 位
        // 然后 k 值减 k/2 或数组剩余长度
        return findKth(nums1, nums2, m, n, start1, j+1, k-minInt(len2, halfK))
    } else {
        return findKth(nums1, nums2, m, n, i+1, start2, k-minInt(len1, halfK))
    }
}
```



## 方法三：

### 思路：

假设有两个有序数组 A 和 B，长度分别为 `m`，`n`。

使用 `i`，`j` 切割数组

```
        left_part        |         right_part

A[0], A[1], ... A[i-1]   |   A[i], A[i+1], ... A[m-1]

B[0], B[1], ... B[j-1]   |   B[j], B[j+1], ... B[m-1]

                         |
```

A 左边部分长度 `len(left_A) = i`，右边部分长度 `len(right_A) = m - i`。`0 <= i <= m`，当 `i = 0` 时，左边部分为空集，`i = m` 时，右边部分为空集

B 左边部分长度 `len(left_B) = j`，右边部分长度 `len(right_B) = n - j`。`0 <= j <= n`，当 `j = 0` 时，左边部分为空集，`i = n` 时，右边部分为空集

要算出中位数，则必须满足以下条件：

1. 切割后左边部分集合内的值永远小于右边部分集合内的值。`A[i-1] < B[j]` 并且 `B[j-1] < A[j]`
2. 偶数情况下，两边集合个数相等 `len(left_part) = len(right_part)`
3. 奇数情况下，左边集合个数比右边多一个 `len(left_part) + 1 = len(right_part)`

因为，`len(left_part) = len(right_part)` 或 `len(left_part) + 1 = len(right_part) `
所以，切割后，`i`、`j` 满足以下条件
`m + n` 是偶数情况下：
 左边部分个数等于右边部分：
  ` i + j = m - i + n - j`
  ` j = (m + n)/2 - i`
`m + n` 是奇数情况下：
 左边部分个数比右边部分多一个：
   `i + j = m - i + n - j + 1`
   `j = (m + n + 1)/2 - i`

因为 `i` 和 `j` 都整数，`m + n` 是偶数情况下 `(m + n)/2` 与 `(m + n + 1)/2` 取整的结果是一样的。

所以计算中可以统一表达为：
  `j = (m + n + 1)/2 - i`
结合前面 `A[i-1] <= B[j]` 并且 `B[j-1] <= A[j]`，找出满足以上条件的 `i` 值，根据 `i` 求出中位数：
偶数情况下，取左边的最大值与右边的最小值的和，除以2。得到中位数。
   `(max(left_part) + min(right_part))/2 -> (max(A[i-1], B[j-1]) + max(A[i] + B[j]))/2`
奇数情况下，取左边的最大值作为中位数。
   `max(left_part) -> max(A[i-1], B[j-1])`

临界值处理，当 `i = 0` 或 `j = 0` 或 `i = m` 或 `j = n` 时，
可以直接代入 `j = (m + n + 1)/2 - i`，即可算出结果。

例如：
当 `i = 0` 时，A[i-1] 不存在，A 是在最左边切割的。
此时可知左边部分最大是 B[j-1]，不需判断 `max(A[i-1], B[j-1])`

`j` 不能是负数，即 `(m + n + 1)/2 - i >= 0`，常数部分忽略，因为偶数的时候，没有常数
简化后得 `m + n - 2i >= 0`，又因为 `0 <= i <= m` 所以，`i` 取 `m` 的时候，`m + n - 2i` 最小，
得 `n - m >= 0`，`n >= m`。

`j` 不能大于 `n`，即 `(m + n + 1)/2 - i <= n`，常数部分忽略，因为偶数的时候，没有常数
简化后得 `m - n - 2i <= 0`，又因为 `0 <= i <= m` 所以，`i` 取 0 的时候，`m + n - 2i` 最大，
得 `m <= n`。 

所以必须确保 `n >= m`。
当 `n >= m`，`0 <= i <= m` 时，确保了 `0 <= j <= n`。`j` 不会超出范围

这样确保了搜索在短的那个数组中进行（[0 ~ m]，性能高）
也确保不会出现下面这样的情况：
例如：
假设 A = [0 2 4 6 8 10 12]，B = [13 15]，正常 `i = 5`，`j = 0`。
但如果不确保 n >= m，那当 i = 6，j 等于 -1 时（异常），`j = (m + n + 1)/2 - i` 也成立。
怎么确保 `n >= m`，当 `m > n` 时，把两个数组互换一下就可以了
A = [9 11]，B = [0 2 4 6 8 10 12]

接下来，按照以下步骤进行二叉树搜索，在 [0 ~ m] 中寻找对应的 i 值，
`i` 满足 `j = (m + n + 1)/2 - i` ，`A[i-1] <= B[j]` 并且 `B[j-1] <= A[j]` 的



> 步骤 1. 设 `iMin = 0`，`imax = m`，然后开始在[iMin, iMax] 中进行搜索。
>
> 步骤 2. 令 `i = (iMin + iMax)/2`，`j = (m + n + 1)/2 − i`
>
> 步骤 3. 满足 `j = (m + n + 1)/2 − i` 后，当 `i` 改变，会出现以下三种情况。
>
> A[i-1] <= B[j] 并且 B[j-1] <= A[i]：
>  	满足条件，找到目标，停止搜索
>
> B[j-1] > A[i]：
> 		说明 A[i] 太小，B[j-1] 太大。
> 		需要增大 i，使 A[i] 增大，使 B[j-1]（j 减小）。
> 		不能减少 i，因为会得到相反的结果。
> 		此时需要把搜索范围缩小到 [i+1, iMax]，把 iMin 设为 i + 1。
> 		转到 “步骤 2”
>
> A[i−1] > B[j]：
> 		说明 A[i - 1] 太大，B[j] 太小。
> 		需要减小 i，使 A[i - 1] 减小，使 B[j] 增大（j 增大）。
> 		此时需要把搜索范围缩小到 [iMin, i-1]，把 iMax 设为 i - 1。
> 		转到 “步骤 2”
> 按上面步骤，每次减少搜索范围的一半，知道搜索到目标结果。



### 代码实现：

```go
package main

import (
	"fmt"
	"math"
)
/*
 * Runtime: 12ms
 * Memory: 5.6MB
 *
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
func minInt(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// 确保 n >= m，如果 m > n，将两个数组互换
	if m > n {
		temp := nums1
		nums1 = nums2
		nums2 = temp
		m = len(nums1)
		n = len(nums2)
	}
	// 步骤 1
	iMin, iMax, halfLen := 0, m, int((m+n+1)/2)
	for iMin <= iMax {
		// 步骤 2
		i := int((iMin + iMax) / 2)
		j := halfLen - i

		// 步骤 3，三种情况分别处理
		if i < iMax && nums2[j-1] > nums1[i] {
			// 说明 nums1[i] 太小，nums2[j-1] 太大。需要增大 i，使 nums1[i] 增大，使 nums2[j-1]（j 减小）。
			// 搜索范围缩小到 [i+1, iMax]
			iMin = i + 1
		} else if i > iMin && nums1[i-1] > nums2[j] {
			// 说明 nums1[i-1] 太大，nums2[j] 太小。需要减小 i，使 nums1[i-1] 减小，使 nums2[j]（j 增大）。
			// 搜索范围缩小到 [iMin, i - 1]
			iMax = i - 1
		} else {
			// 除了上面俩种情况，最后就是搜索到目标
			// 搜索到后，需要找出最边最大，跟右边最小，求出中位数
			maxLeft := 0
			// 临界情况：i = 0，此时分割线在 A 的最左边，左边部分没有 A 的值，所以左边最大的是 B[j-1]。j = 0 同理

			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = maxInt(nums1[i-1], nums2[j-1])
			}
			// 如果总个数是奇数的情况下，直接返回最边最大值
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			minRight := 0
			// 临界情况：i = m，此时分割线在 A 的最右边，右边部分没有 A 的值，所以右边最小的是 B[j]。j = n 同理
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = minInt(nums1[i], nums2[j])
			}

			return (float64(maxLeft) + float64(minRight)) / 2.0
		}
	}
	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1}, []int{2, 3, 4, 5, 6})) // 3.5
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3})) // 2.0
	fmt.Println(findMedianSortedArrays([]int{2}, []int{1, 3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4})) // 2.5
	fmt.Println(findMedianSortedArrays([]int{}, []int{1, 2})) // 1.5
}
```

