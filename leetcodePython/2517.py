class Solution(object):
    def maximumTastiness(self, price, k):
        """
        :type price: List[int]
        :type k: int
        :rtype: int
        """
        sorted_price = sorted(price)
        max_price = sorted_price[-1]
        min_price = sorted_price[0]
        # max_testiness = (max_price - min_price) // (k - 1)
        max_testiness = max_price
        min_testiness = 0
        ans = 0
        prev_maxes = list()

        cur_testiness = (max_testiness - min_testiness) // 2

        while True:
            cur_maxes = [max_testiness, min_testiness]
            pointer = min_price
            # cur_k = k - 1 # осталось собрать конфет до полной коробки
            nums_in = 1
            
            # while cur_k > 0:
            while nums_in <= k and pointer <= max_price:
                pointer += cur_testiness
                
                while pointer not in sorted_price and pointer <= max_price:
                    pointer += 1
                
                if pointer in sorted_price: 
                    nums_in += 1

                if nums_in >= k:
                    ans = cur_testiness
                    break
            
            if pointer > max_price: # нужна меньшая половина
                max_testiness = cur_testiness
                cur_testiness = min_testiness + (max_testiness - min_testiness) // 2
            else:
                min_testiness = cur_testiness
                cur_testiness = min_testiness + (max_testiness - min_testiness) // 2

            if cur_maxes == prev_maxes:
                return ans
            else:
                prev_maxes = cur_maxes
            

        return ans
    
s = Solution()

print(s.maximumTastiness([13,5,1,8,21,2], 3))