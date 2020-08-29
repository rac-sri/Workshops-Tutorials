### Initializing

```
- vargrant init bento/ubuntu-20.04
- vagrant up jenkins_master
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
- sudo apt-get install openjdk-8-jdk
- ssh-keygen        // on master node
- ssh-copy-
```

### Google SMTP Config

- Username : Gmail Id
- Password : Gmail Password
- PORT: 465
- SMTP Server: smtp.gmail.com
- email suffix : @gmail (or any other)
