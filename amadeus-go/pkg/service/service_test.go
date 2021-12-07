package service

import (
	"amadeus-go/config/env"
	"amadeus-go/pkg/repository"
	"github.com/hooklift/gowsdl/soap"
	"strings"
	"testing"
)

func setup(){
	env.Init("./config.json")
}

func TestAmadeusService_GenerateControlNumber(t *testing.T) {
	setup()
	uri := env.GetString("api.base_url")
	client := soap.NewClient(uri)
	
	client.AddHeader(
		&repository.AuthenticationSoapHeader{
			WSUserName: env.GetString("api.auth.user"),
			WSPassword: env.GetString("api.auth.password"),
		},
	
	)
	
	soapSvc := repository.NewEPowerServiceSoap(client)
	svc := New(nil,soapSvc)
	
	str,_:=svc.GenerateControlNumber("fef578e5-90d1-46cc-a19f-0592ba1e3515")
	 if !strings.EqualFold("2961ebf68fb27da0139939c040190f32696989fd",str){
	 	t.FailNow()
	 }
	println(str)
	
}
