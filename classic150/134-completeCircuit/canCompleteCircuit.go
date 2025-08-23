package _34_completeCircuit

// 134. 加油站
// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。
// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。
// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。
//
// 示例 1:
//
// 输入: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
// 输出: 3
// 解释:
// 从 3 号加油站(索引为 3 处)出发，可获得 4 升汽油。此时油箱有 = 0 + 4 = 4 升汽油
// 开往 4 号加油站，此时油箱有 4 - 1 + 5 = 8 升汽油
// 开往 0 号加油站，此时油箱有 8 - 2 + 1 = 7 升汽油
// 开往 1 号加油站，此时油箱有 7 - 3 + 2 = 6 升汽油
// 开往 2 号加油站，此时油箱有 6 - 4 + 3 = 5 升汽油
// 开往 3 号加油站，你需要消耗 5 升汽油，正好足够你返回到 3 号加油站。
// 因此，3 可为起始索引。
// 示例 2:
//
// 输入: gas = [2,3,4], cost = [3,4,3]
// 输出: -1
// 解释:
// 你不能从 0 号或 1 号加油站出发，因为没有足够的汽油可以让你行驶到下一个加油站。
// 我们从 2 号加油站出发，可以获得 4 升汽油。 此时油箱有 = 0 + 4 = 4 升汽油
// 开往 0 号加油站，此时油箱有 4 - 3 + 2 = 3 升汽油
// 开往 1 号加油站，此时油箱有 3 - 3 + 3 = 3 升汽油
// 你无法返回 2 号加油站，因为返程需要消耗 4 升汽油，但是你的油箱只有 3 升汽油。
// 因此，无论怎样，你都不可能绕环路行驶一周。
//
// 提示:
//
// gas.length == n
// cost.length == n
// 1 <= n <= 105
// 0 <= gas[i], cost[i] <= 104

// 直观的暴力解法 O(n^2)
func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	for i := 0; i < length; i++ {
		startPos := i
		gasLeft := gas[i] - cost[i]
		if gasLeft < 0 {
			continue
		}

		j := 1
		for ; j <= length; j++ {
			gasLeft += gas[(i+j)%length] - cost[(i+j)%length]
			if gasLeft < 0 {
				break
			}
		}
		if j == length+1 {
			return startPos
		}
	}
	return -1
}

// 解法2
// 若所有gas的和 大于等于  所有cost的和 就一定存在能走完一圈的起始位置
// 考虑 sum(gas) == sum(cost)
// 假设 在 0 - length 区间总共分了三段 S1[0,j], S2[j+1,k], S3[k+1, last]
// (1) S1 段范围内 gasSum 和 costSum 在 0 ~ j-1 范围内都满足 gasSum >= costSum 知道 j位置出现了 gasSum < costSum
//
//	由此可以知道 gasSum[0 ~ j-1] >= costSum[0 ~ j-1], 而 gasSum[0 ~ j] < costSum[0 ~ j]
//	可以推导知道 对于任意 m ∈ (0, j) gasSum[m ~ j] < costSum[m ~ j]。 因为前面的 gas[0]-cost[0],gas[1] - cost[1] 等
//	都在给后面留存非负油量, 如果没有前面部分的留存, 后面的[m ~ j] 区间内累积的有量只可能会更少，不可能会更多
//	因此，如果在 j 的位置出现了 gasSum < costSum , 下次开始位置就可以直接从 j+1 开始，
//	否则中间任务位置开始都会在j处出现 gasSum < costSum
//
// (2) 前面出现了S1 段范围内 gasSum[i ~ j] < costSum[i ~ j], 同样的在 S2段[j+1, k] 范围内也出现了 gasSum < costSum
//
//	假设 costSum[i - j] - gasSum[i ~ j]  = gap1
//	    costSum[j+1, k] - gasSum[j+1, k] = gap2
//	由于总油量满足  sum(gas) == sum(cost)  可以知道 gasSum[k+1, last] = costSum[k+1, last] + gap1 +gap2
//	那么从 k+1 位置开始，就一定能累及足够多的油量 让加油车可以正好走上一圈

// 时间复杂度O(n), 空间复杂度O(1) 击败: 91.42%
func canCompleteCircuit1(gas []int, cost []int) int {
	length := len(gas)
	sumGas, sumCost := 0, 0
	for i := 0; i < length; i++ {
		sumGas, sumCost = sumGas+gas[i], sumCost+cost[i]
	}
	if sumGas < sumCost {
		return -1
	}

	gasLeft, pos := 0, 0
	for i := 0; i < length; i++ {
		gasLeft += gas[i] - cost[i]
		if gasLeft < 0 {
			gasLeft, pos = 0, i+1
		}
	}
	return pos
}
