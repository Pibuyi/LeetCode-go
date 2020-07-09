func numSubmat(mat [][]int) int {
    ans := 0
    row, col := len(mat), len(mat[0])
    for i := 0; i < row; i++ {
        // 统计一行中矩形个数
        for j := i; j < row; j++ {
            c := 0
            for k := 0; k < col; k++ {
                if mat[j][k] == 1 {
                    c++
                }else {
                    c = 0
                }
                ans += c
            }
        }

        //进行压缩
        for j := row - 1; j > i; j-- {
            for k := 0; k < col; k++ {
                mat[j][k] &= mat[j - 1][k]
            }
        }
    }
    return ans
}
