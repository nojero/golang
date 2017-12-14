#!/bin/bash
curl -XPOST -d'{"stack":100}' localhost:8080/v1/computers
curl -XPATCH -d'{"addr":50}' localhost:8080/v1/computers/0/stack/pointer
curl -XPOST localhost:8080/v1/computers/0/stack/insert/MULT
curl -XPOST localhost:8080/v1/computers/0/stack/insert/PRINT
curl -XPOST localhost:8080/v1/computers/0/stack/insert/RET
# The start of the main function
curl -XPATCH -d'{"addr": 0}' localhost:8080/v1/computers/0/stack/pointer
curl -XPOST -d'{"arg":1009}' localhost:8080/v1/computers/0/stack/insert/PUSH
curl -XPOST localhost:8080/v1/computers/0/stack/insert/PRINT
# Return address for when print_tenten function finishes
curl -XPOST -d'{"arg":6}' localhost:8080/v1/computers/0/stack/insert/PUSH
# Setup arguments and call print_tenten
curl -XPOST -d'{"arg":101}' localhost:8080/v1/computers/0/stack/insert/PUSH
curl -XPOST -d'{"arg":10}' localhost:8080/v1/computers/0/stack/insert/PUSH
curl -XPOST -d'{"arg":50}' localhost:8080/v1/computers/0/stack/insert/CALL
# Stop the program
curl -XPOST localhost:8080/v1/computers/0/stack/insert/STOP
# Execute the program
curl -XPATCH -d'{"addr":0}' localhost:8080/v1/computers/0/stack/pointer
curl -XPOST localhost:8080/v1/computers/0/exec
