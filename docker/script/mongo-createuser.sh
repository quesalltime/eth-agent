#!/usr/bin/env bash

NORMAL="\\033[0;39m"
RED="\\033[1;31m"
BLUE="\\033[1;34m"

# Create user who is dbOwner
echo -e -n "$BLUE"
echo "Creating mongo users(dbOWner)... $MONGO_BART_DB $MONGO_BART_USERNAME $MONGO_BART_PASSWORD"
echo -e -n "$NORMAL"

mongo << EOF
use $MONGO_BART_DB 
db.createUser({user: '$MONGO_BART_USERNAME', pwd: '$MONGO_BART_PASSWORD', roles:[{role:'dbOwner',db:'$MONGO_BART_DB'}]})
EOF

echo "Mongo users(dbOWner) created."
echo ""
echo ""

# Create user who read only
echo -e -n "$BLUE"
echo "Creating mongo users(read)... $MONGO_BART_DB $MONGO_BART_READONLY_USERNAME $MONGO_BART_READONLY_PASSWORD"
echo -e -n "$NORMAL"

mongo << EOF
use $MONGO_BART_DB 
db.createRole({role:'bartread', privileges:[{resource:{db:'$MONGO_BART_DB',collection:''},actions:['collStats','dbHash','dbStats','find','killCursors','listIndexes','listCollections']}],roles:[]})
db.createUser({user:'$MONGO_BART_READONLY_USERNAME', pwd:'$MONGO_BART_READONLY_PASSWORD', roles:[{role:'bartread',db:'$MONGO_BART_DB'}]})
EOF

echo "Mongo users(read) created."
