```go
GcConfig Fields to .NET Core GC Configuration

type GcConfig struct {
Enable               *bool
Concurrent           *bool
HeapCount            *uint16
HeapHardLimit        *uint64
HeapHardLimitPercent *uint8
NoAffinity           *bool
HeapAffinityMask     *uint16
HeapAffinityRanges   *string
CpuGroup             *bool
HighMemPercent       *uint8
RetainVM             *bool
}

```
### Field Descriptions

* [Enable](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#workstation-vs-server)
  * .NET Core Equivalent: System.GC.Server
    * Description: Controls whether the server garbage collector is enabled or not.
    * Configuration: Set System.GC.Server to true or false in the .runtimeconfig.json.
    * Example: true

* [Concurrent](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#background-gc)
  * .NET Core Equivalent: System.GC.Concurrent
    * Description: Controls whether the GC runs concurrently or non-concurrently with the application.
    * Configuration: Set System.GC.Concurrent to true or false in the .runtimeconfig.json.
    * Example: true

* [HeapCount](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#heap-count)
  * .NET Core Equivalent: System.GC.HeapCount
    * Description: Limits the number of heaps created by the server garbage collector.
    * Configuration: Set System.GC.HeapCount to a decimal value in the .runtimeconfig.json file.
    * Example: 16

* [HeapHardLimit](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#heap-limit)
  * .NET Core Equivalent: System.GC.HeapHardLimit
    * Description: Specifies the maximum commit size, in bytes, for the GC heap and GC bookkeeping.
    * Configuration: Set System.GC.HeapHardLimit to a decimal value in the .runtimeconfig.json file. 
    * Example: 209715200 (200 mebibytes)

* [HeapHardLimitPercent](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#heap-limit-percent)
  * .NET Core Equivalent: System.GC.HeapHardLimitPercent
    * Description: Specifies the allowable GC heap usage as a percentage of the total physical memory.
    * Configuration: Set System.GC.HeapHardLimitPercent to a decimal value in the .runtimeconfig.json file.
    * Example: 30

* [NoAffinity](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#affinitize)
  * .NET Core Equivalent: System.GC.NoAffinitize
    * Description: Specifies whether to affinitize garbage collection threads with processors.
    * Configuration: Set System.GC.NoAffinitize to true or false.
    * Example: true

* [HeapAffinityMask](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#affinitize-mask)
  * .NET Core Equivalent: System.GC.HeapAffinitizeMask
    * Description: Specifies the exact processors that garbage collector threads should use.
    * Configuration: Set System.GC.HeapAffinitizeMask to a decimal value in the .runtimeconfig.json file.
    * Example: 1023

* [HeapAffinityRanges](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#affinitize-ranges)
  * .NET Core Equivalent: System.GC.HeapAffinitizeRanges
    * Description: Specifies the list of processors to use for garbage collector threads.
    * Configuration: Set System.GC.HeapAffinitizeRanges as comma-separated list of processor numbers or ranges of processor numbers in .runtimeconfig.json.
    * Example: Unix example: "1-10,12,50-52,70"

* [CpuGroup](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#cpu-groups)
  * .NET Core Equivalent: System.GC.CpuGroup
    * Description: Configures whether the garbage collector uses CPU groups or not.
    * Configuration: Set System.GC.CpuGroup to true or false in the .runtimeconfig.json file.
    * Example: true

* [HighMemPercent](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#high-memory-percent)
  * .NET Core Equivalent: System.GC.HighMemoryPercent
    * Description: Memory load is indicated by the percentage of physical memory in use.
    * Configuration: Set System.GC.HighMemoryPercent to a decimal value in the .runtimeconfig.json file.
    * Example: 75

* [RetainVM](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/garbage-collector#retain-vm)
  * .NET Core Equivalent: System.GC.RetainVM
    * Description: Configures whether segments that should be deleted are put on a standby list for future use or are released back to the operating system (OS).
    * Configuration: Set System.GC.RetainVM to true or false in the .runtimeconfig.json file.
    * Example: true

  
```go
GlobalizationConfig to.NET Core Globalization Configuration

type GlobalizationConfig struct {
Invariant              *bool
UseNls                 *bool
PredefinedCulturesOnly *bool
}

```
### Field Descriptions

* [Invariant](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/globalization#invariant-mode)
  * .NET Core Equivalent: System.Globalization.Invariant
    * Description: Determines whether a .NET Core app runs in globalization-invariant mode without access to culture-specific data and behavior.
    * Configuration: Set System.Globalization.Invariant to true or false in the .runtimeconfig.json file.
    * Example: true

* [UseNls](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/globalization#nls)
  * .NET Core Equivalent: System.Globalization.UseNls
    * Description: Determines whether .NET uses National Language Support (NLS) or International Components for Unicode (ICU) globalization APIs for Windows apps.
    * Configuration: Set System.Globalization.UseNls to true or false in the .runtimeconfig.json file.
    * Example: true

* [PredefinedCulturesOnly](https://learn.microsoft.com/en-us/dotnet/core/runtime-config/globalization#predefined-cultures)
  * .NET Core Equivalent: System.Globalization.PredefinedCulturesOnly
    * Description: Configures whether apps can create cultures other than the invariant culture.
    * Configuration: Set System.Globalization.PredefinedCulturesOnly to true or false in the .runtimeconfig.json file.
    * Example: true
