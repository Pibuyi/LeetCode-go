var (
    res int
    M int
    N int
    K int
    dx = []int{1,-1,0,0}
    dy = []int{0,0,-1,1}
)

func movingCount(m int, n int, k int) int {
    res, M, N, K = 0, m, n, k
    g := make([][]int, m)
    for i := 0; i < m; i++ {
        g[i] = make([]int, n)
    }
    helper(0,0,g)
    return res
}

func helper(x, y int, g [][]int) {
    if x < 0 || y < 0 || x >= M || y >= N || g[x][y] == 1 || (toD(x) + toD(y) > K){
        return
    }
    g[x][y] = 1
    res++
    for i := 0; i < 4; i++ {
        nx, ny := x + dx[i], y + dy[i] 
        helper(nx, ny, g)
    }
}

func toD(n int) int {
    t  := 0
    for n != 0 {
        t += n % 10
        n /= 10
    }
    return t
}
