### simple-concurrency
> 一个简单的Go并发管理封装

主要的结构如下：
<img src="framework.svg" width=500>

`dispatcher` 接收到`Job`数据，将对应的Job发送到自身的一个`worker`的`Job`队列中，`worker`从队列中取出`Job`进行处理

使用方式参考 [example](./examples)