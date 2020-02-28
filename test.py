def printUnsorted(arr, n): 
    e = n-1
    # step 1(a) of above algo 
    for s in range(0,n-1): 
        if arr[s] > arr[s+1]: 
            break
          
    if s == n-1: 
        print ("The complete array is sorted") 
        exit() 
  
    # step 1(b) of above algo 
    e= n-1
    while e > 0: 
        if arr[e] < arr[e-1]: 
            break
        e -= 1
  
    # step 2(a) of above algo 
    max = arr[s] 
    min = arr[s] 
    for i in range(s+1,e+1): 
        if arr[i] > max: 
            max = arr[i] 
        if arr[i] < min: 
            min = arr[i] 
              
    # step 2(b) of above algo 
    for i in range(s): 
        if arr[i] > min: 
            s = i 
            break
  
    # step 2(c) of above algo 
    i = n-1
    while i >= e+1: 
        if arr[i] < max: 
            e = i 
            break
        i -= 1
      
    # step 3 of above algo
    print e, s
    print ("The unsorted subarray which makes the given array") 
    print ("sorted lies between the indexes %d and %d"%( s, e))

# printUnsorted([1, 2, 5, 4, 3], 5)


import math

def printSubsequences(arr, n) : 
  
    # Number of subsequences is (2**n -1) 
    opsize = math.pow(2, n) 
  
    # Run from counter 000..1 to 111..1 
    for counter in range( 1, (int)(opsize)) : 
        for j in range(0, n) : 
              
            # Check if jth bit in the counter 
            # is set If set then print jth  
            # element from arr[]  
            if (counter & (1<<j)) : 
                print arr[j],
          
        print

printSubsequences([1, 2, 3], 3)