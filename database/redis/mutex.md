# 分布式锁

分布式锁实现要点

- 互斥。任意时刻，只能有一个客户端能持有锁
- 不会出现死锁。持有锁的客户端crash之后，不会导致整个系统死锁
- 容错性。只要redis大部分节点能工作，则分布式锁就能正常工作
- 解铃还需系铃人。解锁操作只能由加锁的人来执行