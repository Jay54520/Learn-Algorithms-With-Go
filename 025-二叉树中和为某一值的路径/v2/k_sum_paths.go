package main

import (
	"reflect"
)

var globalPaths [][]int

// KSumPaths 题意见 https://www.geeksforgeeks.org/print-k-sum-paths-binary-tree/
// 相比于 KSumPathsFromRoot，这里变为对所有节点求 KSumPathsFromRoot
func KSumPaths(root *Node, k int) (paths [][]int) {
	if root == nil {
		return
	} else {
		kSumPathsFromRoot(root, k)

		if root.left != nil {
			kSumPathsFromRoot(root.left, k)
		}

		if root.right != nil {
			kSumPathsFromRoot(root.right, k)
		}

		return globalPaths
	}

}

func kSumPathsFromRoot(root *Node, k int) {
	if root == nil {
		return
	} else {
		path := []int{root.value}
		search(k, path, root.left)
		search(k, path, root.right)
	}

	if root.left != nil {
		kSumPathsFromRoot(root.left, k)
	}
	if root.right != nil {
		kSumPathsFromRoot(root.right, k)
	}
}

func search(k int, path []int, root *Node) {

	hasSamePath := false
	for _, existedPath := range globalPaths {
		// 不统计已存在的路径，由于左右节点为空会造成相同的路径
		if reflect.DeepEqual(existedPath, path) {
			hasSamePath = true
			break
		}
	}

	if !hasSamePath {
		sum := sumSlice(path)
		if sum == k {
			globalPaths = append(globalPaths, path)
		}
	}

	if root == nil {
		return
	} else {
		// 创建 newPath，与 path 使用不同的数组
		newPath := make([]int, len(path))
		copy(newPath, path)

		newPath = append(newPath, root.value)
		search(k, newPath, root.left)
		search(k, newPath, root.right)
	}

}

func sumSlice(aSlice []int) (sum int) {
	for _, value := range aSlice {
		sum += value
	}
	return
}
