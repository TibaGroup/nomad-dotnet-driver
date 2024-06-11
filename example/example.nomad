group "example" {
  task "dotnet_test" {
    driver = "dotnet"
    config {
      dll_path = "${NOMAD_TASK_DIR}/TestNomadTask.dll"
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
}