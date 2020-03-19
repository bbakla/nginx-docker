# Docker
name of the docker image: docker_learn


#Docker compose
* I had a big problem with reaching to the web application from nginx. The reason was port mapping.

    In _docker-compose.yml_ I used the port mapping of `8098:8099`. That means when i want to 
    reach to this service, I should call `localhost:8098`. Therefore, in my _default.conf_ 
    file I used the upstream server address as `localServerAlias:8098;` **That was the problem!**
    
    I actually run nginx and the wep application in the same network. In that case,
    port exposed to the external world, `8098` in my case, should not be used. Changing the 
    configuration in _default.conf_ to    `localServerAlias:8099;` solved the issue. Thanks to [1]

* if you want to create your own interface, take a look to [to that link][2]

[1]: https://forums.docker.com/t/connection-refused-from-upstream/49229
[2]: https://forums.docker.com/t/accessing-host-machine-from-within-docker-container/14248