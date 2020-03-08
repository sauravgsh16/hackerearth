def max_distance(arr):
    n = len(arr)

    if n <= 1:
        return 0
    
    pair = []
    for i in range(n):
        pair.append([arr[i], i])
    
    pair.sort(key= lambda x: x[0])

    ans = 0
    rmax = pair[-1][1]
    for i in range(n-2, -1, -1):
        if rmax-pair[i][1] > ans:
            ans = rmax-pair[i][1]
        
        if pair[i][1] > rmax:
            rmax = pair[i][1]
    
    return ans


print max_distance([3, 5, 4, 2])