#!/bin/bash

# turn on bash's job control
set -m

# Start the primary process and put it in the background
#./tcp_server -interactive=false $1 $NUMRUNS &

#sleep 1

# Start the helper process
./tcp_client -interactive=false $1 tcp-server:8080 $NUMBYTES $NUMRUNS

# the my_helper_process might need to know how to wait on the
# primary process to start before it does its work and returns


# now we bring the primary process back into the foreground
# and leave it there
#fg %1
#echo Your container args are : "$@"
