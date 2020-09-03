### Initializing

```
- vagrant init bento/ubuntu-20.04
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
- ssh-copy-id
```

### Google SMTP Config

- Username : Gmail Id
- Password : Gmail Password
- PORT: 465
- SMTP Server: smtp.gmail.com
- email suffix : @gmail (or any other)

### UpStream and DownStream

##### Workspace

> \${WORKSPACE}

One of the default environment variables. We can share the same workspace among different jobs.

You can pass this as an env variable, in post build actions, in "Trigger Parameterized Build on Other Projects" and then use that in the next Project by setting custom workspace and parameters.

### Docker

- If installing docker on a node, disconnect and reconnect the node for jenkins to run it.
