# проходимся по массиву с последнего элемента, паралельно меняя значения массива на значения суммы элементов алгаритма после. Далее находим самое большое число в массиве.
from typing import List

class Solution:
    def maximumEnergy(self, energy: List[int], k: int) -> int:
        energy_len = len(energy)
        for backward_i in range(energy_len, -1, -1): # Проходимся по массиву с конца 
            forward_i = energy_len - backward_i

            if forward_i > k:
                energy[backward_i] += energy[backward_i + k]
        
        print(energy)

        return max(energy)
    

s = Solution()
print(s.maximumEnergy([5,2,-10,-5,1], 3))