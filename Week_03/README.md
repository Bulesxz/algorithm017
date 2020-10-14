学习笔记

# 回溯模板
``` go
func backtrace(路径,选择列表){
    //1. 满足条件
    if 满足条件 {
        //处理结果
        result.add(路径)
        return
    }

    for 选择i in 选择列表 {
        选择i
        backtrace(路径+i,选择列表)
        撤销i
    }
}
```