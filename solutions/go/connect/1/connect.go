package connect

import (
	"fmt"
)

type Pair struct {
	m, n int
}

func ResultOf(lines []string) (winner string, err error) {
	var board [][]int
	board, err = getBoard(lines)
	if err != nil {
		return
	}
	M := len(board)
	var N int
	if len(board) == 0 {
		N = 0
	} else {
		N = len(board[0])
	}

	valid := func(p Pair) bool {
		return 0 <= p.m && p.m < M && 0 <= p.n && p.n < N
	}

	next := func(p Pair) (r []Pair) {
		return []Pair{
			{m: p.m, n: p.n - 1},
			{m: p.m, n: p.n + 1},
			{m: p.m - 1, n: p.n},
			{m: p.m + 1, n: p.n},
			{m: p.m - 1, n: p.n + 1},
			{m: p.m + 1, n: p.n - 1},
		}
	}

	bfs := func(queue []Pair, flag int, reachable, visited [][]bool) {
		var cur Pair
		var ne Pair
		for len(queue) > 0 {
			cur, queue = queue[0], queue[1:]
			for _, ne = range next(cur) {
				if valid(ne) && !visited[ne.m][ne.n] {
					if flag == board[ne.m][ne.n] {
						reachable[ne.m][ne.n] = true
						queue = append(queue, ne)
					}
					visited[ne.m][ne.n] = true
				}
			}
		}
	}

	connectO := func(owin chan<- bool) {
		visited := initVisited(M, N)
		reachable := initVisited(M, N)
		for cn := 0; cn < N; cn++ {
			if board[0][cn] == 1 {
				visited[0][cn], reachable[0][cn] = true, true
				bfs([]Pair{{m: 0, n: cn}}, 1, reachable, visited)
			}
		}
		for cn := 0; cn < N; cn++ {
			if reachable[M-1][cn] {
				owin <- true
				return
			}
		}
		close(owin)
		return
	}

	connectX := func(xwin chan<- bool) {
		visited := initVisited(M, N)
		reachable := initVisited(M, N)
		for rn := 0; rn < M; rn++ {
			if board[rn][0] == 2 {
				visited[rn][0], reachable[rn][0] = true, true
				bfs([]Pair{{m: rn, n: 0}}, 2, reachable, visited)
			}
		}
		for rn := 0; rn < M; rn++ {
			if reachable[rn][N-1] {
				xwin <- true
				return
			}
		}
		close(xwin)
		return
	}

	owin := make(chan bool)
	xwin := make(chan bool)
	go connectO(owin)
	go connectX(xwin)
	var won, open bool
loop:
	for done := 0; done < 2; {
		select {
		case won, open = <-owin:
			if !open {
				done++
				owin = nil
			} else if won {
				winner = "O"
				break loop
			}
		case won, open = <-xwin:
			if !open {
				done++
				xwin = nil
			} else if won {
				winner = "X"
				break loop
			}
		}
	}
	return
}

func getBoard(lines []string) (board [][]int, err error) {
	rows := len(lines)
	var cols int
	board = make([][]int, rows)
	if rows == 0 {
		return
	}
	fields := []rune(lines[0])
	cols = len(fields)
	matrix := make([]int, rows*cols)
	for i := 0; i < rows; i++ {
		if i != 0 {
			fields = []rune(lines[i])
		}
		if len(fields) != cols {
			err = fmt.Errorf("uneven row length %v", lines)
			return
		}
		for j := 0; j < cols; j++ {
			switch fields[j] {
			case '.':
				matrix[i*cols+j] = 0
			case 'O':
				matrix[i*cols+j] = 1
			case 'X':
				matrix[i*cols+j] = 2
			default:
				err = fmt.Errorf("invalid character in board '%c' %v %d", fields[j], fields, len(fields))
				return
			}
		}
		board[i] = matrix[i*cols : (i+1)*cols]
	}
	return
}

func initVisited(rows, cols int) [][]bool {
	v := make([][]bool, rows)
	m := make([]bool, rows*cols)
	for i := 0; i < rows; i++ {
		v[i] = m[i*cols : (i+1)*cols]
	}
	return v
}
