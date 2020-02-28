def arrayNesting(arr):
    maxlength = 0
    visited = [False] * len(arr)

    for i in range(len(visited)):
        if not visited[i]:
            maxlength = max(maxlength, dfs(arr, i, visited))
    return maxlength


def dfs(arr, start, visited):
    count = 0
    i = start

    while count == 0 or i != start:
        visited[i] = True
        i = arr[i]
        count += 1
    return count


print arrayNesting([5,4,0,3,1,6,2])