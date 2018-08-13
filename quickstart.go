package main

import (
        "encoding/json"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "time"

        "golang.org/x/net/context"
        "golang.org/x/oauth2"
        "golang.org/x/oauth2/google"
        "google.golang.org/api/calendar/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
        tokFile := "token.json"
        tok, err := tokenFromFile(tokFile)
        if err != nil {
                tok = getTokenFromWeb(config)
                saveToken(tokFile, tok)
        }
        return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
        authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
        fmt.Printf("Go to the following link in your browser then type the "+
                "authorization code: \n%v\n", authURL)

        var authCode string
        if _, err := fmt.Scan(&authCode); err != nil {
                log.Fatalf("Unable to read authorization code: %v", err)
        }

        tok, err := config.Exchange(oauth2.NoContext, authCode)
        if err != nil {
                log.Fatalf("Unable to retrieve token from web: %v", err)
        }
        return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
        f, err := os.Open(file)
        defer f.Close()
        if err != nil {
                return nil, err
        }
        tok := &oauth2.Token{}
        err = json.NewDecoder(f).Decode(tok)
        return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
        fmt.Printf("Saving credential file to: %s\n", path)
        f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
        defer f.Close()
        if err != nil {
                log.Fatalf("Unable to cache oauth token: %v", err)
        }
        json.NewEncoder(f).Encode(token)
}

func main() {
        b, err := ioutil.ReadFile("credentials.json")
        if err != nil {
                log.Fatalf("Unable to read client secret file: %v", err)
        }

        // If modifying these scopes, delete your previously saved token.json.
        config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
        if err != nil {
                log.Fatalf("Unable to parse client secret file to config: %v", err)
        }
        client := getClient(config)

        srv, err := calendar.New(client)
        if err != nil {
                log.Fatalf("Unable to retrieve Calendar client: %v", err)
        }

        t := time.Now().Format(time.RFC3339)
        events, err := srv.Events.List("primary").ShowDeleted(false).
                SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
        if err != nil {
                log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
        }
        fmt.Println("Upcoming events:")
        if len(events.Items) == 0 {
                fmt.Println("No upcoming events found.")
        } else {
                for _, item := range events.Items {
                        date := item.Start.DateTime
                        if date == "" {
                                date = item.Start.Date
                        }
                        fmt.Printf("%v (%v)\n", item.Summary, date)
                }
        }


        fmt.Println("Testing calendar event create")
        event := &calendar.Event{
          Summary: "Test",
          Location: "Laura's Virtual Desk",
          Description: "A test event created with the go client",
          Start: &calendar.EventDateTime{
            Date: "2018-08-09",
          },
          End: &calendar.EventDateTime{
            Date: "2018-08-09",
          },
          Attendees: []*calendar.EventAttendee{
            &calendar.EventAttendee{Email:"test@test.com"},
          },
        }

        calendarId := "primary"
        event, err = srv.Events.Insert(calendarId, event).Do()
        if err != nil {
          log.Fatalf("Unable to create event. %v\n", err)
        }
        fmt.Printf("Event created: %s\n", event.HtmlLink)

        // events.insert("primary",)
}
