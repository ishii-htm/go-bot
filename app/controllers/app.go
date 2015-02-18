package controllers

import "github.com/revel/revel"
import "strings"

type App struct {
	*revel.Controller
}

type JsonResponse struct {
  Text string `json:"text"`
  Username string `json:"username"`
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Incoming(
  token string,
  team_id string,
  channel_id string,
  channel_name string,
  timestamp int,
  user_id string,
  user_name string,
  text string,
  trigger_word string) revel.Result {

  // TODO トークンのチェックぐらいはするべきですよ

  response := JsonResponse{strings.TrimSpace(strings.Replace(text, trigger_word, "", -1))  + "よねー　わかるー", "socrates-go"}

  return c.RenderJson(response)
}
