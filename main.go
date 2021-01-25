package main

import (	
	"encoding/json"
	"io/ioutil"
	 "fmt"
	 "reflect"
	 "net/http"
	 "strings"
)

// PostAuthor type
type PostAuthor struct {
	name string
	title string
	body string
}

func main () {

	author := PostAuthor{}

	getPost(&author)

	fmt.Println(author)

}

func getPost(postAuthor *PostAuthor) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if(err != nil){
		// log error
	}else{
		
		json.NewDecoder(response.Body).Decode(&postAuthor)
			body, readErr := ioutil.ReadAll(response.Body)
			if(readErr != nil){
				//log error
			}else{
				//sb := string(body)
				//json.NewDecoder(response.Body).Decode(&postAuthor)
				// fmt.Println(sb)
				// postAuthor.title = sb.postTitle
				// postAuthor.body = sb.body
				//If you define a struct named Foo, the kind is struct and the type is Foo.
				typ := reflect.TypeOf(body)
				examiner(typ, 0)
		// for i := 0; i < typ.NumIn(); i++ {
		// 	fmt.Printf("Param %d: %v\n", i, typ.In(i))
		// }
		}
	}
}

func examiner(t reflect.Type, depth int) {
	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "and kind is", t.Kind())
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type:")
		examiner(t.Elem(), depth+1)
	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, "type is", f.Type.Name(), "and kind is", f.Type.Kind())
			if f.Tag != "" {
				fmt.Println(strings.Repeat("\t", depth+2), "Tag is", f.Tag)
				fmt.Println(strings.Repeat("\t", depth+2), "tag1 is", f.Tag.Get("tag1"), "tag2 is", f.Tag.Get("tag2"))
			}
		}
	}
}