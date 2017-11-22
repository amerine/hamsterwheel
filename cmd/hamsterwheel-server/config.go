package main

type config struct {
	//DatabaseURL is
	DatabaseURL string `env:"DATABASE_URL,required"`
}
