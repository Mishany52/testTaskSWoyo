package handler

import (
	"io"
	"net/http"
	"regexp"

	"github.com/Mishany52/testTaskSWoyo/config"
	pathstore "github.com/Mishany52/testTaskSWoyo/pathStore"
	"github.com/Mishany52/testTaskSWoyo/utils"
	"github.com/go-chi/chi/v5"
)

var conf = config.New()

type ShortPath struct {
	pathStore *pathstore.PathStoreMap
}

func NewShortPath(ps *pathstore.PathStoreMap) *ShortPath {
	return &ShortPath{pathStore: ps}
}
func (s *ShortPath) Create(w http.ResponseWriter, r *http.Request){
	//Регулярное выражение для проверки на валидность url
	re := regexp.MustCompile("((((https?|ftps?|gopher|telnet|nntp)://)|(mailto:|news:))([-%()_.!~*';/?:@&=+$,A-Za-z0-9])+)")
	
	//Чтение из body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
	}
	defer r.Body.Close() // Закрываем r.Body после чтения

	//Валидация
	bodyStr := string(body)
	matchUrl := re.FindString(bodyStr)
	if(matchUrl == ""){
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	shortPath, ok := s.pathStore.GetShortPath(matchUrl)
	if !ok {
		shortPath = conf.FullServerAddr + utils.RandomString(7)
		s.pathStore.Add(matchUrl, shortPath)
		// log.Println("New create a shortPath:", shortPath)
	} else {
		// log.Println("Old shortPath:", shortPath)
	}
	// log.Println("Received body:", bodyStr)
	// fmt.Fprintf(w, "Received body: %s", bodyStr)
	// log.Println("shortPath:", shortPath)
	// longPath, _ := s.pathStore.GetLongPath(shortPath)
	// log.Println("longPath:", longPath)
	w.Write([]byte(shortPath))
}

func (s *ShortPath) GetByShortPath(w http.ResponseWriter, r *http.Request){
	shortPath := conf.FullServerAddr + chi.URLParam(r,"urlKey")
	longPath, ok := s.pathStore.GetLongPath(shortPath)
	if ok {
		w.Write([]byte(longPath))
	}
}
