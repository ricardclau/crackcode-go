10.1 Basketball hoop
====================

Game 1: p -> win, 1 - p -> fail

Game 2: To win: Make the 3 shots or fail one -> p ^ 3 + 3 * p ^ 2 * (1 - p) -> 3 * p ^ 2 -  2 * p ^ 3

Choose Game 1 for 

p > 3 * p ^ 2 - 2 * p ^ 3 
1 > 3 * p - 2 * p ^ 2
2 * p ^ 2 - 3 * p + 1 > 0
(p - 1) * (2p - 1) > 0 && p < 1 -> p - 1 < 0, 2p - 1 < 0 -> p < 1 && 2p < 1

p < 1 / 2 -> play game 1!!!!

10.2 Ants walking
=================

For a triangle 
--------------

The only option to avoid collision is that all 3 ants start moving to the same direction, and this direction can have 2 possibilites
Solution: 1 - 2 * (1/2) ^ 3

For a polygon of N
------------------

Solution: 1 - 2 * (1/2) ^ N -> 1 - (1/2) ^ (N - 1)

10.3 Lines intersection
=======================

2 lines intersect on a cartesian plane unless they are 100% parallel
But if they are the same line exactly they are considered to intercept

So, if 2 lines are 

y = a1 * x + b1
y = a2 * x + b2

They don´t intersect if a1 = a2 && b1 != b2
They intersect if a1 != a2 || b1 != b2
