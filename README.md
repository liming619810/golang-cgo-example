# golang-cgo-example
golang. use cgo example 

## 目录

### test-so 
```
c 语言代码，编译生成 test_so.so 
```

### test-go 
```
golang 语言代码，通过cgo 调用 test_so.so 

需要先定义一个 loac_so.c 来加载 test_so.so
再 通过cgao，再golang的main 中调用/Users/liming/workspace/aland/code/business-library-service/gateway/.gitignore

```

### purego 
```
 golang purego 项目 调用 test_so.so 
 需要先注册，才能调用
```