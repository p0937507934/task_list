# Task List
用 gin + gorm 簡單的 task api
採用 clean archtecture 架構, 使用 wire 當依賴注入

使用方式
```
        git clone repo
        <!-- 啟動服務 -->
        sudo docker-compose up
        <!-- 測試 -->
        curl localhost:8000/tasks
```
or
```
        git clone repo
        <!-- get required package -->
        go mod tidy
        <!-- 執行 -->
        go run main.go
        <!-- 單元測試 -->
        go test ./...

```

# Swagger
啟動服務後, http://localhost:8000/swagger/index.html#/

# API

GET 獲取所有(此處看規格有加s
/tasks

POST 新增
/task

PUT 修改
/task/:id

DELETE 刪除
/task/:id

# 專案目錄
#### api - 放 api handler
#### api_test - api handler test
#### config - 初始化配置
#### docs - swagger file
#### driver - 初始化 db,logger 等
#### dto - 資料傳輸層定義(放API_Request 及 API_Response)
#### internal - Service 實作 業務邏輯, Repository 實作與 db 交互
#### middlware - 中間件
#### migration - init database
#### models - 定義 db schema 及用來操作 db 交互的 struct
#### pkg - helper 工具
#### route - api 路由
#### wire - 實現依賴注入


# 測試連結
```
<!-- 服務跑在 EC2 -->

curl 35.175.150.170:8000/tasks
```