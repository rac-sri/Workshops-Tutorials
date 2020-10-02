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
- ANSIBLE_KEEP_REMOTE_FILES <above command>
- ansible-doc -l // list all modules
- in config: fork=number of servers to run in parallel
- ansbile group/ip -m copy -a "src=.. dest=..."
- ansbile group/ip -m copy -a "content=...  dest=.../content.txt backup=yes"
- ansible group/ip -m fetch -a "src=.. dest=.."   // download a file
- - ansible group/ip -m fetch -a "src=.. dest=./newdemo/{{inventory_hostname}} flat=yes"   // download a file if the same file has diff values


// imp

- ansible group/ip -m file -a "path=... state=touch" // create file (touch->mkdir for dir) -b or -k // to get elevated priviledges
```
