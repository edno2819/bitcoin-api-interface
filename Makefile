install:
	go get "github.com/joho/godotenv"

start_bitcoin:
	bitcoind -daemon