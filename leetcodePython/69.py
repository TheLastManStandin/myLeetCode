class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 0:
            return 0
        
        left = 0
        right = x
        ans = -1
        while left <= right:
            mid = (right + left) // 2

            if mid * mid < x:
                ans = mid
                left = mid + 1
            elif mid * mid > x:
                right = mid - 1
            elif mid * mid == x:
                return mid

        return ans

s = Solution()
print(s.mySqrt(9))