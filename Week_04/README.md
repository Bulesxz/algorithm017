学习笔记

# 广度优先
- 1. 一层一层遍历
- 2. 使用队列
- 3. 多叉树模板
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
- 4. 对于图，需要增加 visit 访问集合去重、


# 深度优先
- 1. 不停的遍历子节点
- 2. 递归法
- 3. 多叉树模板
``` go
func dfs(root *Node){
    visit(root)
    for j:=0;j<root.childs;j++{
        dfs(root.childs[j])
    }
}
```
- 4. 对于图，需要增加 visit 访问集合去重