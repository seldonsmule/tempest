package tempest

import (
  "fmt"
  "encoding/json"
  "github.com/seldonsmule/logmsg"
  "github.com/seldonsmule/restapi"
)

type Tempest struct {

  sToken string
  sStationId string

}

func New(token string, stationid string) *Tempest {

  logmsg.Print(logmsg.Info, "NewTempest()")

  t := new(Tempest)

  t.sToken = token
  t.sStationId = stationid

  return t

}

func (pT *Tempest) Full_BetterForecastUrl() string{

  return BetterForecastUrl() + "?station_id=" + pT.sStationId + "&token=" + pT.sToken

}

func (pT *Tempest) HelloWorld() {

  fmt.Println("Hello World")
  fmt.Println("Urls.....")
  fmt.Println("\t BetterForecastUrl: " + pT.Full_BetterForecastUrl())

} 


func (pT *Tempest) GetBetterForecast() (bool, BetterForecast){

  var forecast BetterForecast

  r := restapi.NewGet("weatherflow", pT.Full_BetterForecastUrl())

  r.JsonOnly()

  if(!r.Send()){

    logmsg.Print(logmsg.Error, "Call to Weatherflow failed")
    return false, forecast

  }

  json.Unmarshal(r.BodyBytes, &forecast)

  return true, forecast

}
