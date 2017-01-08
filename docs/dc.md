# Instructions for the Datacenter Cook

## Facebook's Data Center Fabric
Link to talk: https://www.youtube.com/watch?v=kcI3fGEait0

* A server pod is a unit of deployment in the datacenter.
* A server pod has 4 fabric switches and 48 rack switches.
* Pods are connected via spine planes (as forwarding capacity between the pods)
* If you need to scale compute capacity, you add server pods.
* If you need to scale forwarding performance, you add spine planes.
