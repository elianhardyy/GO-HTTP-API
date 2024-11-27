package ws

import (
	"encoding/json"
	"net/http"
	"server/utils"

	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request){
	var req CreateRoomReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"error")
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"error")
	}
	roomId := r.URL.Query().Get("roomId")
	userId := r.URL.Query().Get("userId")
	username := r.URL.Query().Get("username")

	cl := &Client{
		Conn: conn,
		Message: make(chan *Message,10),
		ID: userId,
		RoomID: roomId,
		Username: username,
	}

	m := &Message{
		Content: "A new user has joined the room",
		RoomID: roomId,
		Username: username,
	}
	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}