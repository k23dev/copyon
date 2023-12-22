# GoShareIt

Open a http server.

# Run it

The **path** arg by default is the same dir where the server ir running
The **port** arg by default is _5656_

    run goshareit [PATH] [PORT]

Now just open your browser and go to:

http://localhost:[PORT]

# Command to make windows exe

    GOOS=windows GOARCH=amd64 go build -o goshareit -ldflags "-w -s"