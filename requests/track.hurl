POST http://localhost:8080/track
{
	"info": {
		"user_id": "989e27fc-ab6a-4d1e-9581-aef237b5421a",
		"device": {
			"manufacturer": "Apple",
			"model": "iPhone X",
			"build_number": "333",
			"os": "ios",
			"os_version": "17",
			"screen_width": 1920,
			"screen_height": 1080
		},
		"app_version": "1.0"
	},
	"data": [
		{
			"action": "open",
			"data": {
				"screen": "first"
			},
			"timestamp": "2024-04-21T02:21:00Z"
		}
	]
}
HTTP 200
