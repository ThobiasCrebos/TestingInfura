you need to have the solidity code for contract 

Generate Ethereum Keystore Wallet using Golang
- chcek the keystore file

need to create a version of contract that go understands
- first go to the solididty documentation and install a solidity compiler for your machine
-install protoc for you machine

to complie the solidity file needed for creatin go contract
- run "solc --bin --abi contract/SimpleStorage.sol -o build"

to complie the solidity file needed for interact with deployed smart contract
- run "abigen --bin=build/SimpleStorage.bin --abi=build/SimpleStorage.abi --pkg=SimpleStorage --out=gen/SimpleStorage.go"