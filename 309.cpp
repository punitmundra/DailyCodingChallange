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

index : 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14
seats : . . . . x . . x x .  .  .  x  .  .


current occupied seat
ind = [0,1,2,3]
arr = [4,7,8,12]
mid = arr.size() / 2 = 2 


start = arr[i];
end =  arr[mid]-mid + i

start = ... arr[] 
end = centerPointIndex - midPoint + i 

start = 4 
end = 8 - 2 + 0 = 6
jump = abs(end-start) = 2

start = 7 
end = 8 - 2 + 1 = 7
jump = abs(end-start) = 0


start = 8 
end = 8 - 2 + 2 = 8
jump = abs(end-start) = 0

start = 12 
end = 8 - 2 + 3 = 9
jump = abs(end-start) = 3

Answer = 2 + 0 + 0 + 3 = 5

*/


