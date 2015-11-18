server {
	    listen       9999;
	    server_name  drive.wps.cn test-drive.wps.cn;

	    location / {
	        #前端项目路径
	        root /Users/xiaotian/Workspace/wps-drive/src/wps.cn/web_driver/dist/drive/;
    }

    # 静态资源不需转跳至https; 支付宝异步通知只使用http
	    if ($request_uri !~* "^/s/|^/api/payment/notify") {
	        rewrite ^(.*)$  https://$host$1 permanent;
	    }
		}

		server {
	    listen       443;
	    server_name  drive.wps.cn;


	    ssl on;
	    # 上述步骤中生成证书的路径
	    ssl_certificate     /usr/local/etc/nginx/csr/server.crt;
	    ssl_certificate_key /usr/local/etc/nginx/csr/server.key;
	    ssl_protocols       TLSv1.2 TLSv1.1 TLSv1;
	    ssl_ciphers         HIGH:!aNULL:!MD5;
	    ssl_session_cache   shared:SSL:500m;
	    ssl_session_timeout 10m;

	    proxy_next_upstream off;
	    proxy_set_header    X-Real-IP           $remote_addr;
	    proxy_set_header    X-Forwarded-For     $proxy_add_x_forwarded_for;
	    proxy_set_header    X-Forwarded-Proto   "https";
	    proxy_set_header    Host                $host;
	    proxy_http_version  1.1;
	    proxy_set_header    Connection  "";

	    location / {
	        #前端项目路径
	        root /Users/xiaotian/Workspace/wps-drive/src/wps.cn/web_driver/dist/drive/;
	        try_files $uri $uri.html $uri/ =404;
	    }
	    location /usercenter/{
	        alias /Users/xiaotian/Workspace/wps-drive/src/wps.cn/web_driver/dist/account/;
	        index usercenter.html;
	    }
	    #############################文件下载相关，兼容ie###########################
	    location /api/download/ {
	        add_header Cache-Control "no-store, no-cache, must-revalidate";
	       # proxy_pass http://store_backend;
	    }
	    #############################文件下载相关，兼容ie###########################

	    location /api {
	        proxy_pass http://qing.wps.cn/api;
	    }
	    
	    location /preview {
	        #preview 服务ip
	        proxy_pass https://10.20.187.98;
	    }

	    location /p {
	        proxy_pass https://drive.wps.cn;
	    }
		}

    }