job "example" {
  group "example" {
    task "dotnet_test" {
      driver = "dotnet"

      config {
        dll_path = "${NOMAD_TASK_DIR}/TestNomadTask.dll"
        runtime_version = "8.0.8"
        threading {
          min_threads = 10
          max_threads = 100
        }
        args = ["9090"]
      }

      artifact {
     source = "http://localhost/test_nomad_task.zip"
     destination = "local"
      }
    }
    task "dotnet_test_2" {
          driver = "dotnet"

          config {
            dll_path = "${NOMAD_TASK_DIR}/TestNomadTask.dll"
            runtime_version = "7.0.20"
            threading {
              min_threads = 10
              max_threads = 100
            }
            args = ["8080"]
          }

          artifact {
         source = "http://localhost/test_nomad_task.zip"
         destination = "local"
          }
        }
  }
}