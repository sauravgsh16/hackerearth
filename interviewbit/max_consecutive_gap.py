def maxSortedAdjacentDiff(arr):
    INT_MIN, INT_MAX = float('-inf'), float('inf')
    n = len(arr)  
    maxVal, minVal = arr[0], arr[0]  
    for i in range(1, n):  
        maxVal = max(maxVal, arr[i])  
        minVal = min(minVal, arr[i])  
  
 
    maxBucket = [INT_MIN] * (n)  
    minBucket = [INT_MAX] * (n)  
      

    delta = (maxVal - minVal) // (n - 1)  
  
    for i in range(0, n): 
        if arr[i] == maxVal or arr[i] == minVal:  
            continue
  

        index = (arr[i] - minVal) // delta  

        if maxBucket[index] == INT_MIN:  
            maxBucket[index] = arr[i]  
        else: 
            maxBucket[index] = max(maxBucket[index], 
                                            arr[i])  

        if minBucket[index] == INT_MAX:  
            minBucket[index] = arr[i]  
        else: 
            minBucket[index] = min(minBucket[index], 
                                            arr[i])

    prev_val, max_gap = minVal, 0
      
    for i in range(0, n - 1):  
        if minBucket[i] == INT_MAX:  
            continue
              
        max_gap = max(max_gap,  
                      minBucket[i] - prev_val)  
        prev_val = maxBucket[i]  
      
    max_gap = max(max_gap, maxVal - prev_val)  
  
    return max_gap

import math

def maximumGap(num):
    if len(num) < 2 or min(num) == max(num):
        return 0
    a, b = min(num), max(num)
    size = math.ceil((b-a)/(len(num)-1))
    bucket = [[None, None] for _ in range((b-a)//size+1)]
    for n in num:
        b = bucket[(n-a)//size]
        b[0] = n if b[0] is None else min(b[0], n)
        b[1] = n if b[1] is None else max(b[1], n)
    bucket = [b for b in bucket if b[0] is not None]
    return max(bucket[i][0]-bucket[i-1][1] for i in range(1, len(bucket)))

print maximumGap([1, 10, 21, 33, 45])
