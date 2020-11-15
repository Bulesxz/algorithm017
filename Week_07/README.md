学习笔记


# 双向BFS 

- 开始 和 结束 入队列
- 循环出队列
- 取出现有队列所有元素 判断 是否满足结束条件  ，然后依此扩展
- 层级+1
- 比较两个队列长度 ，优先使用短的 然后交换

queue = []
queue.append(begin)

queue2 = []
queue2.append(end)

level = 0

while queue:
    for  top range queue:
        if top in queue2:
            return level
        for i in top:
            queue.append(i)
    if len(queue)  > len(queue2)
        queue ,queue2 = queue2,queue
    level++