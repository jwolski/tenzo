# Instructions for the Docker Cook

## What is Docker?
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

# How long have cgroups and namespaces been around?

# What makes Docker a "containerization platform"?

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

# What is a Docker image?
An image is a filesystem and parameters to use at runtime. A container is a
running instance of an image.

You can list images with:

```
$ docker images
```

# References
1. https://www.docker.com/
2. https://en.wikipedia.org/wiki/LXC