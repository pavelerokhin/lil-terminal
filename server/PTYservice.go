package server

//type SocketService struct {
//	Socket
//	PTY
//}
//
//func GetSocketService() *SocketService {
//	return &SocketService{}
//}
//
//func (s *SocketService) AttachServer() error {
//
//}

//startPtyProcess() {
//this.ptyProcess = pty.spawn(this.shell, [], {
//name: "xterm-color",
//cwd: process.env.HOME, // Which path should terminal start
//env: process.env // Pass environment variables
//});
//
//// Add a "data" event listener.
//this.ptyProcess.on("data", data => {
//// Whenever terminal generates any data, send that output to socket.io client to display on UI
//this.sendToClient(data);
//});
//}
//
///**
// * Use this function to send in the input to Pseudo Terminal process.
// * @param {*} data Input from user like command sent from terminal UI
// */
//
//write(data) {
//this.ptyProcess.write(data);
//}
//
//sendToClient(data) {
//// Emit data to socket.io client in an event "output"
//this.socket.emit("output", data);
//}
//}

//func main() {
//	c := exec.Command("grep", "--color=auto", "bar")
//	f, err := pty.Start(c)
//	if err != nil {
//		panic(err)
//	}
//
//	go func() {
//		f.Write([]byte("foo\n"))
//		f.Write([]byte("bar\n"))
//		f.Write([]byte("baz\n"))
//		f.Write([]byte{4}) // EOT
//	}()
//	io.Copy(os.Stdout, f)
//}
