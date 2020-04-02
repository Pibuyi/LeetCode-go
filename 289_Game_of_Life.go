var (
    row int
    col int
    arr [][]int
)
/*
    共有四种状态：
        0： 死的死
        1： 活得死 
        2： 死得活
        3： 活得活
        右移动更新地图
*/
func countAlive(x, y int) int {
    cnt := 0
    for i := -1;  i <= 1; i++ {
        for j := -1; j <= 1; j++ {
            nx := x + i
            ny := y + j
            if nx < 0 || ny < 0 || nx >= row || ny >= col {
                continue
            }
            cnt += (arr[nx][ny] & 1)
        }
    }
    return cnt
}

func gameOfLife(board [][]int)  {
    arr = board
    row = len(board)
    if row == 0 {
        return
    }
    col = len(board[0])
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            cnt := countAlive(i, j)
            //fmt.Println(cnt)
            // 周围有两个或三个加本身一个 下一轮继续活
            if board[i][j] == 1 && (cnt == 3 || cnt == 4) {
                board[i][j] = 3     
            }
            //死的复活
            if board[i][j] == 0 && cnt == 3 {
                board[i][j] = 2
            }
        }
    }

    //更新地图
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            board[i][j] >>= 1
        }
    }
}
