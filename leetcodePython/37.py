from pprint import pprint

class Solution:
    def __init__(self):
        self.sudoku_solved = False
        self.ans_board = list()

    def find_the_best_cell_for_brutforce(self, board):
        candidates = self.calc_condidates(board)
        min_candidates = 9
        min_x = 9
        min_y = 9

        for y in range(9):
            for x in range(9):
                candidates_num_of_cell = len(candidates[y][x])
                if candidates_num_of_cell < min_candidates and candidates_num_of_cell > 1:
                    min_candidates = len(candidates[y][x])
                    min_x = x
                    min_y = y

        return [min_y, min_x]

    def debug_log(self, message):
        if self.debug == True:
            print(message)

    def calc_condidates(self, board):
        candidates = list()
        for i in range(9):
            candidates.append([])
            for j in range(9):
                candidates[i].append([])
        emptys = False
        
        for y in range(9):
            for x in range(9):
                val = board[y][x]

                if val == ".":
                    emptys = True
                    # print(x, y)
                    # if candidates[y][x] == []:
                    candidates[y][x] = ['1','2','3','4','5','6','7','8','9']

                    for x_candidates in range(9): # проходимся по ряду
                        x_candidate = board[y][x_candidates]
                        cur_candidates = candidates[y][x]

                        if x_candidate in cur_candidates:
                            candidates[y][x].remove(x_candidate)

                    for y_candidates in range(9): # проходимся по колонне 
                        y_candidate = board[y_candidates][x]

                        if y_candidate in candidates[y][x]:
                            candidates[y][x].remove(y_candidate)

                    box_x_start = (x // 3) * 3
                    box_y_start = (y // 3) * 3

                    for i in range(3): # проходимся по квадрату
                        for j in range(3):
                            box_x = box_x_start + j
                            box_y = box_y_start + i
                            # print(box_y, box_x)

                            if board[box_y][box_x] in candidates[y][x]:
                                candidates[y][x].remove(board[box_y][box_x])
                else:
                    candidates[y][x] = ['0']

        # if emptys == False:
        #     self.ans_board = board
        #     self.sudoku_solved = True

        return candidates
    
    def open_nums(self, board): # возвращает false если хотя бы в одной ячейке есть один единственный кандидат
        candidates = list()
        for i in range(9):
            candidates.append([])
            for j in range(9):
                candidates[i].append([])

        candidates = self.calc_condidates(board)
        local_board = self.copy_board(board)

        for y in range(9):
            for x in range(9):
                if len(candidates[y][x]) == 1 and candidates[y][x][0] != '0':
                    local_board[y][x] = candidates[y][x][0]

        return local_board
        
    def hidden_nums(self, board): # ищем уникальные цифры из кандидатов в ряду, колонне, и блоке 3 на 3
        candidates = list()
        for i in range(9):
            candidates.append([])
            for j in range(9):
                candidates[i].append([])

        candidates = self.calc_condidates(board)
        local_board = self.copy_board(board)

        hidden_num = 0

        def check_row(candidates, x, y):
            candidates_last = [1, 2, 3, 4, 5, 6, 7, 8, 9]
            for x_coord_candidates_index in range(9):
                if x_coord_candidates_index != x:
                    for candidate in candidates[y][x_coord_candidates_index]: # вычитаем кандидатов ячейки из ряда из candidates_last
                        if candidate in candidates_last:
                            candidates_last.remove(candidate)

            for candidate in candidates_last:
                if candidate in candidates[y][x]: # если есть уникальная цифра в кандидатах, это и есть наш ответ
                    nonlocal hidden_num
                    hidden_num = candidate

        def check_col(candidates, x, y):
            candidates_last = [1, 2, 3, 4, 5, 6, 7, 8, 9]
            for y_coord_candidates_index in range(9):
                if y_coord_candidates_index != y:
                    for candidate in candidates[y_coord_candidates_index][x]: # вычитаем кандидатов ячейки из ряда из candidates_last
                        if candidate in candidates_last:
                            candidates_last.remove(candidate)

            for candidate in candidates_last:
                if candidate in candidates[y][x]: # если есть уникальная цифра в кандидатах, это и есть наш ответ
                    nonlocal hidden_num
                    hidden_num = candidate

        def check_box(candidates, x, y):
            candidates_last = [1, 2, 3, 4, 5, 6, 7, 8, 9]
            box_x_start = (x // 3) * 3
            box_y_start = (y // 3) * 3

            for i in range(3): # проходимся по квадрату
                for j in range(3):
                    box_x = box_x_start + j
                    box_y = box_y_start + i

                    for candidate in candidates[box_y][box_x]:
                        if candidate in candidates_last:
                            candidates_last.remove(candidate)
            
            for candidate in candidates_last:
                if candidate in candidates[y][x]: # если есть уникальная цифра в кандидатах, это и есть наш ответ
                    nonlocal hidden_num
                    hidden_num = candidate

        for y in range(9):
            for x in range(9):
                if local_board[y][x] == ".":
                    check_row(candidates, x, y)
                    check_col(candidates, x, y)
                    check_box(candidates, x, y)

                    if hidden_num != 0:
                        local_board[y][x] = hidden_num

        return local_board

    def check_valid(self, board):
        candidates = list()
        for i in range(9):
            candidates.append([])
            for j in range(9):
                candidates[i].append([])
        candidates = self.calc_condidates(board)
        # if board[3][1] == '1' and board[3][-1] == '1': pass

        board_have_emptys = False

        for y in range(9):
            for x in range(9):
                if candidates[y][x] == []:
                    self.debug_log(str(y) + " " + str(x) + " no candidates")
                    return False
                
        # проверяем строки
        for row in board:
            nums = [x for x in row if x != '.']
            if len(nums) != len(set(nums)):
                self.debug_log(str(y) + " " + str(x) + " in row")
                return False

        # проверяем столбцы
        for col in range(9):
            nums = [board[row][col] for row in range(9) if board[row][col] != '.']
            if len(nums) != len(set(nums)):
                self.debug_log(str(y) + " " + str(x) + " in col")
                return False

        # проверяем квадраты 3x3
        for box_y in range(0, 9, 3):
            for box_x in range(0, 9, 3):
                nums = []
                for y in range(box_y, box_y + 3):
                    for x in range(box_x, box_x + 3):
                        if board[y][x] != '.':
                            nums.append(board[y][x])
                        else:
                            board_have_emptys = True
                if len(nums) != len(set(nums)):
                    self.debug_log(str(y) + " " + str(x) + " in box")
                    return False
        
        if board_have_emptys == False:
            self.ans_board = board
            self.sudoku_solved = True

        return True


    def copy_board(self, board):
        new_board = list()
        for row in range(9):
            new_board.append(board[row].copy())
        return new_board
    
    def recursion(self, board):
        recursion_board_copy = self.copy_board(board)

        while True:
            while_board_copy = self.copy_board(recursion_board_copy)
            while_board_copy = self.open_nums(while_board_copy)
            while_board_copy = self.hidden_nums(while_board_copy)

            if while_board_copy == recursion_board_copy:
                break
            
            recursion_board_copy = while_board_copy
        
        # print("вышел из while") #

        if self.check_valid(recursion_board_copy):
            if self.sudoku_solved:
                return
            
            # for y in range(9):
            #     flag = False
            #     for x in range(9):
            #         if recursion_board_copy[y][x] == ".":
            #             flag = True
            #             break
            #     if flag: break
            y_and_x = self.find_the_best_cell_for_brutforce(recursion_board_copy)
            y = y_and_x[0]
            x = y_and_x[1]

            in_cycle_candidates = self.calc_condidates(recursion_board_copy)
            for candidat_to_empty in in_cycle_candidates[y][x]:
                recursion_board_copy[y][x] = candidat_to_empty
                self.recursion(recursion_board_copy)
                if self.sudoku_solved == True:
                    # self.ans_board = recursion_board_copy
                    return
                recursion_board_copy[y][x] = "."
                    
        else:
            return

    def solveSudoku(self, board: list[list[str]]) -> None:
        """
        Do not return anything, modify board in-place instead.
        """
        # ЗАДАНИЕ СО ЗВЕЗДОЧКОЙ Переделать, сделать отдельный класс Sudocu в котором будут сразу все показатели о конкретном состоянии поля
        candidates = list()
        for i in range(9):
            candidates.append([])
            for j in range(9):
                candidates[i].append([])

        candidates = self.calc_condidates(board)

        self.debug = True

        self.recursion(board)
        
        # board = self.copy_board(self.ans_board)
        for y in range(9):
            for x in range(9):
                board[y][x] = self.ans_board[y][x]
        self.debug_log(str(board))
        # return board
    
# board = [["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]
# board = [[".",".","9","7","4","8",".",".","."],["7",".",".",".",".",".",".",".","."],[".","2",".","1",".","9",".",".","."],[".",".","7",".",".",".","2","4","."],[".","6","4",".","1",".","5","9","."],[".","9","8",".",".",".","3",".","."],[".",".",".","8",".","3",".","2","."],[".",".",".",".",".",".",".",".","6"],[".",".",".","2","7","5","9",".","."]]
# board = [[".",".",".","2",".",".",".","6","3"],["3",".",".",".",".","5","4",".","1"],[".",".","1",".",".","3","9","8","."],[".",".",".",".",".",".",".","9","."],[".",".",".","5","3","8",".",".","."],[".","3",".",".",".",".",".",".","."],[".","2","6","3",".",".","5",".","."],["5",".","3","7",".",".",".",".","8"],["4","7",".",".",".","1",".",".","."]]
board = [[".",".",".",".",".",".",".",".","."],[".","9",".",".","1",".",".","3","."],[".",".","6",".","2",".","7",".","."],[".",".",".","3",".","4",".",".","."],["2","1",".",".",".",".",".","9","8"],[".",".",".",".",".",".",".",".","."],[".",".","2","5",".","6","4",".","."],[".","8",".",".",".",".",".","1","."],[".",".",".",".",".",".",".",".","."]]

a = Solution()
a.solveSudoku(board)

