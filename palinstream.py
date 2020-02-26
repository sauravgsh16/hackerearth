def checkPalindrome(s):
    n = len(s)

    for i in range(1, n+1):

        if i == 1:
            print s[0], " --> YES"
            continue
        
        mid = i / 2
        stack = []
        if i % 2 == 0:
            while mid <= i-1:
                stack.append(s[mid])
                mid += 1
        else:
            while mid + 1 <= i-1:
                stack.append(s[mid+1])
                mid += 1

        j = 0
        while len(stack) > 0 and j <= mid and stack[-1] == s[j]:
            stack.pop() != s[j]
            j += 1
        if len(stack) > 0:
            print s[0:i], " --> NO"
        else:  
            print s[0:i], " --> YES"

checkPalindrome("aabaacaabaa")