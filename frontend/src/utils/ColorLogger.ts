class ColorLogger {
  private red = '\x1b[31m'
  private green = '\x1b[32m'
  private yellow = '\x1b[33m'
  private blue = '\x1b[34m'
  private reset = '\x1b[0m'

  fatal(message: string | undefined) {
    console.error(`${this.red}[FATAL]: ${this.reset} ${message}`)
  }

  info(message: string) {
    console.info(`${this.green}[INFO]: ${this.reset} ${message}`)
  }

  error(message: string) {
    console.error(`${this.red}[ERROR]: ${this.reset} ${message}`)
  }

  debug(message: string) {
    console.debug(`${this.blue}[DEBUG]: ${this.reset} ${message}`)
  }

  warning(message: string) {
    console.info(`${this.yellow}[WARNING]: ${this.reset} ${message}`)
  }
}

export const clogger = new ColorLogger()
