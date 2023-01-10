#!/bin/sh
mysql -u root -p < db/db.sql
go build main.go 
sudo ss -lptn 'sport = :5003'
kill $(pgrep shab)
./hand_held > /dev/null 2>&1 & 
echo "runnin on prccedd id :" + $(pgrep main)

In the X project, I illustrated a social media campaign for a local clothes brand. The goal of the campaign was to improve their visibility and build an engaged audience within their target market’s age range. To achieve this goal, I implemented unique strategies such as targeted Facebook ads, promotional products for sale at local events. These strategies were extremely effective because they delivered results consistently over time - by the end of the month-long campaign, this local clothing brand became fully aware of its market and had increased sales by 15%.