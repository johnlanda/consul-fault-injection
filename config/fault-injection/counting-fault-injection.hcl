Kind = "service-defaults"
Name = "client"
MutualTLSMode = "strict"
EnvoyExtensions = [
  {
    Name = "builtin/fault-injection"
    Arguments = {
      Config = {
        # Comment out this block to disable the abort injection
        Abort = {
          HttpStatus = 500
          Percentage = 50
        },
        # Uncomment to enable delay injection
#        Delay = {
#          Duration = "500"
#          Percentage = 50
#        }
      }
    }
  }
]
