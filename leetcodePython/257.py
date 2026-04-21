class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def binaryTreePaths(self, root: TreeNode) -> list[str]:
        out = list()
        def recursion(node: TreeNode, path: str):
            node.val = str(node.val)
            path = str(path)
            if path == node.val:
                new_path = path
            else:
                new_path = path + "->" + node.val

            if node.left == None and node.right == None:
                nonlocal out
                out.append(new_path)
            else:
                if node.left:
                    recursion(node.left, new_path)
                if node.right:
                    recursion(node.right, new_path)

        recursion(root, root.val)
        return out

five = TreeNode("5")
three = TreeNode('3')
two = TreeNode('2', five)
one = TreeNode('1', two, three)

a = Solution()
a.binaryTreePaths(one)