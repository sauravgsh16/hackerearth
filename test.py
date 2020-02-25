def printPairs(arr, arr_size, sum): 
      
    # Create an empty hash set 
    s = set() 
      
    for i in range(0, arr_size): 
        temp = sum-arr[i] 
        if (temp in s): 
            print "Pair with given sum "+ str(sum) + " is (" + str(arr[i]) + ", " + str(temp) + ")"
        s.add(arr[i])
        print s        

l = [3, 1, 4, 1, 5]
printPairs(l, len(l), 2)