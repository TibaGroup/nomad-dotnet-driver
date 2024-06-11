package dotnet

type GcConfig struct {
	Enable               *bool   `codec:"enable"`
	Concurrent           *bool   `codec:"concurrent"`
	HeapCount            *uint16 `codec:"heap_count"`
	HeapHardLimit        *uint64 `codec:"heap_limit"`
	HeapHardLimitPercent *uint8  `codec:"heap_limit_percent"`
	NoAffinity           *bool   `codec:"no_affinity"`
	HeapAffinityMask     *uint16 `codec:"heap_affinity_mask"`
	HeapAffinityRanges   *string `codec:"heap_affinity_ranges"`
	CpuGroup             *bool   `codec:"cpu_group"`
	HighMemPercent       *uint8  `codec:"high_mem_percent"`
	RetainVM             *bool   `codec:"retain_vm"`
}

type GlobalizationConfig struct {
	Invariant              *bool `codec:"invariant"`
	UseNls                 *bool `codec:"use_nls"`
	PredefinedCulturesOnly *bool `codec:"predefined_cultures_only"`
}

type ThreadingConfig struct {
	ThreadPoolMinThreads   *uint32 `codec:"min_threads"`
	ThreadPoolMaxThreads   *uint32 `codec:"max_threads"`
	UseWindowsThreadPool   *bool   `codec:"windows_thread_pool"`
	AutoReleasePoolSupport *bool   `codec:"enable_autorelease_pool"`
}

type ConfigFile struct {
	RuntimeOptions struct {
		ConfigProperties struct {
			Enable                 *bool   `json:"System.GC.Server,omitempty"`
			Concurrent             *bool   `json:"System.GC.Concurrent,omitempty"`
			HeapCount              *uint16 `json:"System.GC.HeapCount,omitempty"`
			HeapHardLimit          *uint64 `json:"System.GC.HeapHardLimit,omitempty,"`
			HeapHardLimitPercent   *uint8  `json:"System.GC.HeapHardLimitPercent,omitempty"`
			NoAffinity             *bool   `json:"System.GC.NoAffinitize,omitempty"`
			HeapAffinityMask       *uint16 `json:"System.GC.HeapAffinitizeMask,omitempty"`
			HeapAffinityRanges     *string `json:"System.GC.HeapAffinitizeRanges,omitempty"`
			CpuGroup               *bool   `json:"System.GC.CpuGroup,omitempty"`
			HighMemPercent         *uint8  `json:"System.GC.HighMemPercent,omitempty"`
			RetainVM               *bool   `json:"System.GC.RetainVM,omitempty"`
			InvariantCulture       *bool   `json:"System.Globalization.InvariantCulture,omitempty"`
			UseNls                 *bool   `json:"System.Globalization.UseNls,omitempty"`
			PredefinedCulturesOnly *bool   `json:"System.Globalization.PredefinedCulturesOnly,omitempty"`
			ThreadPoolMinThreads   *uint32 `json:"System.Threading.ThreadPool.MinThreads,omitempty"`
			ThreadPoolMaxThreads   *uint32 `json:"System.Threading.ThreadPool.MaxThreads,omitempty"`
			UseWindowsThreadPool   *bool   `json:"System.Threading.ThreadPool.UseWindowsThreadPool,omitempty"`
			AutoReleasePoolSupport *bool   `json:"System.Threading.Thread.EnableAutoreleasePool,omitempty"`
		} `json:"configProperties"`
	} `json:"runtimeOptions"`
}

func addGcConfig(hcl *GcConfig, conf *ConfigFile) {
	if hcl == nil {
		return
	}
	conf.RuntimeOptions.ConfigProperties.Enable = hcl.Enable
	conf.RuntimeOptions.ConfigProperties.Concurrent = hcl.Concurrent
	conf.RuntimeOptions.ConfigProperties.HeapCount = hcl.HeapCount
	conf.RuntimeOptions.ConfigProperties.HeapHardLimit = hcl.HeapHardLimit
	conf.RuntimeOptions.ConfigProperties.HeapHardLimitPercent = hcl.HeapHardLimitPercent
	conf.RuntimeOptions.ConfigProperties.HeapHardLimit = hcl.HeapHardLimit
	conf.RuntimeOptions.ConfigProperties.RetainVM = hcl.RetainVM
	conf.RuntimeOptions.ConfigProperties.NoAffinity = hcl.NoAffinity
	conf.RuntimeOptions.ConfigProperties.HeapAffinityMask = hcl.HeapAffinityMask
	conf.RuntimeOptions.ConfigProperties.HeapAffinityRanges = hcl.HeapAffinityRanges
	conf.RuntimeOptions.ConfigProperties.CpuGroup = hcl.CpuGroup
	conf.RuntimeOptions.ConfigProperties.HighMemPercent = hcl.HighMemPercent
	conf.RuntimeOptions.ConfigProperties.RetainVM = hcl.RetainVM
}

func addGlobalizationConfig(hcl *GlobalizationConfig, conf *ConfigFile) {
	if hcl == nil {
		return
	}
	conf.RuntimeOptions.ConfigProperties.InvariantCulture = hcl.Invariant
	conf.RuntimeOptions.ConfigProperties.UseNls = hcl.UseNls
	conf.RuntimeOptions.ConfigProperties.PredefinedCulturesOnly = hcl.PredefinedCulturesOnly
}

func addThreadingConfig(hcl *ThreadingConfig, conf *ConfigFile) {
	if hcl == nil {
		return
	}
	conf.RuntimeOptions.ConfigProperties.ThreadPoolMinThreads = hcl.ThreadPoolMinThreads
	conf.RuntimeOptions.ConfigProperties.ThreadPoolMaxThreads = hcl.ThreadPoolMaxThreads
	conf.RuntimeOptions.ConfigProperties.UseWindowsThreadPool = hcl.UseWindowsThreadPool
	conf.RuntimeOptions.ConfigProperties.AutoReleasePoolSupport = hcl.AutoReleasePoolSupport
}
