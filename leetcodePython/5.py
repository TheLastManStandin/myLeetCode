class Solution(object):
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        len_of_longest = 0
        longest_palindrom = s[0]
        index = 0

        # for index in range(len(s)):
        while index < len(s):
            cur_len = 1
            symbol = s[index]
            l_pointer = index
            r_pointer = index
            flag1 = True
            flag2 = True
            
            while l_pointer > 0 and (flag1 or flag2): # перемещаем указатели до первого и последнего повторяющихся элементов
                if s[l_pointer - 1] == symbol:
                    cur_len += 1
                    l_pointer -= 1
                else:
                    flag1 = False

                if r_pointer < len(s) - 1:
                    if s[r_pointer + 1] == symbol:
                        cur_len += 1
                        r_pointer += 1
                        index += 1
                    else:
                        flag2 = False
                else:
                    flag2 = False

            while l_pointer > 0 and r_pointer < len(s) - 1:
                l_pointer -= 1
                r_pointer += 1

                if s[l_pointer] == s[r_pointer]:
                    cur_len += 2
                else:
                    l_pointer += 1
                    r_pointer -= 1
                    break
            
            if cur_len > len_of_longest:
                len_of_longest = cur_len
                longest_palindrom = s[l_pointer:r_pointer + 1]
            
            index += 1
        return longest_palindrom
    
s = Solution()
print(s.longestPalindrome("babad"))