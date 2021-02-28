package env
import (

    
)

 
 
type Redis struct {
	Default string      `default:"Rone"`
 
	Rone struct {
		Addr     string   `default:"127.0.0.1:6379"`
		Password string `default:""`
		Db       int   `default:"0"`
		PoolSize int   `default:"500"`
    }

}