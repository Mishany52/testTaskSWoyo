package pathstore

import "fmt"

type PathStoreMap struct{ 
	shortToLong map[string]string
	longToShort map[string]string

}

//Конструктор
func NewPathStoreMap() *PathStoreMap{
	return &PathStoreMap{
		shortToLong: make(map[string]string),
		longToShort: make(map[string]string),
	}
}

func (p *PathStoreMap) GetShortPath(longPath string) (string, bool){
	fmt.Printf("%#v \n", p.longToShort)
	shortPath, ok := p.longToShort[longPath]
	return shortPath, ok
}
func (p *PathStoreMap) Add(longPath string, shortPath string){
	if(p.shortToLong[shortPath] == "" && p.longToShort[longPath] == ""){
		p.shortToLong[shortPath] = longPath
		p.longToShort[longPath] = shortPath
	}
}
func (p *PathStoreMap) GetLongPath(shortPath string) (string, bool){
	longPath, ok := p.shortToLong[shortPath]
	return longPath, ok
}