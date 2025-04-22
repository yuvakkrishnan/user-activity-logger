package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/yuvakkrishnan/user-activity-logger/internal/db"
	"github.com/yuvakkrishnan/user-activity-logger/internal/kafka"
	"github.com/yuvakkrishnan/user-activity-logger/pkg/models"
)

func UserActivityHandler(w http.ResponseWriter, r *http.Request) {
	var activity models.UserActivity // Use a local variable for thread safety

	err := json.NewDecoder(r.Body).Decode(&activity)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest) // Corrected status code
		return
	}

	// Send to Kafka
	activityJson, _ := json.Marshal(activity)
	kafka.SendMessage("user-activity-topic", string(activityJson))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User activity received....."))

	// Save to DB
	err = db.SaveActivity(models.UserActivity{
		UserID:    activity.UserID,
		TimeStamp: activity.TimeStamp,
		Action:    activity.Action,
	})
	if err != nil {
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
		log.Println("User activity processed")
		w.Write([]byte("Activity logged"))
	}
}
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
