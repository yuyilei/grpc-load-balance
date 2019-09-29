# grpc-load-balance

This document displays several examples of gRPC Load Balancing. 

## Nginx 

[Nginx 1.13.10](https://www.nginx.com/blog/nginx-1-13-10-grpc/) adds new feature to support gRPC connections. With simple configuration, 
we can use Nginx for reverse proxy to load balance a cluster of gRPC services. 

To test it out, I have a simple gRPC Hello service that will return message to declare the gRPC server's name. Then, I deploy 20 gRPC server instances behind Nginx server, 
and use the random load balance policy. And I have a gRPC client that will send many gRPC requests concurrently, collect all gRPC responses and add up all the gRPC server's name. 

As a result, the gRPC cluster can handle about 800 concurrent gRPC requests at most. 

![](https://github.com/yuyilei/grpc-load-balance/pictrue/ng.jpg) 

