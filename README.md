# EasyProxy
A simple proxy server.


# Introduction
```
+----------+                           +------------+
|  client  +<----+                     |            |
+----------+     +--->+-----+<-------->+   server   |
                      |     |          |            |
+----------+          |     |          +------------+
|  client  +<-------->|     |
+----------+          |     |
                      |     |          +------------+
+----------+          |  L  |          |            |
|  client  +<-------->|  B  |<-------->+   server   |
+----------+          |  S  |          |            |
                      |     |          +------------+
+----------+          |     |
|  client  +<-------->|     |
+----------+          |     |          +------------+
                      |     |          |            |
+----------+     +--->+-----+<-------->+   server   |
|  client  +<----+                     |            |
+----------+                           +------------+            
```  
As the chart above, easyproxy act as the LBS server. If you meet the situation that cannot access some
service directly, you may use this to proxy.  

# Usage

```bash
go get github.com/xsank/EasyProxy
go run build.go build
```
The sample config is about mysql proxy, you can easily modify the `default.json` to change the service.  
And you can `curl localhost:7259/statistic` to get the current connection status.


# License
Easyproxy is distributed under MIT License.
