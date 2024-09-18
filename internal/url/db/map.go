package url

import (
	"context"
	"errors"
	"fmt"
)

type Url struct{ 
	shortToLong map[string]string
	longToShort map[string]string
}

//Конструктор
func NewRepositoryMap() *Url{
	return &Url{
		shortToLong: make(map[string]string),
		longToShort: make(map[string]string),
	}
}

func (p *Url) Create(ctx context.Context, longPath string, shortPath string) error {
	if(p.shortToLong[shortPath] == "" && p.longToShort[longPath] == ""){
		p.shortToLong[shortPath] = longPath
		p.longToShort[longPath] = shortPath
		} else {
			return errors.New("Fail add long and short url to map")
		}
		return nil
	}

func (p *Url) FindOneLongByShort(ctx context.Context, shortPath string) (longUrl string, err error){
	longPath, ok := p.shortToLong[shortPath]
	if !ok {
		return "", errors.New("Don't find longUrl")
	}
	return longPath, nil
}

func (p *Url) FindOneShortByLong(ctx context.Context, longPath string) ( shortUrl string, err error){
	fmt.Printf("%#v \n", p.longToShort)
	shortPath, ok := p.longToShort[longPath]
	if !ok {
		return "", errors.New("Don't find short")
	}
	return shortPath, nil
}