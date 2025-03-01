#!/bin/bash

SERVER_IP="$1";
SOLUTION="$2";

cd "$SOLUTION"
go build -o "$SOLUTION".out
scp ./"$SOLUTION".out root@66.228.43.192:/root/apps/protohackers-go/"$SOLUTION/"

ssh -tt -o StrictHostKeyChecking=no -l root "$SERVER_IP" <<ENDSSH
cd /root/apps/protohackers-go/"$SOLUTION"/

./"$SOLUTION".out

exit
ENDSSH
