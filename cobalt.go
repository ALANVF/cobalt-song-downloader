package main

const (
	CobaltCodec_h264 = "h264"
	CobaltCodec_av1 = "av1"
	CobaltCodec_vp9 = "vp9"
)

const (
	CobaltFormat_best = "best"
	CobaltFormat_mp3 = "mp3"
	CobaltFormat_ogg = "ogg"
	CobaltFormat_wav = "wav"
	CobaltFormat_opus = "opus"
)

const (
	CobaltFilePattern_classic = "classic"
	CobaltFilePattern_pretty = "pretty"
	CobaltFilePattern_basic = "basic"
	CobaltFilePattern_nerdy = "nerdy"
)

type CobaltRequest struct {
	Url             string `json:"url"`
	Codec           string `json:"vCodec,omitempty"`          // default: h264
	Quality         string `json:"vQuality,omitempty"`        // default: 720
	Format          string `json:"aFormat,omitempty"`         // default: mp3
	FilePattern     string `json:"filenamePattern,omitempty"` // default: classic
	IsAudioOnly     bool   `json:"isAudioOnly,omitempty"`     // default: false
	IsTTFullAudio   bool   `json:"isTTFullAudio,omitempty"`   // default: false
	IsAudioMuted    bool   `json:"isAudioMuted,omitempty"`    // default: false
	DubLang         bool   `json:"dubLang,omitempty"`         // default: false
	DisableMetadata bool   `json:"disableMetadata,omitempty"` // default: false
	TwitterGif      bool   `json:"twitterGif,omitempty"`      // default: false
	TiktokH265      bool   `json:"tiktokH265,omitempty"`      // default: false
}


const (
	CobaltStatus_error = "error"
	CobaltStatus_redirect = "redirect"
	CobaltStatus_stream = "stream"
	CobaltStatus_success = "success"
	CobaltStatus_rate_limit = "rate-limit"
	CobaltStatus_picker = "picker"
)

const (
	CobaltPickerType_various = "various"
	CobaltPickerType_images = "images"
)

type CobaltPickerItem struct {
	Type  string `json:"type"`
	Url   string `json:"url"`
	Thumb string `json:"thumb"`
}

type CobaltResponse struct {
	Status     string             `json:"status"`
	Text       string             `json:"text"`
	Url        string             `json:"url"`
	PickerType string             `json:"pickerType"`
	Picker     []CobaltPickerItem `json:"picker"`
	Audio      string             `json:"audio"`
}