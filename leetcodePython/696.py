# class Border:
#     def __init__(self, left):
#         self.left = left
#         self.right = 0
#     def process(self, massive :list, direction: int):
#         """
#         direction - слева нули или еденицы
#         """
#         # if self.right == self.left:
#         a = min(self.left, self.right)
#         for i in range(a, 0, -1):
#             s = ("0" if direction else "1") * i + ("1" if direction else "0") * i

#             massive.append(s)

class Solution(object):
    def countBinarySubstrings(self, s):
        """
        :type s: str
        :rtype: int
        """
        out = 0
        flag = int(s[0])
        count_l = 0
        count_r = 0
        for start_index in range(len(s)):
            if s[start_index] != str(flag):
                out += min(count_l, count_r)
                count_r = count_l
                flag = 1 if not flag else 0
                count_l = 0
            count_l += 1

        out += min(count_l, count_r)

        return out

s = Solution()
print(s.countBinarySubstrings("00110011"))
print(s.countBinarySubstrings("00110"))