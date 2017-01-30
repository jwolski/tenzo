# What is Docker?
Docker is the world's leading software containerization platform. [1]

# What's a (Linux) container?
Linux Containers (abbreviated LXC) "is an operating-system-level virtualization
method for running multiple isolated Linux systems (containers) on a control
host using a single Linux kernel." It goes on to say, "the Linux kernel provides
the cgroups functionality that allows limitation and prioritization of resources
(CPU, memory, block I/O, network, etc), without the need for starting any
virtual machines." They also provide "namespace isolation functionality that
allows complete isolation of an applicatons' view of the operating environment,
including process trees, networking, user IDs and mounted file systems."
Finally, "LXC combines the kernel's cgroups and support for isolated namespaces
to provide an isolated environment for applications." [2]

# Why are limitation, prioritization and isolation "nice to have" properties?

# What makes virtual machines "bad"?
* Containers are lightweight. Because they do not require the extra load of
a hypervisor, you can run more containers on a given hardware combination than
if you were using virtual machines. [3]

# What is a hypervisor?

# How long have cgroups and namespaces been around?
Initial introduction of cgroups came in 2007. cgroups came from the work
originally done by engineers at Google, originally called "process containers."
cgroups were merged into the mainline kernel in version 2.6.24. The cgroups
implementation merged into the kernel at this time are now referred to v1. This
implementation was eventually rewritten. The current version is referred to as
v2.

cgroups provide:
* resource limiting: e.g., groups cannot exceed memory limit.
* prioritization: assigning shares (larger/smaller) of CPU utilization or I/O.
* accounting: measuring group's resource usage
* control: freezing, check pointing, and restarting.

# What makes Docker a "containerization platform"?
* Docker provides tooling and a platform to manage the lifecycle of your
containers: encapsulating applications into containers, distributing/shipping
containers to your team for development and testing, deploying applications to
your production environment (in private or public cloud). [3]

# What are the Docker tools?
* Docker Engine
* Docker Compose - allows you to define an application's components, e.g.
containers, configuration, links and volumes, in a single file.
* Docker Hub - hosts public Docker images
* Docker Machine - sets up hosts for Docker Engines on your computer, on cloud
providers, and/or in your data center.
* Docker Swarm - pools Docker Engines together and allows you to treat them as a
single virtual Docker Engine. It supports the standard Docker API thus allowing
you to control multiple hosts as one.

# What is Docker Engine?
It is a client-server application with the following components: a server
daemon, a REST API, and a CLI. The CLI is used to talk to the server daemon
through the REST API. The daemon creates and manages Docker objects like:
images, containers, networks, and data volumes.

## Where does the server daemon run?

# What is a Docker image?
An image is a filesystem and parameters to use at runtime. A container is a
running instance of an image.

You can list images with:

```
$ docker images
```

# How is Docker implemented?
It is implemented using namespaces, control groups, and union file systems.

## What is a union file system?
Union file systems, or UnionFS, are file systems that operate by creating
layers, making them very lightweight and fast. [3]

# What is the Docker architecture?
Docker uses a client-server architecture. The CLI talks to the daemon using the
REST API over Unix domain sockets or a network interface.

## Is Docker Swarm just a set of interconnected Docker server daemons?

# dockerd
Dockerd is the persistent process that manages containers. It is the Docker "server
daemon." [4]

# What are Docker's competitors?
* Rkt, runC

# References
1. https://www.docker.com/
2. https://en.wikipedia.org/wiki/LXC
3. https://docs.docker.com/engine/understanding-docker/
4. https://docs.docker.com/engine/reference/commandline/dockerd/
5. https://en.wikipedia.org/wiki/UnionFS
