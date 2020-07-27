# Plan

## docker-compose 的使用

目前使用 --link 的方式让 golang 容器访问 redis 和 mysql 容器, 无法在本地进行调试, 需要考虑使用其他的方式