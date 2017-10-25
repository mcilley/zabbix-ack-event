package main

import (
	"os"
	"log"
	"flag"
	"github.com/AlekSi/zabbix"
)

func Ack( events []string, name *string, api *zabbix.API ) {
	response, err := api.CallWithError( "event.acknowledge", zabbix.Params{ "eventids": events, "message":"Event has been acknowledged by "+*name } )
	
	//if there's an error print it out, yerp
	if err != nil {
		log.Fatalf( "There has been an error acknowledging the provided eventid(s)" )
	}
	log.Println( response.Result )
}

func EventGet( triggerId *string, api *zabbix.API ) []string {
	response, err := api.CallWithError( "event.get", zabbix.Params{ "filter": map[string]string{"acknowledged":"0","value":"1"}, "objectids":triggerId, "withUnacknowledgedEvents":"True" } )

	//if there's an error print it out, yerp
	if err != nil {
		log.Fatalf( "There has been an error gathering the eventid(s) associated with the given triggerid" )
	}else if len(response.Result.( []interface{} ) ) == 0 {
		log.Fatalf( "No Events Retrieved for the noted triggerid")
	}

	eventMap := response.Result.( []interface{} )
	var events []string

	for _, v := range eventMap {
		n := v.( map[string]interface{} )
		events = append( events, n["eventid"].(string) )
	}
	return events
}

func main() {

	//throw an error if we don't have any args
	if len(os.Args) < 1{
		log.Fatalf( "No args provided, nothing to do here!" )
	}

	//setup and parse our arguments
	name 			:=	flag.String( "name","Name Not Provided","Name of Person that is acknowledging alert" )
	triggerId 		:=	flag.String( "triggerId","","zabbix trigger ID - numeric ID for trigger used for acking alerts via api" )
	zabbixServer 	:=	flag.String( "zabbixServer","<default ip>", "zabbix server address" )
	zabbixUser 		:= 	flag.String( "zabbixUser","<default user>", "zabbix user for acking alerts" )
	zabbixPw   		:= 	flag.String( "zabbixPw","", "zabbix user's password" )
	flag.Parse()

	//use a default if none provided
	if *zabbixPw == ""{
		*zabbixPw = "<default password>"
	}

	//bail if we don't have a trigger id provided
	if *triggerId == "" {
		log.Fatalf( "No triggerd provided, nothing to do here!" )
	}

	//setup our zabbix connection
	api := zabbix.NewAPI("http://"+*zabbixServer+"/zabbix/api_jsonrpc.php")
	api.Login(*zabbixUser, *zabbixPw)
	
	//Ack Our Crap
	Ack( EventGet( triggerId, api ), name, api )
}
