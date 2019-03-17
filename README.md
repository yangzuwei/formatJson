# formatJson
调节一下节奏，写个小工具

# 说明

公司的文档中直接都是给出 `json` 样子的输入和输出的，大概像这样
```
{
    "name":"张三",//姓名
    "age":17,//年龄
}

```
有时候为了拿样例去接口试试，然后就要去掉注释，并且要规范一下末尾的那个逗号，如果遇到几百行的那种实在手速堪忧，而且还容易错，于是就写了这个命令行工具。

# 用法

把文档中的内容拷贝到一个文件里备用。
把该程序编译成自己想要的名称，例如 `fj`,然后执行 `./fj 文件名`，就会得到打印屏幕的信息，和一个新的 json 文件（以json+原文件名）

:runner:
