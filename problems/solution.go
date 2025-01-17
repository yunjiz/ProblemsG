package problems

import (
	"container/heap"
	"github.com/golang-collections/collections/queue"
	"github.com/golang-collections/collections/stack"
	"math"
	"sort"
	"strings"
)

func twoSum(nums []int, target int) []int {
	ans := make([]int, 2)
	check_list := make(map[int]int)
	for i, n := range nums {
		complement := target - nums[i]
		if val, ok := check_list[complement]; ok {
			ans[0] = val
			ans[1] = i
			break
		} else {
			check_list[n] = i
		}
	}
	return ans
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummuyNode := new(ListNode)
	previousNode := dummuyNode

	carry := 0

	for l1 != nil || l2 != nil {
		var val1, val2 int
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		} else {
			val1 = 0
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		} else {
			val2 = 0
		}

		sum := val1 + val2 + carry
		val := sum % 10
		carry = sum / 10

		newNode := ListNode{Val: val}
		previousNode.Next = &newNode
		previousNode = &newNode
	}

	if carry != 0 {
		newNode := ListNode{Val: carry}
		previousNode.Next = &newNode
	}

	return dummuyNode.Next
}

func lengthOfLongestSubstring(s string) int {
	result := 0
	left, right := 0, 0
	set := make(map[rune]bool)

	for right < len(s) {
		if set[rune(s[right])] == true {
			delete(set, rune(s[left]))
			left++
		} else {
			set[rune(s[right])] = true
			right++
			if right-left > result {
				result = right - left
			}
		}
	}
	return result
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if (m+n)%2 == 0 {
		return float64(getKth(nums1, nums2, (m+n)/2+1)+getKth(nums1, nums2, (m+n)/2)) * 0.5
	} else {
		return float64(getKth(nums1, nums2, (m+n)/2+1)) * 1.0
	}
}

func getKth(A []int, B []int, k int) int {
	m := len(A)
	n := len(B)

	if m > n {
		return getKth(B, A, k)
	}
	if m == 0 {
		return B[k-1]
	}
	if k == 1 {
		return min(A[0], B[0])
	}

	pa := min(k/2, m)
	pb := k - pa
	if A[pa-1] <= B[pb-1] {
		return getKth(A[pa:], B, pb)
	} else {
		return getKth(A, B[pb:], pa)
	}

}

func longestPalindrome(s string) string {
	max := -1
	result := ""
	for i := 0; i < len(s); i++ {
		length, value := expandCheck(s, i, i)
		if length > max {
			result = value
			max = length
		}
		length, value = expandCheck(s, i, i+1)
		if length > max {
			result = value
			max = length
		}
	}
	return result
}

func expandCheck(s string, i, j int) (length int, result string) {
	for i >= 0 && j < len(s) {
		if s[i] != s[j] {
			break
		}
		i--
		j++
	}
	i++
	j--
	if i <= j {
		return j - i + 1, s[i : j+1]
	}
	return 0, ""
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	rows := make([]strings.Builder, 0)
	for i := 0; i < min(len(s), numRows); i++ {
		rows = append(rows, strings.Builder{})
	}

	curRow := 0
	goingDown := false

	for _, c := range s {
		rows[curRow].WriteString(string(c))
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow -= 1
		}
	}

	var result strings.Builder
	for _, row := range rows {
		result.WriteString(row.String())
	}
	return result.String()
}

func reverse(x int) int {
	var result int
	for x != 0 {
		result = result*10 + x%10
		x = x / 10
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}

func myAtoi(str string) int {
	str = strings.Trim(str, " ")
	result := 0
	if str == "" {
		return result
	}
	sign := 1
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			result = result*10 + int(str[i]-'0')
		} else {
			break
		}

		if result > math.MaxInt32 && sign == 1 {
			return math.MaxInt32
		} else if result > math.MaxInt32+1 && sign == -1 {
			result = math.MinInt32
			return math.MinInt32
		}
	}

	result = sign * result

	return result
}

func isPalindrome(x int) bool {
	origin := x
	if x < 0 {
		return false
	}
	result := 0
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	return result == origin
}

func isMatch1(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	var first_match bool
	first_match = len(s) > 0 && (s[0] == p[0] || p[0] == '.')

	if len(p) >= 2 && p[1] == '*' {
		return isMatch1(s, p[2:]) || (first_match && isMatch1(s[1:], p))
	} else {
		return first_match && isMatch1(s[1:], p[1:])
	}
}

var memo [][]int

const (
	TRUE  = 1
	FALSE = 2
)

func isMatch2(s string, p string) bool {
	memo = make([][]int, len(s)+1)
	for i := range memo {
		memo[i] = make([]int, len(p)+1)
	}
	return isMatch2DP(0, 0, s, p)
}

func isMatch2DP(i, j int, text, pattern string) bool {
	if memo[i][j] != 0 {
		return memo[i][j] == TRUE
	}
	var ans bool
	if j == len(pattern) {
		ans = i == len(text)
	} else {
		first_match := (i < len(text) && (pattern[j] == text[i] || pattern[j] == '.'))
		if j+1 < len(pattern) && pattern[j+1] == '*' {
			ans = (isMatch2DP(i, j+2, text, pattern) || first_match && isMatch2DP(i+1, j, text, pattern))
		} else {
			ans = first_match && isMatch2DP(i+1, j+1, text, pattern)
		}
	}
	if ans {
		memo[i][j] = TRUE
	} else {
		memo[i][j] = FALSE
	}
	return ans
}

func isMatch3(text string, pattern string) bool {
	dp := make([][]bool, len(text)+1)
	for i := range dp {
		dp[i] = make([]bool, len(pattern)+1)
	}
	dp[len(text)][len(pattern)] = true

	for i := len(text); i >= 0; i-- {
		for j := len(pattern) - 1; j >= 0; j-- {
			first_match := (i < len(text) &&
				(pattern[j] == text[i] || pattern[j] == '.'))
			if j+1 < len(pattern) && pattern[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || first_match && dp[i+1][j]
			} else {
				dp[i][j] = first_match && dp[i+1][j+1]
			}
		}
	}
	return dp[0][0]
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	tmplist := make([]int, 0)
	backtrackSubsetsRef(&result, &tmplist, nums, 0)
	return result
}

func backtrackSubsetsRef(result *[][]int, tmplist *[]int, nums []int, start int) {
	clone := make([]int, len(*tmplist))
	copy(clone, *tmplist)
	*result = append(*result, clone)

	for i := start; i < len(nums); i++ {
		*tmplist = append(*tmplist, nums[i])
		backtrackSubsetsRef(result, tmplist, nums, i+1)
		*tmplist = (*tmplist)[:len(*tmplist)-1]
	}
}

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	tmplist := make([]int, 0)
	sort.Ints(nums)
	backtrackSubsetsWithDupRef(&result, &tmplist, nums, 0)
	return result
}

func backtrackSubsetsWithDupRef(result *[][]int, tmplist *[]int, nums []int, start int) {
	tmplistClone := append([]int{}, *tmplist...)
	*result = append(*result, tmplistClone)

	for i := start; i < len(nums); i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		*tmplist = append(*tmplist, nums[i])
		backtrackSubsetsWithDupRef(result, tmplist, nums, i+1)
		*tmplist = (*tmplist)[:len(*tmplist)-1]
	}
}

func maxProfit(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	if 2*k > len(prices) {
		result := 0
		for i := 0; i < len(prices)-1; i++ {
			if prices[i+1] > prices[i] {
				result += prices[i+1] - prices[i]
			}
		}
		return result
	}

	dp := make([][][]int, len(prices)+1)
	for i := range dp {
		dp[i] = make([][]int, k+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	dp[0][0][1] = math.MinInt32

	for i := 1; i < len(prices)+1; i++ {
		dp[i][0][1] = max(dp[i-1][0][1], -prices[i-1])
	}

	for i := 1; i <= k; i++ {
		dp[0][i][1] = -prices[0]
	}

	for i := 1; i < len(prices)+1; i++ {
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j-1][1]+prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j][0]-prices[i-1])
		}
	}

	return dp[len(prices)][k][0]
}

func recoverFromPreorder(S string) *TreeNode {
	nodeQueue := queue.New()
	depthQueue := queue.New()
	generateNodeQueues(S, nodeQueue, depthQueue)
	if nodeQueue.Len() == 0 {
		return nil
	}

	nodeStack := stack.New()
	depthStack := stack.New()

	root := nodeQueue.Dequeue().(*TreeNode)
	depth := depthQueue.Dequeue().(int)
	nodeStack.Push(root)
	depthStack.Push(depth)

	for nodeQueue.Len() > 0 {
		node := nodeQueue.Dequeue().(*TreeNode)
		depth = depthQueue.Dequeue().(int)

		var parent *TreeNode
		parentDepth := -1
		for true {
			parent = nodeStack.Pop().(*TreeNode)
			parentDepth = depthStack.Pop().(int)
			if parentDepth == depth-1 {
				break
			}
		}
		if parent.Left == nil {
			parent.Left = node
			nodeStack.Push(parent)
			depthStack.Push(parentDepth)
		} else {
			parent.Right = node
		}
		nodeStack.Push(node)
		depthStack.Push(depth)
	}
	return root

}

func generateNodeQueues(S string, nodeQueue, depthQueue *queue.Queue) {
	if len(S) == 0 {
		return
	}
	isDigit := false
	depth := 0
	value := 0
	for i := range S {
		if '-' != S[i] {
			if isDigit {
				value = value*10 + int(S[i]-'0')
			} else {
				isDigit = true
				value = int(S[i] - '0')
			}
		} else {
			if !isDigit {
				depth++
			} else {
				nodeQueue.Enqueue(&TreeNode{Val: value})
				depthQueue.Enqueue(depth)
				isDigit = false
				depth = 1
			}
		}
	}
	nodeQueue.Enqueue(&TreeNode{Val: value})
	depthQueue.Enqueue(depth)
}

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	result := 0
	for left < right {
		minH := -1
		width := right - left
		if height[left] <= height[right] {
			minH = height[left]
			left++
		} else {
			minH = height[right]
			right--
		}
		result = max(result, minH*width)
	}
	return result
}

func intToRoman(num int) string {
	M := []string{"", "M", "MM", "MMM"}
	C := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	X := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	I := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

	return M[num/1000] + C[(num%1000)/100] + X[(num%100)/10] + I[num%10]
}

func romanToInt(s string) int {
	roman := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	result := 0
	for i := 0; i < len(s)-1; i++ {
		if roman[rune(s[i])] < roman[rune(s[i+1])] {
			result -= roman[rune(s[i])]
		} else {
			result += roman[rune(s[i])]
		}
	}
	return result + roman[rune(s[len(s)-1])]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	return longestCommonPrefixDC(strs, 0, len(strs)-1)
}

func longestCommonPrefixDC(strs []string, l, r int) string {
	if l == r {
		return strs[l]
	}

	mid := (l + r) / 2
	left := longestCommonPrefixDC(strs, l, mid)
	right := longestCommonPrefixDC(strs, mid+1, r)
	return commonPrefix(left, right)
}

func commonPrefix(left, right string) string {
	minLength := min(len(left), len(right))
	for i := 0; i < minLength; i++ {
		if left[i] != right[i] {
			return left[:i]
		}
	}
	return left[:minLength]
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			left := i + 1
			right := len(nums) - 1
			for left < right {
				if nums[i]+nums[left]+nums[right] > 0 {
					right--
				} else if nums[i]+nums[left]+nums[right] < 0 {
					left++
				} else {
					ans := []int{nums[i], nums[left], nums[right]}
					result = append(result, ans)
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				}
			}
		}
	}
	return result
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	result := math.MaxInt32
	distance := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if abs(sum-target) < distance {
				result = sum
				distance = abs(sum - target)
			}
			if sum == target {
				return sum
			}
			if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}

func letterCombinations(digits string) []string {
	phone := map[uint8][]string{
		'2': []string{"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': []string{"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	letterCombinationsRecursive(&result, phone, digits, 0, "")
	return result
}

func letterCombinationsRecursive(result *[]string, phone map[uint8][]string, digits string, pos int, temp string) {
	if pos == len(digits) {
		*result = append(*result, temp)
		return
	}

	for _, s := range phone[digits[pos]] {
		temp += s
		letterCombinationsRecursive(result, phone, digits, pos+1, temp)
		temp = temp[:len(temp)-1]
	}
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			for j := i + 1; j < len(nums)-2; j++ {
				if j == i+1 || nums[j] != nums[j-1] {
					left := j + 1
					right := len(nums) - 1
					for left < right {
						sum := nums[i] + nums[j] + nums[left] + nums[right]
						if sum > target {
							right--
						} else if sum < target {
							left++
						} else {
							ans := []int{nums[i], nums[j], nums[left], nums[right]}
							result = append(result, ans)
							for left < right && nums[left] == nums[left+1] {
								left++
							}
							for left < right && nums[right] == nums[right-1] {
								right--
							}
							left++
							right--
						}
					}
				}
			}
		}
	}
	return result
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	first := dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	second := dummy
	for first.Next != nil {
		first = first.Next
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func isValid(s string) bool {
	myStack := stack.New()
	for i := range s {
		if s[i] == '(' {
			myStack.Push(')')
		} else if s[i] == '[' {
			myStack.Push(']')
		} else if s[i] == '{' {
			myStack.Push('}')
		} else if myStack.Len() == 0 || myStack.Peek().(rune) != rune(s[i]) {
			return false
		} else {
			myStack.Pop()
		}
	}
	return myStack.Len() == 0
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			prev.Next = l2
			l2 = l2.Next
		} else {
			prev.Next = l1
			l1 = l1.Next
		}
		prev = prev.Next
	}
	if l1 != nil {
		prev.Next = l1
	} else {
		prev.Next = l2
	}
	return dummy.Next
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	generateParenthesisRecursive(&result, "", 0, 0, n)
	return result
}

func generateParenthesisRecursive(result *[]string, temp string, left, right, n int) {
	if len(temp) == 2*n {
		*result = append(*result, temp)
		return
	}
	if left < n {
		generateParenthesisRecursive(result, temp+"(", left+1, right, n)
	}
	if right < left {
		generateParenthesisRecursive(result, temp+")", left, right+1, n)
	}
}

func mergeKLists(lists []*ListNode) *ListNode {
	var pq ListNodePriorityQueue
	pq = lists

	//from end
	for i := pq.Len() - 1; i >= 0; i-- {
		if pq[i] == nil {
			pq.Remove(i)
		}
	}

	heap.Init(&pq)

	dummy := &ListNode{}
	point := dummy

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*ListNode)
		point.Next = node
		point = point.Next
		node = node.Next
		if node != nil {
			heap.Push(&pq, node)
		}
	}

	return dummy.Next
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	prev := dummy
	prev.Next = head
	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next
		prev.Next = second
		first.Next = second.Next
		second.Next = first
		prev = first
	}
	return dummy.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	count := 0
	for curr != nil && count != k {
		curr = curr.Next
		count++
	}
	if count == k {
		curr = reverseKGroup(curr, k)
		for count > 0 {
			tmp := &head.Next
			head.Next = curr
			curr = head
			head = *tmp
			count--
		}
		head = curr
	}
	return head
}

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	slow := 1
	fast := 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 0
	fast := 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		j := 0
		for ; j < len(needle); j++ {
			if needle[j] != haystack[i+j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}

func findSubstring(s string, words []string) []int {
	var result []int
	if len(words) == 0 {
		return result
	}

	wordLength := len(words[0])
	hashmap := make(map[string]int)
	for _, w := range words {
		hashmap[w]++
	}

	for i := 0; i < len(s); i++ {
		copymap := make(map[string]int)
		for k, v := range hashmap {
			copymap[k] = v
		}
		found := true
		for j := 1; j <= len(words); j++ {
			end := j*wordLength + i
			start := end - wordLength
			if end <= len(s) {
				sub := s[start:end]
				if val, ok := copymap[sub]; ok && val > 0 {
					copymap[sub]--
				} else {
					found = false
					break
				}
			} else {
				return result
			}
		}
		if found {
			result = append(result, i)
		}
	}

	return result
}

func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i+1] <= nums[i] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		swap(nums, i, j)
	}
	reverseInSlice(nums, i+1)
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func reverseInSlice(nums []int, start int) {
	i := start
	j := len(nums) - 1
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	iMap := make(map[int]int)
	for index, value := range inorder {
		iMap[value] = index
	}
	length := len(preorder)

	return build(preorder, 0, length, inorder, 0, length, iMap)
}

func build(preorder []int, pStart, pEnd int, inorder []int, iStart, iEnd int, iMap map[int]int) *TreeNode {
	if pStart == pEnd {
		return nil
	}
	root := &TreeNode{Val: preorder[pStart]}
	iIndex := iMap[preorder[pStart]]
	leftCount := iIndex - iStart
	left := build(preorder, pStart+1, pStart+leftCount+1, inorder, iStart, iIndex-1, iMap)
	right := build(preorder, pStart+leftCount+1, pEnd, inorder, iIndex+1, iEnd, iMap)
	root.Left = left
	root.Right = right
	return root
}

func movingCount(m int, n int, k int) int {
	if m == 0 || n == 0 {
		return 0
	}
	result := 0
	var queue [][]int
	queue = append(queue, []int{0, 0})
	visited := make(map[int]bool)
	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]
		result++
		row := pos[0]
		col := pos[1]
		visited[0] = true
		for i := 0; i < 4; i++ {
			newRow := row + directs[i][0]
			newCol := col + directs[i][1]
			if newRow < 0 || newRow >= m || newCol < 0 || newCol >= n || visited[newRow*n+newCol] {
				continue
			}
			if numberValue(newRow)+numberValue(newCol) <= k {
				queue = append(queue, []int{newRow, newCol})
				visited[newRow*n+newCol] = true
			}
		}
	}
	return result
}

var directs = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func numberValue(number int) int {
	var result int
	for number > 0 {
		result += number % 10
		number /= 10
	}
	return result
}

func verifyPostorder(postorder []int) bool {
	if len(postorder) < 3 {
		return true
	}

	root := postorder[len(postorder)-1]
	leftSplit := -1
	for i := len(postorder) - 2; i >= 0; i-- {
		if leftSplit == -1 {
			if postorder[i] < root {
				leftSplit = i
			}
		} else {
			if postorder[i] > root {
				return false
			}
		}
	}

	if verifyPostorder(postorder[:leftSplit+1]) && verifyPostorder(postorder[leftSplit+1:len(postorder)-1]) {
		return true
	}

	return false
}

func copyRandomList(head *Node) *Node {
	hashMap := make(map[*Node]*Node)

	var dummy Node
	dummy.Next = head
	var prev *Node
	for head != nil {
		copyHead := Node{Val: head.Val}
		if prev != nil {
			prev.Next = &copyHead
		}
		prev = &copyHead
		hashMap[head] = &copyHead
		head = head.Next
	}

	head = dummy.Next
	copyHead := hashMap[head]
	for head != nil {
		copyHead.Random = hashMap[head.Random]
		head = head.Next
		copyHead = copyHead.Next
	}

	return hashMap[dummy.Next]
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1

	//found most left
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	var mostLeft int
	if nums[left] == target {
		mostLeft = left
	} else {
		if left+1 < len(nums) && nums[left+1] == target {
			mostLeft = left + 1
		} else {
			return 0
		}
	}

	left, right = 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	var mostRight int
	if nums[right] == target {
		mostRight = right
	} else {
		if right-1 >= 0 && nums[right-1] == target {
			mostRight = right - 1
		} else {
			return 0
		}
	}

	return mostRight - mostLeft + 1
}

func reversePairs(nums []int) int {
	var count int
	mergeSort(nums, &count)
	return count
}

func mergeSort(nums []int, count *int) {
	if len(nums) <= 1 {
		return
	}
	left := nums[:len(nums)/2]
	right := nums[len(nums)/2:]
	mergeSort(left, count)
	mergeSort(right, count)

	index, leftIdx, rightIdx := 0, 0, 0
	for leftIdx < len(left) && rightIdx < len(right) {
		if right[rightIdx] < left[leftIdx] {
			nums[index] = right[rightIdx]
			rightIdx++
			(*count) += len(left) - leftIdx
		} else {
			nums[index] = left[leftIdx]
			leftIdx++
		}
		index++
	}

	for leftIdx < len(left) {
		nums[index] = left[leftIdx]
		index++
		leftIdx++
	}

	for rightIdx < len(right) {
		nums[index] = right[rightIdx]
		index++
		rightIdx++
	}
}

func lengthOfLongestSubstring2(s string) int {
	hashmap := make(map[byte]int)
	var result int
	left, right := 0, len(s)
	for right < len(s) {
		hashmap[s[right]]++
		if hashmap[s[right]] < 2 {
			right++
			continue
		}
		result = max(result, right-left)
		for {
			hashmap[s[left]]--
			if hashmap[s[left]] == 1 {
				left++
				break
			}
			left++
		}
		right++
	}
	result = max(result, right-left)
	return result
}

func oddEvenList(head *ListNode) *ListNode {
	dummyOdd := &ListNode{}
	dummyEven := &ListNode{}
	oddPrev, evenPrev := dummyOdd, dummyEven
	isOdd := true

	for head != nil {
		if isOdd {
			oddPrev.Next = head
			oddPrev = head
		} else {
			evenPrev.Next = head
			evenPrev = head
		}
		isOdd = !isOdd
		head = head.Next
	}
	oddPrev.Next = dummyEven.Next
	evenPrev.Next = nil
	return dummyOdd.Next

}
