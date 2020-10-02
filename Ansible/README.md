```
- ansible all -m ping
- ansible <ip>:<secondip> -m ping   // ip,secondip should be in host file
- ansbile group1 -m ping
- ansible -i <invetary file>  // max priority in inventary file selection
```

Disable Host Checking:

- export ANSIBLE_HOST_KEY_CHECKING=False
-

```
[defaults]
host_key_checking = False
```

Ad-hoc

```
ansible [-i <inventary file>] servername:group:group2 -m module [-1 argument_value]
```

```
- ansible <gp>/ip -m shell -a "uptime"
- ansible-doc -l // list all modules
```
