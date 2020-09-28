#!/usr/bin/env python

class Stack(object):
    def __init__(self):
        self.stack = []
    def push(self ,item):
        self.stack.append(item)
    
    def pop(self):
        return self.stack.pop()
    def top(self):
        return self.stack[len(self.stack)-1]
    
    def is_empty(self):
        if len(self.stack) == 0:
            return True
        else:
            return False

if __name__ == "__main__":
    stack = Stack()
    stack.push(1)
    print (stack.pop())
    stack.push(2)
    print (stack.is_empty())
    stack.push(3)
    print (stack.pop())
    print (stack.pop())
    print (stack.is_empty())