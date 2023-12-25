Implement Priority Queue using Linked Lists. 

push(): This function is used to insert a new data into the queue.
pop(): This function removes the element with the highest priority from the queue.
peek() / top(): This function is used to get the highest priority element in the queue without removing it from the queue.
Priority Queues can be implemented using common data structures like arrays, linked-lists, heaps and binary trees.

Prerequisites : 
Linked Lists, Priority Queues

The list is so created so that the highest priority element is always at the head of the list. The list is arranged in descending order of elements based on their priority. This allow us to remove the highest priority element in O(1) time. To insert an element we must traverse the list and find the proper position to insert the node so that the overall order of the priority queue is maintained. This makes the push() operation takes O(N) time. The pop() and peek() operations are performed in constant time.

