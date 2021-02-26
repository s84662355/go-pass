package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/timest/env"
	"os"
	"time"
)

func init() {

	godotenv.Load(".env")

}
