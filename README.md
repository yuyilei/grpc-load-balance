# grpc-load-balance

This document displays several examples of gRPC Load Balancing. 

## Nginx 

[Nginx 1.13.10](https://www.nginx.com/blog/nginx-1-13-10-grpc/) adds new feature to support gRPC connections. With simple configuration, we can use Nginx for reverse proxy to load balance a cluster of gRPC services. 

To test it out, I have a simple gRPC Hello service that will return message to declare the gRPC server's name. Then, I deploy 20 gRPC server instances behind Nginx server, and use the random load balance policy. 

And I have a gRPC client that will send many gRPC requests concurrently, collect all gRPC responses and add up all the gRPC server's name. As a result, the gRPC cluster can handle about **800** concurrent gRPC requests at most. As we can see, the gRPC requests are randomly dispersed to 20 gRPC server instances.

![](https://github.com/yuyilei/grpc-load-balance/blob/master/pictrue/ng.jpg?raw=true) 

## Load balance in Kubernetes

### L4 Load Balancing

Kubernetes supports load balancing in Layer-4: Transport Layer. 

gRPC is built on HTTP/2, and HTTP/2 is designed to have a single long-lived TCP connection, thus, Once the gRPC connection is established, there’s no more balancing to be done. All requests will get pinned to a single destination pod.

To sum up, gPRC connections can not be automatically balanced in Kubernetes in Layer-4.

### Istio

Istio is a Service Mesh that essentially deploys an Envoy proxy per microservice instance. Istio automatically intercepts the requests and forward the request to a sidecar proxy (i.e., there is a proxy instance running along side of every microservice instance). The proxy can automatically discover backend instances and perform L7 load balancing. 

#### Prerequisites
I deploy a Kubernetes cluster with only one node(as master and work node as well), then I install Istio and bootstrap the Kubernetes cluster with [Istio](https://istio.io/docs/setup/install/kubernetes/). 
And to make sidecars can be automatically added to applicable Kubernetes pods, I enable the [automatic sidecar injection](https://istio.io/docs/setup/additional-setup/sidecar-injection/#automatic-sidecar-injection). 

#### Deployment 

##### configuration file:

[server configuration](https://github.com/yuyilei/grpc-load-balance/blob/master/kubernetes/istio/raw_server.yaml)  --> [server with sidecar](https://github.com/yuyilei/grpc-load-balance/blob/master/kubernetes/istio/server.yaml)

[client configuration](https://github.com/yuyilei/grpc-load-balance/blob/master/kubernetes/istio/client.yaml)

I deploy the gRPC server with configuration that declare there are 20 gRPC server instances(replica) and each pod have a sidecar proxy. 

Then I make the same gRPC client to test the cluster, and the result is it can handle **500** gRPC requests concurrently. 



 