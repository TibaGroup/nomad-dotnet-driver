package dotnet

type RuntimeConfig struct {
	ConfigProperties struct {
		Gc struct {
			Enable               bool `json:"Server"`
			Concurrent           bool
			HeapCount            uint16
			HeapHardLimit        uint64
			HeapHardLimitPercent uint8
			NoAffinity           bool   `json:"NoAffinitize"`
			HeapAffinityMask     uint16 `json:"HeapAffinitizeMask"`
			HeapAffinityRanges   string `json:"HeapAffinitizeRanges"`
			CpuGroup             bool
			HighMemPercent       uint8
			RetainVM             bool
		} `json:"System.GC"`
		Globalization struct {
			Invariant              bool
			UseNls                 bool
			PredefinedCulturesOnly bool
		} `json:"System.Globalization"`
		Threading struct {
			ThreadPoolMinThreads   uint32 `json:"ThreadPool.MinThreads"`
			ThreadPoolMaxThreads   uint32 `json:"ThreadPool.MaxThreads"`
			UseWindowsThreadPool   bool   `json:"ThreadPool.UseWindowsThreadPool"`
			AutoReleasePoolSupport bool   `json:"Thread.EnableAutoreleasePool"`
		}
	} `json:"configProperties"`
}
