package chat

// import (
// 	"context"
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"

// 	e "main/domain/errors"
// 	proto "main/src/proto"

// 	"golang.org/x/oauth2"
// 	"golang.org/x/oauth2/google"
// 	"google.golang.org/api/calendar/v3"
// 	"google.golang.org/api/option"
// )

// func getClient(tokFile string, config *oauth2.Config) (*http.Client, error) {
// 	tok, err := tokenFromFile(tokFile)
// 	if err != nil {
// 		return nil, e.StacktraceError(err)
// 	}
// 	tokenSource := config.TokenSource(context.Background(), tok)
// 	newToken, err := tokenSource.Token()
// 	if err != nil {
// 		return nil, e.StacktraceError(err)
// 	}
// 	if newToken.AccessToken != tok.AccessToken {
// 		if err := saveToken(tokFile, newToken); err != nil {
// 			return nil, e.StacktraceError(err)
// 		}
// 		log.Println("Saved new token:", newToken.AccessToken)
// 	}
// 	return config.Client(context.Background(), tok), nil
// }

// // Retrieves a token from a local file.
// func tokenFromFile(file string) (*oauth2.Token, error) {
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return nil, e.StacktraceError(err)
// 	}
// 	defer f.Close()
// 	tok := &oauth2.Token{}
// 	if err := json.NewDecoder(f).Decode(tok); err != nil {
// 		return nil, e.StacktraceError(err)
// 	}
// 	return tok, nil
// }

// // Saves a token to a file path.
// func saveToken(path string, token *oauth2.Token) error {
// 	log.Println("Saving credential file to: ", path)

// 	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
// 	if err != nil {
// 		return e.StacktraceError(err)
// 	}
// 	defer f.Close()

// 	if err := json.NewEncoder(f).Encode(token); err != nil {
// 		return e.StacktraceError(err)
// 	}

// 	return nil
// }

// func (cm *ChatManager) getCalendarServiceClient() (*calendar.Service, error) {
// 	ctx := context.Background()
// 	b, err := os.ReadFile(cm.credentialsFile)
// 	if err != nil {
// 		log.Println("Unable to read client secret file: ", err)
// 		return nil, e.StacktraceError(err)
// 	}

// 	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
// 	if err != nil {
// 		log.Println("Unable to parse client secret file to config: ", err)
// 		return nil, e.StacktraceError(err)
// 	}
// 	client, err := getClient(cm.tokenFile, config)
// 	if err != nil {
// 		log.Println("Unable to get client from token: ", err)
// 		return nil, e.StacktraceError(err)
// 	}

// 	return calendar.NewService(ctx, option.WithHTTPClient(client))

// }

// func (uc *ChatManager) GetCalendarEvents(teacherID int, classID int) (*proto.GetEventsResponse, error) {
// 	srv, err := uc.getCalendarServiceClient()
// 	if err != nil {
// 		log.Println("Unable to retrieve calendar Client: ", err)
// 		return nil, e.StacktraceError(err)
// 	}
// 	calendarDB, err := uc.store.GetCalendarDB(teacherID)
// 	if err != nil {
// 		log.Println("DB err: ", err)
// 		return nil, e.StacktraceError(err)
// 	}
// 	t := time.Now().Format(time.RFC3339)
// 	events, err := srv.Events.List(calendarDB.IDInGoogle).ShowDeleted(false).
// 		SingleEvents(true).TimeMin(t).MaxResults(100).OrderBy("startTime").Do()
// 	if err != nil {
// 		log.Println("Unable to retrieve next ten of the user's events: ", err)
// 		return nil, e.StacktraceError(err)
// 	}

// 	//ans := []model.CalendarEvent{}
// 	ans := proto.GetEventsResponse{}
// 	for _, item := range events.Items {
// 		s := strings.Split(item.Summary, " ")
// 		clID := 0
// 		if len(s) > 2 && s[len(s)-2] == "Class" {
// 			clIDs := s[len(s)-1]
// 			clID, err = strconv.Atoi(clIDs)
// 			if err != nil {
// 				log.Println("error: ", err)
// 				return nil, e.StacktraceError(err)
// 			}
// 		}
// 		if clID != classID {
// 			continue
// 		}
// 		time1, err := time.Parse(time.RFC3339, item.Start.DateTime)
// 		if err != nil {
// 			log.Println("Error while parsing date :", err)
// 			return nil, e.StacktraceError(err)
// 		}
// 		time2, err := time.Parse(time.RFC3339, item.End.DateTime)
// 		if err != nil {
// 			log.Println("Error while parsing date :", err)
// 			return nil, e.StacktraceError(err)
// 		}
// 		// tmp := model.CalendarEvent{Title: item.Summary, Description: item.Description,
// 		// 	StartDate: time1, EndDate: time2, ClassID: classID, ID: item.Id}
// 		tmp := proto.EventData{Title: item.Summary, Description: item.Description,
// 			StartDate: time1.String(), EndDate: time2.String(), Id: item.Id}

// 		ans.Events = append(ans.Events, &tmp)
// 	}

// 	return &ans, nil
// }
