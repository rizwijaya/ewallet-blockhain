.PHONY : format install build

#Website/Restfull API
run:
	@echo "Running server..."
	go run main.go

init:
	@echo "Initializing dependencies..."
	go mod init
	go mod tidy
	
install:
	@echo "Downloading dependencies..."
	go mod download

build:
	@echo "building binary..."
	go build main.go

start:
	@echo "Starting server..."
	./main

clean:
	@echo "Cleaning..."
	rm -rf main.exe
# live reload using nodemon: npm -g i nodemon
run-nodemon:
	@echo "Running server with nodemon..."
	nodemon --exec go run main.go

sol:
	solcjs --optimize --abi ./app/contracts/wallet.sol -o build
	solcjs --optimize --bin ./app/contracts/wallet.sol -o build
	abigen --abi=./build/app_contracts_wallet_sol_Wallet.abi --bin=./build/app_contracts_wallet_sol_Wallet.bin --pkg=api --out=./app/contracts/wallet.go