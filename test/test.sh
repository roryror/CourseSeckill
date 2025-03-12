#!/bin/bash

echo "start test"

# 预热缓存
curl -X POST "http://localhost:8080/warmup/" -H "Content-Type: application/json"

# 创建一个空的post_data.txt文件
touch post_data.txt

# 使用ab工具进行压力测试
ab -n 10000 -c 100 -k -r -p post_data.txt -T 'application/x-www-form-urlencoded' "http://localhost:8080/seckill/0/0"

echo "test done"