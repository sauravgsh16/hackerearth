class Solution:
    def possible(self, arrive, depart, k):
        flat = []
        n = len(arrive)

        arrive.sort()
        depart.sort()

        i, j = 0, 0
        while i < n and j < n:
            if arrive[i] < depart[j]:
                flat.append((arrive[i], "a"))
                i += 1
            else:
                flat.append((depart[j], "d"))
                j += 1
        
        while i < n:
            flat.append((arrive[i], "a"))
            i += 1
        
        while j < n:
            flat.append((depart[j], "d"))
            j += 1

        for var in flat:
            if var[1] == "a":
                k -= 1
            else:
                k += 1
            
            if k < 0:
                return False
        
        return True

s = Solution()
print s.possible([13, 14, 36, 19, 44, 1, 45, 4, 48, 23, 32, 16, 37, 44, 47, 28, 8, 47, 4, 31, 25, 48, 49, 12, 7, 8],
                 [28, 27, 61, 34, 73, 18, 50, 5, 86, 28, 34, 32, 75, 45, 68, 65, 35, 91, 13, 76, 60, 90, 67, 22, 51, 53],
                 2)
