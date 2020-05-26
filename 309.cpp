/*
There are M people sitting in a row of N seats, where M < N. 
Your task is to redistribute people such that there are no gaps between any of them, 
while keeping overall movement to a minimum.

For example, 
suppose you are faced with an input of [0, 1, 1, 0, 1, 0, 0, 0, 1], 
where 0 represents an empty seat and 1 represents a person. 
In this case, one solution would be to place the person on the right in the fourth seat. 
We can consider the cost of a solution to be the sum of the absolute distance each person must move, 
so that the cost here would be five.


in simple words move all 1 together so that relative displacement is minimal. 
*/
