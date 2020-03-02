def combinationSum(arr, target):
    result = []
    temp = []

    backtrack(result, temp, arr, target, 0)
    return result


def backtrack(result, temp, arr, remain, start):
    if remain < 0:
        return
    
    if remain == 0:
        result.append(list(temp))
        return
    
    for i in range(start, len(arr)):
        temp.append(arr[i])
        backtrack(result, temp, arr, remain-arr[i], i) # not i + 1, since we can reuse same element
        temp.pop()


print combinationSum([2, 3, 6, 7], 7)
