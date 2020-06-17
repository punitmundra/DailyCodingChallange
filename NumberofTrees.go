/*
 kattia

Number of Trees

Number of trees

X is a great mathematician and is fond of calculating large numbers. His friend is fond of making numbers with trees, and so he asks X to give him
n consecutive numbers so that he could make a binary search tree using them. Recall that a binary search tree is one in which the values in the left subtree
are smaller than the root node, while the values in the right subtree are greater than the root node.

The friend was wondering, if given n numbers, if X could find the total number of possible binary search trees using them.

X is feeling lazy and so he asks you for help. Write a program to help him out!

Note: A binary search tree is a tree that can have maximum of two children, with the value of the left child node always less than
the value of the node and the value of right child always greater than the value of the node.

Input Format

The first line contains t which represents the number of test cases. Then t lines follow each of which contains an integer n. This n is
the number of consecutive numbers.

Output Format

The output contains t lines. Each line contains the number of trees that can be made from the given n consecutive numbers.
Constraints

    1<=t<=10000

    1<=n<=1000

Sample Input

2
2
3

Sample Output

2
5

Explanation

For the 1st test case, n=2. 2 binary search trees can be formed using two consecutive numbers as shown in the figure.
	


*/
