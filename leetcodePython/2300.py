class Solution(object):
    def find_min_successful_potion(self, potions, enough_success, l, r):
        """
        Находит индекс ближайшего элемента, который <= enough_success.
        Если все элементы больше — возвращает -1.
        """
        result = -1
        while l <= r:
            mid = (l + r) // 2
            if potions[mid] < enough_success:
                result = mid      # запоминаем возможный ответ
                l = mid + 1       # идём вправо — ищем больший, но всё ещё подходящий
            else:
                r = mid - 1       # идём влево — текущий элемент слишком большой
        return result


    def successfulPairs(self, spells, potions, success):
        """
        :type spells: List[int]
        :type potions: List[int]
        :type success: int
        :rtype: List[int]
        """
        out = list()
        potions.sort()
        potions_len = len(potions)

        for spell in spells:
            potion_index = self.find_min_successful_potion(potions, success / spell, 0, potions_len - 1)

            
            n = potions_len - potion_index

            out.append(n - 1)

        return out

s = Solution()
print(s.successfulPairs([5,1,3], [1,2,3,4,5],7))
print(s.successfulPairs([3,1,2], [8,5,8], 16))