package main

import (
	"reflect"
	"fmt"
)

// KSumPaths 题意见 https://www.geeksforgeeks.org/print-k-sum-paths-binary-tree/
// 相比于 KSumPathsFromRoot，这里变为对所有节点求 KSumPathsFromRoot
func KSumPaths(root *Node, k int) (paths [][]int) {
	if root == nil {
		return
	} else {
		kSumPathsFromRoot(root, k, &paths)

		if root.left != nil {
			kSumPathsFromRoot(root.left, k, &paths)
		}

		if root.right != nil {
			kSumPathsFromRoot(root.right, k, &paths)
		}

		return
	}

}

func kSumPathsFromRoot(root *Node, k int, tmpPaths *[][]int) {
	if root == nil {
		return
	} else {
		path := []int{root.value}
		search(k, tmpPaths, path, root.left)
		search(k, tmpPaths, path, root.right)
	}

	if root.left != nil {
		kSumPathsFromRoot(root.left, k, tmpPaths)
	}
	if root.right != nil {
		kSumPathsFromRoot(root.right, k, tmpPaths)
	}
}

func search(k int, paths *[][]int, path []int, root *Node) {

	hasSamePath := false
	for _, existedPath := range *paths {
		// 不统计已存在的路径，由于左右节点为空会造成相同的路径
		if reflect.DeepEqual(existedPath, path) {
			hasSamePath = true
			break
		}
	}

	if !hasSamePath {
		sum := sumSlice(path)
		if sum == k {
			fmt.Printf("Current paths: %v \n", *paths)
			*paths = append(*paths, path)
			fmt.Printf("Current paths: %v \n", *paths)
		}
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
