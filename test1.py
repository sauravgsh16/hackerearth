def findBestValue(A, target):
    A.sort(reverse=1)
    maxA = A[0]
    print A[-1] * len(A)
    while A and target >= A[-1] * len(A):
        ele = A.pop()
        print 'Ele', ele
        target -= ele
        print 'target', target
    return int(round((target - 0.0001) / len(A))) if A else maxA


def combinationSum3(K, N):
    def dfs(min_num, path):
        if len(path)==K and sum(path)==N:
            opt.append(path)
        for num in xrange(min_num, 10):
            dfs(num+1, path+[num])
    opt = []
    dfs(1, [])
    return opt


from itertools import product

def maxSideLength(mat, threshold):
    """
    prefix[i][j] = prefix[i-1][j] + prefix[i][j-1] - prefix[i-1][j-1] + mat[i][j]
    """
    m, n = len(mat), len(mat[0])
    #build prefix sum 
    prefix = [[0]*(n+1) for _ in range(m+1)]
    
    for i in range(m):
        for j in range(n):
            prefix[i+1][j+1] = prefix[i+1][j] + prefix[i][j+1] - prefix[i][j] + mat[i][j]
    
    print prefix
    
    def below(k): 
        """reture true if there is such a sub-matrix of length k"""
        for i in range(m+1-k):
            for j in range(n+1-k):
                print i, j
                print prefix[i+k][j+k], prefix[i][j+k], prefix[i+k][j],  prefix[i][j]
                if prefix[i+k][j+k] - prefix[i][j+k] - prefix[i+k][j] + prefix[i][j] <= threshold:
                    return True
        return False 
            
    #binary search
    lo, hi = 1, min(m, n)
    while lo <= hi: 
        mid = (lo + hi)//2

        print "Calling below for mid: ", mid
        if below(mid):
            lo = mid + 1
        else:
            print "Here"
            hi = mid - 1
        
    return hi


print maxSideLength([[1,1,3],[1,1,3],[1,1,3]], 4)