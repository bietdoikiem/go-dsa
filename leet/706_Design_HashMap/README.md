<p align="center"><img width="350" src="https://i.ibb.co/4dyGXwW/image.png" /></p>

# No. 706 - Design HashMap

## Core Concept

**Precaution**:

> In this problem, the Hash Map is going to be designed for the unoptimally in
> order for a workable and valid submission. For optimal solution, you can check other
> solutions in the LeetCode's discussion section. Alternatively, you can check for
> the implemenetation of Hash Map in Java, in which talented engineers behind Java
> has implemented a scalable HashMap using Re-hashing technique based upon Load
> Factor.

The Hash Map is implemented with an underneath array. Since the key is
defined as integer (non-negative) in this problem, we're going to use modulo
operator to calculate the index based on the given key in order to prevent out
of bounds situation.

Moreover, the array's size would not be as big as the number of elements in the
map because it would waste a lot of unused resource upon declaration. For that
reason, a small/appropriate array's size is used here.

To prevent collision in the hashed integer key of a limited array's size, a linked list is used
as an element in the array. Each of the node in the list contains key and value
for retrieval purpose. However, there's one drawback in using this approach is
that if the number of collided items are high, the constant look-up time of the
Hash Map is no longer preserved.

<p align="center"><img width="350" src="https://i.ibb.co/rGSLcCJ/Screenshot-20220213-005309-Samsung-Notes.jpg" /></p>

## Pseudocode

(Check out the source code, since the implementation is quite simple xD)
