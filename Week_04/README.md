学习笔记

# 广度优先
- 1. 一层一层遍历
- 2. 使用队列
- 3. 模板
``` go

func bfs(root *Node) {
    result := []
    queue:=[]*Node{root}
    for len(queue) >0 {
        levelSize:= len(queue) //当前层元素个数
        for i:=0;i<levelSize;i++{
            top:=queue[0] //头节点
            queue = queue[1:]//出队列
            for j:=0;j<top.childs;j++{
                queue = apppend(queue,top.childs[j])//加入子节点
            }
        }
    }
}

```