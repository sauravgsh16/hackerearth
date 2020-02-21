# function find Next greater element 
def nextGreater(arr, n, order): 
    arr1 = [None]*n
    S = []
    for i in range(n-1,-1,-1): 
  
        # Keep removing top element from S while the top 
        # element is smaller then or equal to arr[i] (if Key is G) 
        # element is greater then or equal to arr[i] (if order is S)
        if order == "G":
            while (len(S)>0 and (arr[S[-1]] <= arr[i])):
                S.pop()
        else:
            while (S!=[] and (arr[S[len(S)-1]] >= arr[i])):
                S.pop()

   
        # If all elements in S were smaller than arr[i]
        if (len(S) == 0): 
            arr1[i] = -1
         # store the next greater element of current element
        else:
            arr1[i] = S[-1] 
   
        # Push this element 
        S.append(i)
    print arr1
    return arr1
   
# Function to find Right smaller element of next greater 
# element 
def nextSmallerOfNextGreater(arr, n):
    NG = nextGreater(arr, n, 'G') 
   
    # Find right smaller element 
    # using same function nextGreater() 
    # Here S indicate right smaller elements 
    RS = nextGreater(arr, n, 'S') 
   
    # If NG[i] == -1 then there is no smaller element 
    # on right side. We can find Right smaller of next 
    # greater by arr[RS[NG[i]]]


    for i in range(n): 
        if (NG[i] != -1 and RS[NG[i]] != -1): 
            print arr[RS[NG[i]]],
        else: 
            print "-1",
   
# Driver program 
if __name__=="__main__": 
    arr = [3, 7, 1, 7, 8, 4, 5, 2] 
    n = len(arr) 
    nextSmallerOfNextGreater(arr, n) 