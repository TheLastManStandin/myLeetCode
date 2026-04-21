# The isBadVersion API is already defined for you.
def isBadVersion(version: int) -> bool:
    if version > 0:
        return True
    return False

class Solution:
    def firstBadVersion(self, n: int) -> int:
        l = 0
        r = n
        ans = -1
        while l <= r:
            mid = (r + l) // 2
            if isBadVersion(mid):
                r = mid - 1
            else:
                ans = mid
                l = mid + 1
        
        return ans + 1

s = Solution()
print(s.firstBadVersion(1))