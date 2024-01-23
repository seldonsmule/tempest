package tempest

import (
  "testing"
  "fmt"
  "time"
  "os"
)

func TestEnvVarsSet(t *testing.T) {

  var envTestToken = os.Getenv("TEMPEST_TEST_TOKEN")
  var envTestStationId = os.Getenv("TEMPEST_TEST_STATIONID")

  if(envTestToken == ""){
    t.Logf("environment variable TEMPEST_TEST_TOKEN is not set")
    t.FailNow()
  }

  if(envTestStationId == ""){
    t.Logf("environment variable TEMPEST_TEST_STATIONID is not set")
    t.FailNow()
  }

  fmt.Println("TEMPEST_TEST_TOKEN= " + envTestToken)
  fmt.Println("TEMPEST_TEST_STATIONID= " + envTestStationId)


}

func TestTempest(t *testing.T) {

  fmt.Println("Testing Tempest")

}

func TestTempestNew(t *testing.T) {

  fmt.Println("Testing Tempest New")

  var envTestToken = os.Getenv("TEMPEST_TEST_TOKEN")
  var envTestStationId = os.Getenv("TEMPEST_TEST_STATIONID")

  if(envTestToken == ""){
    t.Logf("environment variable TEMPEST_TEST_TOKEN is not set")
    t.FailNow()
  }

  if(envTestStationId == ""){
    t.Logf("environment variable TEMPEST_TEST_STATIONID is not set")
    t.FailNow()
  }
  

  Temp := New(envTestToken, envTestStationId)

  Temp.HelloWorld()

}

func TestTempestPrintUrls(t *testing.T) {

//  fmt.Println(BetterForecastUrl("stationid123", "token123"))
  fmt.Println(BetterForecastUrl())

}


func TestBetterForecast(t *testing.T) {

  var sunrise int
  var sunset int


  fmt.Println("Testing Calling the BetterForecast Api")

  var envTestToken = os.Getenv("TEMPEST_TEST_TOKEN")
  var envTestStationId = os.Getenv("TEMPEST_TEST_STATIONID")

  if(envTestToken == ""){
    t.Logf("environment variable TEMPEST_TEST_TOKEN is not set")
    t.FailNow()
  }

  if(envTestStationId == ""){
    t.Logf("environment variable TEMPEST_TEST_STATIONID is not set")
    t.FailNow()
  }
  

  Temp := New(envTestToken, envTestStationId)

  worked, forecast := Temp.GetBetterForecast()

  if(!worked){
    t.Logf("tempest.GetBetterForecast() failed")
    t.FailNow()
  }

  //fmt.Println(forecast)

  sunrise = forecast.Forecast.Daily[0].Sunrise
  sunset = forecast.Forecast.Daily[0].Sunset

  fmt.Println("LocationName: ", forecast.LocationName)

  fmt.Println("Sunrise: ", time.Unix(int64(sunrise), 0))
  fmt.Println("Sunset: ", time.Unix(int64(sunset), 0))
  fmt.Println("AirTemperature-c: ", forecast.CurrentConditions.AirTemperature)
  fmt.Println("AirTemperature-f: ", ((forecast.CurrentConditions.AirTemperature * 1.8) + 32))
  fmt.Println("Conditions: ", forecast.CurrentConditions.Conditions)
  fmt.Println("Brightness: ", forecast.CurrentConditions.Brightness)
  fmt.Println("SolarRadiation: ", forecast.CurrentConditions.SolarRadiation)
  fmt.Println("Uv: ", forecast.CurrentConditions.Uv)



}
