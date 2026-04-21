class Solution(object):
    def fill_recursive(self, y, x, t):
        def check_if_can_swim_in_cell(y, x):
            height = self.grid[y][x]
            can_reach_start = self.can_reach_start_grid[y][x]

            if height <= t and can_reach_start == 0:
                self.fill_recursive(y, x, t)

        if y == self.n - 1 and x == self.n - 1: 
            self.ans = t
            # print('tuta')

        self.can_reach_start_grid[y][x] = 1

        if x > 0:
            check_if_can_swim_in_cell(y, x - 1)
        if x < self.n - 1: # возможно n - 1
            check_if_can_swim_in_cell(y, x + 1)
        if y > 0:
            check_if_can_swim_in_cell(y - 1, x)
        if y < self.n - 1: # возможно n - 1
            check_if_can_swim_in_cell(y + 1, x)


    def swimInWater(self, grid):
        """
        :type grid: List[List[int]]
        :rtype: int
        """
        self.grid = grid
        self.ans = 0
        self.hashmap = dict()
        self.n = len(grid)
        if self.n == 1:
            return 0
        self.can_reach_start_grid = [[0 for j in range(self.n)] for i in range(self.n)]
        # self.can_reach_start_grid[0][0] = 1

        for y in range(self.n):
            for x in range(self.n):
                str_val = str(grid[y][x])
                self.hashmap[str_val] = [y, x]
        
        for t in range(self.n ** 2):
            if t > self.grid[0][0]:
                y, x = self.hashmap[str(t)]

                if x > 0:
                    if self.can_reach_start_grid[y][x - 1] == 1:
                        self.fill_recursive(y, x, t)
                if x < self.n - 1: # возможно n - 1
                    if self.can_reach_start_grid[y][x + 1] == 1:
                        self.fill_recursive(y, x, t)
                if y < self.n - 1: # возможно n - 1
                    if self.can_reach_start_grid[y + 1][x] == 1:
                        self.fill_recursive(y, x, t)
                if y > 0:
                    if self.can_reach_start_grid[y - 1][x] == 1:
                        self.fill_recursive(y, x, t)
                
            elif t == self.grid[0][0]:
                y, x = self.hashmap[str(t)]
                self.fill_recursive(y, x, t)

            if self.ans: return self.ans
                
            

s = Solution()
print(s.swimInWater([[0]]))
print(s.swimInWater([[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6]]))
print(s.swimInWater([[3,2],[0,1]]))