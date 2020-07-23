package message

type messageJSON struct {
	Participants       []participant `json:"participants"`
	Messages	       []rawMessage  `json:"messages"`
	Title              string        `json:"title"`
    IsStillParticipant bool          `json:"is_still_participant"`
	ThreadType         string        `json:"thread_type"`
	ThreadPath         string        `json:"thread_path"`
}

type participant struct {
	Name string `json:"name"`
}

type photo struct {
	URI               string `json:"uri"`
	CreationTimestamp float64 `json:"creation_timestamp"`
}

type reaction struct {
	Reaction string `json:"reaction"`
	Actor    string `json:"actor"`
}

type video struct {
	URI               string    `json:"uri"`
	CreationTimestamp float64   `json:"creation_timestamp"`
	Thumbnail         thumbnail `json:"thumbnail,omitempty"`
}

type thumbnail struct {
	URI string `json:"uri"`
}

type share struct {
	Link string `json:"link,omitempty"`
}

type rawMessage struct {
	SenderName  string     `json:"sender_name"`
	TimestampMs float64    `json:"timestamp_ms"`
	Content     string     `json:"content,omitempty"`
	Videos      []video    `json:"videos,omitempty"`
	Share       share      `json:"share,omitempty"`
	Photos      []photo    `json:"photos,omitempty"`
	Reactions   []reaction `json:"reactions,omitempty"`
	Types       string     `json:"type"`
}