# Notes from Articles, Papers, and Talks

## Building a Billion User Load Balancer Link:
https://www.youtube.com/watch?v=MKgJeqF1DHw

* All requests coming into Facebook outside of DNS are TCP.  IP header has
* source and destination address.  TCP header has source and destination port.
* When you send an HTTP request, that request is encapsulated within a TCP
segment, and that segment is encapsulated within a IP packet.
* Facebook uses proxygen as its layer 7 load-balancer.  They use shiv as their
* layer 4 load-balancer which sits in front of many
proxygen instances.
* When request reaches Facebook datacenter, shiv will consistently hash the
request to one of the proxygen instances. It consistently hashes src ip/port and
dst ip/port.
* Then, there's another load-balancer in front of that, which is actually
a top-of-rack switch.
* A front-end cluster at Facebook may have 10s of shiv instances (IPVS), 100s of
proxygens (terminating TCP and SSL) and 1000s of HHVM web servers.
* A datacenter has a bunch of front-end clusters.  A l4 load-balancer could live
* in the same rack as the l7 load-balancer.  They don't pin a user to a
* particular instance of an HHVM web server.  L4 load-balancers get ECMP'd from
* the TOR switch.  ExaBGP (a Python daemon) is used in (l4) load-balancers to
* peer with TOR
switch in order for all load-balancers to have the same IP address.
* ECMP hash is used in between TOR switch and l4 load-balancers and then other
hashing scheme is used between l4 and l7 load-balancers.
* They do keep state in the l4 load-balancer.  IPVS is a kernel module that is
* doing the l4 load-balancing.  Shiv is a Python thing on top of IPVS that is
* doing health-checking.  Packets do not go back through l4 load-balancer layer
* when direct server
return (DSR) is being used.
* (23:11) DSR works by: packet hits l4 load-balancer (shiv) which puts another
* IP
packet around it -- destination is an instance of l7 load-balancer (proxygen),
proxygen sees a wrapped IP packet within the packet shiv added. proxygen drops
that packet on the networking stack which is a tunnel interface that proxygen is
listening on. when request is forwarded to HHVM web server, the response back
from HHVM to proxygen will be sent back directly to the client.
* There are 1 or 2 shivs (l4) per rack.

### Edge POPs and Reducing Latency
* Found that there's 75ms latency between Oregon DC and Seoul... therefore 150ms
* just to
establish a TCP connection. 5-way SSL handshake adds another 300ms
* They put an edge POP in Tokyo to reduce initial latency.  TCP and SSL to Tokyo
* only took only 90ms. Rest of HTTP request still had to traverse Tokyo
to Oregon.
* (26:43) They have pre-established (aka persisjtent) connections between edge
* and facebook There is dedicated transit between Tokyo and Oregon.  Edge POP
* locations: See ./images/building3.tiff.  Only difference between what's
* deployed in the POP and the datacenter is that their is no
HHVM web servers deployed in the POP. In the POP, there are only shivs and
proxygens.

### Links
* https://github.com/Exa-Networks/exabgp
* http://www.linuxvirtualserver.org/software/ipvs.html

## Filing a Good Bug Report
When filing a bug report make sure to include the following details: version
numbers or git SHAs, configuration values, logs, sample code to reproduce,
problem description.

## HTTP/2 and http2 in Go 1.6 Link:
https://www.youtube.com/watch?v=FARQMJndUn0#t=0m0s

* HTTP1.0 and 1.1 only allowed you to do one thing at a time, setting up a new
TCP connection is slow, no way to abort request without closing connection, no
way for servers to gracefully shutdown.
* HTTP 1.1 does have pipelining, but most browsers have it disabled because
proxies break it. Proxies don't usually speak full, compliant HTTP.
* HTTP hasn't changed much between 1999-2013. Pipelining was a fail. So
eventually browsers just started to allow more connections to be opened up (e.g.
2 or 6)
* People started doing hacks with 6 connections per hostname by using hostnames
e.g., tiles{1,2,N}.googlemaps.com
* More hacks: (image) spriting, concatenating css/js together.  So in 2009,
* Google came up with SPDY, which became starting point of HTTP/2 Every HTTP2
* request has a: 9 byte header which includes frame length, frame
type, frame flags, and stream ID (aka a unique request ID). Then, finally, you
have the payload.
* (5:40) http2 clients/servers can exchange frames in any order.  Can have
* multiple requests outstanding Server responses can be interleaved There are
* different types of frames: settings, ping (aka heartbeat), headers,
data, goaway (graceful shutdown), rst_stream (abort request), window_update (for
flow control), priority, push_promise
* hpack is a new compression format that was built for http2
* http2 is a binary protocol
* http2 "just works". It's on by default.

## A Technical Overview of Kubernetes Link:
https://www.youtube.com/watch?v=WwBdNXt6wO4

* Kubernetes is about decoupling components that are ordinarily tightly-coupled.
* The new stack: Application Ops (microservices), Cluster Ops (Kubernetes),
Kernel/OS Ops (CoreOS), Hardware Ops (IaaS: public clouds)
* You interact with Kubernetes over its API server via kubectl,
programmatically via REST, etc:
* Kubernetes is backed by etcd.  Kubernetes runs two other binaries (other than
* the API server): scheduler and
a controller manager.
* Scheduler is responsible for scheduling containers out onto machines or pods.
* Controller manager is responsible for doing health maintenance.
* Self-healingness of Kubernetes is about what comes after deployment and
orchestration of getting an app deployed
* Kubernetes is intended to be an online system that keeps your apps alive Is
* responsible for "creating the world," but also healing and maintaining the
world.
* On each machine is a kubelet (a small daemon for managing the controllers
themselves) and the service proxy provides load-balancing
* Pods are the atomic unit of a Kubernetes cluster. They are 3 things: 1) They
are a collection of one or more containers that work well together. 2) They are
a set of data volumes. 3) They are a set of namespaces that all of these
containers share (e.g. IPC, network, etc). The namespaces are shared by all
containers within the pod. 4) They have labels. Every object inside Kubernetes
has labels. Labels are just key/values pairs. Labels provide a way to
slice/query across the state of running things.
* The lifespan of a data volume is independent of the lifespan of a container.
* Anything in the container itself should be considered transient. Anything that
you want persisted probably should be network attached storage (NAS).
* Pod is the atomic unit of scheduling where it doesn't make sense for
"these two containers" to land on multiple machines. Pods and containers offer
us the ability to separate infrastructure/deployment concerns.
* Different containers a pod can all see each other on `localhost` because they
share the same network namespace.
* Kubernetes employs a reconciliation pattern (aka goal-orientedness) whereby
there is a desired/goal state and the actual state. Kubernetes' reconciliation
loop's responsibility is to figure out the difference between the two states and
take action to get the world into the desired state.
* (17:32) A replication controller is a combination of a template (of the
* desired
state) and labels that are used to query out to get the current state of the
world and reconcile the differences.
* The concept of a service is a logical grouping like frontend, backend,
* database, etc.  Service have a known fixed IP address. They have a DNS entry
* associated with that
IP address. The IP address is actually a fake address. It doesn't actually exist
other than in the routing tables of the machines in the cluster.
