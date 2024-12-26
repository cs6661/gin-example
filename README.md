# gin-example
gin脚手架

鉴权: jwt  
orm： Gorm  
监控： Prometheus  
日志： zap  
数据库： Mysql  
缓存： Redis  

大数据量可 使用go接受到channel中 再将数据批量写入kafka(高吞吐量,可接受延迟) 减少网络交互 再将kafka中的数据定时备份到对象存储和ClickHouse
计算框架的选型Spark/Flink/Hadoop 或者直接写入带有聚合计算功能的olap数据库

kafka/rabbitMQ/rocketMQ的选型(成本/数据)

| 特性/系统         | Kafka                              | RabbitMQ                           | RocketMQ                           |
|------------------|------------------------------------|------------------------------------|------------------------------------|
| **消息传输模型**   | 发布/订阅、队列                    | 发布/订阅、队列                    | 发布/订阅、队列、顺序消息、事务消息 |
| **吞吐量**         | 高吞吐量                            | 适中，适合低到中等吞吐量场景        | 高吞吐量                            |
| **延迟**           | 相对较高                            | 低延迟，适合实时性较强的场景        | 低延迟，适合高吞吐量场景           |
| **存储模型**       | 磁盘存储，持久化，消息可重放         | 基于内存，消息持久化               | 消息持久化，支持顺序存储             |
| **事务支持**       | 不支持事务                          | 支持事务                            | 支持事务消息                        |
| **集群和扩展性**   | 分布式架构，水平扩展性好              | 集群模式，水平扩展性差              | 分布式架构，支持水平扩展             |
| **协议支持**       | 自定义协议                          | 支持 AMQP、STOMP、MQTT 等协议      | 支持 JMS、REST、OpenMessaging 等协议 |
| **场景**           | 高吞吐量流式处理、大数据分析         | 任务队列、微服务、低延迟消息传递     | 高吞吐量、高可用、分布式事务处理    |
| **运维难度**       | 中等，集群管理相对复杂                | 较低，易于使用和管理                | 较高，运维和配置需要精细化管理       |


kafka数据同步方案  
Kafka 消费者：通过 Kafka 消费者消费 Kafka 主题中的消息。  
对象存储 SDK：使用云服务提供的 SDK 将数据上传到对象存储（例如 AWS S3、阿里云 OSS、Azure Blob Storage 等）。  
批量上传：将多个消息聚合在一起并批量上传，减少网络开销和存储操作。  
流式上传：直接将数据写入对象存储，实时将消息同步到云存储。  
官方提供的工具Kafka Connect  