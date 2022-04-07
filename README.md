# Go Botnet Manager

# WIP
This project is still a WIP, and its only done for research and experimentation porpuses do not use this code for any malicious intent

## FEATURES
- Central Server thats able to communicate with different nodes from the same logical net/botnet
- Agent binary thats able to interpretate the commands from the Central Server 
- Central Server thats able to manage the nodes from the botnet
- Central Server can perform healthchecks on the different stations
- The connections between the Central Server and its stations is encrypted
---
###Sequence Diagram
                    
```seq
Central_Node->stub_1: Remote_Command 
Central_Node->stub_2:Remote_Command
Central_Node->stub_2:HealthStatus
stub_2->>Central_Node:Health_OK
```