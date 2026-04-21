
class Solution:
    def missingNumber(self, nums) -> int:
        l = 0
        nums.sort()
        if nums[-1] == len(nums) - 1:
            return nums[-1] + 1
        n = len(nums)
        r = n
        ans = -1
        while l <= r:
            mid = (l + r) // 2
            # print(mid)
            if nums[mid] == mid:
                ans = mid
                l = mid + 1
            elif nums[mid] > mid:
                r = mid - 1
            
        return ans + 1
    
s = Solution()
print(s.missingNumber([9,6,4,2,3,5,7,0,1]))
print(s.missingNumber([0,1,2,3,4]))
print(s.missingNumber([1]))