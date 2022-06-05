# Design of websocket implementation

Draw description of the logic of the implementation, TODO: More descriptive documentation

It will be tested with [tiingo api](https://api.tiingo.com/documentation/websockets)

In the main.go file, in the init function, it will be initialized the **websocket api**, and connected to the url of the websocket. This happened here because, in the first moment when the project is up, it will be necessary get the data

It will be created the way for subscribe more users to the application, to be notified in the callback of the websocket api 