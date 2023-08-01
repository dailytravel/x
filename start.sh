#!/bin/bash

function cleanup {
    kill "$ACCOUNT_PID"
    kill "$CONFIG_PID"
    kill "$CMS_PID"
    kill "$SALES_PID"
    kill "$HRM_PID"
    kill "$MARKETING_PID"
    kill "$FINANCE_PID"
    kill "$SERVICE_PID"
    kill "$REPORTING_PID"
}

trap cleanup EXIT

go build -o /tmp/srv-account ./account
go build -o /tmp/srv-config ./config
go build -o /tmp/srv-cms ./cms
go build -o /tmp/srv-sales ./sales
go build -o /tmp/srv-hrm ./hrm
go build -o /tmp/srv-marketing ./marketing
go build -o /tmp/srv-finance ./finance
go build -o /tmp/srv-service ./service
go build -o /tmp/srv-reporting ./reporting

/tmp/srv-account & ACCOUNT_PID=$!
/tmp/srv-config & CONFIG_PID=$!
/tmp/srv-cms & CMS_PID=$!
/tmp/srv-sales & SALES_PID=$!
/tmp/srv-hrm & HRM_PID=$!
/tmp/srv-marketing & MARKETING_PID=$!
/tmp/srv-finance & FINANCE_PID=$!
/tmp/srv-service & SERVICE_PID=$!
/tmp/srv-reporting & REPORTING_PID=$!

sleep 1

node gateway/index.js
