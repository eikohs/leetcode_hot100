/*
 * @lc app=leetcode id=399 lang=golang
 *
 * [399] Evaluate Division
 */

// @lc code=start
/*type multi struct {
	symbol string
	val    float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	hash := make(map[string]multi)
	mark := make(map[string]bool)
	for key, equation := range equations {
		hash[equation[0]] = multi{equation[1], values[key]}
		mark[equation[0]] = true
		mark[equation[1]] = true
	}

	var convert func(string) (multi, bool)
	convert = func(src string) (multi, bool) {
		_, exists := mark[src]
		if !exists {
			return multi{}, false
		}
		rst, exists := hash[src]
		if !exists {
			return multi{src, 1.0}, true
		}
		final, _ := convert(rst.symbol)
		rst.symbol = final.symbol
		rst.val = rst.val * final.val

		return rst, true
	}

	result := []float64{}
	for _, query := range queries {
		up, stat1 := convert(query[0])
		down, stat2 := convert(query[1])
		if !stat1 || !stat2 || up.symbol != down.symbol {
			result = append(result, -1)
		} else {
			result = append(result, up.val/down.val)
		}
	}

	return result
}*/

// 实现一个并查集
type unionFind struct {
	parent []int
	weight []float64
}

// 初始化并查集
func newUnionFind(size int) *unionFind {
	uf := &unionFind{
		parent: make([]int, size),
		weight: make([]float64, size),
	}

	// 将每个节点初始化
	for i := 0; i < size; i++ {
		uf.parent[i] = i
		uf.weight[i] = 1.0
	}

	return uf
}

// 实现并查集的查找 同时进行路径压缩
func (uf *unionFind) find(x int) int {
	if x != uf.parent[x] {
		origin := uf.parent[x]
		// 路径压缩
		uf.parent[x] = uf.find(origin)
		// 权值需要同步进行修改
		uf.weight[x] = uf.weight[x] * uf.weight[origin]
	}
	return uf.parent[x]
}

// 在并查集中添加一个关系
func (uf *unionFind) union(x int, y int, val float64) {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootX == rootY {
		return
	}

	// 添加一个关系，但不进行路径压缩（避免特殊处理 x == rootX 的情况）
	// x -> rootX | x -> y -> rootY => y -> rootY | x -> rootX -> rootY
	uf.parent[rootX] = rootY
	// 先压缩 rootX，避免值被覆盖
	uf.weight[rootX] = val * uf.weight[y] / uf.weight[x]
}

// 在并查集中查找两个节点的关系
func (uf *unionFind) caculate(x int, y int) float64 {
	rootX := uf.find(x)
	rootY := uf.find(y)

	if rootX != rootY {
		return -1.0
	}

	return uf.weight[x] / uf.weight[y]
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// 1. 初始化并查集，将 string 映射到 int 后在并查集中添加关系
	uf := newUnionFind(2 * len(equations))
	hash := make(map[string]int)
	flag := 0
	for key, equation := range equations {
		x, exists := hash[equation[0]]
		if !exists {
			x = flag
			hash[equation[0]] = flag
			flag++
		}
		y, exists := hash[equation[1]]
		if !exists {
			y = flag
			hash[equation[1]] = flag
			flag++
		}
		uf.union(x, y, values[key])
	}

	// 2. 查找并查集并计算结果
	result := make([]float64, len(queries))
	for key, query := range queries {
		x, stat1 := hash[query[0]]
		y, stat2 := hash[query[1]]
		if !stat1 || !stat2 {
			result[key] = -1.0
		} else {
			result[key] = uf.caculate(x, y)
		}
	}

	return result
}

// @lc code=end

