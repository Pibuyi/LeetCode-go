var dx = [4]int{-1, 1, 0, 0}
var dy = [4]int{0, 0, 1, -1}
var row, col int

func numIslands(grid [][]byte) int {
    row = len(grid)
    if row == 0{
        return 0
    }
    
    col = len(grid[0])
    count := 0
    for i := 0; i < row; i++{
        for j := 0; j < col; j++{
            if grid[i][j] == '1'{
                count++
                DFS(grid, i, j)
            }
        }
    }
    return count
}

func DFS(grid [][]byte, x int, y int) {
    grid[x][y] = '0'
    for i := 0; i < 4; i++ {
        nx := x + dx[i]
        ny := y + dy[i]
        if nx >= 0 && ny >= 0 && nx < row && ny < col && grid[nx][ny] == '1' {
            DFS(grid, nx, ny)
        }
    }
}
