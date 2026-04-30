package graph

//207. 课程表
//你这个学期必须选修 numCourses 门课程，记为 0 到 numCourses - 1 。
//在选修某些课程之前需要一些先修课程。 先修课程按数组 prerequisites 给出，其中 prerequisites[i] = [ai, bi] ，表示如果要学习课程 ai 则 必须 先学习课程  bi 。
//例如，先修课程对 [0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
//请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false
//链接：https://leetcode.cn/problems/course-schedule/description/?envType=study-plan-v2&envId=top-100-liked
//题解：https://leetcode.cn/problems/course-schedule/solutions/1181465/che-di-gao-dong-tuo-bu-pai-xu-zui-qiang-vmec6/?envType=study-plan-v2&envId=top-100-liked
//kahn 算法：拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	//统计节点的指向以及每个节点的入度
	adj := make(map[int][]int)
	indegree := make(map[int]int)
	for _, row := range prerequisites {
		adj[row[1]] = append(adj[row[1]], row[0])
		indegree[row[0]]++
	}

	//找出入度为0的节点
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topo := make([]int, 0)
	for len(queue) > 0 {
		node := queue[0]
		topo = append(topo, node) //topo中记录入度为0的点
		queue = queue[1:]
		//更新当前节点指向的节点入度
		for _, neighbor := range adj[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return len(topo) == numCourses
}
