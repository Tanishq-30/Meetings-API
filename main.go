package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
)

type Meeting struct {
  Id				 string `json:"Id"`
  Title				 string `json:"Title"`
  Participants 	 	 string `json:"Participants"`
  StartTime    	 string `json:"Start_Time"`
  EndTime      	 string `json:"End_Time"`
  CreationTimestamp string `json:"Creation_Timestamp"`
}

type Participant struct {
  Name string `json:"Name"`
  Email string `json:"Email"`
  RSVP string `json:RSVP`
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllMeetings(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllMeetings")
    json.NewEncoder(w).Encode(Meetings)
}

func returnAllParticipants(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllParticipants")
    json.NewEncoder(w).Encode(Participants)
}

var Meetings []Meeting
var Participants []Participant

func main() {
    Meetings = []Meeting{
      Meeting{Id: "1", Title: "Lecture CS", Participants: "100", StartTime:"17:00", EndTime:"18:00", CreationTimestamp:"12:00"},
    }
    Participants = []Participant{
      Participant{Name: "xyz", Email:"abc@d.com", RSVP:"Yes"},
    }
    http.HandleFunc("/", homePage)
    http.HandleFunc("/meetings", returnAllMeetings)
    http.HandleFunc("/participants", returnAllParticipants)
    log.Fatal(http.ListenAndServe(":5000", nil))
}