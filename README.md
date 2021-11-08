# sblog

### 概述

* 后端基于golang fiber 参考文档 https://docs.gofiber.io/
* 后台的前端基于element-ui 参考文档 https://element.eleme.cn/#/zh-CN/component/installation
* 外层是后端部分，webvue目录是后台的前端element-ui部分
* vue对seo不太友好，前台部分用的是模板渲染
* golang
    ```shell
    go build -o sblog ./cmd/sblog/main.go
    ./sblog
    ```
* yarn
    ```shell
    cd webvue
    # 引入第三方库
    yarn add xxx
    # 引入本项目所需要的库
    yarn install
    # 运行
    yarn serve
    # 打包
    yarn build
    ```

### 部署

  * golang nginx部分
    ```shell
      # 需要代理一下
      location / {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;
        proxy_pass http://127.0.0.1:2021;
      }
    
    #websocket相关备份
    #location /websocket {
    #    proxy_pass http://0.0.0.0:6666;
    #    proxy_http_version 1.1;
    #    proxy_set_header Upgrade $http_upgrade;
    #    proxy_set_header Connection "upgrade";
    #}
    ```
  * element-ui nginx部分
    ```shell
    location / {
      index index.html index.htm;
      try_files $uri $uri/ @router;
    }
      
    location @router{
      rewrite ^.*$ /index.html last;
    }
      
    # 如果vue部分开启了代理则nginx需要代理一下
    location /api/{
      rewrite ^/b/(.*)$ /$1 break;
      proxy_pass http://www.xxxxx.com/;
    }
    ```