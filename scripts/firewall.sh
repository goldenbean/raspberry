iptables -A INPUT -p tcp -m tcp --sport 80 -m string --string "Location: http://111." --algo bm --to 65535 -j DROP
iptables -A INPUT -p tcp -m tcp --sport 80 -m string --string "Location: http://117." --algo bm --to 65535 -j DROP
iptables -A INPUT -p tcp -m tcp --sport 80 -m string --string "Location: http://91." --algo bm --to 65535 -j DROP
