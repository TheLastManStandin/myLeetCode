class Solution(object):
    def numWaterBottles(self, numBottles, numExchange):
        """
        :type numBottles: int
        :type numExchange: int
        :rtype: int
        """
        ans = 0
        remainder = 0
        while numBottles + remainder >= numExchange:
            ans += numBottles
            numBottles += remainder
            remainder = numBottles % numExchange
            numBottles = numBottles // numExchange

        return ans + numBottles
    
s = Solution()
print(s.numWaterBottles(15, 4))