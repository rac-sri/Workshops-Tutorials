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

- ansible group/ip -m yum -a "name=git state=latest" -b

- ansbile ip/group db -m setup // get server information  - default facts
```

Dynamic Inventory

```
- Download ec2.py and ec2.ini from official repo
- ansible -i ec2.py -m ping
```

Raw Module

```
ansible group/ip -m raw -a "scp "
```

Playbook:

```
- ansible-playbook <playbook>  // -v verbosity
```

Command Line Arg:

```
- ansible-playbook cmd_line.yml -e "x=56 y='ansible playbooks'"
- ansbile-playbook cmd_line.yml -e "{'x':[3,4,5,6]}"
- ansbile-playbook cmd_line.yml -e "{'x':{'one':1,'two':2}}"
- ansbile-playbook cmd_line.yml -e "@variable_values.yml"
- ansible-playbook install_uninstall_httpd.yml -e "pkg="nginx" req_state="present"
```

#### Notes

The order of precedence for vars in the current version of Ansible:

- Vars set on the command line -e foo=set_on_cmd_line
- Vars set in the vars_files: block in the play
- Vars set in the vars: block in the play
- Vars set in host_vars/
- Vars set in group_vars/
  Role default vars roles/.../defaults/main.yml
