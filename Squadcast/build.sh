cd APIServer
go build -o apiserver
cd ..
cd EventHandler
go build -o eventhandler
cd ..
cd Worker/Jira
go build -o jira
cd ../..
cd Worker/Slack
go build -o slack
cd ../..
cd Worker/Reddit
go build -o reddit
cd ../..

./NatsServer/gnatsd --addr 127.0.0.1 --port 4222 &
./APIServer/apiserver &
./EventHandler/eventhandler &
./Worker/Jira/jira &
./Worker/Slack/slack &
./Worker/Reddit/reddit &