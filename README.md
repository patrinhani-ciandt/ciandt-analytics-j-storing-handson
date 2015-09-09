#CI&T Analytics Journey - Big Table Hands On

## Initializing the container

    for Windows:
    - Create a folder called "DockerHome" on  c:\Users\[username]
    - Run the following command on your "Boo2Docker Shell"
    $ docker run --rm -t -i -e WHOAMI=$(whoami) --name=analytics-j-img-test -v //c/Users/[username]/DockerHome:/home  patrinhani/ciandt-analytics-j-storing-handson

    for Linux / MacOS:
    - Create a folder called "DockerHome" on  ~/
    - Run the following command on your Console if you are into a Linux distribution or "Boo2Docker Shell" if you are on MacOS.
    $ docker run --rm -t -i -e WHOAMI=$(whoami) --name=analytics-j-img-test -v ~/DockerHome:/home  patrinhani/ciandt-analytics-j-storing-handson
