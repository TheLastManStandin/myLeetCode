# from typing import List

# class Solution:
#     def insert_spell(self, spell_index):
#         l = 1
#         r = len(self.sorted_spells_by_total_power)
#         spell_power = self.hashmap[str(spell_index)]
#         ans = -1

#         # if r == -1:
#         #     self.sorted_spells_by_power.append(spell)

#         while l <= r:
#             mid = (l + r) // 2
#             power_mid = self.hashmap[str(self.sorted_spells_by_total_power[mid - 1])]
#             if power_mid > spell_power:
#                 r = mid - 1
#             elif power_mid <= spell_power:
#                 ans = mid
#                 l = mid + 1

#         self.sorted_spells_by_total_power.insert(ans, spell_index)

#     def process_sequence(self, sequence) -> int: # Получаем максимальный урон из этой последовательности
#         mid = len(sequence) / 2

#         for i in range(len(self.sorted_spells_by_total_power), -1, -1):
#             spell_i_in_sorted = self.sorted_spells_by_total_power[i]
#             total_power_of_i_spell = self.hashmap[spell_i_in_sorted]
#             if spell_i_in_sorted in sequence:
#                 j = i # ищем все спелы с такой же тотальной силой
#                 while self.sorted_spells_by_total_power[j] == spell_i_in_sorted:
#                     pass


#     def maximumTotalDamage(self, power: List[int]) -> int:
#         # словарь
#         # power[i], i - спелл
#         # power[i] - сила спела
#         # hashmap[str(i)] - сила спела * количество спелов = total_power
#         # 
#         self.hashmap = dict()
#         self.sorted_spells_by_total_power = list()

#         power.sort()

#         for spell in power:
#             if str(spell) in self.hashmap:
#                 # print(spell)
#                 self.hashmap[str(spell)] += spell
#             else:
#                 self.hashmap[str(spell)] = spell

#         # print("hashmap = ", hashmap)

#         for spell in self.hashmap: # формирую sorted_spells_by_power
#             self.insert_spell(spell)

#         # прохожусь по ssbp отмечая
#         current_continuous_sequence = [power[0]]
#         prew_num = power[0]
#         ans = 0

#         for i in range(1, len(power)):
#             cur_num = power[i]

#             if cur_num - prew_num > 2:
#                 ans += self.process_sequence(current_continuous_sequence)
#                 current_continuous_sequence = [cur_num]
#             elif cur_num != prew_num:
#                 current_continuous_sequence.append(cur_num)

#             prew_num = cur_num

#         ans += self.process_sequence(current_continuous_sequence)

#         return ans
from bisect import bisect_left

class Solution:
    def maximumTotalDamage(self, power):
        if not power:
            return 0

        power.sort()
        vals, earn = [], []

        i = 0
        while i < len(power):
            v = power[i]
            total = 0
            while i < len(power) and power[i] == v:
                total += v
                i += 1
            vals.append(v)
            earn.append(total)

        n = len(vals)
        dp = [0] * (n + 1)

        for i in range(n - 1, -1, -1):
            next_idx = bisect_left(vals, vals[i] + 3)
            take = earn[i] + dp[next_idx]
            skip = dp[i + 1]
            dp[i] = max(take, skip)

        return dp[0]        

s = Solution()

# print(s.maximumTotalDamage([1,2,2,2,3,3,4,5,6,6,6,6,7,8,8,8])) # худший вариант
print(s.maximumTotalDamage([1,2,2,2,3,3,4,5,6,6,6,6,7,8,8,8])) # худший вариант

# 5,1,6,6,1,6,6,1,6,6,1,4
# 1,6,6,4,5,24,7,24