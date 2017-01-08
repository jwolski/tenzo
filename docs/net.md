# Instructions for the Networking Cook

## Tech
Facebook Cartographer

## What is Direct Server Return?
* The load-balancer routes packets without changing anything in it but the
destination MAC address. The backends process requests and answer directly
to the clients, without passing through the load-balancer.
- http://blog.haproxy.com/2011/07/29/layer-4-load-balancing-direct-server-return-mode/
