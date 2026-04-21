class Solution(object):
    def __init__(self):
        self.ans = list()

    def backtrack(self, n, n_last, balance, string):
        def plus_open():
            nonlocal string, balance, n_last

            string += "("
            balance += 1
            n_last -= 1
        
        def plus_closed():
            nonlocal string, balance, n_last

            string += ")"
            balance -= 1
            # n_last -= 1
        
        def minus():
            nonlocal string, balance, n_last

            if string[-1:] == "(":
                balance -= 1
                n_last += 1
            else:
                balance += 1

            string = string[:-1]

        if n_last or balance:
            if balance == 0:
                if n_last > 0:
                    plus_open()    
                    self.backtrack(n, n_last, balance, string)
                    minus()
                else:
                    pass
            else:
                if n_last > 0:
                    plus_open()
                    self.backtrack(n, n_last, balance, string)
                    minus()

                    plus_closed()
                    self.backtrack(n, n_last, balance, string)
                    minus()
                else:
                    plus_closed()
                    self.backtrack(n, n_last, balance, string)
        else:
            self.ans.append(string)
            
    def generateParenthesis(self, n):
        """
        :type n: int
        :rtype: List[str]
        """
        self.backtrack(n, n, 0, "")

        print(self.ans)
        # return self.ans


s = Solution()
s.generateParenthesis(6)
