package helpers

import (
	"amadeus-go/config/env"
	"reflect"
	"testing"
	"time"
)

func setup() {
	env.Init("./config.json")
}

func TestGetClassByLetter(t *testing.T) {
	setup()
	
	letters := []string{"N", "W", "Z", "R"}
	expected := []string{
		"economy", "premium", "business", "first",
	}
	
	got := make([]string, 0)
	
	for _, l := range letters {
		class := GetClassByLetter(l)
		got = append(got, class)
	}
	
	if !reflect.DeepEqual(expected, got) {
		t.FailNow()
	}
	
}

func TestCalcDuration(t *testing.T) {
	ArrivalDateTime:= "2020-01-10T03:00:00"
	DepartureDay := "2020-01-10T06:05:00"
	t1,_:=time.Parse("2006-01-02T15:04:05",DepartureDay)
	t2,_:=time.Parse("2006-01-02T15:04:05",ArrivalDateTime)
	// 02:45
	dep:=t1.UTC().Format("2006-01-02T15:04:05")
	arr:=t2.UTC().Format("2006-01-02T15:04:05")
	// dep:=t2.Local().Location().String()
	// arr:=t1.Local().Location().String()
	
	println(dep)
	println(arr)
	println(time.Now().UTC().String())
	println(time.Now().Local().String())
	
	
}
