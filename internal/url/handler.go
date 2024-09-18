package url

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/Mishany52/testTaskSWoyo/internal/config"
	"github.com/Mishany52/testTaskSWoyo/internal/handlers"
	"github.com/Mishany52/testTaskSWoyo/utils"

	"github.com/go-chi/chi/v5"
)

var conf = config.GetConfig()


type handler struct {
	repository Repository
}

func NewHandler( repository Repository) handlers.Handler{
	return &handler{
		repository: repository,
	}
}

func (h *handler) Register(router *chi.Mux){
	router.Post("/", h.Create)
	router.Get("/{urlKey}", h.GetByShortPath)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request){
	
	var shortUrl string;
	
	var dto CreateShortUrl
	//Чтение из body
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	//Валидация
	bodyStr := string(dto.LongUrl)
	validUrl := h.validateUrl(bodyStr)
	if(validUrl == ""){
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	//Проверяем есть ли уже короткая ссылка в хранилище
	shortUrl, _ = h.repository.FindOneShortByLong(context.TODO(), validUrl)
	if shortUrl == ""{
		//Генерируем короткий url
		shortUrl = conf.FullServerAddr + utils.RandomString(7)

		//Добавляем в базу или map в зависимости от запуска
		err = h.repository.Create(context.TODO(), validUrl, shortUrl)

		if err != nil {
			log.Println(err)
		}
	}
	
	w.Write([]byte(shortUrl))
	w.WriteHeader(201)
}

func (h *handler) GetByShortPath(w http.ResponseWriter, r *http.Request){
	//извлекаем из query ключ 
	shortUrl := conf.FullServerAddr + chi.URLParam(r,"urlKey")

	longUrl, err := h.repository.FindOneLongByShort(context.TODO(), shortUrl)
	if err != nil || longUrl == ""{
		log.Println(err)
		http.Error(w, "Not found longUrl for shortUrl", http.StatusNotFound)
	}
	w.Write([]byte(longUrl))
	w.WriteHeader(200)
}

func (h* handler) validateUrl(url string) (validUrl string) {
	//Регулярное выражение для проверки на валидность url
	re := regexp.MustCompile("((((https?|ftps?|gopher|telnet|nntp)://)|(mailto:|news:))([-%()_.!~*';/?:@&=+$,A-Za-z0-9])+)")
	validUrl = re.FindString(url)
	return validUrl
}
