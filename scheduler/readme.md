# Scheduler
调度器,用于执行异步，周期性，不会立刻又结果的任务，此处用于异步删除视频为例
## Scheduler Server
接收任务用，用计时器触发，利用消费者/生产者模型
## 任务下发过程
* 1 user-> api serice -> delete video
* 2 api service -> scheduler -> write video deletion record
* 3 timer
* 4 timer -> runner ->read write->exec->de.ete