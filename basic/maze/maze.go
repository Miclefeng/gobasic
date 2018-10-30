package main

import (
	"fmt"
	"os"
)

// 读取地图到二维数组
func readMaze(filename string) (maze [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze = make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return
}

// 定义坐标点,i 竖轴、j 横轴
type endPoint struct{
	i, j int
}

// 当前坐标点的4个相邻点位
var pointSet = [4]endPoint{
	// 坐标点的4个相邻点(也是距离当前坐标点的移动距离)，方向上左下右，逆时针
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

// 移动坐标点
func (p endPoint) move(r endPoint) endPoint {
	return endPoint{p.i + r.i, p.j + r.j}
}

// 判断坐标点是否超出边界
func (p endPoint) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	// 返回当前坐标点的值，即当前走的最短步数
	return grid[p.i][p.j], true
}

// 开始走迷宫
func walk(maze [][]int, start, end endPoint) [][]int {
	// 对迷宫地图进行copy，并初始化为0
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []endPoint{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		// 判断是否走的终点
		if cur == end {
			break
		}

		for _, p := range pointSet {
			next := cur.move(p)
			// 不能向回走
			if next == start {
				continue
			}

			// 1为墙壁，无法通行
			val, ok := next.at(maze)
			if !ok || 1 == val {
				continue
			}

			// 为0的才可以通行,不为0说明已经走过
			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			// 当前走的步数
			curStep, _ := cur.at(steps)

			// 往前走了一步
			steps[next.i][next.j] = curStep + 1
			Q = append(Q, next)
		}
	}
	return steps
}

func main()  {
	filename := "basic/maze/maze.in"
	maze := readMaze(filename)
	//for _, row := range maze {
	//	for _, col := range row {
	//		fmt.Printf("%d ", col)
	//	}
	//	fmt.Println()
	//}
	steps := walk(maze, endPoint{0, 0}, endPoint {len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, col := range row {
			fmt.Printf("%4d", col)
		}
		fmt.Println()
	}
}
