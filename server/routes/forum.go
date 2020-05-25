package routes

import (
	"net/http"
	"strconv"
	"time"

	db "github.com/gittoks/diplom/server/database"
)

func ForumHandlerGET(w http.ResponseWriter, r *http.Request) {
	buyerCookie := GetCookie(w, r)
	if CheckLoginByCookie(buyerCookie) {

		topic_id, val := 0, r.URL.Query()
		if s := val["topic"]; s != nil {
			topic_id, _ = strconv.Atoi(s[0])
		}

		if topic_id == 0 {
			topics, _ := db.GetTopics()
			interfaces := make([]interface{}, len(topics))
			for i, v := range topics {
				interfaces[i] = v
			}
			Answer(w, GetNavBar(buyerCookie), CommentPage{interfaces, buyerCookie, db.Topic{}}, "forum.html", "", "", 4)
		} else {
			topic, _ := db.GetTopic(uint(topic_id))
			comments, _ := db.GetComments(uint(topic_id))
			interfaces := make([]interface{}, len(comments))
			for i, v := range comments {
				interfaces[i] = v
			}
			Answer(w, GetNavBar(buyerCookie), CommentPage{interfaces, buyerCookie, topic}, "forum_topic.html", "", "", 4)
		}

	} else {
		Answer(w, GetNavBar(buyerCookie), nil, "info.html", "вы уже авторизованы", "warning", 0)
	}
}

func ForumHandlerPOST(w http.ResponseWriter, r *http.Request) {
	buyerCookie := GetCookie(w, r)
	if CheckLoginByCookie(buyerCookie) {

		topic_id, val := 0, r.URL.Query()
		if s := val["topic"]; s != nil {
			topic_id, _ = strconv.Atoi(s[0])
		}

		switch r.PostFormValue("action") {
		case "topic":
			topic := db.Topic{
				Name:        r.PostFormValue("name"),
				Description: r.PostFormValue("text"),
				BuyerID:     buyerCookie.ID,
			}
			db.CreateTopic(topic)
			break
		case "comment":
			comment := db.Comment{
				Text:    r.PostFormValue("name"),
				BuyerID: buyerCookie.ID,
				TopicID: uint(topic_id),
				Date:    time.Now().Format("15:04 02.01.2006"),
			}
			db.CreateComment(comment)
			break
		case "delete_comment":
			id, _ := strconv.Atoi(r.PostFormValue("id"))
			db.DeleteComment(uint(id), buyerCookie.ID)
			break
		case "delete_topic":
			if buyerCookie.Role == 1 {
				id, _ := strconv.Atoi(r.PostFormValue("id"))
				db.DeleteTopic(uint(id))
				db.DeleteCommentByTopic(uint(id))
			}
			break
		}
	}
	ForumHandlerGET(w, r)
}
