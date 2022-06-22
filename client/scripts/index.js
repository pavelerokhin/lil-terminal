// index.js
debugger;

// IMPORTANT: Make sure you replace this address with your server address.

const serverAddress = "http://localhost:1323/ws";

//Server sandbox available at https://codesandbox.io/s/web-terminal-tutorial-server-g2ihu

function connectToSocket() {
  // return new Promise(res => {
  //   const socket = io(serverAddress);
  //   res(socket);
  // });

  let loc = window.location;
  let uri = "ws:";

  if (loc.protocol === "https:") {
    uri = "wss:";
  }
  uri += "//127.0.0.1:1323/ws";

  return new WebSocket(uri);
}

function startTerminal(container, socket) {
  // Create an xterm.js instance (TerminalUI class is a wrapper with some utils. Check that file for info.)
  const terminal = new TerminalUI(socket);

  // Attach created terminal to a DOM element.
  terminal.attachTo(container);

  // When terminal attached to DOM, start listening for input, output events.
  // Check TerminalUI startListening() function for details.
  terminal.startListening();
}

function start() {
  const container = document.getElementById("terminal-container");
  // Connect to socket and when it is available, start terminal.
  let ws = connectToSocket()
  startTerminal(container, ws);
}

// Better to start on DOMContentLoaded. So, we know terminal-container is loaded
start();
