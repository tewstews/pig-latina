# pig-latina

## Instalations
```
go get github.com/tewstews/pig-latina
```

## Usage
``` golang
func New(vowelSfx, consonantsSfx string, ignoreE bool) *translator // return a translator type

// translator 
type translator struct { 
	VowelsSfx     string // sufix which is substituted for the first vowels
	ConsonantsSfx string // sufix which is substituted for the first consonants
	IgnoreE       bool   // remove last vowel "e"
}

// translate 
func (cfg *translator) Translate(textfields ...string) []string 
``` 

### example
```golang
main()  {
	trs := pig.New("hay", "ay", true)

	translation := trs.Translate("There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form")
  ```
  ``` 
 //output
 [Erthay arhay anymay ariationsvay ofhay assagespay ofhay Oremlay Umipsay availablhay, utbay thay ajoritymay avhay ufferedsay alterationhay inhay omsay ormfay]
 ```
 
