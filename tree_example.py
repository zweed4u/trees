#!/usr/bin/python3
"""
            5
    3               8   
 1      4       7       9
    2

5->3(left) and 8(right)
3->1(left) and 4(right)
1->2(right)
8->7(left) and 9(right)
https://pythonspot.com/wp-content/uploads/2015/09/tree.jpg.webp
"""


class Node:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


# Setup similar to breadth first search (level-by-level) horizontally
bfs_root = Node(5)
bfs_root.left = Node(3)
bfs_root.right = Node(8)
bfs_root.left.left = Node(1)
bfs_root.left.right = Node(4)
bfs_root.right.left = Node(7)
bfs_root.right.right = Node(9)
bfs_root.left.left.right = Node(2)


# Setup similar to depth first search
dfs_root = Node(5)
dfs_root.left = Node(3)
dfs_root.left.left = Node(1)
dfs_root.left.left.right = Node(2)
dfs_root.left.right = Node(4)
dfs_root.right = Node(8)
dfs_root.right.left = Node(7)
dfs_root.right.right = Node(9)
