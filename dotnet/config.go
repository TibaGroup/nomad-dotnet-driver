package dotnet

type RuntimeConfig struct {
	ConfigProperties struct {
		Gc struct {
			Enable               bool   `json:"Server" codec:"enable"`
			Concurrent           bool   `codec:"concurrent"`
			HeapCount            uint16 `codec:"heap_count"`
			HeapHardLimit        uint64 `codec:"heap_limit"`
			HeapHardLimitPercent uint8  `codec:"heap_limit_percent"`
			NoAffinity           bool   `json:"NoAffinitize" codec:"no_affinity"`
			HeapAffinityMask     uint16 `json:"HeapAffinitizeMask" codec:"heap_affinity_mask"`
			HeapAffinityRanges   string `json:"HeapAffinitizeRanges" codec:"heap_affinity_ranges"`
			CpuGroup             bool   `codec:"cpu_group"`
			HighMemPercent       uint8  `codec:"high_mem_percent"`
			RetainVM             bool   `codec:"retain_vm"`
		} `json:"System.GC" codec:"gc"`
		Globalization struct {
			Invariant              bool `codec:"invariant"`
			UseNls                 bool `codec:"use_nls"`
			PredefinedCulturesOnly bool `codec:"predefined_cultures_only"`
		} `json:"System.Globalization" codec:"globalization"`
		Threading struct {
			ThreadPoolMinThreads   uint32 `json:"ThreadPool.MinThreads" codec:"min_threads"`
			ThreadPoolMaxThreads   uint32 `json:"ThreadPool.MaxThreads" codec:"max_threads"`
			UseWindowsThreadPool   bool   `json:"ThreadPool.UseWindowsThreadPool" codec:"windows_thread_pool"`
			AutoReleasePoolSupport bool   `json:"Thread.EnableAutoreleasePool" codec:"enable_autorelease_pool"`
		} `json:"System.Threading" codec:"threading"`
	} `json:"configProperties"`
}
