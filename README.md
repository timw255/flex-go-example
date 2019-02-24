# Flex Go Example

An example Flex service that's used to test [Flex Go](https://github.com/timw255/flex-go).

This project is an experiment. The goal is to write a Flex services in an entirely different language, compile it to WebAssembly, and then run it in Kinvey.

We're not there yet...

## WebAssembly

This example compiles to WebAssembly and can be run in a browser.

If you'd like to test that out:

* Clone or download the project
* Run `npm install`
* Run `npm run build` to compile the code to WebAssembly (main.wasm)
* Run `npm run web` to start the test web server
* Navigate to [http://localhost:8080](http://localhost:8080)
* Click the "Run" button to send a test task to the service