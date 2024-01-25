Kind = "service-defaults"
Name = "heartbeat"
MutualTLSMode = "strict"
EnvoyExtensions = [
  {
    Name = "builtin/fault-injection"
    Arguments = {
      Config = {
        Delay = {
          Duration = 500
          Percentage = 50
        }
      }
    }
  }
]
