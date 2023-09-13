# [Access Control (AC)](https://en.wikipedia.org/wiki/Access_control)

What can you do in an information system (IS)?
---
- [FIPS 201-3 definition](https://csrc.nist.gov/glossary/term/access_control)
  - *granting or denying specific requests* to 
    - obtain and use information and related services 
    - enter specific physical facilities
- [RFC 4949 definition](https://www.rfc-editor.org/rfc/rfc4949)
  - regulates use of system resources according to a security policy 
  - permits only to authorized entities


## Access Control Principles

🖊️ Practice
---
- Computer security is maily about access control
- Explore [TABLE D-1: MAPPING ACCESS CONTROL REQUIREMENTS TO CONTROLS](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-171r2.pdf) in NIST SP 800-171 for
  - Basic AC security requirements
  - Derived AC security requirements


Access Control Context
---
```mermaid
flowchart LR
  SA(Security Administrator);
  U(User)
  R[IS resources]
  AC[Access Control]
  AT[Authentication]
  AD[Authorization]
  U -->AT
  AT -->|Authenticated| AC
  AC -->|Granted operations|R
  SA -->AD
  AD -->AC
```
- All activities are under audit


## Subjects access objects with certain access rights


Subject
---
- an entity capable of accessing objects
- the process that represents user or application
  - takes on the attributes of the user, such as access rights
-  is typically held accountable for its actions
   -  the actions may be recorded by an audit trail 
- belongs to one of three classes with different access rights
  - owner of the IS resources
  - group with identical access rights for each member
    - a user may belong to multiple groups
  - world (others) includes anyone other than owner and group
    - have least amount of access


Objects
---
- an IS resource under access control
- contains information, such as
  - files, directories
  - records, blocks, tables, databases
  - pages, segments, devices
  - messages, mailboxes, programs


Access Rights
---
- the operations that a subject can do on an object
- typical access rights
  - create, delete
  - read, write
  - search, execute


Access Control Policies
---
- integrated in authorization systems
- dictate what types of access are permitted, under what circumstances, and by whom
- include four typical categories
  - Discretionary access control (DAC)
    - by the requestor's volition
  - Mandatory access control (MAC)
    - by the requestor's clearance
  - Role-based access control (RBAC)
    - by the requestor's roles
  - Attribute-based access control (ABAC)
    - by related attributes of the requestor, resources and environmental conditions
- these categories are not mutually exclusive



[Discretionary Access Control (DAC)](https://en.wikipedia.org/wiki/Discretionary_access_control)
---
- restricts access to objects based on 
  - the identity of subjects
  - groups to which they belong
- allow pass of access rights to other subjects
- implemented in [access matrix (table)](https://en.wikipedia.org/wiki/Access_control_matrix) $A[s_i,r_j]$ with
  - left-most column of subjects $s_i$
  - head row of resources $r_j$
  - cells of access rights $a_{s_i,r_j}$
- can be extended to
  - a general model [Graham–Denning model](https://en.wikipedia.org/wiki/Graham%E2%80%93Denning_model)
  - associate capabilities with protection domains, such as
    - user domain (user mode)
    - kernel domain (kernel mode)


A access matrix
---
| S/A/R | Resource 1 | Resource 2 | Resource 3 |
| --- | --- | --- | --- |
| Subject 1 |  | read | own |
| Subject 2 | read<br>write | | write |
| Subject 3 | read | | read<br>write |


Access control list(ACL) and capability ticket
---
- each column $r_j$ is an **access control list(ACL)**
  - lists subjects and their access rights on this resource
  - unlisted subjects are assigned a default set of rights
  - convenient to determine which subjects have which access rights to a specified resource
  - hard to determine all access rights for a particular subject
- each row $s_i$ is a **capability ticket**
  - specifies authorized objects and operations for a particular subject
  - convenient to determine the set of access rights for a given subject
  - can be given to other subjects
  - must be unforgeable. How?
    - let the OS hold  all tickets on behalf of subjects
    - add a message authentication code (MAC) to each ticket
      - the MAC is verified by the relevant resource whenever access is requested
  - hard to determine subjects and their access rights on a given resource


Authorization table
---
- like access matrix but not sparse
- each row specifies one access right of one subject to one resource
  - sorting the table by subject gives capability lists
  - sorting the table by resource gives access control lists
- can be easily implemented with relational database


💡 Demo 
---
- [Traditional UNIX File Access Control](https://en.wikipedia.org/wiki/File-system_permissions)
- subjects: user (owner), group, and other
- objects: files and directories
- access rights: read, write, and execute

```bash
# 1. list permissions of a file
ls -l afile

# 2. list permissions of a directory
ls -ld adirectory

# 3. change permissions
chmod {options} filename
```


🔭 Explore
--- 
- Explore [Access Control Lists (ACLs) in UNIX](https://help.ubuntu.com/community/FilePermissionsACLs)


[Role-Based Access Control (RBAC)](https://en.wikipedia.org/wiki/Role-based_access_control)
---
- simplifies DAC by abstracting the access rights shared by a group of subjects into role
  - roles have hierarchies
- decouples access rights from subjects
- typically defines a role as a job function within an organization
- a user can be assigned more than one roles statically or dynamically
  - according to the principle of least privileges
- widely used now and required by [FIPS PUB 140-3: Security Requirements for Cryptographic Modules](https://csrc.nist.gov/pubs/fips/140-3/final)
- modelled with two matrixes
  - role assignment matrix assigns subjects with roles 
  - role access rights matrix specifies the access rights of each role on the resources


### RBAC Reference Models

RBAC models
---
```mermaid
flowchart LR
U((Users))
S((Sessions))
R((Roles))
subgraph OPO
Op((Operations))
Ob((Objects))
Op <-->|Permissions| Ob
end
U <-->|user_sessions| S
S <-->|session_roles|R
U <-->|"User assignments (UA) "|R
R <-->|"Role hierarchy (RH)"| R
R <-->|"Permission assignment (PA)"| OPO
```

Relationship among RBAC models
---
```mermaid
flowchart LR
  R0("RBAC0 
  Base model")
  R1("RBAC1
  Role hierarchies")
  R2("RBAC2
  Constraints")
  R3("RBAC3
  Consolidated model")
  R1---R3
  R3---R2
  R1---R0
  R0---R2
```


## Attribute-Based Access Control

### Attributes

### ABAC Logical Architecture

### ABAC Policies

## Identity, Credential, and Access Management

### Identity Management

### Credential Management

### Access Management

### Identity Federation

## Trust Frameworks

### Traditional Identity Exchange Approach

### Open Identity Trust Framework

# References
- [CS 509U: Access Control: Theory and Practice](https://www.cs.purdue.edu/homes/ninghui/courses/Spring06/)
- [Competitive Programming Courses at Purdue Computer Science](https://www.cs.purdue.edu/homes/ninghui/courses/cpx_index.html)