package main

import "reflect"

// KSumPathsFromRoot 题意见 https://www.geeksforgeeks.org/print-paths-root-specified-sum-binary-tree/
// 某个节点的路径 = 当前路径 + 这个节点的值 + 子节点的路径
// 循环：如上，并将符合条件的 path 加入到 paths，由于节点的值可以为负数，所以不能提前终止
// 终止条件：检验后，如果当前节点为空则终止函数
func KSumPathsFromRoot(root *Node, k int) (paths [][]int) {
	if root == nil {
		return
	}
	path := []int{root.value}
	search(k, &paths, path, root.left)
	search(k, &paths, path, root.right)
	return
}

func search(k int, paths *[][]int, path []int, root *Node) {

	hasSamePath := false
	for _, existedPath := range *paths {
		// 由于左右节点为空或节点值相同会造成相同的路径
		if reflect.DeepEqual(existedPath, path) {
			hasSamePath = true
			break
		}
	}

	if !hasSamePath {
		// ----------------检验与加入----------------
		sum := sumSlice(path)
		if sum == k {
			*paths = append(*paths, path)
		}
		// ----------------检验与加入----------------
	}

	if root == nil {
		return
	} else {
		path = append(path, root.value)
		search(k, paths, path, root.left)
		search(k, paths, path, root.right)
	}
}

func sumSlice(aSlice []int) (sum int) {
	for _, value := range aSlice {
		sum += value
	}
	return
}
