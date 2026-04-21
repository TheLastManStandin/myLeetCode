class Solution(object):
    def __init__(self):
        self.ans_list = list()

    def findRepeatedDnaSequences(self, s):
        """
        :type s: str
        :rtype: List[str]
        """
        len_of_dna = len(s)
        if len_of_dna < 10:
            return [] # Если меншье 10 элементов на вход
        
        current_sequence = s[:10]
        pointer = 10 # индекс первого символа последовательности
        hashtable = dict()

        while len_of_dna >= pointer:
            if current_sequence not in hashtable:
                hashtable[current_sequence] = 1
            else:
                hashtable[current_sequence] += 1
            
            current_sequence = current_sequence[1:]
            if pointer != len_of_dna:
                current_sequence = current_sequence + s[pointer]

            pointer += 1

        for sequence in hashtable:
            if hashtable[sequence] > 1:
                self.ans_list.append(sequence)
        
        print(self.ans_list)
        return self.ans_list
        
string = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

sol = Solution()
sol.findRepeatedDnaSequences(string)