# zabbix-ack-event
#### Script to ack zabbix event(s), created for use as an agi script for asterisk
Used in https://github.com/mcilley/asterisk-zabbix-phone-escalation
As this was created for use with phone alerting agi script above, this script will ack all prior unacknoleged alerts (otherwise current alerting event not noted as being acknoleged).


#### Example Usage in Asterisk - extensions.conf:
```
 ;;PLayback the name of the digit and ack alert in zabbix
 exten => 1,1,Set(RESPOND=1)
 exten => 1,n,agi(ackEvents.agi,"--name=${NAME}","--triggerId=${ALERTID}")
 exten => 1,n,agi(ttsGoogle.agi,"You have Pressed ${EXTEN} to Accept this issue, Goodbye!")
 exten => 1,n,Hangup()
 ```
