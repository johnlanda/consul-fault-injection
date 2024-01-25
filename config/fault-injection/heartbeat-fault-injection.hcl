Kind = "service-defaults"
Name = "heartbeat"
MutualTLSMode = "strict"
EnvoyExtensions = [
  {
    Name = "builtin/fault-injection"
    Arguments = {
      Config = {
        Abort = {
          HttpStatus = 500
          Percentage = 50
        },
      }
    }
  }
]
