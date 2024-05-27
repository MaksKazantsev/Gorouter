package gorouter

// Middleware is a function, that accepts HandlerFunc and return another HandlerFunc
type Middleware func(hFunc HandlerFunc) HandlerFunc
