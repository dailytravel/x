#!/bin/bash

# Define an array of port numbers 4001 - 4020 to kill processes on
ports=(4000 4001 4002 4003 4004 4005 4006 4007 4008 4009 4010 4011 4012 4013 4014 4015 4016 4017 4018 4019 4020 50051)

# Loop through each port and kill the process (if exists)
for port in "${ports[@]}"; do
  echo "Trying to kill process on port $port..."
  pid=$(lsof -t -i :$port)  # Get the PID of the process using the port
  if [[ -n "$pid" ]]; then
    echo "Killing process with PID: $pid"
    kill $pid  # Terminate the process
  else
    echo "No process found on port $port"
  fi
done

echo "Process termination completed."
