def maxProfit(prices):
    minPrice = float('inf')
    profitArr = []
    maxProfit = 0
    i = 0
    while i < len(prices):
        if (prices[i] < minPrice):
            minPrice = prices[i]
        else:
            profitArr.append(prices[i]-minPrice)
            minPrice = prices[i]
        i += 1

    print profitArr

    return sum(profitArr)

maxProfit([3, 3, 5, 0, 0, 3, 1, 4])