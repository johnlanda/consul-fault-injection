Kind = "mesh"
AllowEnablingPermissiveMutualTLS = false
TransparentProxy {
  # In a production env, traces should be sent through a terminating gateway
  # so that t-proxy can operate without exceptions.
  MeshDestinationsOnly = false
}