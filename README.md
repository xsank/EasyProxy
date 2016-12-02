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
|  client  |<----------------->|  B  |<--------------->|   server   |
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
* The sample config is about mysql proxy, so you can access the mysql database actually at backend
through the `localhost 9527`, you can easily modify the `default.json` to change the service. Any protocol based on `TCP` would work well such as:mysql, ssh, svn...    
* For keep simple, easyproxy only support `poll`,`random`,`iphash` strategies, and the default strategy is `poll`,the easyproxy would add more if there is a demand.  
* Besides, you can `curl localhost:7259/statistic` to get the current connection status. If you want to know more runtime debug info about the proxy, you can `curl http://localhost:7259/debug/pprof` to get more details.  
* You can change the number of MAX processorsand and the heartbeat time based on your requirements.


# License
Easyproxy is distributed under MIT License.
