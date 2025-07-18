# Ascii Art Web

## Description
This project builds upon Ascii Art which is a CLI tool that takes input from the user in the terminal and outputs the text in different art styles (standard, shadow and thinkertoy).

Ascii Art Web implements a server and a GUI. The server runs on port :8080. The server is written in golang and for the frontend, HTML and CSS is used.

## Authors
Georgina and Masoud

## Usage: how to run
1. Make sure you have the [latest version of Go](https://go.dev/doc/install)\
Go to your terminal and run the following:
2. `git clone <repo-link>`
3. `cd ascii-art-web`
4. `go run main.go`
5. Follow link [http://localhost:8080](http://localhost:8080) displayed in terminal
6. On the webpage, enter text where prompted and click on the button

## Implementation details
In order to modify the the code from Ascii Art, the logic for the terminal arguments where removed. The rest of the code was written in the generateAscii function, where the the lines splitting at '\n' was moved into the function. The loadBannerfromURL function also largely stayed the same, except there was more error handling, including http status codes.

The new additions were due to the handlers for the two endpoints "/" and "/ascii-art". Each handler followed the method signature outlined in [the official docs](https://go.dev/doc/articles/wiki/). The first endpoint loads the html template, which contains a form. The asciiArtHandler passes the form, extracts the form values, uses these values to load the right banner from the API and generate the text in the correct font.