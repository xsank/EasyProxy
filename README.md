# EasyProxy
A simple proxy server.


# Introduction
```
+----------+                                           +------------+
|  client  |<--------+                                 |            |
+----------+         +-------->+-----+<--------------->|   server   |
                               |     |                 |            |
+----------+                   |     |                 +------------+
|  client  |<----------------->|     |
+----------+                   |     |
                               |     |                 +------------+
+----------+                   |  L  |                 |            |
|  client  |^----------------->|  B  |<--------------->|   server   |
+----------+                   |  S  |                 |            |
                               |     |                 +------------+
+----------+                   |     |
|  client  |<----------------->|     |
+----------+                   |     |                 +------------+
                               |     |                 |            |
+----------+        +--------->+-----+<--------------->|   server   |
|  client  |<-------+                                  |            |
+----------+                                           +------------+

```  
As the chart above, easyproxy act as the LBS server. If you meet the situation that cannot access some
service directly or need a server as a gateway, you may use this to proxy.

# Usage

```bash
go get github.com/xsank/EasyProxy
go run build.go build
```
The sample config is about mysql proxy, so you can access the mysql database actually at backend
through the `localhost 9527`, you can easily modify the `default.json` to change the service.
Also you can `curl localhost:7259/statistic` to get the current connection status. If you want to know
more runtime debug info about the proxy, you can `curl http://localhost:7259/debug/pprof` to get more details.
For keep simple, easyproxy only support `poll`,`random`,`iphash` strategies, it would add more if there is a demand.

# License
Easyproxy is distributed under MIT License.
