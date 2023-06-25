# go加载动态链接库

## 测试C代码
```c
#include <stdio.h>
#include <stdlib.h>
void fun1()
{
    printf("fun1\n");
}
void fun2()
{
    exit(0);
    printf("fun2\n");
}
```
## 编译成库
```
gcc -shared -fPIC -o libtest.so libtest.c
```

## 编译
```
go build
```