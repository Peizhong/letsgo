# 需求
socket.io，集群模式，双人通话(同一个roomId)，数据量较大，希望两个人能连同一个pod/进程。否则连到不同的pod，消息通过redis转发会有较大延迟
直接访问k8s的service，不能保证同一个房间的人进入到pod
用户A先连接到一个服务器，并将房间号通过app推送给用户B，用户B通过房间号，连接同一个pod

通话服务数据量：
单个连接
进：256B*30/s
出：256B*30/s
共：15KB/s

2MB：136个连接
10MB：682个连接
单个代理：500个连接，需要8MB左右

# 设计
增加一个网关服务，定时查询通话service的endpoints，进行代理// 希望无状态，可以横向扩展
接收http/ws请求，根据请求是否带参数room，判断使用一般负载规则 或 根据room指定pod
与客户端建立socket连接，并建立pod的socket连接
外部缓存，保存room和podid的关系
内部缓存，room的连接都断开后，删除外部缓存对应的数据
500个线程
api服务，查询负载状态