from typing import List

class Solution:
    def toLetters(self, word) -> List[str]:
        ans = list()

        for letter in word:
            ans.append(letter)
        
        return sorted(ans)

    def removeAnagrams(self, words: List[str]) -> List[str]:
        i = 0
        while True:
            word = words[i]
            letters = self.toLetters(word)
            while len(words) - 1 > i:
                next_word = words[i + 1]
                next_word_letters = self.toLetters(next_word)

                if letters == next_word_letters:
                    words.pop(i + 1)
                else:
                    break
            
            if len(words) - 1 == i:
                break

            i += 1
        
        return words
    
s = Solution()
print(s.removeAnagrams(["abba","baba","bbaa","cd","cd"]))