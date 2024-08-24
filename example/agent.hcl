# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

log_level = "TRACE"

client {
  options = {
    "driver.allowlist" = "dotnet"
  }
}

plugin "nomad-dotnet-driver" {
    config {
        sdk_path = "/usr/local/bin/dotnet"
    }
}
