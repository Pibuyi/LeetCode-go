func superEggDrop(K int, N int) int {
    times := 1
    for {
        if getTimes(times, K) >= N {
            break
        }
        times++
    }
    return times
}

// 如果鸡蛋没有碎，那么对应的是 f(T - 1, K)f(T−1,K)，也就是说在这一层的上方可以有 f(T - 1, K)f(T−1,K) 层；
// 如果鸡蛋碎了，那么对应的是 f(T - 1, K - 1)f(T−1,K−1)，也就是说在这一层的下方可以有 f(T - 1， K - 1)f(T−1，K−1) 层。
// f(T, K) = 1 + f(T-1, K-1) + f(T-1, K)
func getTimes(times, eggCount int) int {
    if times == 1 || eggCount == 1 {
        return times
    }
    return getTimes(times - 1, eggCount - 1) + 1 + getTimes(times - 1, eggCount)
}
