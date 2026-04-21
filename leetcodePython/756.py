def calc_possible_letters_above(row: str, allowed: list()):
  # Проходим по основанию по 2 буквы, сравниваем их с каждым основанием allowed
  possible_letters_above = list()

  for bottom_two_index in range(len(row) - 1):
    row_two = row[bottom_two_index:bottom_two_index + 2]
    list_of_possible_letters_above_bottom_two = list()

    for allowed_two_index in range(len(allowed)):
      allowed_two = allowed[allowed_two_index][:2]

      if row_two == allowed_two:
        if allowed[allowed_two_index][2] not in list_of_possible_letters_above_bottom_two:
          list_of_possible_letters_above_bottom_two.append(allowed[allowed_two_index][2])

    possible_letters_above.append(list_of_possible_letters_above_bottom_two)

  return possible_letters_above


def calc_possible_rows_above(possible_letters_above: list(), allowed: list()):
  def backtrack_to_find_possible_rows(pos_in_row, possible_letters_above, path, allowed, result):
    if pos_in_row < possible_letters_above.__len__():
      for letter in possible_letters_above[pos_in_row]:
        two_ = path[-1] + letter
        
        for allowed_two_index in range(len(allowed)):
          if allowed[allowed_two_index][:2] == two_:
            path.append(letter)
            backtrack_to_find_possible_rows(pos_in_row+1, possible_letters_above, path, allowed, result)
            path.pop()
            break
    else:
      # print(path)
      result.append(path.copy())
    
  result = list()
  for letter in possible_letters_above[0]:
    path = list()
    path.append(letter)
    backtrack_to_find_possible_rows(1, possible_letters_above, path, allowed, result)

  for i in range(len(result)):
    result[i] = ''.join(result[i])

  return (result)


def main_backtrack(bottom, allowed):
  possible_letters_above = calc_possible_letters_above(bottom, allowed)

  for possible_letter in possible_letters_above:
    if possible_letter == []:
      return 1

  possible_rows_above = calc_possible_rows_above(possible_letters_above, allowed)

  if len(possible_rows_above) == 0:
    return 1

  if len(bottom) == 2:
    if len(possible_rows_above) > 0:
      return 0
    else:
      return 1
  else:
    for new_bottom in possible_rows_above:
      return main_backtrack(new_bottom, allowed)


class Solution(object):
  def pyramidTransition(self, bottom, allowed):
    """
    :type bottom: str
    :type allowed: List[str]
    :rtype: bool
    """
    if main_backtrack(bottom, allowed) == 0:
      return True
    else:
      return False
        

s = Solution()
print(s.pyramidTransition(bottom = "FDBACE", allowed = ["EEF","BFE","EDD","EFB","EFC","DCE","CCE","ABB","BBB","EDC","ADD","AFE","CAF","DEF","ABE","BBD","CBB","ADB","ABD","EDF","FAE","CAA","CFB","BCA","BDC","EAB","FFE","FBF","EFF","AFD","DFA","BED","BDD","ABA","FCB","CBD","CDC","CEC","ECC","ECA","EBC","DFD","FFB","BDE","DBD","EBB","DEB","BEF","FFA","AEA","CCC","BFF","DCD","BBA","CFF","ECD","CBF","CCD","FAA","EDA","ADF","ECE","FCF","FFF","FCE","BFC","CCF","ACD","FDB","DBA","AED","FDD","BDF","FBE","DCB","ACE","FBC","FEF","FDF","AEF","DDB","CFA","BCB","EFA","EAC","FBD","CFC","AEE","CEB","AFA","CCA","BFD","BAC","BAA","CEE","DAB","AFC","DBE","BEE","DAF","DCA","EEA","BDB","EEB","EAA","BEC","DED","CDE","CDB","EEE","DAC","EBF","EBD","FDE","ABC","ACB","DBC","FBA","BAE","EFE","BBC","CBC","FED","FEA","ACF","DCF","FDA","BCC","ADE","DAE","DCC","EDB","AFB","CEA","DFE","BAD","FEC","EEC","EBE","CEF","EAD","ABF","EFD","AAB","AAD","FAB","FEE","CBE","BBE","ADC","FAD","DBB","CAB","CDA","AAF","DBF","FCA","DEE","EDE","FFC","DDD","DDA","DEC","DFF","BCD","ECF","DDF","AEB","BDA","FBB","BCE","DAA","ACC","CCB","FAC","BAF","BEA","CFD","EBA","ACA","DAD","BFB","ECB","CAD","DDC","FCC","BEB","FAF","BBF","AAA","AAC","CED","AAE","CDD","DDE","DEA","FFD","DFC","CFE","FEB","FDC","ADA","BCF","AFF","EAE","AEC","CAC","CDF","BAB","EED","CAE","FCD","BFA","EAF","CBA","DFB"]))