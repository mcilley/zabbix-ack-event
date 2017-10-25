# zabbix-ack-event
#### Script to ack zabbix event(s), created for use as an agi script for asterisk

#### Example Usage in Asterisk - extensions.conf:
```
 ;;PLayback the name of the digit and ack alert in zabbix
 exten => 1,1,Set(RESPOND=1)
 exten => 1,n,agi(ackEvents.agi,"--name=${NAME}","--triggerId=${ALERTID}")
 exten => 1,n,agi(ttsGoogle.agi,"You have Pressed ${EXTEN} to Accept this issue, Goodbye!")
 exten => 1,n,Hangup()
 ```
