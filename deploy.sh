#!/bin/bash

SERVER_IP="$1";
SOLUTION="$2";

git add "$SOLUTION/";
git commit -m "$SOLUTION solution iteration";
git push;

ssh -tt -o StrictHostKeyChecking=no -l root "$SERVER_IP" <<ENDSSH
cd /root/apps/protohackers;

git checkout "$SOLUTION";
git pull origin "$SOLUTION";

cd /root/apps/protohackers/"$SOLUTION"/
go build
./"$SOLUTION"

git checkout .

exit
ENDSSH
