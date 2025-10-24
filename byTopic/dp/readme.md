参考LeetCode灵茶山艾府 [分享题单](https://leetcode.cn/discuss/post/tXLS3i/) 重点攻破DP算法 


## 一、入门 DP

### §1.1 爬楼梯

[70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/)  
[746. 使用最小花费爬楼梯](https://leetcode.cn/problems/min-cost-climbing-stairs/) 约 1500  
[3693. 爬楼梯 II](https://leetcode.cn/problems/climbing-stairs-ii/) 1560  
[377. 组合总和 Ⅳ](https://leetcode.cn/problems/combination-sum-iv/) 约 1600 这类题目的本质就是爬楼梯    
[2466. 统计构造好字符串的方案数](https://leetcode.cn/problems/count-ways-to-build-good-strings/) 1694    
[2266. 统计打字方案数](https://leetcode.cn/problems/count-number-of-texts/) 1857
[2533. 好二进制字符串的数量](https://leetcode.cn/problems/number-of-good-binary-strings/)（会员题）同 2466 题

### §1.2 打家劫舍

问：在 1:1 翻译的过程中，如何根据记忆化搜索，确定递推数组（DP 数组）的大小？为什么有时候要开 n+1 大小的数组，有时候要开 n+2 大小的数组？
> 答：看记忆化搜索的参数的范围（最小值和最大值）。例如 i 最小是 −1（递归边界也算），最大是 n−1（递归入口），那么一共有 n+1 个不同的 i，就需要开 n+1 大小的 DP 数组。如果 i 最小是 −2，最大是 n−1，一共有 n+2 个不同的 i，就需要开 n+2 大小的 DP 数组。

[198. 打家劫舍](https://leetcode.cn/problems/house-robber/)
[213. 打家劫舍 II](https://leetcode.cn/problems/house-robber-ii/) 环形
[2320. 统计放置房子的方式数](https://leetcode.cn/problems/count-number-of-ways-to-place-houses/) 1608
[740. 删除并获得点数](https://leetcode.cn/problems/delete-and-earn/)
[3186. 施咒的最大总伤害](https://leetcode.cn/problems/maximum-total-damage-with-spell-casting/) 1841
思维扩展：
[2140. 解决智力问题](https://leetcode.cn/problems/solving-questions-with-brainpower/) 1709

### §1.3 最大子数组和（最大子段和）

有两种做法：
> 1. 定义状态 f[i] 表示以 a[i] 结尾的最大子数组和，不和 i 左边拼起来就是 f[i]=a[i]，和 i 左边拼起来就是 f[i]=f[i−1]+a[i]，取最大值就得到了状态转移方程 f[i]=max(f[i−1],0)+a[i]，答案为 max(f)。这个做法也叫做 Kadane 算法。

> 2. 用 前缀和，转化成 121. 买卖股票的最佳时机。

[53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/) 可以看看我题解中的思考题
[2606. 找到最大开销的子字符串](https://leetcode.cn/problems/find-the-substring-with-maximum-cost/) 1422
[1749. 任意子数组和的绝对值的最大值](https://leetcode.cn/problems/maximum-absolute-sum-of-any-subarray/) 1542
[1191. K 次串联后最大子数组之和](https://leetcode.cn/problems/k-concatenation-maximum-sum/) 1748
[918. 环形子数组的最大和](https://leetcode.cn/problems/maximum-sum-circular-subarray/) 1777
[2321. 拼接数组的最大分数](https://leetcode.cn/problems/maximum-score-of-spliced-array/) 1791
思维扩展：
[152. 乘积最大子数组](https://leetcode.cn/problems/maximum-product-subarray/)
[1186. 删除一次得到子数组最大和](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/) 1799

思考题  
> 完成本章后，请思考：什么时候要返回 f[n]，什么时候要返回 max(f) ？

## 二、网格图 DP

> 对于一些二维 DP（例如背包、最长公共子序列），如果把 DP 矩阵画出来，其实状态转移可以视作在网格图上的移动。所以在学习相对更抽象的二维 DP 之前，做一些形象的网格图 DP 会让后续的学习更轻松（比如 0-1 背包的空间优化写法为什么要倒序遍历）。


### §2.1 基础

[64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)  
[62. 不同路径](https://leetcode.cn/problems/unique-paths/)  
[63. 不同路径 II](https://leetcode.cn/problems/unique-paths-ii/)  
[120. 三角形最小路径和](https://leetcode.cn/problems/triangle/)  
[3393. 统计异或值为给定值的路径数目](https://leetcode.cn/problems/count-paths-with-the-given-xor-value/)  
[931. 下降路径最小和](https://leetcode.cn/problems/minimum-falling-path-sum/)  
[2684. 矩阵中移动的最大次数](https://leetcode.cn/problems/maximum-number-of-moves-in-a-grid/)  
[3603. 交替方向的最小路径代价 II](https://leetcode.cn/problems/minimum-cost-path-with-alternating-directions-ii/)  
[2304. 网格中的最小路径代价](https://leetcode.cn/problems/minimum-path-cost-in-a-grid/)  
[1289. 下降路径最小和 II](https://leetcode.cn/problems/minimum-falling-path-sum-ii/)  
[3418. 机器人可以获得的最大金币数](https://leetcode.cn/problems/maximum-amount-of-money-robot-can-earn/)  

思维扩展：  
[1824. 最少侧跳次数](https://leetcode.cn/problems/minimum-sideway-jumps/) 1778  

### §2.2 进阶

[1594. 矩阵的最大非负积](https://leetcode.cn/problems/maximum-non-negative-product-in-a-matrix/) 1807  
[1301. 最大得分的路径数目](https://leetcode.cn/problems/number-of-paths-with-max-score/) 1853  
[3665. 统计镜子反射路径数目](https://leetcode.cn/problems/twisted-mirror-path-count/) 1883  
[2435. 矩阵中和能被 K 整除的路径](https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/) 1952  
[174. 地下城游戏](https://leetcode.cn/problems/dungeon-game/)  
[329. 矩阵中的最长递增路径](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/)  
[2328. 网格图中递增路径的数目](https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid/) 2001  
[2267. 检查是否有合法括号字符串路径](https://leetcode.cn/problems/check-if-there-is-a-valid-parentheses-string-path/) 2085  
[1937. 扣分后的最大得分](https://leetcode.cn/problems/maximum-number-of-points-with-cost/) 2106  
[3363. 最多可收集的水果数目](https://leetcode.cn/problems/find-the-maximum-number-of-fruits-collected/) 实际难度 2200  
[1463. 摘樱桃 II](https://leetcode.cn/problems/cherry-pickup-ii/)  
[741. 摘樱桃](https://leetcode.cn/problems/cherry-pickup/)  
[3459. 最长 V 形对角线段的长度](https://leetcode.cn/problems/length-of-longest-v-shaped-diagonal-segment/) 2337 状态设计  
[2510. 检查是否有路径经过相同数量的 0 和 1](https://leetcode.cn/problems/check-if-there-is-a-path-with-equal-number-of-0s-and-1s/)（会员题）  

## 三、背包


### §3.1 0-1背包

每个物品只能选一次

[416. 分割等和子集](https://leetcode.cn/problems/partition-equal-subset-sum/)  
[494. 目标和](https://leetcode.cn/problems/target-sum/)  
[2915. 和为目标值的最长子序列的长度](https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/) 1659  
[2787. 将一个数字表示成幂的和的方案数](https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/) 1818  
[3180. 执行操作可获得的最大总奖励 I](https://leetcode.cn/problems/maximum-total-reward-using-operations-i/) 1849  
进阶：

[474. 一和零](https://leetcode.cn/problems/ones-and-zeroes/) 二维背包  
[3489. 零数组变换 IV](https://leetcode.cn/problems/zero-array-transformation-iv/) 2068    
[3685. 含上限元素的子序列和](https://leetcode.cn/problems/subsequence-sum-after-capping-elements/) 2073  
[1049. 最后一块石头的重量 II](https://leetcode.cn/problems/last-stone-weight-ii/) 2092  
[1774. 最接近目标价格的甜点成本](https://leetcode.cn/problems/closest-dessert-cost/)  
[879. 盈利计划](https://leetcode.cn/problems/profitable-schemes/) 2204  
[3082. 求出所有子序列的能量和](https://leetcode.cn/problems/find-the-sum-of-the-power-of-all-subsequences/) 2242  
[956. 最高的广告牌](https://leetcode.cn/problems/tallest-billboard/) 2381  
[2518. 好分区的数目](https://leetcode.cn/problems/number-of-great-partitions/) 2415  
[2742. 给墙壁刷油漆](https://leetcode.cn/problems/painting-the-walls/) 2425  
[3287. 求出数组中最大序列值](https://leetcode.cn/problems/find-the-maximum-sequence-value-of-array/) 2545  
[3181. 执行操作可获得的最大总奖励 II](https://leetcode.cn/problems/maximum-total-reward-using-operations-ii/) 2688 bitset 优化  
[LCP 47. 入场安检](https://leetcode.cn/problems/oPs9Bm/)  
[2291. 最大股票收益](https://leetcode.cn/problems/maximum-profit-from-trading-stocks/)（会员题）  
[2431. 最大限度地提高购买水果的口味](https://leetcode.cn/problems/maximize-total-tastiness-of-purchased-fruits/)（会员题）  
[3647. 两个袋子中的最大重量](https://leetcode.cn/problems/maximum-weight-in-two-bags/)（会员题）多背包  
[2189. 建造纸牌屋的方法数](https://leetcode.cn/problems/number-of-ways-to-build-house-of-cards/)（会员题）  

### §3.2 完全背包

> 物品可以重复选，无个数限制。

[322. 零钱兑换](https://leetcode.cn/problems/coin-change/)  
[518. 零钱兑换 II](https://leetcode.cn/problems/coin-change-ii/)  
[279. 完全平方数](https://leetcode.cn/problems/perfect-squares/)  
[3610. 目标和所需的最小质数个数](https://leetcode.cn/problems/minimum-number-of-primes-to-sum-to-target/)（会员题）  
[3183. 达到总和的方法数量](https://leetcode.cn/problems/the-number-of-ways-to-make-the-sum/)（会员题）混合背包  

思维扩展：

[3592. 硬币面值还原](https://leetcode.cn/problems/inverse-coin-change/) 1701  
[1449. 数位成本和为目标值的最大数字](https://leetcode.cn/problems/form-largest-integer-with-digits-that-add-up-to-target/) 1927  

### §3.3 多重背包（选做）

> 物品可以重复选，有个数限制。 求方案数

[2585. 获得分数的方法数](https://leetcode.cn/problems/number-of-ways-to-earn-points/) 1910  
[3333. 找到初始输入字符串 II](https://leetcode.cn/problems/find-the-original-typed-string-ii/) 2629  
[2902. 和带限制的子多重集合的数目](https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/) 2759    

二进制优化  
[3489. 零数组变换 IV](https://leetcode.cn/problems/zero-array-transformation-iv/)  

### §3.4 分组背包

> 同一组内的物品至多/恰好选一个。

[1155. 掷骰子等于目标和的方法数](https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/) 1654  
[1981. 最小化目标值与所选元素的差](https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/) 2010  
[2218. 从栈中取出 K 个硬币的最大面值和](https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/) 2158  

### §3.5 树形背包（选做）

> 也叫树上背包、依赖背包等。

注：目前力扣只有无依赖的背包，时间复杂度为 O(nW 2)。如果有依赖，可以优化到 O(nW)。

[3562. 折扣价交易股票的最大利润](https://leetcode.cn/problems/maximum-profit-from-trading-stocks-with-discounts/) 2458

## 四、经典线性 DP

### §4.1 最长公共子序列（LCS）

讲解：最长公共子序列 编辑距离【基础算法精讲 19】

一般定义 f[i][j] 表示对 (s[:i],t[:j]) 的求解结果。

### §4.1.1 基础

[1143. 最长公共子序列](https://leetcode.cn/problems/longest-common-subsequence/)  
[583. 两个字符串的删除操作](https://leetcode.cn/problems/delete-operation-for-two-strings/)  
[712. 两个字符串的最小 ASCII 删除和](https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/)  
[72. 编辑距离](https://leetcode.cn/problems/edit-distance/)  
[1035. 不相交的线](https://leetcode.cn/problems/uncrossed-lines/) 1806  
[1458. 两个子序列的最大点积](https://leetcode.cn/problems/max-dot-product-of-two-subsequences/) 1824  

思维扩展：

[718. 最长重复子数组](https://leetcode.cn/problems/maximum-length-of-repeated-subarray/)  

### §4.1.2 进阶

[3290. 最高乘法得分](https://leetcode.cn/problems/maximum-multiplication-score/) 1692  
[115. 不同的子序列](https://leetcode.cn/problems/distinct-subsequences/)  
[3628. 插入一个字母的最大子序列数](https://leetcode.cn/problems/maximum-number-of-subsequences-after-one-inserting/) 1754  
[3316. 从原字符串里进行删除操作的最多次数](https://leetcode.cn/problems/find-maximum-removals-from-source-string/) 2062  
[1639. 通过给定词典构造目标字符串的方案数](https://leetcode.cn/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/) 2082  
[97. 交错字符串](https://leetcode.cn/problems/interleaving-string/)  
[1092. 最短公共超序列](https://leetcode.cn/problems/shortest-common-supersequence/)  
[44. 通配符匹配](https://leetcode.cn/problems/wildcard-matching/)  
[10. 正则表达式匹配](https://leetcode.cn/problems/regular-expression-matching/)  

思考题：
> 115 题的扩展。给定字符串 s 和 t，你可以在 s 的任意位置插入一个字母，插入后，s 最多有多少个子序列等于 t ？

## §4.2 最长递增子序列（LIS）

做法有很多：

* 枚举选哪个。（见讲解）a
* 二分。（见讲解）
* 计算 a 和把 a 排序后的数组 sortedA 的最长公共子序列。（用 LCS 求 LIS）
* 数据结构优化。（见 [2407 题](https://leetcode.cn/problems/longest-increasing-subsequence-ii/) ）

### §4.2.1 基础

[300. 最长递增子序列](https://leetcode.cn/problems/longest-increasing-subsequence/)  
[2826. 将三个组排序](https://leetcode.cn/problems/sorting-three-groups/) 1721  
[1671. 得到山形数组的最少删除次数](https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/) 1913  
[1964. 找出到每个位置为止最长的有效障碍赛跑路线](https://leetcode.cn/problems/find-the-longest-valid-obstacle-course-at-each-position/) 1933  
[2111. 使数组 K 递增的最少操作次数](https://leetcode.cn/problems/minimum-operations-to-make-the-array-k-increasing/) 1941  

### §4.2.2 进阶

[354. 俄罗斯套娃信封问题](https://leetcode.cn/problems/russian-doll-envelopes/)  
[1626. 无矛盾的最佳球队](https://leetcode.cn/problems/best-team-with-no-conflicts/) 2027  
[1691. 堆叠长方体的最大高度](https://leetcode.cn/problems/maximum-height-by-stacking-cuboids/) 2172  
[960. 删列造序 III](https://leetcode.cn/problems/delete-columns-to-make-sorted-iii/) 2247  
[2407. 最长递增子序列 II](https://leetcode.cn/problems/longest-increasing-subsequence-ii/) 2280  
[673. 最长递增子序列的个数](https://leetcode.cn/problems/number-of-longest-increasing-subsequence/)  
[1187. 使数组严格递增](https://leetcode.cn/problems/make-array-strictly-increasing/) 2316  
[1713. 得到子序列的最少操作次数](https://leetcode.cn/problems/minimum-operations-to-make-a-subsequence/) 2351 用 LIS 求 LCS  
[3288. 最长上升路径的长度](https://leetcode.cn/problems/length-of-the-longest-increasing-path/) 2450  

思维扩展：

[368. 最大整除子集](https://leetcode.cn/problems/largest-divisible-subset/) 约 1800  
[2901. 最长相邻不相等子序列 II](https://leetcode.cn/problems/longest-unequal-adjacent-groups-subsequence-ii/) 1899

思考题：
> 给定整数 k，构造一个数组 a，使得 a 恰好有 k 个最长递增子序列。

## 五、划分型 DP

### §5.1 判定能否划分

> 一般定义 f[i] 表示长为 i 的前缀 a[:i] 能否划分。  

> 枚举最后一个子数组的左端点 L，从 f[L] 转移到 f[i]，并考虑 a[L:i] 是否满足要求。

[2369. 检查数组是否存在有效划分](https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/) 1780  
[139. 单词拆分](https://leetcode.cn/problems/word-break/)  

### §5.2 最优划分

计算最少（最多）可以划分出多少段、最优划分得分等。  
一般定义 f[i] 表示长为 i 的前缀 a[:i] 在题目约束下，分割出的最少（最多）子数组个数（或者定义成分割方案数）。  
枚举最后一个子数组的左端点 L，从 f[L] 转移到 f[i]，并考虑 a[L:i] 对最优解的影响。  

[132. 分割回文串 II](https://leetcode.cn/problems/palindrome-partitioning-ii/)  
[2707. 字符串中的额外字符](https://leetcode.cn/problems/extra-characters-in-a-string/) 1736  
[3196. 最大化子数组的总成本](https://leetcode.cn/problems/maximize-total-cost-of-alternating-subarrays/) 1847 也有状态机 DP 做法  
[2767. 将字符串分割为最少的美丽子字符串](https://leetcode.cn/problems/partition-string-into-minimum-beautiful-substrings/) 1865  
[91. 解码方法](https://leetcode.cn/problems/decode-ways/)  
[639. 解码方法 II](https://leetcode.cn/problems/decode-ways-ii/)  
[LCR 165. 解密数字](https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/)
[1043. 分隔数组以得到最大和](https://leetcode.cn/problems/partition-array-for-maximum-sum/) 1916  
[3144. 分割字符频率相等的最少子字符串](https://leetcode.cn/problems/minimum-substring-partition-of-equal-character-frequency/) 1917  
[1416. 恢复数组](https://leetcode.cn/problems/restore-the-array/) 1920  
[2472. 不重叠回文子字符串的最大数目](https://leetcode.cn/problems/maximum-number-of-non-overlapping-palindrome-substrings/) 2013  
[1105. 填充书架](https://leetcode.cn/problems/filling-bookcase-shelves/) 2014  
[2547. 拆分数组的最小代价](https://leetcode.cn/problems/minimum-cost-to-split-an-array/) 2020  
[3578. 统计极差最大为 K 的分割方式数](https://leetcode.cn/problems/count-partitions-with-max-min-difference-at-most-k/) 2033 优化  
[2430. 对字母串可执行的最大删除数](https://leetcode.cn/problems/maximum-deletions-on-a-string/) 2102  
[2463. 最小移动总距离](https://leetcode.cn/problems/minimum-total-distance-traveled/) 2454  
[3579. 字符串转换需要的最小操作数](https://leetcode.cn/problems/minimum-steps-to-convert-string-with-operations/) 2493 可以做到   
[3500. 将数组分割为子数组的最小代价](https://leetcode.cn/problems/minimum-cost-to-divide-array-into-subarrays/) 2569 式子变形  
[2977. 转换字符串的最小成本 II](https://leetcode.cn/problems/minimum-cost-to-convert-string-ii/) 2696  
[3441. 变成好标题的最少代价](https://leetcode.cn/problems/minimum-cost-good-caption/) 2765  
[2052. 将句子分隔成行的最低成本](https://leetcode.cn/problems/minimum-cost-to-separate-sentence-into-rows/)（会员题）  
[2464. 有效分割中的最少子数组数目](https://leetcode.cn/problems/minimum-subarrays-in-a-valid-split/)（会员题）

### §5.3 约束划分个数

将数组分成（恰好/至多）k 个连续子数组，计算与这些子数组有关的最优值。  
一般定义 f[i][j] 表示将长为 j 的前缀 a[:j] 分成 i 个连续子数组所得到的最优解。
枚举最后一个子数组的左端点 L，从 f[i−1][L] 转移到 f[i][j]，并考虑 a[L:j] 对最优解的影响。  

> 注：对于恰好型划分 DP，可以通过控制内层循环的上下界，把时间复杂度从 O(nk) 优化至 O((n−k)k)。例如 3473 题。

[813. 最大平均值和的分组](https://leetcode.cn/problems/largest-sum-of-averages/) 1937  
[3599. 划分数组得到最小 XOR](https://leetcode.cn/problems/partition-array-to-minimize-xor/) 1955 做法不止一种  
[410. 分割数组的最大值](https://leetcode.cn/problems/split-array-largest-sum/)  
[1278. 分割回文串 III](https://leetcode.cn/problems/palindrome-partitioning-iii/) 1979  
[1745. 分割回文串 IV](https://leetcode.cn/problems/palindrome-partitioning-iv/)  
[1335. 工作计划的最低难度](https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/) 2035  
[1473. 粉刷房子 III](https://leetcode.cn/problems/paint-house-iii/) 2056  
[2209. 用地毯覆盖后的最少白色砖块](https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets/) 2106  
[1478. 安排邮筒](https://leetcode.cn/problems/allocate-mailboxes/) 2190  
[3473. 长度至少为 M 的 K 个子数组之和](https://leetcode.cn/problems/sum-of-k-subarrays-with-length-at-least-m/) 2274 优化  
[1959. K 次调整数组大小浪费的最小总空间](https://leetcode.cn/problems/minimum-total-space-wasted-with-k-resizing-operations/) 2310  
[2478. 完美分割的方案数](https://leetcode.cn/problems/number-of-beautiful-partitions/) 2344  
[3538. 合并得到最小旅行时间](https://leetcode.cn/problems/merge-operations-for-minimum-travel-time/) 2461 相邻相关  
[3505. 使 K 个子数组内元素相等的最少操作数](https://leetcode.cn/problems/minimum-operations-to-make-elements-within-k-subarrays-equal/) 2539  
[3077. K 个不相交子数组的最大能量值](https://leetcode.cn/problems/maximum-strength-of-k-disjoint-subarrays/) 2557 优化  
[2911. 得到 K 个半回文串的最少修改次数](https://leetcode.cn/problems/minimum-changes-to-make-k-semi-palindromes/) 2608  
[3117. 划分数组得到最小的值之和](https://leetcode.cn/problems/minimum-sum-of-values-by-dividing-array/) 2735  
      
## 六、状态机 DP

一般定义 f[i][j] 表示前缀 a[:i] 在状态 j 下的最优值。j 一般很小。

### §6.1 买卖股票

讲解【基础算法精讲 21】
[121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/) 交易一次  
[122. 买卖股票的最佳时机 II](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/) 交易次数不限  
[123. 买卖股票的最佳时机 III](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/) 交易两次  
[188. 买卖股票的最佳时机 IV](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/) 交易   
[3573. 买卖股票的最佳时机 V](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-v/) 1777 交易   
[309. 买卖股票的最佳时机含冷冻期](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/) 交易次数不限  
[714. 买卖股票的最佳时机含手续费](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/) 交易次数不限

### §6.2 基础

[3259. 超级饮料的最大强化能量](https://leetcode.cn/problems/maximum-energy-boost-from-two-drinks/) 1484  
[2222. 选择建筑的方案数](https://leetcode.cn/problems/number-of-ways-to-select-buildings/) 1657  
[2708. 一个小组的最大实力值](https://leetcode.cn/problems/maximum-strength-of-a-group/) 做到   
[1567. 乘积为正数的最长子数组长度](https://leetcode.cn/problems/maximum-length-of-subarray-with-positive-product/) 1710  
[2786. 访问数组中的位置使分数最大](https://leetcode.cn/problems/visit-array-positions-to-maximize-score/) 1733  
[1911. 最大交替子序列和](https://leetcode.cn/problems/maximum-alternating-subsequence-sum/) 1786  
[376. 摆动序列](https://leetcode.cn/problems/wiggle-subsequence/)  
[3466. 最大硬币收集量](https://leetcode.cn/problems/maximum-coin-collection/)（会员题）

### §6.3 进阶

[3628. 插入一个字母的最大子序列数](https://leetcode.cn/problems/maximum-number-of-subsequences-after-one-inserting/) 1754  
[2771. 构造最长非递减子数组](https://leetcode.cn/problems/longest-non-decreasing-subarray-from-two-arrays/) 1792  
[1186. 删除一次得到子数组最大和](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/) 1799  
[1594. 矩阵的最大非负积](https://leetcode.cn/problems/maximum-non-negative-product-in-a-matrix/) 1807  
[3196. 最大化子数组的总成本](https://leetcode.cn/problems/maximize-total-cost-of-alternating-subarrays/) 1847 也有划分型 DP 做法  
[935. 骑士拨号器](https://leetcode.cn/problems/knight-dialer/)  
[1537. 最大得分](https://leetcode.cn/problems/get-the-maximum-score/) 1961  
[2919. 使数组变美的最小增量运算数](https://leetcode.cn/problems/minimum-increment-operations-to-make-array-beautiful/) 2031  
[801. 使序列递增的最小交换次数](https://leetcode.cn/problems/minimum-swaps-to-make-sequences-increasing/) 2066  
[3434. 子数组操作后的最大频率](https://leetcode.cn/problems/maximum-frequency-after-subarray-operation/) 2094  
[1955. 统计特殊子序列的数目](https://leetcode.cn/problems/count-number-of-special-subsequences/) 2125  
[3068. 最大节点价值之和](https://leetcode.cn/problems/find-the-maximum-sum-of-node-values/) 2268  
[3640. 三段式数组 II](https://leetcode.cn/problems/trionic-array-ii/) 2278  
[2272. 最大波动的子字符串](https://leetcode.cn/problems/substring-with-largest-variance/) 2516  
[3661. 可以被机器人摧毁的最大墙壁数目](https://leetcode.cn/problems/maximum-walls-destroyed-by-robots/) 2525  
[LCP 19. 秋叶收藏集](https://leetcode.cn/problems/UlBDOe/)  
[276. 栅栏涂色](https://leetcode.cn/problems/paint-fence/)（会员题）  
[1746. 经过一次操作后的最大子数组和](https://leetcode.cn/problems/maximum-subarray-sum-after-one-operation/)（会员题）  
[2036. 最大交替子数组和](https://leetcode.cn/problems/maximum-alternating-subarray-sum/)（会员题）  
[2361. 乘坐火车路线的最少费用](https://leetcode.cn/problems/minimum-costs-using-the-train-line/)（会员题）  
[3269. 构建两个递增数组](https://leetcode.cn/problems/constructing-two-increasing-arrays/)（会员题）

## 七、其他线性 DP

### §7.1 一维 DP

> 发生在前缀/后缀之间的转移，例如从 f[i−1] 转移到 f[i]，或者从 f[j] 转移到 f[i]。

[3147. 从魔法师身上吸取的最大能量](https://leetcode.cn/problems/taking-maximum-energy-from-the-mystic-dungeon/) 1460  
[2944. 购买水果需要的最少金币数](https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/) 1709  
[2140. 解决智力问题](https://leetcode.cn/problems/solving-questions-with-brainpower/) 1709 可以用刷表法  
[983. 最低票价](https://leetcode.cn/problems/minimum-cost-for-tickets/) 1786 有   
[368. 最大整除子集](https://leetcode.cn/problems/largest-divisible-subset/) 约 1800  
[2901. 最长相邻不相等子序列 II](https://leetcode.cn/problems/longest-unequal-adjacent-groups-subsequence-ii/) 1899  
[650. 两个键的键盘](https://leetcode.cn/problems/2-keys-keyboard/) 约 2000  
[871. 最低加油次数](https://leetcode.cn/problems/minimum-number-of-refueling-stops/) 2074  
[32. 最长有效括号](https://leetcode.cn/problems/longest-valid-parentheses/)  
[2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/) 2219  
[2188. 完成比赛的最少时间](https://leetcode.cn/problems/minimum-time-to-finish-the-race/) 2315  
[2896. 执行操作使两个字符串相等](https://leetcode.cn/problems/apply-operations-to-make-two-strings-equal/) 做到   
[818. 赛车](https://leetcode.cn/problems/race-car/) 2392  
[3389. 使字符频率相等的最少操作次数](https://leetcode.cn/problems/minimum-operations-to-make-character-frequencies-equal/) 2940  
[3464. 正方形上的点之间的最大距离](https://leetcode.cn/problems/maximize-the-distance-between-points-on-a-square/)  
[3205. 最大数组跳跃得分 I](https://leetcode.cn/problems/maximum-array-hopping-score-i/)（会员题）有   
[1259. 不相交的握手](https://leetcode.cn/problems/handshakes-that-dont-cross/)（会员题）  
[2597. 美丽子集的数目](https://leetcode.cn/problems/the-number-of-beautiful-subsets/) 用 DP 解决  
[2638. 统计 K-Free 子集的总数](https://leetcode.cn/problems/count-the-number-of-k-free-subsets/)（会员题）上面这题的加强版  
      
### §7.2 不相交区间

[2830. 销售利润最大化](https://leetcode.cn/problems/maximize-the-profit-as-the-salesman/) 1851  
[2008. 出租车的最大盈利](https://leetcode.cn/problems/maximum-earnings-from-taxi/) 1872  
[2054. 两个最好的不重叠活动](https://leetcode.cn/problems/two-best-non-overlapping-events/) 1883  
[1235. 规划兼职工作](https://leetcode.cn/problems/maximum-profit-in-job-scheduling/) 2023 做法不止一种  
[1751. 最多可以参加的会议数目 II](https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii/) 2041  
[3414. 不重叠区间的最大得分](https://leetcode.cn/problems/maximum-score-of-non-overlapping-intervals/) 2723  

### §7.3 子数组 DP

[53. 最大子数组和](https://leetcode.cn/problems/maximum-subarray/)  
[152. 乘积最大子数组](https://leetcode.cn/problems/maximum-product-subarray/)  
[3524. 求出数组的 X 值 I](https://leetcode.cn/problems/find-x-value-of-array-i/) 2008 刷表法  
[3448. 统计可以被最后一个数位整除的子字符串数目](https://leetcode.cn/problems/count-substrings-divisible-by-last-digit/) 2387 刷表法  
      
思维扩展：

[2262. 字符串的总引力](https://leetcode.cn/problems/total-appeal-of-a-string/) 2033  
[828. 统计子串中的唯一字符](https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/) 2034  
[467. 环绕字符串中唯一的子字符串](https://leetcode.cn/problems/unique-substrings-in-wraparound-string/)  
     
### §7.4 合法子序列 DP

计算合法子序列的最长长度、个数、元素和等。  
一般定义 f[x] 表示以元素 x 结尾的合法子序列的最长长度/个数/元素和，从子序列的倒数第二个数转移过来。
注意这里的 x 不是下标，是元素值。如果 x 不是整数，或者值域范围很大，可以用哈希表代替数组。

[2501. 数组中最长的方波](https://leetcode.cn/problems/longest-square-streak-in-an-array/) 1480  
[1218. 最长定差子序列](https://leetcode.cn/problems/longest-arithmetic-subsequence-of-given-difference/) 1597  
[2826. 将三个组排序](https://leetcode.cn/problems/sorting-three-groups/) 1721 做法不止一种  
[1027. 最长等差数列](https://leetcode.cn/problems/longest-arithmetic-subsequence/) 1759  
[2370. 最长理想子序列](https://leetcode.cn/problems/longest-ideal-subsequence/) 1835  
[873. 最长的斐波那契子序列的长度](https://leetcode.cn/problems/length-of-longest-fibonacci-subsequence/) 1911  
[3686. 稳定子序列的数量](https://leetcode.cn/problems/number-of-stable-subsequences/) 1969  
[3202. 找出有效子序列的最大长度 II](https://leetcode.cn/problems/find-the-maximum-length-of-valid-subsequence-ii/) 1974  
[446. 等差数列划分 II - 子序列](https://leetcode.cn/problems/arithmetic-slices-ii-subsequence/) 求个数  
[3351. 好子序列的元素之和](https://leetcode.cn/problems/sum-of-good-subsequences/) 2086 求和  
[3041. 修改数组后最大化数组中的连续元素数目](https://leetcode.cn/problems/maximize-consecutive-elements-in-an-array-after-modification/) 2231  
[3409. 最长相邻绝对差递减子序列](https://leetcode.cn/problems/longest-subsequence-with-decreasing-adjacent-difference/) 2500 状态设计  
[3098. 求出所有子序列的能量和](https://leetcode.cn/problems/find-the-sum-of-subsequence-powers/) 2553  
[2901. 最长相邻不相等子序列 II](https://leetcode.cn/problems/longest-unequal-adjacent-groups-subsequence-ii/) 做到线性复杂度  
[3299. 连续子序列的和](https://leetcode.cn/problems/sum-of-consecutive-subsequences/)（会员题）  
      
思维扩展：  
[1048. 最长字符串链](https://leetcode.cn/problems/longest-string-chain/) 1599  
[940. 不同的子序列 II](https://leetcode.cn/problems/distinct-subsequences-ii/) 1985  
[1987. 不同的好子序列数目](https://leetcode.cn/problems/number-of-unique-good-subsequences/) 2422 思考方式同 940 题  

### §7.5 子矩形 DP

[3148. 矩阵中的最大得分](https://leetcode.cn/problems/maximum-difference-score-in-a-grid/) 1820  
[221. 最大正方形](https://leetcode.cn/problems/maximal-square/)  
[1277. 统计全为 1 的正方形子矩阵](https://leetcode.cn/problems/count-square-submatrices-with-all-ones/)  
[2088. 统计农场中肥沃金字塔的数目](https://leetcode.cn/problems/count-fertile-pyramids-in-a-land/) 2105  
[3197. 包含所有 1 的最小矩形面积 II](https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-ii/) 做到 O(mn)

### §7.6 多维 DP

> 不会设计状态？那你可要好好刷这一节了。

[2222. 选择建筑的方案数](https://leetcode.cn/problems/number-of-ways-to-select-buildings/) 1657 把   
[2826. 将三个组排序](https://leetcode.cn/problems/sorting-three-groups/) 1721 做法不止一种  
[2400. 恰好移动 k 步到达某一位置的方法数目](https://leetcode.cn/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/) 1751  
[1262. 可被三整除的最大和](https://leetcode.cn/problems/greatest-sum-divisible-by-three/) 1762  
[3332. 旅客可以得到的最多点数](https://leetcode.cn/problems/maximum-points-tourist-can-earn/) 1828  
[3176. 求出最长好子序列 I](https://leetcode.cn/problems/find-the-maximum-length-of-a-good-subsequence-i/) 1849  
[1269. 停在原地的方案数](https://leetcode.cn/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/) 1854  
[3250. 单调数组对的数目 I](https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-i/) 1898  
[3218. 切蛋糕的最小总开销 I](https://leetcode.cn/problems/minimum-cost-for-cutting-cake-i/) 也有贪心做法  
[3122. 使矩阵满足条件的最少操作次数](https://leetcode.cn/problems/minimum-number-of-operations-to-satisfy-conditions/) 1905  
[576. 出界的路径数](https://leetcode.cn/problems/out-of-boundary-paths/)  
[403. 青蛙过河](https://leetcode.cn/problems/frog-jump/)  
[1223. 掷骰子模拟](https://leetcode.cn/problems/dice-roll-simulation/) 2008  
[1320. 二指输入的的最小距离](https://leetcode.cn/problems/minimum-distance-to-type-a-word-using-two-fingers/) 2028 状态优化  
[3366. 最小数组和](https://leetcode.cn/problems/minimum-array-sum/) 2040  
[1575. 统计所有可行路径](https://leetcode.cn/problems/count-all-possible-routes/) 2055  
[3154. 到达第 K 级台阶的方案数](https://leetcode.cn/problems/find-number-of-ways-to-reach-the-k-th-stair/) 2071  
[2318. 不同骰子序列的数目](https://leetcode.cn/problems/number-of-distinct-roll-sequences/) 2090  
[3469. 移除所有数组元素的最小代价](https://leetcode.cn/problems/find-minimum-cost-to-remove-array-elements/) 2112 状态设计  
[2746. 字符串连接删减字母](https://leetcode.cn/problems/decremental-string-concatenation/) 2126  
[1444. 切披萨的方案数](https://leetcode.cn/problems/number-of-ways-of-cutting-a-pizza/) 2127  
[3320. 统计能获胜的出招序列数](https://leetcode.cn/problems/count-the-number-of-winning-sequences/) 2153  
[3429. 粉刷房子 IV](https://leetcode.cn/problems/paint-house-iv/) 2166  
[2896. 执行操作使两个字符串相等](https://leetcode.cn/problems/apply-operations-to-make-two-strings-equal/) 2172  
[1420. 生成数组](https://leetcode.cn/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/) 2176  
[3193. 统计逆序对的数目](https://leetcode.cn/problems/count-the-number-of-inversions/) 2266  
[1079. 活字印刷](https://leetcode.cn/problems/letter-tile-possibilities/) 计数 DP，做到   
[1866. 恰有 K 根木棍可以看到的排列数目](https://leetcode.cn/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/) 2333  
[2312. 卖木头块](https://leetcode.cn/problems/selling-pieces-of-wood/) 2363  
[3177. 求出最长好子序列 II](https://leetcode.cn/problems/find-the-maximum-length-of-a-good-subsequence-ii/) 2365  
[1884. 鸡蛋掉落-两枚鸡蛋](https://leetcode.cn/problems/egg-drop-with-2-eggs-and-n-floors/)  
[887. 鸡蛋掉落](https://leetcode.cn/problems/super-egg-drop/) 2377  
[920. 播放列表的数量](https://leetcode.cn/problems/number-of-music-playlists/) 2400  
[514. 自由之路](https://leetcode.cn/problems/freedom-trail/) 做到   
[3336. 最大公约数相等的子序列数量](https://leetcode.cn/problems/find-the-number-of-subsequences-with-equal-gcd/) 2403  
[1388. 3n 块披萨](https://leetcode.cn/problems/pizza-with-3n-slices/) 2410  
[903. DI 序列的有效排列](https://leetcode.cn/problems/valid-permutations-for-di-sequence/) 2433  
[1900. 最佳运动员的比拼回合](https://leetcode.cn/problems/the-earliest-and-latest-rounds-where-players-compete/) 2455  
[1531. 压缩字符串 II](https://leetcode.cn/problems/string-compression-ii/) 2576  
[1883. 准时抵达会议现场的最小跳过休息次数](https://leetcode.cn/problems/minimum-skips-to-arrive-at-meeting-on-time/) 2588 避免浮点运算的技巧  
[964. 表示数字的最少运算符](https://leetcode.cn/problems/least-operators-to-express-number/) 2594  
[3343. 统计平衡排列的数目](https://leetcode.cn/problems/count-number-of-balanced-permutations/) 2615 计数 DP  
[1787. 使所有区间的异或结果为零](https://leetcode.cn/problems/make-the-xor-of-all-segments-equal-to-zero/) 2640  
[3509. 最大化交错和为 K 的子序列乘积](https://leetcode.cn/problems/maximum-product-of-subsequences-with-an-alternating-sum-equal-to-k/) 2703  
[3441. 变成好标题的最少代价](https://leetcode.cn/problems/minimum-cost-good-caption/) 2765  
[3539. 魔法序列的数组乘积之和](https://leetcode.cn/problems/find-sum-of-array-product-of-magical-sequences/) 2694 计数 DP  
[2060. 同源字符串检测](https://leetcode.cn/problems/check-if-an-original-string-exists-given-two-encoded-strings/) 2804  
[2809. 使数组和小于等于 x 的最少时间](https://leetcode.cn/problems/minimum-time-to-make-array-sum-at-most-x/) 2979  
[3003. 执行操作后的最大分割数量](https://leetcode.cn/problems/maximize-the-number-of-partitions-after-operations/) 3039 如何分析时间复杂度  
[3225. 网格图操作后的最大分数](https://leetcode.cn/problems/maximum-score-from-grid-operations/) 3028 IOI2022 原题  
[LCP 57. 打地鼠](https://leetcode.cn/problems/ZbAuEH/)  
[LCP 43. 十字路口的交通](https://leetcode.cn/problems/Y1VbOX/)  
[LCP 65. 舒适的湿度](https://leetcode.cn/problems/3aqs1c/)  
[LCP 36. 最多牌组数](https://leetcode.cn/problems/Up5XYM/)  
[LCP 38. 守卫城堡](https://leetcode.cn/problems/7rLGCR/)  
[256. 粉刷房子](https://leetcode.cn/problems/paint-house/)（会员题）  
[265. 粉刷房子 II](https://leetcode.cn/problems/paint-house-ii/)（会员题）  
[3339. 查找 K 偶数数组的数量](https://leetcode.cn/problems/find-the-number-of-k-even-arrays/)（会员题）  
[568. 最大休假天数](https://leetcode.cn/problems/maximum-vacation-days/)（会员题）  
[1692. 计算分配糖果的不同方式](https://leetcode.cn/problems/count-ways-to-distribute-candies/)（会员题）  
[2143. 在两个数组的区间中选取数字](https://leetcode.cn/problems/choose-numbers-from-two-arrays-in-range/)（会员题）  
[3269. 构建两个递增数组](https://leetcode.cn/problems/constructing-two-increasing-arrays/)（会员题）

思维扩展：
[638. 大礼包](https://leetcode.cn/problems/shopping-offers/)

## 八、区间 DP

> 从数组的左右两端不断缩短，求解关于某段下标区间的最优值。  
> 一般定义 f[i][j] 表示下标区间 [i,j] 的最优值。

§8.1 最长回文子序列

[516. 最长回文子序列](https://leetcode.cn/problems/longest-palindromic-subsequence/)  
[730. 统计不同回文子序列](https://leetcode.cn/problems/count-different-palindromic-subsequences/)  
[1312. 让字符串成为回文串的最少插入次数](https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/) 1787  
[3472. 至多 K 次操作后的最长回文子序列](https://leetcode.cn/problems/longest-palindromic-subsequence-after-at-most-k-operations/) 1884  
[1771. 由子序列构造的最长回文串的长度](https://leetcode.cn/problems/maximize-palindrome-length-from-subsequences/) 2182  
[1682. 最长回文子序列 II](https://leetcode.cn/problems/longest-palindromic-subsequence-ii/)（会员题）  
[1216. 验证回文串 III](https://leetcode.cn/problems/valid-palindrome-iii/)（会员题）  
[1246. 删除回文子数组](https://leetcode.cn/problems/palindrome-removal/)（会员题）  

### §8.2 区间 DP

对于类似合法括号字符串（RBS）的消除问题，通常根据题意，会有如下性质：

可以消除相邻的匹配字符。
相邻匹配字符消除后，原本不相邻的字符会变成相邻，可以继续消除。换句话说，设子串 A=x+B+y，如果 x 和 y 是匹配的（可以消除），且子串 B 可以完全消除，那么子串 A 可以完全消除。
设子串 A=B+C，如果子串 B 和 C 可以完全消除，那么子串 A 可以完全消除。
满足上述性质的题目（例如 3563 题），可以用区间 DP 解决。

定义 f(i,j) 表示消除 s[i] 到 s[j] 的最优值。

根据性质 2，可以把 f(i,j) 缩小成子问题 f(i+1,j−1)。
根据性质 3，可以枚举子串 B 的右端点，即枚举 k=i+1,i+3,i+5,…,j−2，把 f(i,j) 划分成子问题 f(i,k) 和 f(k+1,j)。注意这里枚举 k 的步长是 2，因为每次消除 2 个字符，被消除的子串长度一定是偶数。
边界：f(i+1,i)，即空串。

答案：f(0,n−1)。

[5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)  
[647. 回文子串](https://leetcode.cn/problems/palindromic-substrings/)  
[3040. 相同分数的最大操作数目 II](https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii/) 1709  
[375. 猜数字大小 II](https://leetcode.cn/problems/guess-number-higher-or-lower-ii/)  
[1130. 叶值的最小代价生成树](https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values/) 1919  
[96. 不同的二叉搜索树](https://leetcode.cn/problems/unique-binary-search-trees/)  
[1770. 执行乘法运算的最大分数](https://leetcode.cn/problems/maximum-score-from-performing-multiplication-operations/) 2068  
[1547. 切棍子的最小成本](https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/) 2116  
[1039. 多边形三角剖分的最低得分](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/) 2130  
[1000. 合并石头的最低成本](https://leetcode.cn/problems/minimum-cost-to-merge-stones/) 2423  
[2019. 解出数学表达式的学生分数](https://leetcode.cn/problems/the-score-of-students-solving-math-expression/) 2584  
[面试题 08.14. 布尔运算](https://leetcode.cn/problems/boolean-evaluation-lcci/)  
[3563. 移除相邻字符后字典序最小的字符串](https://leetcode.cn/problems/lexicographically-smallest-string-after-adjacent-removals/) 2585  
[3277. 查询子数组最大异或值](https://leetcode.cn/problems/maximum-xor-score-subarray-queries/) 2693  
[87. 扰乱字符串](https://leetcode.cn/problems/scramble-string/)  
[312. 戳气球](https://leetcode.cn/problems/burst-balloons/)  
[664. 奇怪的打印机](https://leetcode.cn/problems/strange-printer/)  
[546. 移除盒子](https://leetcode.cn/problems/remove-boxes/) 同 CF1107E  
[471. 编码最短长度的字符串](https://leetcode.cn/problems/encode-string-with-shortest-length/)（会员题）  
[3018. 可处理的最大删除操作数 I](https://leetcode.cn/problems/maximum-number-of-removal-queries-that-can-be-processed-i/)（会员题）  
      
## 九、状态压缩 DP（状压 DP）

### §9.1 排列型状压 DP ① 相邻无关

学习指南：
从集合论到位运算，常见位运算技巧分类总结
教你一步步思考状压 DP：从记忆化搜索到递推
暴力做法是枚举所有排列，对每个排列计算和题目有关的值，时间复杂度（通常来说）是 O(n!)。可以解决 n≤10 的问题。

状压 DP 可以把时间复杂度（通常来说）优化至 O(n⋅2 n)。可以解决 n≤20 的问题。

一般有两种定义方式：

> 1.定义 f[S] 表示已经排列好的元素（下标）集合为 S 时，和题目有关的最优值。通过枚举当前位置要填的元素（下标）来转移。  
> 2.定义 f[S] 表示可以选的元素（下标）集合为 S 时，和题目有关的最优值。通过枚举当前位置要填的元素（下标）来转移。  
注：部分题目由于暴搜+剪枝也能过，难度分仅供参考。  

[526. 优美的排列](https://leetcode.cn/problems/beautiful-arrangement/)  
[3376. 破解锁的最少时间 I](https://leetcode.cn/problems/minimum-time-to-break-locks-i/)  
[1879. 两个数组最小的异或值之和](https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays/) 2145  
[2850. 将石头分散到网格图的最少移动次数](https://leetcode.cn/problems/minimum-moves-to-spread-stones-over-grid/)  
[1947. 最大兼容性评分和](https://leetcode.cn/problems/maximum-compatibility-score-sum/)  
[1799. N 次操作后的最大分数和](https://leetcode.cn/problems/maximize-score-after-n-operations/)  
[3533. 判断连接可整除性](https://leetcode.cn/problems/concatenated-divisibility/) 2257 字典序 暴搜  
[3530. 有向无环图中合法拓扑排序的最大利润](https://leetcode.cn/problems/maximum-profit-from-valid-topological-order-in-dag/) 2353 拓扑序  
[2172. 数组的最大与和](https://leetcode.cn/problems/maximum-and-sum-of-array/) 2392  
[2992. 自整除排列的数量](https://leetcode.cn/problems/number-of-self-divisible-permutations/)（会员题）  
[2403. 杀死所有怪物的最短时间](https://leetcode.cn/problems/minimum-time-to-kill-all-monsters/)（会员题）同 3376 题  
[1066. 校园自行车分配 II](https://leetcode.cn/problems/campus-bikes-ii/)（会员题）  
      
### §9.2 排列型状压 DP ② 相邻相关

一般定义 f[S][i] 表示未选（或者已选）的集合为 S，且上一个填的元素（下标）为 i 时，和题目有关的最优值。通过枚举当前位置要填的元素（下标）来转移。

时间复杂度（通常来说）是 O(n^2⋅2^n)。

[2741. 特别的排列](https://leetcode.cn/problems/special-permutations/) 2021  
[996. 平方数组的数目](https://leetcode.cn/problems/number-of-squareful-arrays/)  
[1681. 最小不兼容性](https://leetcode.cn/problems/minimum-incompatibility/) 2390 见我在官解下的评论  
[3615. 图中的最长回文路径](https://leetcode.cn/problems/longest-palindromic-path-in-graph/) 2463  
[3283. 吃掉所有兵需要的最多移动次数](https://leetcode.cn/problems/maximum-number-of-moves-to-kill-all-pawns/) 2473  
[3149. 找出分数最低的排列](https://leetcode.cn/problems/find-the-minimum-cost-array-permutation/) 2642  
      
### §9.3 旅行商问题（TSP）

本质上就是排列型 ②。

[943. 最短超级串](https://leetcode.cn/problems/find-the-shortest-superstring/) 2186  
[847. 访问所有节点的最短路径](https://leetcode.cn/problems/shortest-path-visiting-all-nodes/) 2201  
[LCP 13. 寻宝](https://leetcode.cn/problems/xun-bao/)  
[2247. K 条高速公路的最大旅行费用](https://leetcode.cn/problems/maximum-cost-of-trip-with-k-highways/)（会员题）  
[3568. 清理教室的最少移动](https://leetcode.cn/problems/minimum-moves-to-clean-the-classroom/) 2143  
[864. 获取所有钥匙的最短路径](https://leetcode.cn/problems/shortest-path-to-get-all-keys/) 2259  
     
### §9.4 子集状压 DP

一般定义 f[S] 表示未选（或者已选）的集合为 S 时，和题目有关的最优值。通过枚举 S（或者 S 的补集 ∁
U ​ S）的子集来转移。

时间复杂度（通常来说）是 O(3^n)，证明：

值得注意的是，枚举子集的子集还可以用「选或不选」来做，对于存在无效状态的情况，可以做到更优的时间复杂度。具体见 1349 题解 最后的写法。

[2305. 公平分发饼干](https://leetcode.cn/problems/fair-distribution-of-cookies/) 1887  
[1986. 完成任务的最少工作时间段](https://leetcode.cn/problems/minimum-number-of-work-sessions-to-finish-the-tasks/) 1995  
[3670. 没有公共位的整数最大乘积](https://leetcode.cn/problems/maximum-product-of-two-integers-with-no-common-bits/) 2234  
[1723. 完成所有工作的最短时间](https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs/) 2284  
[1655. 分配重复整数](https://leetcode.cn/problems/distribute-repeating-integers/) 2307  
[3444. 使数组包含目标值倍数的最少增量](https://leetcode.cn/problems/minimum-increments-for-target-multiples-in-an-array/) 2337  
[3575. 最大好子树分数](https://leetcode.cn/problems/maximum-good-subtree-score/) 2360  
[1349. 参加考试的最大学生数](https://leetcode.cn/problems/maximum-students-taking-exam/) 2386  
[1681. 最小不兼容性](https://leetcode.cn/problems/minimum-incompatibility/) 2390 有   
[2572. 无平方子集计数](https://leetcode.cn/problems/count-the-number-of-square-free-subsets/) 2420  
[1994. 好子集的数目](https://leetcode.cn/problems/the-number-of-good-subsets/) 2465  
[1494. 并行课程 II](https://leetcode.cn/problems/parallel-courses-ii/)  
[LCP 04. 覆盖](https://leetcode.cn/problems/broken-board-dominoes/)  
[LCP 53. 守护太空城](https://leetcode.cn/problems/EJvmW4/)  
[465. 最优账单平衡](https://leetcode.cn/problems/optimal-account-balancing/)（会员题）  
[2152. 穿过所有点的所需最少直线数量](https://leetcode.cn/problems/minimum-number-of-lines-to-cover-points/)（会员题）  

相关问题：  
[3594. 所有人渡河所需的最短时间](https://leetcode.cn/problems/minimum-time-to-transport-all-individuals/) 2604

### §9.5 SOS DP

子集和 DP（Sum Over Subsets DP，SOS DP），国内算法竞赛圈一般叫高维前缀和。

[3670. 没有公共位的整数最大乘积](https://leetcode.cn/problems/maximum-product-of-two-integers-with-no-common-bits/) 2234  
[2732. 找到矩阵中的好子集](https://leetcode.cn/problems/find-a-good-subset-of-the-matrix/)

### §9.6 其他状压 DP

[1411. 给 N x 3 网格图涂色的方案数](https://leetcode.cn/problems/number-of-ways-to-paint-n-3-grid/) 1845  
[698. 划分为k个相等的子集](https://leetcode.cn/problems/partition-to-k-equal-sum-subsets/)  
[2002. 两个回文子序列长度的最大乘积](https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-subsequences/) 1869  
[473. 火柴拼正方形](https://leetcode.cn/problems/matchsticks-to-square/)  
[1931. 用三种不同颜色为网格涂色](https://leetcode.cn/problems/painting-a-grid-with-three-different-colors/) 2170 关联题目：1411 题  
[1125. 最小的必要团队](https://leetcode.cn/problems/smallest-sufficient-team/) 2251  
[1434. 每个人戴不同帽子的方案数](https://leetcode.cn/problems/number-of-ways-to-wear-different-hats-to-each-other/) 2273  
[464. 我能赢吗](https://leetcode.cn/problems/can-i-win/)  
[691. 贴纸拼词](https://leetcode.cn/problems/stickers-to-spell-word/)  
[3276. 选择矩阵中单元格的最大得分](https://leetcode.cn/problems/select-cells-in-grid-with-maximum-score/) 2403 横看成岭侧成峰  
[1595. 连通两组点的最小成本](https://leetcode.cn/problems/minimum-cost-to-connect-two-groups-of-points/) 2538  
[1815. 得到新鲜甜甜圈的最多组数](https://leetcode.cn/problems/maximum-number-of-groups-getting-fresh-donuts/) 2559  
[1659. 最大化网格幸福感](https://leetcode.cn/problems/maximize-grid-happiness/) 2655  
[980. 不同路径 III](https://leetcode.cn/problems/unique-paths-iii/)   
[LCP 69. Hello LeetCode!](https://leetcode.cn/problems/rMeRt2/)  
[LCP 76. 魔法棋盘](https://leetcode.cn/problems/1ybDKD/)  
[LCP 82. 万灵之树](https://leetcode.cn/problems/cnHoX6/)  
[351. 安卓系统手势解锁](https://leetcode.cn/problems/android-unlock-patterns/)（会员题）  
[2184. 建造坚实的砖墙的方法数](https://leetcode.cn/problems/number-of-ways-to-build-sturdy-brick-wall/)（会员题）  
      
## 十、数位 DP

数位 DP v1.0 模板讲解  
数位 DP v2.0 模板讲解 上下界数位 DP  
下面是数位 DP v2.1 模板。相比 v2.0，不需要写 isNum 参数。  

注：只有上界约束的题目，相当于 low=0 或者 low=1。

```
# 代码示例：返回 [low, high] 中的恰好包含 target 个 0 的数字个数
# 比如 digitDP(0, 10, 1) == 2
# 要点：我们统计的是 0 的个数，需要区分【前导零】和【数字中的零】，前导零不能计入，而数字中的零需要计入

    def digitDP(low: int, high: int, target: int) -> int:
        low_s = list(map(int, str(low)))  # 避免在 dfs 中频繁调用 int()
        high_s = list(map(int, str(high)))
        n = len(high_s)
        diff_lh = n - len(low_s)

    @cache
    def dfs(i: int, cnt0: int, limit_low: bool, limit_high: bool) -> int:
        if cnt0 > target:
            return 0  # 不合法
        if i == n:
            return 1 if cnt0 == target else 0

        lo = low_s[i - diff_lh] if limit_low and i >= diff_lh else 0
        hi = high_s[i] if limit_high else 9

        res = 0
        d = lo
        # 通过 limit_low 和 i 可以判断能否不填数字，无需 is_num 参数
        # 如果前导零不影响答案，去掉这个 if block
        if limit_low and i < diff_lh:
            # 不填数字，上界不受约束
            res = dfs(i + 1, cnt0, True, False)
            d = 1
        for d in range(d, hi + 1):
            res += dfs(i + 1,
                       cnt0 + (1 if d == 0 else 0),  # 统计 0 的个数
                       limit_low and d == lo,
                       limit_high and d == hi)

        # res %= MOD
        return res

    return dfs(0, 0, True, True)
```

[2719. 统计整数数目](https://leetcode.cn/problems/count-of-integers/) 数位和  
[1399. 统计最大组的数目](https://leetcode.cn/problems/count-largest-group/) 数位和 非暴力做法  
[1742. 盒子中小球的最大数量](https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/) 非暴力做法  
[788. 旋转数字](https://leetcode.cn/problems/rotated-digits/) 非暴力做法  
[902. 最大为 N 的数字组合](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/) 1990  
[600. 不含连续 1 的非负整数](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/)  
[2376. 统计特殊整数](https://leetcode.cn/problems/count-special-integers/) 2120  
[357. 统计各位数字都不同的数字个数](https://leetcode.cn/problems/count-numbers-with-unique-digits/) 同 2376 题  
[1012. 至少有 1 位重复的数字](https://leetcode.cn/problems/numbers-with-repeated-digits/) 2230  
[3519. 统计逐位非递减的整数](https://leetcode.cn/problems/count-numbers-with-non-decreasing-digits/) 2246 进制转换  
[2827. 范围中美丽整数的数目](https://leetcode.cn/problems/number-of-beautiful-integers-in-the-range/) 2324  
[2999. 统计强大整数的数目](https://leetcode.cn/problems/count-the-number-of-powerful-integers/) 2351  
[2801. 统计范围内的步进数字数目](https://leetcode.cn/problems/count-stepping-numbers-in-range/) 2367  
[2843. 统计对称整数的数目](https://leetcode.cn/problems/count-symmetric-integers/) 非暴力做法  
[3352. 统计小于 N 的 K 可约简整数](https://leetcode.cn/problems/count-k-reducible-numbers-less-than-n/) 2451  
[3621. 位计数深度为 K 的整数数目 I](https://leetcode.cn/problems/number-of-integers-with-popcount-depth-equal-to-k-i/) 同 3352 题  
[3490. 统计美丽整数的数目](https://leetcode.cn/problems/count-beautiful-numbers/) 2502  
[1397. 找到所有好字符串](https://leetcode.cn/problems/find-all-good-strings/) 2667  
[1215. 步进数](https://leetcode.cn/problems/stepping-numbers/)（会员题）  
[1067. 范围内的数字计数](https://leetcode.cn/problems/digit-count-in-range/)（会员题）  
[3032. 统计各位数字都不同的数字个数 II](https://leetcode.cn/problems/count-numbers-with-unique-digits-ii/)（会员题）  
      
从低到高：
[3704. 统计和为 N 的无零数对](https://leetcode.cn/problems/count-no-zero-pairs-that-sum-to-n/) 2419

思维扩展：

[233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/)  
[面试题 17.06. 2 出现的次数](https://leetcode.cn/problems/number-of-2s-in-range-lcci/)  
[3677. 统计二进制回文数字的数目](https://leetcode.cn/problems/count-binary-palindromic-numbers/) 2223  
[3007. 价值和小于等于 K 的最大数字](https://leetcode.cn/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/) 2258 做法不止一种  
[3646. 下一个特殊回文数](https://leetcode.cn/problems/next-special-palindrome-number/) 2445 字典序 暴搜  
[3348. 最小可整除数位乘积 II](https://leetcode.cn/problems/smallest-divisible-digit-product-ii/) 3101 字典序 暴搜  
      
## 十一、优化 DP

### §11.1 前缀和优化 DP

[1871. 跳跃游戏 VII](https://leetcode.cn/problems/jump-game-vii/) 1896 也可以用滑动窗口优化  
[2327. 知道秘密的人数](https://leetcode.cn/problems/number-of-people-aware-of-a-secret/) 做到  O(n)  
[3699. 锯齿形数组的总数 I](https://leetcode.cn/problems/number-of-zigzag-arrays-i/) 2123  
[1997. 访问完所有房间的第一天](https://leetcode.cn/problems/first-day-where-you-have-been-in-all-the-rooms/) 2260  
[629. K 个逆序对数组](https://leetcode.cn/problems/k-inverse-pairs-array/)  
[3193. 统计逆序对的数目](https://leetcode.cn/problems/count-the-number-of-inversions/) 2266  
[3473. 长度至少为 M 的 K 个子数组之和](https://leetcode.cn/problems/sum-of-k-subarrays-with-length-at-least-m/) 2274 划分型  
[3251. 单调数组对的数目 II](https://leetcode.cn/problems/find-the-count-of-monotonic-pairs-ii/) 2323 也有组合数学做法  
[2478. 完美分割的方案数](https://leetcode.cn/problems/number-of-beautiful-partitions/) 2344  
[837. 新 21 点](https://leetcode.cn/problems/new-21-game/) 2350 也可以用滑动窗口优化  
[3077. K 个不相交子数组的最大能量值](https://leetcode.cn/problems/maximum-strength-of-k-disjoint-subarrays/) 2557 划分型  
[3333. 找到初始输入字符串 II](https://leetcode.cn/problems/find-the-original-typed-string-ii/) 2629 也有生成函数做法  
[2902. 和带限制的子多重集合的数目](https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/) 2759  
[1977. 划分数字的方案数](https://leetcode.cn/problems/number-of-ways-to-separate-numbers/) 2817  
[3130. 找出所有稳定的二进制数组 II](https://leetcode.cn/problems/find-all-possible-stable-binary-arrays-ii/) 2825 也有组合数学做法
    
### §11.2 单调栈优化 DP

前置题单：单调栈（矩形系列/字典序最小/贡献法）

[1335. 工作计划的最低难度](https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/) 2035  
[2866. 美丽塔 II](https://leetcode.cn/problems/beautiful-towers-ii/) 2072  
[2617. 网格图中最少访问的格子数](https://leetcode.cn/problems/minimum-number-of-visited-cells-in-a-grid/) 2582  
[2355. 你能拿走的最大图书数量](https://leetcode.cn/problems/maximum-number-of-books-you-can-take/)（会员题）  
      
### §11.3 单调队列优化 DP

一般用来维护一段转移来源的最值。  
前提：区间右端点变大时，左端点也在变大（同滑动窗口）。  
转移前，去掉队首无用数据。  
计算转移（直接从队首转移）。  
把数据（一般是 f[i]）插入队尾前，去掉队尾无用数据。  

[2944. 购买水果需要的最少金币数](https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/) 1709 做到O(n)   
[1696. 跳跃游戏 VI](https://leetcode.cn/problems/jump-game-vi/) 1954  
[1425. 带限制的子序列和](https://leetcode.cn/problems/constrained-subsequence-sum/) 2032  
[3578. 统计极差最大为 K 的分割方式数](https://leetcode.cn/problems/count-partitions-with-max-min-difference-at-most-k/) 2033  
[375. 猜数字大小 II](https://leetcode.cn/problems/guess-number-higher-or-lower-ii/) 做到做到 O(n2)   
[1687. 从仓库到码头运输箱子](https://leetcode.cn/problems/delivering-boxes-from-storage-to-ports/) 2610  
[2463. 最小移动总距离](https://leetcode.cn/problems/minimum-total-distance-traveled/) 做到O(nm)（注：还有 O((n+m)log(n+m)) 的反悔贪心做法）
[3117. 划分数组得到最小的值之和](https://leetcode.cn/problems/minimum-sum-of-values-by-dividing-array/) 2735  
[2945. 找到最大非递减数组的长度](https://leetcode.cn/problems/find-maximum-non-decreasing-array-length/) 2943  
[2969. 购买水果需要的最少金币数 II](https://leetcode.cn/problems/minimum-number-of-coins-for-fruits-ii/)（会员题）
      
### §11.4 树状数组/线段树优化 DP

[1626. 无矛盾的最佳球队](https://leetcode.cn/problems/best-team-with-no-conflicts/) 2027  
[2407. 最长递增子序列 II](https://leetcode.cn/problems/longest-increasing-subsequence-ii/) 2280  
[2770. 达到末尾下标所需的最大跳跃次数](https://leetcode.cn/problems/maximum-number-of-jumps-to-reach-the-last-index/) 见我题解下的   
[2926. 平衡子序列的最大和](https://leetcode.cn/problems/maximum-balanced-subsequence-sum/) 2448  
[2547. 拆分数组的最小代价](https://leetcode.cn/problems/minimum-cost-to-split-an-array/) 做到 O(nlogn)
[3671. 子序列美丽值求和](https://leetcode.cn/problems/sum-of-beautiful-subsequences/) 2647  
[2916. 子数组不同元素数目的平方和 II](https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-ii/) 2816 思路类似 2547 题  
      
### §11.5 字典树优化 DP

[139. 单词拆分](https://leetcode.cn/problems/word-break/)  
[2707. 字符串中的额外字符](https://leetcode.cn/problems/extra-characters-in-a-string/)  
[面试题 17.13. 恢复空格](https://leetcode.cn/problems/re-space-lcci/)  
[472. 连接词](https://leetcode.cn/problems/concatenated-words/) 约 2300  
[2977. 转换字符串的最小成本 II](https://leetcode.cn/problems/minimum-cost-to-convert-string-ii/) 2696  
      
### §11.6 矩阵快速幂优化 DP

有两种类型的矩阵快速幂优化 DP：
> 1. 线性 DP（常系数齐次线性递推），转移系数写在第一行，其余行 m  i+1,i =1，例如 70，1137。讲解。
> 2. 多维 DP 或者状态机 DP，转移系数可表示为一个 k×k 大小的矩阵，例如 935。讲解。

[70](https://leetcode.cn/problems/climbing-stairs/)，  
[935](https://leetcode.cn/problems/knight-dialer/)。  
[70. 爬楼梯](https://leetcode.cn/problems/climbing-stairs/)  
[509. 斐波那契数](https://leetcode.cn/problems/fibonacci-number/)  
[1137. 第 N 个泰波那契数](https://leetcode.cn/problems/n-th-tribonacci-number/)  
[1220. 统计元音字母序列的数目](https://leetcode.cn/problems/count-vowels-permutation/)  
[552. 学生出勤记录 II](https://leetcode.cn/problems/student-attendance-record-ii/)  
[935. 骑士拨号器](https://leetcode.cn/problems/knight-dialer/)  
[790. 多米诺和托米诺平铺](https://leetcode.cn/problems/domino-and-tromino-tiling/)  
[1411. 给 N x 3 网格图涂色的方案数](https://leetcode.cn/problems/number-of-ways-to-paint-n-3-grid/)  
[1931. 用三种不同颜色为网格涂色](https://leetcode.cn/problems/painting-a-grid-with-three-different-colors/)  
[3337. 字符串转换后的长度 II](https://leetcode.cn/problems/total-characters-in-string-after-transformations-ii/) 2412  
[3700. 锯齿形数组的总数 II](https://leetcode.cn/problems/number-of-zigzag-arrays-ii/) 2435  
[2851. 字符串转换](https://leetcode.cn/problems/string-transformation/) 2858  
[2912. 在网格上移动到目的地的方法数](https://leetcode.cn/problems/number-of-ways-to-reach-destination-in-the-grid/)（会员题）

### §11.7 斜率优化 DP

> 也叫凸包优化/凸壳优化（CHT，Convex Hull Trick）。

[3494. 酿造药水需要的最少总时间](https://leetcode.cn/problems/find-the-minimum-amount-of-time-to-brew-potions/)  
[3500. 将数组分割为子数组的最小代价](https://leetcode.cn/problems/minimum-cost-to-divide-array-into-subarrays/)  
[3693. 爬楼梯 II](https://leetcode.cn/problems/climbing-stairs-ii/) 与 K=3 无关的做法
      
### §11.8 WQS 二分优化 DP

把最多选 k 个物品的问题（时间复杂度高）转换成选任意个物品的问题（时间复杂度低）。

[188. 买卖股票的最佳时机 IV](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/) 做到 O(nlogU)
[2209. 用地毯覆盖后的最少白色砖块](https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets/) 做到 O(nlogn)
      
### §11.9 其他优化 DP

[3654. 删除可整除和后的最小数组和](https://leetcode.cn/problems/minimum-sum-after-divisible-sum-deletions/) 2039  
[1937. 扣分后的最大得分](https://leetcode.cn/problems/maximum-number-of-points-with-cost/) 2106  
[2713. 矩阵中严格递增的单元格数](https://leetcode.cn/problems/maximum-strictly-increasing-cells-in-a-matrix/) 2387  
[3651. 带传送的最小路径成本](https://leetcode.cn/problems/minimum-cost-path-with-teleportations/) 2411  
[2318. 不同骰子序列的数目](https://leetcode.cn/problems/number-of-distinct-roll-sequences/) 容斥优化（类似 CF1943D2）  
[3181. 执行操作可获得的最大总奖励 II](https://leetcode.cn/problems/maximum-total-reward-using-operations-ii/) 2688 bitset 优化  
[2267. 检查是否有合法括号字符串路径](https://leetcode.cn/problems/check-if-there-is-a-valid-parentheses-string-path/) 可以进一步地用 bitset 优化  
[3213. 最小代价构造字符串](https://leetcode.cn/problems/construct-string-with-minimum-cost/) 字符串哈希 / 后缀数组 / AC 自动机优化  
[3292. 形成目标字符串需要的最少字符串数 II](https://leetcode.cn/problems/minimum-number-of-valid-strings-to-form-target-ii/) 字符串哈希 / AC 自动机优化  
[LCP 14. 切分数组](https://leetcode.cn/problems/qie-fen-shu-zu/)  
[LCP 59. 搭桥过河](https://leetcode.cn/problems/NfY1m5/) Slope Trick  
[2263. 数组变为有序的最小操作次数](https://leetcode.cn/problems/make-array-non-decreasing-or-non-increasing/)（会员题）Slope Trick

## 十二、树形 DP

> 注：可能有同学觉得树形 DP 没有重复访问同一个状态（重叠子问题），并不能算作 DP，而是算作普通的递归。这么说也有一定道理，不过考虑到思维方式和 DP 是一样的自底向上，所以仍然叫做树形 DP。此外，如果是自顶向下的递归做法，是存在重叠子问题的，一般要结合记忆化搜索实现。

### §12.1 树的直径

[543. 二叉树的直径](https://leetcode.cn/problems/diameter-of-binary-tree/)  
[687. 最长同值路径](https://leetcode.cn/problems/longest-univalue-path/)  
[124. 二叉树中的最大路径和](https://leetcode.cn/problems/binary-tree-maximum-path-sum/)  
[2385. 感染二叉树需要的总时间](https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/) 1711  
[2246. 相邻字符不同的最长路径](https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/) 2126  
[3203. 合并两棵树后的最小直径](https://leetcode.cn/problems/find-minimum-diameter-after-merging-two-trees/) 2266  
[1617. 统计子树中城市之间最大距离](https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities/) 2309  
[2538. 最大价值和与最小价值和的差值](https://leetcode.cn/problems/difference-between-maximum-and-minimum-price-sum/) 2398  
[1522. N 叉树的直径](https://leetcode.cn/problems/diameter-of-n-ary-tree/)（会员题）  
[1245. 树的直径](https://leetcode.cn/problems/tree-diameter/)（会员题）  
[549. 二叉树最长连续序列 II](https://leetcode.cn/problems/binary-tree-longest-consecutive-sequence-ii/)（会员题）  
[3372. 连接两棵树后最大目标节点数目 I](https://leetcode.cn/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/) 算法优化  
      
> 注：求直径也有两次 DFS 的做法。

### §12.2 树上最大独立集

讲解：树形 DP：打家劫舍 III【基础算法精讲 24】

[337. 打家劫舍 III](https://leetcode.cn/problems/house-robber-iii/)（没有上司的舞会）  
[2646. 最小化旅行的价格总和](https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/) 2238  
[3544. 子树反转和](https://leetcode.cn/problems/subtree-inversion-sum/) 2545 记忆化搜索 / 树上刷表法  
[2378. 选择边来最大化树的得分](https://leetcode.cn/problems/choose-edges-to-maximize-score-in-a-tree/)（会员题）

### §12.3 树上最小支配集

讲解：树形 DP：监控二叉树【基础算法精讲 25】，包含 968 的变形题。

[968. 监控二叉树](https://leetcode.cn/problems/binary-tree-cameras/) 2124
     
### §12.4 换根 DP

也叫二次扫描法。

[834. 树中距离之和](https://leetcode.cn/problems/sum-of-distances-in-tree/) 2197  
[2581. 统计可能的树根数目](https://leetcode.cn/problems/count-number-of-possible-root-nodes/) 2228  
[2858. 可以到达每一个节点的最少边反转次数](https://leetcode.cn/problems/minimum-edge-reversals-so-every-node-is-reachable/) 2295  
[310. 最小高度树](https://leetcode.cn/problems/minimum-height-trees/) 也可以用拓扑排序做  
[3241. 标记所有节点需要的时间](https://leetcode.cn/problems/time-taken-to-mark-all-nodes/) 2522  
[3372. 连接两棵树后最大目标节点数目 I](https://leetcode.cn/problems/maximize-the-number-of-target-nodes-after-connecting-trees-i/) 做到 O((n+m)k)，解决 n 和 m 很大但 k 很小的场景

> 注：前后缀分解，可以视作一条链上的换根 DP。

### §12.5 其他树形 DP

[2925. 在树上执行操作以后得到的最大分数](https://leetcode.cn/problems/maximum-score-after-applying-operations-on-a-tree/) 1940  
[3068. 最大节点价值之和](https://leetcode.cn/problems/find-the-maximum-sum-of-node-values/) 2268  
[2920. 收集所有金币可获得的最大积分](https://leetcode.cn/problems/maximum-points-after-collecting-coins-from-all-nodes/) 2351 记忆化搜索  
[3575. 最大好子树分数](https://leetcode.cn/problems/maximum-good-subtree-score/) 2360  
[3562. 折扣价交易股票的最大利润](https://leetcode.cn/problems/maximum-profit-from-trading-stocks-with-discounts/) 2458 背包  
[1916. 统计为蚁群构筑房间的不同顺序](https://leetcode.cn/problems/count-ways-to-build-rooms-in-an-ant-colony/) 2486  
[3367. 移除边之后的权重最大和](https://leetcode.cn/problems/maximize-sum-of-weights-after-edge-removals/) 2602  
[LCP 10. 二叉树任务调度](https://leetcode.cn/problems/er-cha-shu-ren-wu-diao-du/)  
[LCP 34. 二叉树染色](https://leetcode.cn/problems/er-cha-shu-ran-se-UGC/)  
[LCP 64. 二叉树灯饰](https://leetcode.cn/problems/U7WvvU/)  
[2313. 二叉树中得到结果所需的最少翻转次数](https://leetcode.cn/problems/minimum-flips-in-binary-tree-to-get-result/)（会员题）  
      
## 十三、图 DP

[3243. 新增道路查询后的最短距离 I](https://leetcode.cn/problems/shortest-distance-after-road-addition-queries-i/) 1568  
[787. K 站中转内最便宜的航班](https://leetcode.cn/problems/cheapest-flights-within-k-stops/) 1786  
[1786. 从第一个节点出发到最后一个节点的受限路径数](https://leetcode.cn/problems/number-of-restricted-paths-from-first-to-last-node/) 2079  
[2050. 并行课程 III](https://leetcode.cn/problems/parallel-courses-iii/) 2084   
[3620. 恢复网络路径](https://leetcode.cn/problems/network-recovery-pathways/)   
[1976. 到达目的地的方案数](https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination/) 2095  
[3543. K 条边路径的最大边权和](https://leetcode.cn/problems/maximum-weighted-k-edge-path/) 2110   
[1857. 有向图中最大颜色值](https://leetcode.cn/problems/largest-color-value-in-a-directed-graph/) 2313   
[1928. 规定时间内到达终点的最小花费](https://leetcode.cn/problems/minimum-cost-to-reach-destination-in-time/) 2413  
[913. 猫和老鼠](https://leetcode.cn/problems/cat-and-mouse/) 2800 拓扑序  
[1728. 猫和老鼠 II](https://leetcode.cn/problems/cat-and-mouse-ii/) 2849 拓扑序  
[LCP 07. 传递信息](https://leetcode.cn/problems/chuan-di-xin-xi/)  
[1548. 图中最相似的路径](https://leetcode.cn/problems/the-most-similar-path-in-a-graph/)（会员题）  
      
另见【题单】图论算法 中的「全源最短路：Floyd」，本质是多维 DP。

## 十四、博弈 DP

[1025. 除数博弈](https://leetcode.cn/problems/divisor-game/) 1435 有数学做法  
[877. 石子游戏](https://leetcode.cn/problems/stone-game/) 1590 有数学做法  
[486. 预测赢家](https://leetcode.cn/problems/predict-the-winner/)  
[1510. 石子游戏 IV](https://leetcode.cn/problems/stone-game-iv/) 1787  
[1690. 石子游戏 VII](https://leetcode.cn/problems/stone-game-vii/) 1951  
[1406. 石子游戏 III](https://leetcode.cn/problems/stone-game-iii/) 2027  
[1140. 石子游戏 II](https://leetcode.cn/problems/stone-game-ii/) 2035  
[1563. 石子游戏 V](https://leetcode.cn/problems/stone-game-v/) 2087  
[464. 我能赢吗](https://leetcode.cn/problems/can-i-win/)  
[1872. 石子游戏 VIII](https://leetcode.cn/problems/stone-game-viii/) 2440  
[913. 猫和老鼠](https://leetcode.cn/problems/cat-and-mouse/) 2800  
[1728. 猫和老鼠 II](https://leetcode.cn/problems/cat-and-mouse-ii/) 2849  
[294. 翻转游戏 II](https://leetcode.cn/problems/flip-game-ii/)（会员题）

## 十五、概率/期望 DP

[688. 骑士在棋盘上的概率](https://leetcode.cn/problems/knight-probability-in-chessboard/)  
[837. 新 21 点](https://leetcode.cn/problems/new-21-game/) 2350  
[1467. 两个盒子中球的颜色数相同的概率](https://leetcode.cn/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/) 2357  
[808. 分汤](https://leetcode.cn/problems/soup-servings/) 2397  
[LCR 185. 统计结果概率](https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/)  
[九坤-04. 筹码游戏](https://leetcode.cn/contest/ubiquant2022/problems/I3Gm2h/)  
[1230. 抛掷硬币](https://leetcode.cn/problems/toss-strange-coins/)（会员题）  

专题：输出具体方案（打印方案）
注意这些题目和回溯的区别，某些回溯题目要求输出所有方案，这里只要求输出一个。

[368. 最大整除子集](https://leetcode.cn/problems/largest-divisible-subset/)  
[1363. 形成三的最大倍数](https://leetcode.cn/problems/largest-multiple-of-three/) 1823 字典序  
[1449. 数位成本和为目标值的最大数字](https://leetcode.cn/problems/form-largest-integer-with-digits-that-add-up-to-target/) 1927 字典序  
[1092. 最短公共超序列](https://leetcode.cn/problems/shortest-common-supersequence/) 1977  
[943. 最短超级串](https://leetcode.cn/problems/find-the-shortest-superstring/) 2186  
[1125. 最小的必要团队](https://leetcode.cn/problems/smallest-sufficient-team/) 2251  
[3533. 判断连接可整除性](https://leetcode.cn/problems/concatenated-divisibility/) 2257 字典序 暴搜  
[3260. 找出最大的 N 位 K 回文数](https://leetcode.cn/problems/find-the-largest-palindrome-divisible-by-k/) 2370 字典序 暴搜  
[3149. 找出分数最低的排列](https://leetcode.cn/problems/find-the-minimum-cost-array-permutation/) 2642 字典序  
[3441. 变成好标题的最少代价](https://leetcode.cn/problems/minimum-cost-good-caption/) 2765 字典序  
[3348. 最小可整除数位乘积 II](https://leetcode.cn/problems/smallest-divisible-digit-product-ii/) 3101 字典序 暴搜  
[656. 金币路径](https://leetcode.cn/problems/coin-path/)（会员题）字典序  
[471. 编码最短长度的字符串](https://leetcode.cn/problems/encode-string-with-shortest-length/)（会员题）  

专题：前后缀分解  
部分题目也可以用状态机 DP 解决。
如果涉及到的只是若干元素，而不是前缀/后缀这样的一段元素，也可以用「枚举右，维护左」思考，详见数据结构题单。  

[724. 寻找数组的中心下标](https://leetcode.cn/problems/find-pivot-index/) 做到   
[1991. 找到数组的中间位置](https://leetcode.cn/problems/find-the-middle-index-in-array/) 同 724 题  
[3707. 相等子字符串分数](https://leetcode.cn/problems/equal-score-substrings/) 1262  
[2270. 分割数组的方案数](https://leetcode.cn/problems/number-of-ways-to-split-array/) 1334  
[2256. 最小平均差](https://leetcode.cn/problems/minimum-average-difference/) 1395  
[1422. 分割字符串的最大得分](https://leetcode.cn/problems/maximum-score-after-splitting-a-string/) 做到   
[238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)  
[1493. 删掉一个元素以后全为 1 的最长子数组](https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/) 1423  
[845. 数组中的最长山脉](https://leetcode.cn/problems/longest-mountain-in-array/) 1437 也可以分组循环  
[2012. 数组美丽值求和](https://leetcode.cn/problems/sum-of-beauty-in-the-array/) 1468  
[2909. 元素和最小的山形三元组 II](https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/) 1479  
[2483. 商店的最少代价](https://leetcode.cn/problems/minimum-penalty-for-a-shop/) 1495  
[1525. 字符串的好分割数目](https://leetcode.cn/problems/number-of-good-ways-to-split-a-string/) 1500  
[3583. 统计特殊三元组](https://leetcode.cn/problems/count-special-triplets/) 1510  
[3354. 使数组元素等于零](https://leetcode.cn/problems/make-array-elements-equal-to-zero/) 做到   
[2874. 有序三元组中的最大值 II](https://leetcode.cn/problems/maximum-value-of-an-ordered-triplet-ii/) 1583  
[123. 买卖股票的最佳时机 III](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/) 拆分成两个 121 题  
[3698. 分割数组得到最小绝对差](https://leetcode.cn/problems/split-array-with-minimum-difference/) 1649  
[3598. 相邻字符串之间的最长公共前缀](https://leetcode.cn/problems/longest-common-prefix-between-adjacent-strings-after-removals/) 1655  
[2222. 选择建筑的方案数](https://leetcode.cn/problems/number-of-ways-to-select-buildings/) 1657  
[1031. 两个无重叠子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-two-non-overlapping-subarrays/) 1680  
[689. 三个无重叠子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays/)  
[2420. 找到所有好下标](https://leetcode.cn/problems/find-all-good-indices/) 1695  
[2100. 适合野炊的日子](https://leetcode.cn/problems/find-good-days-to-rob-the-bank/) 1702  
[926. 将字符串翻转到单调递增](https://leetcode.cn/problems/flip-string-to-monotone-increasing/)  
[334. 递增的三元子序列](https://leetcode.cn/problems/increasing-triplet-subsequence/)  
[1653. 使字符串平衡的最少删除次数](https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced/) 1794  
[1186. 删除一次得到子数组最大和](https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/) 1799  
[42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)
[2711. 对角线上不同值的数量差](https://leetcode.cn/problems/difference-of-number-of-distinct-values-on-diagonals/) 做到   
[1477. 找两个和为目标值且不重叠的子数组](https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/) 1851  
[2680. 最大或值](https://leetcode.cn/problems/maximum-or/) 1912  
[1671. 得到山形数组的最少删除次数](https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/) 1913  
[2906. 构造乘积矩阵](https://leetcode.cn/problems/construct-product-matrix/) 2075  
[3334. 数组的最大因子得分](https://leetcode.cn/problems/find-the-maximum-factor-score-of-array/) 非暴力做法  
[2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/) 2219  
[2484. 统计回文子序列数目](https://leetcode.cn/problems/count-palindromic-subsequences/) 2223  
[2163. 删除元素后和的最小差值](https://leetcode.cn/problems/minimum-difference-in-sums-after-removal-of-elements/) 2225  
[2565. 最少得分子序列](https://leetcode.cn/problems/subsequence-with-the-minimum-score/) 2432  
[1995. 统计特殊四元组](https://leetcode.cn/problems/count-special-quadruplets/) 四个数  
[2552. 统计上升四元组](https://leetcode.cn/problems/count-increasing-quadruplets/) 2433 四个数  
[3302. 字典序最小的合法序列](https://leetcode.cn/problems/find-the-lexicographically-smallest-valid-sequence/) 2474 第 2565 题的进阶版本  
[3404. 统计特殊子序列的数目](https://leetcode.cn/problems/count-special-subsequences/) 2445 四个数  
[3303. 第一个几乎相等子字符串的下标](https://leetcode.cn/problems/find-the-occurrence-of-first-almost-equal-substring/) 2509  
[3287. 求出数组中最大序列值](https://leetcode.cn/problems/find-the-maximum-sequence-value-of-array/) 2545  
[3257. 放三个车的价值之和最大 II](https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-ii/) 2553  
[3410. 删除所有值为某个元素后的最大子数组和](https://leetcode.cn/problems/maximize-subarray-sum-after-removing-all-occurrences-of-one-element/) 2844  
[3003. 执行操作后的最大分割数量](https://leetcode.cn/problems/maximize-the-number-of-partitions-after-operations/) 3039  
[487. 最大连续 1 的个数 II](https://leetcode.cn/problems/max-consecutive-ones-ii/)（会员题）  
[1746. 经过一次操作后的最大子数组和](https://leetcode.cn/problems/maximum-subarray-sum-after-one-operation/)（会员题）
补充题目：

输入一个长为 n 的 prices 数组，你需要返回一个长为 n 的 answer 数组，其中 answer[i] 表示删除 prices[i]，也就是禁止在第 i 天买卖股票，在此约束下 121. 买卖股票的最佳时机 的答案。
专题：把 X 变成 Y

部分题目也可以用 BFS 解决。

[397. 整数替换](https://leetcode.cn/problems/integer-replacement/)  
[2998. 使 X 和 Y 相等的最少操作次数](https://leetcode.cn/problems/minimum-number-of-operations-to-make-x-and-y-equal/) 1795  
[2059. 转化数字的最小运算数](https://leetcode.cn/problems/minimum-operations-to-convert-number/) 1850  
[991. 坏了的计算器](https://leetcode.cn/problems/broken-calculator/) 1909  
[1553. 吃掉 N 个橘子的最少天数](https://leetcode.cn/problems/minimum-number-of-days-to-eat-n-oranges/) 2048  

另见 图论题单 中的 Dijkstra 算法，例如：

[3377. 使两个整数相等的数位操作](https://leetcode.cn/problems/digit-operations-to-make-two-integers-equal/) 2186

专题：跳跃游戏

[1306. 跳跃游戏 III](https://leetcode.cn/problems/jump-game-iii/) 1397  
[2770. 达到末尾下标所需的最大跳跃次数](https://leetcode.cn/problems/maximum-number-of-jumps-to-reach-the-last-index/) 1533  
[403. 青蛙过河](https://leetcode.cn/problems/frog-jump/)  
[1340. 跳跃游戏 V](https://leetcode.cn/problems/jump-game-v/) 1866  
[1871. 跳跃游戏 VII](https://leetcode.cn/problems/jump-game-vii/) 1896  
[1696. 跳跃游戏 VI](https://leetcode.cn/problems/jump-game-vi/) 1954  
[975. 奇偶跳](https://leetcode.cn/problems/odd-even-jump/) 2079  
[1654. 到家的最少跳跃次数](https://leetcode.cn/problems/minimum-jumps-to-reach-home/) 2124  
[LCP 09. 最小跳跃次数](https://leetcode.cn/problems/zui-xiao-tiao-yue-ci-shu/)  
[LCP 20. 快速公交](https://leetcode.cn/problems/meChtZ/)  
[656. 金币路径](https://leetcode.cn/problems/coin-path/)（会员题）  
[2297. 跳跃游戏 VIII](https://leetcode.cn/problems/jump-game-viii/)（会员题）  

其他

[1387. 将整数按权重排序](https://leetcode.cn/problems/sort-integers-by-the-power-value/) 1507  
[823. 带因子的二叉树](https://leetcode.cn/problems/binary-trees-with-factors/) 1900  
[756. 金字塔转换矩阵](https://leetcode.cn/problems/pyramid-transition-matrix/) 1990  
[2930. 重新排列后包含指定子字符串的字符串数目](https://leetcode.cn/problems/number-of-strings-which-can-be-rearranged-to-contain-substring/) 2227  
[1896. 反转表达式值的最少操作次数](https://leetcode.cn/problems/minimum-cost-to-change-the-final-value-of-expression/) 2532

