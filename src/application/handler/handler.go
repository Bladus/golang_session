package handler

import (
    "application/config"
    "application/route"
    "application/session"
    "net/http"
    "time"
    "fmt"
    "log"
)

const (
    COOKIE_NAME = "sessionId"
)

type Handler struct {
    Route     *route.Route
    Config    *config.Config
    Session   *session.Session
}

func (h *Handler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
    switch path := req.URL.Path; path {
    case "/":
        cookie, _ := req.Cookie(COOKIE_NAME)
        if cookie != nil {
            log.Printf("%s\n", h.Session.Get(cookie.Value))
        }
    case "/login":
        username := req.FormValue("username")
        password := req.FormValue("password")
        fmt.Printf("%s\n", username)
        fmt.Printf("%s\n", password)

        sessionId := h.Session.Init(username)

        cookie := &http.Cookie{
            Name: COOKIE_NAME,
            Value: sessionId,
            Expires: time.Now().Add(5 * time.Minute),
        }
        http.SetCookie(res, cookie)

    default:
    }
}