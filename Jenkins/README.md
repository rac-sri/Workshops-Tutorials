### Initializing

```
- vargrant init bento/ubuntu-20.04
- vagrant up vagrant up jenkins_master
- vagrant ssh
- Install Jenkins following instructions ( + install java jdk seprately if needed)
- sudo usermod -aG sudo jenkins
- /etc/init.d/jenkins start
```

### Adding a build Node

```
- vagrant up jenkins_node1
- vagrant ssh jenkings
- sudo adduser --gecos "" jenkins
- sudo usermod -aG sudo jenkins
- sudo su jenkins
- ssh-keygen        // on master node
- ssh-copy-
```
