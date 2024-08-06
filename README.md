# http-client-cli
perform simple request on CLI built with Cobra-CLI

### Supported Commands
- `go run main.go router post --url "http://localhost:4000/example/login" --data "{}"`
- `go run main.go router get  --url ""`

### json body format
- `"{username: jondoe, password:123doe}"`
- or drag and drop a json file on the terminal for the data flag like this `go run main.go router post --data C:\User\example\Desktop\cred.json --url []"`
- you can choose to type in the path to a `.json` file
  
### TODO
- [] support `Get` method with extra params like headers
- [] implement `Delete` request
- [] send `Update` request