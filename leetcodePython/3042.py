class Solution(object):
    def countPrefixSuffixPairs(self, words):
        """
        :type words: List[str]
        :rtype: int
        """
        # num_of_words_fix_in
        for word_index in range(len(words) - 1, -1, -1):
            print(word_index)

s = Solution()
s.countPrefixSuffixPairs(["a","b"])