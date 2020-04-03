/** 
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
    var ans [][]int
    x, y :=  1, 1000
    for x <= z && y >= 1 {
        res := customFunction(x, y)
        if res < z {
            x += 1
        } else if res > z {
            y -= 1
        }
        if res == z {
            temp := []int{x, y}
            ans = append(ans, temp)
            x += 1
            y -= 1
        }
    }
    return ans
}
