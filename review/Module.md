# Using Go Modules

[Using Go Modules](https://blog.golang.org/using-go-modules)

go.mod 文件定义了模块路径(也即用于根目录的导入路径)，

| 参数           | 值     |
| -------------- | ------ |
| Batch-size     | 16     |
| 迭代次数       | 60000  |
| 学习率         | 0.01   |
| 学习率策略     | Step   |
| Stepsize       | 15000  |
| 动量           | 0.9    |
| 权值衰减       | 0.0005 |
| clip_gradients | 40     |

| 方法           | HMDB51 | UCF101 |
| -------------- | ------ | ------ |
| DT[8]          | 55.9%  | 83.5%  |
| iDT[9]         | 57.2%  | 85.9%  |
| two-stream[16] | 59.4%  | 88.0%  |
| C3D[26]        | 56.8%  | 85.2%  |
| TSN[18]        | 68.5%  | 93.5%  |
| 本章方法       | 68.9%  | 94.6%  |

| Layer    | Input size   | Output size   | Kernels  |
| -------- | ------------ | ------------- | -------- |
| Convl    | $224*224*3$  | $224*24*512$  | $3*3*64$ |
| Max poll | $224*224*64$ | $112*112*512$ |
| Conv2_x  | $112*112*64$ | $56*56*256$   |
