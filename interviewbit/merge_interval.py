class Solution:
    def insert(self, intervals, new_interval):
        if len(intervals) == 0:
            return [new_interval]

        new = [0, 0]

        begin = []
        end = []

        for curr in intervals:
            if new_interval[0] > curr[0] and new_interval[1] < curr[1]:
                return intervals
            
            if new_interval[0] > curr[1]:
                begin.append(curr)
            
            elif new_interval[1] < curr[0]:
                end.append(curr)
            
            else :
            
                if new_interval[0] > curr[0] and new_interval[0] < curr[1]:
                    new[0] = curr[0]

                if new_interval[1] > curr[0] and new_interval[1] < curr[1]:
                    new[1] = curr[1]      
    
        print "New", new

        if new[0] == 0:
            new[0] = new_interval[0]
        
        if new[1] == 0:
            new[1] = new_interval[1]

        begin.append(new)
        begin.extend(end)

        return begin


s = Solution()
print s.insert([(3,4), (5, 6)], [1, 2])

#(29823256, 32060921), (33950165, 36418956)