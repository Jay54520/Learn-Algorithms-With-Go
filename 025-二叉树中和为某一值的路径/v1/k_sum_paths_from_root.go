package main

import "reflect"

// KSumPathsFromRoot 题意见 https://www.geeksforgeeks.org/print-paths-root-specified-sum-binary-tree/
// 某个节点的路径 = 当前路径 + 这个节点的值 + 子节点的路径
// 循环：如上，并将符合条件的 path 加入到 paths，由于节点的值可以为负数，所以不能提前终止
// 终止条件：当前节点为空，检验路径是否符合要求，根据检验结果决定是否加入 paths
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
	if root == nil {
		for _, existedPath := range *paths {
			// 不统计已存在的路径，由于左右节点为空会造成相同的路径
			if reflect.DeepEqual(existedPath, path) {
				return
			}
		}

		sum := sumSlice(path)
		if sum == k {
			*paths = append(*paths, path)
		}
		return
	}

	path = append(path, root.value)
	search(k, paths, path, root.left)
	search(k, paths, path, root.right)
}

func sumSlice(aSlice []int) (sum int) {
	for _, value := range aSlice {
		sum += value
	}
	return
}
