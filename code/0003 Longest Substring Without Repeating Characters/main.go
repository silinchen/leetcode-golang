package main

import (
	"fmt"
	"math"
	"strings"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

// Given a string, find the length of the longest substring without repeating characters.

// Example 1:
// Input: "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.

// Example 2:
// Input: "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.

// Example 3:
// Input: "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
//              Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

/**
 * 备注：滑动窗口
 *
 * 例如字符串 “abcacbbac”，接下来假设有一个窗口，窗口只显示字符串的一部分，过程如下：
 * -> "a" // 窗口内字符串不重复的情况下，窗口右边线向右扩大，下同
 * -> "ab"
 * -> "abc"
 * -> "abca" // a 重复，所以窗口左边线向左移动到 1 位，直到字符串内字符不重复
 * -> "bca"
 * -> "bcac" // c 重复，所以窗口左边线向左移动 1 位，直到字符串内不重复
 * -> "cac" // c 重复，所以窗口左边线向左移动 1 位，直到字符串内不重复
 * -> "ac" -> "acb" -> "acbb" -> "cbb" -> "bb" -> "b" -> "ba"
 * -> "bac" // 到最后窗口移动到最右边，结束循环。取过程中出现的最长的不重复字符 3
 *
 * 例如字符串 “pwwkew”，过程如下：
 * -> "p"       // (p)wwkew
 * -> "pw"      // (pw)wkew
 * -> "pww"     // (pww)kew
 * -> "ww"      // p(ww)kew
 * -> "w"       // pw(w)kew
 * -> "wk"      // pw(wk)ew
 * -> "wke"     // pw(wke)w
 * -> "wkew"    // pw(wkew)
 * -> "kew"     // pww(kew)
 *
 * 这个过程，就好像是通过一个会向右滑动的窗口看长字符串的一个局部字符串。
 * 每次都只移动一个位置，如果窗口内有字符重复，就移动左边线，如果没有，就移动右边线。
 */

/**
 * 滑动窗口优化：
 * 例如： “abcacbbac”
 * -> "a"
 * -> "ab"
 * -> "abc"
 * -> "abca"
 * // 按上面普通滑动窗口的流程，接下来应该是  “bca” -> “bcac”，
 * // 但是其实这个过程可以省略，既然我们可以知道下一个是 c，然后我们又想窗口内字符串不重复，
 * // 那左边就可以直接跳过 “abc”，从 “c” 的位置开始，变成 "cac"，但是 “c” 是不要的，所以下一步就是
 * -> "ac"
 * -> "acb"
 * -> "b" // "acb" 的下一位是“b”，b 重复了，直接跳到 b 的位置 “bb”，但不要重复的那个 “b”，所以 “acbb” 就剩 “b”
 * -> "ba"
 * -> "bac" // 结束
 *
 * 例如：“pwwkew”
 * -> "p"      // (p)wwkew
 * -> "pw"     // (pw)wkew
 * -> "w"      // pw(w)kew
 * -> "wk"     // pw(wk)ew
 * -> "wke"    // pw(wke)w
 * -> "kew"    // pww(kew)
 *
 * 优化后，右边依然每次移动位，但左边变成一次移动一或多位，移动的距离根据新进入的字符判断。
 * 如果新进入的字符使得窗口内字符有重复了，左边线直接移动到重复字符的位置（+1 位），使窗口内字符重新变成不重复。
 */

//abb      a->ab->abb->bb->b
//abb      a->ab->b

/**
 * 对比：假设长度是 n，左边位置是 i，右边位置是 j
 * 优化前，每次移动是 i+1，或者 j+1，最坏的情况下，需要循环 2n-1 次
 * 优化后，每次移动 j+1，i 则判断新增加的字符在窗口内的位置，直接设置i 的位置就行。循环次数是 n 次
 */

func maxInt(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

// 方法三：优化滑动窗口
func lengthOfLongestSubstring(s string) int {
	// 记录最长的不重复子字符串的长度
	var ans = 0
	// 记录窗口开始位置
	var start = 0
	// 保存窗口内每个字符，以及字符对应的位置，方便滑动窗口把开始的位置移动到重复字符的下一位
	var sMap = make(map[rune]int)
	// 循环 j，相当于窗口每次都会向右扩展一个位置
	// for j, str := range []rune(s) { // 支持中文
	for j, str := range s {
		// 如果新加入的字符在窗口内已经存在
		if i, ok := sMap[str]; ok {
			// 把开始位置设置为该字符在窗口内的位置，使窗口内的字符再次变成不重复
			start = maxInt(i, start)
		}
		// 设置最大长度值
		ans = maxInt(ans, j-start+1)
		// 把新加入的字符的位置存起来，方便当字符出现重复的时候，快速跳到该字符的位置。
		// + 1 是为了把位置向右移动一位，也就是把该字符也排除在窗口外。
		// 例如 “abcdb”，第一个 b 是 s[1]，那开始位置就要设置为 s[1+1] -> s[2] -> "cdb"，把
		sMap[str] = j + 1
	}
	return ans
}

// 方法二：滑动窗口
func lengthOfLongestSubstring2(s string) int {
	substr, ans, i, j, n := "", 0, 0, 0, len(s)
	// i 或者 j，不超过 n 就一直循环，向右滑动窗口
	for i < n && j < n {
		ch := s[j : j+1]
		if index := strings.Index(substr, ch); index == -1 {
			j = j + 1
			substr = s[i:j]
			ans = maxInt(ans, j-i)
		} else {
			i = i + 1
			substr = substr[1:]
		}
	}
	return ans
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))    // 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))       // 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))      // 3
	fmt.Println(lengthOfLongestSubstring("你好，世界！世界，你好")) // 6

	// fmt.Println(lengthOfLongestSubstring2("abcabcbb")) // 3
	// fmt.Println(lengthOfLongestSubstring2("pwwkew"))   // 3
}

/**
 * Runtime: 0~8ms
 * Memory: 2.6~3.6MB
 */
