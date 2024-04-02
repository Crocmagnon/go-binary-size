package a

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type SomeStruct struct {
	field int
}

type SomeInt interface {
	Func() json.Decoder
}

func Func1() *gin.Engine {
	fmt.Println("func1")
	return gin.New()
}
