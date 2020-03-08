class Interval:
    def __init__(self, s=0, e=0):
        self.start = s
        self.end = e

class Solution:
    
    def insert(self, intervals):
        intervals.sort(key= lambda x: x.start)

        result = [intervals[0]]
        for i in range(1, len(intervals)):
            curr = intervals[i]
            prev = result[-1]

            if curr.start == 2 and curr.end == 5:
                print "here"

            if curr.start > prev.start and curr.start > prev.end:
                result.append(curr)
            elif curr.start >= prev.start and curr.start < prev.end:
                if curr.end > prev.end:
                    prev.end = curr.end
            elif curr.start < prev.start and curr.end < prev.end:
                prev.start = curr.start
            elif curr.start <= prev.start and curr.end > prev.end:
                prev.start = curr.start
                prev.end = curr.end
        
        for i in result:
            print i.start, i.end
        

s = Solution()
#[1,3],[2,6],[8,10],[15,18]
#2, 6, 8, 10
#intervals =[Interval(2, 6 8, 10), Interval(1, 7)]#, Interval(5, 6), Interval(6, 6)]
#intervals = [Interval(54, 75), Interval(56, 60), Interval(61, 86), Interval(22, 43), Interval(56, 87), Interval(32, 53), Interval(14, 81), Interval(64, 65), Interval(9, 42), Interval(12, 33), Interval(22, 58), Interval(84, 90), Interval(27, 59), Interval(41, 48), Interval(43, 47), Interval(22, 29), Interval(16, 23), Interval(41, 72), Interval(15, 87), Interval(20, 59), Interval(45, 84), Interval(14, 77), Interval(72, 93), Interval(20, 58), Interval(47, 53), Interval(25, 88), Interval(5, 89), Interval(34, 97), Interval(14, 47)]
intervals = [Interval(9, 51), Interval(27, 95), Interval(53, 92), Interval(35, 76), Interval(26, 37), Interval(44, 49), Interval(9, 58), Interval(21, 61), Interval(32, 85), Interval(45, 55), Interval(20, 86), Interval(16, 47), Interval(11, 81), Interval(25, 95), Interval(19, 43), Interval(69, 98), Interval(45, 50), Interval(34, 87), Interval(24, 36), Interval(30, 55), Interval(25, 27), Interval(62, 96), Interval(59, 86), Interval(47, 70), Interval(20, 82), Interval(51, 97), Interval(31, 31), Interval(19, 78), Interval(37, 72), Interval(2, 5), Interval(31, 65), Interval(25, 45), Interval(35, 43), Interval(100, 100), Interval(37, 72), Interval(32, 71), Interval(93, 98), Interval(24, 80), Interval(40, 79), Interval(20, 29), Interval(80, 99), Interval(31, 81), Interval(20, 83), Interval(76, 94), Interval(68, 72), Interval(2, 84), Interval(21, 34), Interval(18, 87), Interval(47, 90), Interval(65, 68), Interval(46, 61), Interval(9, 51), Interval(53, 62), Interval(22, 54), Interval(14, 81), Interval(33, 50), Interval(28, 43), Interval(74, 91), Interval(21, 88), Interval(9, 44), Interval(87, 96)]
print s.insert(intervals)