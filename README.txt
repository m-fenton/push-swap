push-swap

This project uses a Non-Comparative Sorting Algorithm. You have at your disposal a list of int values, two stacks (a and b) and a set of instructions.


This project calculates and displays on the standard output the smallest program using push-swap instruction language that sorts integer arguments received.
The program displays the smallest list of instructions possible to sort the stack a, with the smallest number being at the top.
Instructions must be separated by a newline \n and nothing else.
The goal is to sort the stack with the minimum possible number of operations.
In case of error, you must display Error followed by a newline \n on the standard error. 
Errors are understood as: some arguments are not integers and/or there are duplicates.
If there are no arguments, the program does not display anything (0 instructions).

Example - 
go run ./push-swap "2 1 3 6 5 8"
pb
pb
ra
sa
rrr
pa
pa
1 2 3 5 6 8 

go run ./push-swap "0 one 2 3"
Error

go run ./push-swap


As said before, you will have two stacks at your disposal. 
The goal is to sort the stack a, that will contain the int values received, in ascending order, using both stacks and a set of instructions.

These are the instructions that you can use to sort the stack :

pa push the top first element of stack b to stack a
pb push the top first element of stack a to stack b
sa swap first 2 elements of stack a
sb swap first 2 elements of stack b
ss execute sa and sb
ra rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
rb rotate stack b
rr execute ra and rb
rra reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
rrb reverse rotate b
rrr execute rra and rrb


Created by 

Martin Fenton
Rupert Cheetham
Nikoi Kwasie