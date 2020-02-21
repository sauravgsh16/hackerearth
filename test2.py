def printNGE(arr, n): 
  
    s = list() 
  
    arr1 = [0 for i in range(n)] 
  
    # iterating from n-1 to 0 
    for i in range(n - 1, -1, -1):  
      
        # We will pop till we get the greater   
        # element on top or stack gets empty 
        while (len(s) > 0 and s[-1] <= arr[i]): 
            s.pop() 
  
        # if stack gots empty means there  
        # is no element on right which is  
        # greater than the current element. 
        # if not empty then the next greater  
        # element is on top of stack 
        if (len(s) == 0): 
            arr1[i] = -1        
        else: 
            arr1[i] = s[-1]      
  
        s.append(arr[i]) 
      
    print arr1

printNGE([3, 7, 1, 7, 8, 4, 5, 2], 8)