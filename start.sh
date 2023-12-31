#!/bin/bash

clear
echo "The server is initializing..."

function cleanup {
  kill "$SERVER_PID"
  kill "$ACCOUNT_PID"
  kill "$BASE_PID"
  kill "$CMS_PID"
  kill "$COMMUNITY_PID"
  kill "$CONFIG_PID"
  kill "$FINANCE_PID"
  kill "$HRM_PID"
  kill "$INSIGHT_PID"
  kill "$MARKETING_PID"
  kill "$PAYMENT_PID"
  kill "$SALES_PID"
}

# Build each Go service
(cd server && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd account && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd base && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd cms && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd community && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd configuration && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd finance && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd hrm && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd insight && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd marketing && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd payment && GOOS=linux GOARCH=amd64 go build -v -o app .)
(cd sales && GOOS=linux GOARCH=amd64 go build -v -o app .)

# Start each Go service
./server/app & SERVER_PID=$!
./account/app & ACCOUNT_PID=$!
./base/app & BASE_PID=$!
./cms/app & CMS_PID=$!
./community/app & COMMUNITY_PID=$!
./configuration/app & CONFIG_PID=$!
./finance/app & FINANCE_PID=$!
./hrm/app & HRM_PID=$!
./insight/app & INSIGHT_PID=$!
./marketing/app & MARKETING_PID=$!
./payment/app & PAYMENT_PID=$!
./sales/app & SALES_PID=$!

trap cleanup EXIT

# Wait for the servers to start (adjust the timeout as needed)
timeout=30
interval=1
elapsed=0
all_pids=($SERVER_PID $ACCOUNT_PID $BASE_PID $CMS_PID $COMMUNITY_PID $CONFIG_PID $FINANCE_PID $HRM_PID $INSIGHT_PID $MARKETING_PID $PAYMENT_PID $SALES_PID)

# Loop until all servers are running or timeout occurs
while true; do
  all_running=true
  for pid in "${all_pids[@]}"; do
    if ! ps -p "$pid" > /dev/null; then
      all_running=false
      break
    fi
  done
  if $all_running; then
    echo "All servers are running."
    break
  fi

  sleep $interval
  elapsed=$((elapsed + interval))

  if [ $elapsed -ge $timeout ]; then
    echo "Timeout: Some servers failed to start."
    exit 1
  fi
done

node ./gateway/index.js

# Do any other necessary setup or start additional services here

# Example of running another service (replace with your actual command if needed)
# ./path/to/your_other_service & OTHER_SERVICE_PID=$!

# Wait for user input to exit (you can remove this part if you don't need it)
read -rp "Press Enter to stop the servers and exit..."
