package types

type DownloadInfo struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Duration   string `json:"duration"`
	Resolution string `json:"resolution"`
	FormatNote string `json:"format_note"`
	Uploader   string `json:"uploader"`
	Status     string `json:"status"`
	Percent    string `json:"percent"`
	Downloaded int64  `json:"downloaded,string"`
	Total      int64  `json:"total,string"`
	Speed      string `json:"speed"`
	ETA        string `json:"eta"`
}
