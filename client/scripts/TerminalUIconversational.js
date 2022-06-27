// TerminalUIconversational.js

// TODO: add backspace functionality

class TerminalUI {
  constructor(socket) {
    this.terminal = new Terminal();

    /* You can make your terminals colorful :) */
    this.terminal.setOption("theme", {
      background: "#202B33",
      foreground: "#F5F8FA"
    });

    this.socket = socket;
  }

  /**
   * Attach event listeners for terminal UI and socket.io client
   */
  startListening() {
    this.terminal.onKey(data => {
      if (!isControl(data.domEvent.keyCode)) {
        this.write(data.key);
      }

      this.sendInput(data.domEvent.keyCode);
    });

    let that = this;
    this.socket.onmessage =  (data) => {
      // When there is data from PTY on server, print that on Terminal.
      debugger;
      that.promptA()
      that.write(data.data);
      that.promptQ()
    };
  }

  /**
   * Print something to terminal UI.
   */
  write(text) {
    this.terminal.write(text);
  }

  /**
   * Utility function to print new line on terminal.
   */
  promptA() {
    this.terminal.write(`\r\n> `);
  }

  promptQ() {
    this.terminal.write(`\r\n$ `);
  }

  /**
   * Send whatever you type in Terminal UI to PTY process in server.
   */
  sendInput(input) {
    this.socket.send(input);
  }

  /**
   *
   * @param {HTMLElement} container HTMLElement where xterm can attach terminal ui instance.
   */
  attachTo(container) {
    this.terminal.open(container);
    // Default text to display on terminal.
    this.terminal.write("Lil Terminal assist");
    this.terminal.write("");
    this.promptQ();
  }

  clear() {
    this.terminal.clear();
  }
}


function isControl(key) {
  return key === 8 || // backspace
      key === 13 || // enter
      key === 37 || // arrow left
      key === 38 || // arrow up
      key === 39 || // arrow right
      key === 40 // arrow down
}
