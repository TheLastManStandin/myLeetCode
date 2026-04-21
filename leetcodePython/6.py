class Solution:
    def convert(self, s: str, numRows: int) -> str:
        if numRows == 1:
            return s
        
        out = list()
        
        letters_in_one_zag = numRows * 2 - 2

        # if len(s) % letters_in_one_zag == 0:
        #     zags_num = len(s) // letters_in_one_zag
        # else:
        #     zags_num = len(s) // letters_in_one_zag + 1
        zags_num = (len(s) + letters_in_one_zag - 1) // letters_in_one_zag

        # print(zags_num, letters_in_one_zag)

        i = 0

        for row_index in range(numRows):
            for zag in range(zags_num):
                if row_index == 0:
                    out.append(s[zag * letters_in_one_zag])
                elif row_index == numRows - 1:
                    if zag * letters_in_one_zag + numRows - 1 < len(s):
                        out.append(s[zag * letters_in_one_zag + numRows - 1])
                else:
                    if zag * letters_in_one_zag + row_index < len(s):
                        out.append(s[zag * letters_in_one_zag + row_index])
                    if zag * letters_in_one_zag + (letters_in_one_zag - row_index) < len(s):
                        out.append(s[zag * letters_in_one_zag + (letters_in_one_zag - row_index)])

        return ''.join(out)
    
s = Solution()
print(s.convert("PAYPALISHIRING", 3))
