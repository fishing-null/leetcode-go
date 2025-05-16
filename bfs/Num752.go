package bfs

func OpenLock(deadends []string, target string) int {
	const start = "0000"
	var step int
	if target == start {
		return step
	}
	deadMap := map[string]bool{}
	for _, dead := range deadends {
		deadMap[dead] = true
	}
	if deadMap[start] { // 检查初始节点是否为死锁
		return -1
	}
	//存储出现过的密码
	saveMap := map[string]bool{}
	//用于bfs
	q := []string{start}
	saveMap[start] = true
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			//出队列
			q = q[1:]
			if deadMap[cur] {
				continue
			}
			if cur == target {
				return step
			}
			for i := 0; i < 4; i++ {
				up := pollUp(cur, i)
				if !saveMap[up] {
					q = append(q, up)
					saveMap[up] = true
				}
				down := pollDOwn(cur, i)
				if !saveMap[down] {
					q = append(q, down)
					saveMap[down] = true
				}
			}
		}
		step++
	}
	return -1
}

func pollUp(s string, i int) string {
	ch := []byte(s)
	if ch[i] == '9' {
		ch[i] = '0'
	} else {
		ch[i]++
	}
	return string(ch)
}

func pollDOwn(s string, i int) string {
	ch := []byte(s)
	if ch[i] == '0' {
		ch[i] = '9'
	} else {
		ch[i]--
	}
	return string(ch)
}
