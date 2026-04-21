# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution(object):
    def calc_depth(self, node):
        depth = 0
        while node:
            node = node.left
            depth += 1
        
        return depth
    
    def countNodes(self, root):
        if not root:
            return 0
        
        left_depth = self.calc_depth(root.left)
        right_depth =  self.calc_depth(root.right)
        
        if left_depth == right_depth:
            return 2 ** left_depth + self.countNodes(root.right)
        else:
            return 2 ** right_depth + self.countNodes(root.left)


five = TreeNode(5)
six = TreeNode(6)
three = TreeNode(3, six)
four = TreeNode(4)
two = TreeNode(2, four, five)
one = TreeNode(1, two, three)

s = Solution()
print(s.countNodes(one))