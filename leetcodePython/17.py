class Solution(object):
    def __init__(self):
        self.phone_dict = {
            '2': ['a', 'b', 'c'],
            '3': ['d', 'e', 'f'],
            '4': ['g', 'h', 'i'],
            '5': ['j', 'k', 'l'],
            '6': ['m', 'n', 'o'],
            '7': ['p', 'q', 'r', 's'],
            '8': ['t', 'u', 'v'],
            '9': ['w', 'x', 'y', 'z']
        }
        
        self.ans_massive = []

    def backtrack(self, digit, digits_last, cur_combination):
        digits_last.pop(0)
        for letter in self.phone_dict[digit]:
            cur_combination += letter

            if digits_last == []:
                self.ans_massive.append(cur_combination)
            else:
        #         # for next_digit_index in range(len(digits_last)):
        #         #     next_digit = digits_last.pop(next_digit_index)
        #         #     self.backtrack(next_digit, digits_last, cur_combination)
        #         #     digits_last.insert(next_digit_index, next_digit)
                    new_digit = digits_last[0]
                    self.backtrack(new_digit, digits_last, cur_combination)
                    # digits_last.insert(0, new_digit)
            cur_combination = cur_combination[:-1]
        digits_last.insert(0, digit)

    def letterCombinations(self, digits):
        """
        :type digits: str
        :rtype: List[str]
        """
        inp = list(digits)
        if inp:
            self.backtrack(inp[0], inp, "")

        # print(self.ans_massive)
        # print(len(self.ans_massive))
        return self.ans_massive

s = Solution()
s.letterCombinations("23")