package main

import (
	"syscall/js" //using the syscall/js package

	"github.com/skip2/go-qrcode" // using the qrcode package
)

func main() {
	js.Global().Set("MakeQRCODE", js.FuncOf(MakeQRCODE)) //registering the function

	select {} //infinite loop
}

func MakeQRCODE(this js.Value, args []js.Value) interface{} {
	document := js.Global().Get("document")                               // selecting the DOM document
	data := document.Call("getElementById", "data").Get("value").String() //equivalent to document.getElementById("data").value
	qrimage, err := qrcode.Encode(data, qrcode.Medium, 256)               //generating the QR code
	if err != nil {
		panic(err) //panic if there is an error
	}
	dst := js.Global().Get("Uint8Array").New(len(qrimage)) //creating a new Uint8Array

	js.CopyBytesToJS(dst, qrimage) //copying the QR code to the Uint8Array

	return dst //returning the Uint8Array
}
