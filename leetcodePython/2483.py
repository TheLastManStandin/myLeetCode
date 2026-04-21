def calc_map(customers):
  mapp = list()

  for penalty_index in range(len(customers)):
    mapp.append(1 if customers[penalty_index] == "Y" else 0)
  return mapp

class Solution(object):
    def bestClosingTime(self, customers):
        """
        :type customers: str
        :rtype: int
        """
        mapp = calc_map(customers)
        mapp_summ = sum(mapp)
        lowest_penalty = mapp_summ
        out_index = 0

        for mapp_index in range(len(mapp)):
          if mapp[mapp_index] == 1:
            mapp_summ -= 1
          else:
            mapp_summ += 1

          if mapp_summ < lowest_penalty:
            lowest_penalty = mapp_summ
            out_index = mapp_index + 1

        return out_index

s = Solution()
print(s.bestClosingTime("YYNY"))