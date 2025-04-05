/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start

/*func canFinish(numCourses int, prerequisites [][]int) bool {
	hash := make(map[int](map[int]bool))
	inveHash := make(map[int](map[int]bool))

	var addReq func(int, int)
	var transReq func(int, int)

	addReq = func(key int, val int) {
		subHash, exist := hash[key]
		if exist {
			subHash[val] = true
		} else {
			tmp := make(map[int]bool)
			tmp[val] = true
			hash[key] = tmp
		}
		// add to inve
		subHash, exist = inveHash[val]
		if exist {
			subHash[key] = true
		} else {
			tmp := make(map[int]bool)
			tmp[key] = true
			inveHash[val] = tmp
		}
		transReq(key, val)
	}

	transReq = func(key int, val int) {
		subHash, exist := inveHash[key]
		if exist {
			for k, _ := range subHash {
				addReq(k, val)
			}
		}
	}

	for _, prerequise := range prerequisites {
		if prerequise[0] == prerequise[1] {
			return false
		}
		subHash, exist := hash[prerequise[1]]
		if exist {
			_, ring := subHash[prerequise[0]]
			if ring {
				return false
			}
		}
		addReq(prerequise[0], prerequise[1])
	}

	return true
}*/

func canFinish(numCourses int, prerequisites [][]int) bool {
	edge := make([](map[int]bool), numCourses)
	degrees := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		edge[i] = make(map[int]bool)
	}

	for _, prerequisity := range prerequisites {
		edge[prerequisity[0]][prerequisity[1]] = true
		degrees[prerequisity[1]]++
	}

	var topsort func() bool
	topsort = func() bool {
		queue := make([]int, 0, numCourses)
		for point, degree := range degrees {
			if degree == 0 {
				queue = append(queue, point)
			}
		}

		for len(queue) != 0 {
			start := queue[0]
			queue = queue[1:]
			points := edge[start]
			for point, _ := range points {
				degrees[point]--
				if degrees[point] == 0 {
					queue = append(queue, point)
				}
			}
		}

		for _, degree := range degrees {
			if degree != 0 {
				return false
			}
		}
		return true
	}

	return topsort()
}

// @lc code=end

