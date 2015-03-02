package controllers

import "github.com/revel/revel"
import "strings"
import "math/rand"
import "time"
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

		var response JsonResponse

  // TODO トークンのチェックぐらいはするべきですよ

	// trigger はいらない子なので殺す
	text = strings.TrimSpace(strings.Replace(text, trigger_word, "", -1))

	// メッセージをバラして
	fields := strings.Fields(text)

	switch strings.TrimSpace(fields[0]) {
	case "meshi":
		// マッチしたらめし
		response = c.Meshi()
	default :
		// マッチしないやつにはとりあえず共感しとく(クズっぽさある)
	  response = c.Empathy(text, trigger_word)
	}

  return c.RenderJson(response)
}

func (c App) Empathy(msg string, trigger string) JsonResponse {
	return JsonResponse { msg + "よねー　わかるー", "socrates-go" }
}

func (c App) Meshi() JsonResponse {
	meshi_list := []string{ "麻布ラーメン", "鶴見屋", "謎のカレー", "GOGOカレー", "大戸屋", "AJITO", "豚組", "新規開拓" }

	rand.Seed(time.Now().Unix())

	return JsonResponse { meshi_list[rand.Intn(len(meshi_list))] + "のめしー", "socrates-go" }
}
