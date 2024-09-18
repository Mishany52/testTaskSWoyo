package url

type Url struct {
	ID int64 `json:"id"`
	LongUrl string `json:"longUrl"`
	ShortUrl string `json:"shortUrl"`
}