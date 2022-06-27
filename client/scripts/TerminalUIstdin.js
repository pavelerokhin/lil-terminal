// TerminalUIconversational.js

class TerminalUIstdin {
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
      this.sendInput(data.domEvent.keyCode);
    });

    let that = this;
    this.socket.onmessage =  (data) => {
      that.clear()
      that.write(data.data);
    };
  }

  /**
   * Print something to terminal UI.
   */
  write(text) {
    this.terminal.write(text);
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
  }

  clear() {
    this.terminal.clear();
  }
}
