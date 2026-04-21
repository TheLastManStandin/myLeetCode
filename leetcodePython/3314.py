class Solution(object):
    def minBitwiseArray(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        ans: list[int] = list()
        for num in nums:
            start: int = 0
            while start < num:
                if (start | start + 1) == num:
                    ans.append(start)
                    break
                start += 1
            if start == num:
                ans.append(-1)

        return ans
            


s = Solution()
print(s.minBitwiseArray([2,3,5,7]))