package classifier

import (
	"math"
	"reflect"
	"testing"
)

const (
	testError = "test error"
	testValue = "test value"
)

func ptrToString(s string) *string {
	return &s
}

func Test_stringValueOrError(t *testing.T) {
	type args struct {
		err *string
		val *string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no error, no value",
			args: args{
				err: nil,
				val: nil,
			},
			want: Undefined,
		},
		{
			name: "has error, no value",
			args: args{
				err: ptrToString(testError),
				val: nil,
			},
			want: testError,
		},
		{
			name: "no error, has value",
			args: args{
				err: nil,
				val: ptrToString(testValue),
			},
			want: testValue,
		},
		{
			name: "has error, has value",
			args: args{
				err: ptrToString(testError),
				val: ptrToString(testValue),
			},
			want: testError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringValueOrError(tt.args.err, tt.args.val); got != tt.want {
				t.Errorf("stringValueOrError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_standardStringValueOrError(t *testing.T) {
	type args struct {
		err *string
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no error, no value",
			args: args{
				err: nil,
				val: "",
			},
			want: "",
		},
		{
			name: "has error, no value",
			args: args{
				err: ptrToString(testError),
				val: "",
			},
			want: testError,
		},
		{
			name: "no error, has value",
			args: args{
				err: nil,
				val: testValue,
			},
			want: testValue,
		},
		{
			name: "has error, has value",
			args: args{
				err: ptrToString(testError),
				val: testValue,
			},
			want: testError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := standardStringValueOrError(tt.args.err, tt.args.val); got != tt.want {
				t.Errorf("standardStringValueOrError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hardwareRelatedCompatibility(t *testing.T) {
	deviceMemory0, deviceMemory1 := 0, 1

	type args struct {
		a Fingerprint
		b Fingerprint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic, empty",
			want: 1,
		},
		{
			name: "special case (a has errors)",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Error: ptrToString(testError),
					},
					Languages: Languages{
						Error: ptrToString(testError),
					},
					ColorDepth: ColorDepth{
						Error: ptrToString(testError),
					},
					DeviceMemory: DeviceMemory{
						Error: ptrToString(testError),
					},
					ScreenResolution: ScreenResolution{
						Error: ptrToString(testError),
					},
					AvailableScreenResolution: AvailableScreenResolution{
						Error: ptrToString(testError),
					},
					HardwareConcurrency: HardwareConcurrency{
						Error: ptrToString(testError),
					},
					TimezoneOffset: TimezoneOffset{
						Error: ptrToString(testError),
					},
					Timezone: Timezone{
						Error: ptrToString(testError),
					},
					SessionStorage: SessionStorage{
						Error: ptrToString(testError),
					},
					LocalStorage: LocalStorage{
						Error: ptrToString(testError),
					},
					IndexedDB: IndexedDB{
						Error: ptrToString(testError),
					},
					OpenDatabase: OpenDatabase{
						Error: ptrToString(testError),
					},
					CPUClass: CPUClass{
						Error: ptrToString(testError),
					},
					Platform: Platform{
						Error: ptrToString(testError),
					},
					Plugins: Plugins{
						Error: ptrToString(testError),
					},
					Canvas: Canvas{
						Error: ptrToString(testError),
					},
					TouchSupport: TouchSupport{
						Error: ptrToString(testError),
					},
					Fonts: Fonts{
						Error: ptrToString(testError),
					},
					Audio: Audio{
						Error: ptrToString(testError),
					},
					PluginsSupport: PluginsSupport{
						Error: ptrToString(testError),
					},
					ProductSub: ProductSub{
						Error: ptrToString(testError),
					},
					EmptyEvalLength: EmptyEvalLength{
						Error: ptrToString(testError),
					},
					ErrorFF: ErrorFF{
						Error: ptrToString(testError),
					},
					Vendor: Vendor{
						Error: ptrToString(testError),
					},
					Chrome: Chrome{
						Error: ptrToString(testError),
					},
					CookiesEnabled: CookiesEnabled{
						Error: ptrToString(testError),
					},
				},
			},
			want: 0,
		},
		{
			name: "special case (b has errors)",
			args: args{
				b: Fingerprint{
					OsCPU: OsCPU{
						Error: ptrToString(testError),
					},
					Languages: Languages{
						Error: ptrToString(testError),
					},
					ColorDepth: ColorDepth{
						Error: ptrToString(testError),
					},
					DeviceMemory: DeviceMemory{
						Error: ptrToString(testError),
					},
					ScreenResolution: ScreenResolution{
						Error: ptrToString(testError),
					},
					AvailableScreenResolution: AvailableScreenResolution{
						Error: ptrToString(testError),
					},
					HardwareConcurrency: HardwareConcurrency{
						Error: ptrToString(testError),
					},
					TimezoneOffset: TimezoneOffset{
						Error: ptrToString(testError),
					},
					Timezone: Timezone{
						Error: ptrToString(testError),
					},
					SessionStorage: SessionStorage{
						Error: ptrToString(testError),
					},
					LocalStorage: LocalStorage{
						Error: ptrToString(testError),
					},
					IndexedDB: IndexedDB{
						Error: ptrToString(testError),
					},
					OpenDatabase: OpenDatabase{
						Error: ptrToString(testError),
					},
					CPUClass: CPUClass{
						Error: ptrToString(testError),
					},
					Platform: Platform{
						Error: ptrToString(testError),
					},
					Plugins: Plugins{
						Error: ptrToString(testError),
					},
					Canvas: Canvas{
						Error: ptrToString(testError),
					},
					TouchSupport: TouchSupport{
						Error: ptrToString(testError),
					},
					Fonts: Fonts{
						Error: ptrToString(testError),
					},
					Audio: Audio{
						Error: ptrToString(testError),
					},
					PluginsSupport: PluginsSupport{
						Error: ptrToString(testError),
					},
					ProductSub: ProductSub{
						Error: ptrToString(testError),
					},
					EmptyEvalLength: EmptyEvalLength{
						Error: ptrToString(testError),
					},
					ErrorFF: ErrorFF{
						Error: ptrToString(testError),
					},
					Vendor: Vendor{
						Error: ptrToString(testError),
					},
					Chrome: Chrome{
						Error: ptrToString(testError),
					},
					CookiesEnabled: CookiesEnabled{
						Error: ptrToString(testError),
					},
				},
			},
			want: 0,
		},
		{
			name: "special case (a and b have errors)",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Error: ptrToString(testError),
					},
					Languages: Languages{
						Error: ptrToString(testError),
					},
					ColorDepth: ColorDepth{
						Error: ptrToString(testError),
					},
					DeviceMemory: DeviceMemory{
						Error: ptrToString(testError),
					},
					ScreenResolution: ScreenResolution{
						Error: ptrToString(testError),
					},
					AvailableScreenResolution: AvailableScreenResolution{
						Error: ptrToString(testError),
					},
					HardwareConcurrency: HardwareConcurrency{
						Error: ptrToString(testError),
					},
					TimezoneOffset: TimezoneOffset{
						Error: ptrToString(testError),
					},
					Timezone: Timezone{
						Error: ptrToString(testError),
					},
					SessionStorage: SessionStorage{
						Error: ptrToString(testError),
					},
					LocalStorage: LocalStorage{
						Error: ptrToString(testError),
					},
					IndexedDB: IndexedDB{
						Error: ptrToString(testError),
					},
					OpenDatabase: OpenDatabase{
						Error: ptrToString(testError),
					},
					CPUClass: CPUClass{
						Error: ptrToString(testError),
					},
					Platform: Platform{
						Error: ptrToString(testError),
					},
					Plugins: Plugins{
						Error: ptrToString(testError),
					},
					Canvas: Canvas{
						Error: ptrToString(testError),
					},
					TouchSupport: TouchSupport{
						Error: ptrToString(testError),
					},
					Fonts: Fonts{
						Error: ptrToString(testError),
					},
					Audio: Audio{
						Error: ptrToString(testError),
					},
					PluginsSupport: PluginsSupport{
						Error: ptrToString(testError),
					},
					ProductSub: ProductSub{
						Error: ptrToString(testError),
					},
					EmptyEvalLength: EmptyEvalLength{
						Error: ptrToString(testError),
					},
					ErrorFF: ErrorFF{
						Error: ptrToString(testError),
					},
					Vendor: Vendor{
						Error: ptrToString(testError),
					},
					Chrome: Chrome{
						Error: ptrToString(testError),
					},
					CookiesEnabled: CookiesEnabled{
						Error: ptrToString(testError),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Error: ptrToString(testError),
					},
					Languages: Languages{
						Error: ptrToString(testError),
					},
					ColorDepth: ColorDepth{
						Error: ptrToString(testError),
					},
					DeviceMemory: DeviceMemory{
						Error: ptrToString(testError),
					},
					ScreenResolution: ScreenResolution{
						Error: ptrToString(testError),
					},
					AvailableScreenResolution: AvailableScreenResolution{
						Error: ptrToString(testError),
					},
					HardwareConcurrency: HardwareConcurrency{
						Error: ptrToString(testError),
					},
					TimezoneOffset: TimezoneOffset{
						Error: ptrToString(testError),
					},
					Timezone: Timezone{
						Error: ptrToString(testError),
					},
					SessionStorage: SessionStorage{
						Error: ptrToString(testError),
					},
					LocalStorage: LocalStorage{
						Error: ptrToString(testError),
					},
					IndexedDB: IndexedDB{
						Error: ptrToString(testError),
					},
					OpenDatabase: OpenDatabase{
						Error: ptrToString(testError),
					},
					CPUClass: CPUClass{
						Error: ptrToString(testError),
					},
					Platform: Platform{
						Error: ptrToString(testError),
					},
					Plugins: Plugins{
						Error: ptrToString(testError),
					},
					Canvas: Canvas{
						Error: ptrToString(testError),
					},
					TouchSupport: TouchSupport{
						Error: ptrToString(testError),
					},
					Fonts: Fonts{
						Error: ptrToString(testError),
					},
					Audio: Audio{
						Error: ptrToString(testError),
					},
					PluginsSupport: PluginsSupport{
						Error: ptrToString(testError),
					},
					ProductSub: ProductSub{
						Error: ptrToString(testError),
					},
					EmptyEvalLength: EmptyEvalLength{
						Error: ptrToString(testError),
					},
					ErrorFF: ErrorFF{
						Error: ptrToString(testError),
					},
					Vendor: Vendor{
						Error: ptrToString(testError),
					},
					Chrome: Chrome{
						Error: ptrToString(testError),
					},
					CookiesEnabled: CookiesEnabled{
						Error: ptrToString(testError),
					},
				},
			},
			want: 1,
		},
		{
			name: "only a has ColorDepth",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
				},
			},
			want: 20. / 21,
		},
		{
			name: "only b has ColorDepth",
			args: args{
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
				},
			},
			want: 20. / 21,
		},
		{
			name: "ColorDepth",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 0,
					},
				},
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
				},
			},
			want: 20. / 21,
		},
		{
			name: "ColorDepth, only a has DeviceMemory",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 0,
					},
					DeviceMemory: DeviceMemory{
						Value: &deviceMemory0,
					},
				},
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
				},
			},
			want: 19. / 21,
		},
		{
			name: "ColorDepth, only b has DeviceMemory",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 0,
					},
				},
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
					DeviceMemory: DeviceMemory{
						Value: &deviceMemory0,
					},
				},
			},
			want: 19. / 21,
		},
		{
			name: "ColorDepth, DeviceMemory",
			args: args{
				a: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 0,
					},
					DeviceMemory: DeviceMemory{
						Value: &deviceMemory0,
					},
				},
				b: Fingerprint{
					ColorDepth: ColorDepth{
						Value: 24,
					},
					DeviceMemory: DeviceMemory{
						Value: &deviceMemory1,
					},
				},
			},
			want: 19. / 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hardwareRelatedCompatibility(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("hardwareRelatedCompatibility() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_max(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "a > b",
			args: args{
				a: math.MaxInt64,
				b: math.MinInt64,
			},
			want: math.MaxInt64,
		},
		{
			name: "a < b",
			args: args{
				a: math.MinInt64,
				b: math.MaxInt64,
			},
			want: math.MaxInt64,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := max(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_similarity(t *testing.T) {
	type args struct {
		a Fingerprint
		b Fingerprint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "basic, empty",
			want: 1,
		},
		{
			name: `special case (a.OsCPU ← "", b.OsCPU ← Undefined)`,
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString(""),
					},
				},
			},
			want: 8. / 9,
		},
		{
			name: `special case (a.OsCPU ← Undefined, b.OsCPU ← "")`,
			args: args{
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString(""),
					},
				},
			},
			want: 8. / 9,
		},
		{
			name: "only a has OsCPU",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Linux x86_64"),
					},
				},
			},
			want: 8. / 9,
		},
		{
			name: "only b has OsCPU",
			args: args{
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Linux x86_64"),
					},
				},
			},
			want: 8. / 9,
		},
		{
			name: "OsCPU",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
				},
			},
			want: 125. / 126,
		},
		{
			name: "OsCPU, only a has Timezone",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
				},
			},
			want: 37. / 42,
		},
		{
			name: "OsCPU, only b has Timezone",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
				},
			},
			want: 37. / 42,
		},
		{
			name: "OsCPU, Timezone",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/Warsaw"),
					},
				},
			},
			want: 1541. / 1638,
		},
		{
			name: "OsCPU, Timezone, only a has CPUClass",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
					CPUClass: CPUClass{
						Value: ptrToString("x86"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/Warsaw"),
					},
				},
			},
			want: 151. / 182,
		},
		{
			name: "OsCPU, Timezone, only b has CPUClass",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/Warsaw"),
					},
					CPUClass: CPUClass{
						Value: ptrToString("x86"),
					},
				},
			},
			want: 151. / 182,
		},
		{
			name: "OsCPU, Timezone, CPUClass",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.2"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/London"),
					},
					CPUClass: CPUClass{
						Value: ptrToString("x86"),
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: ptrToString("Windows NT 6.3"),
					},
					Timezone: Timezone{
						Value: ptrToString("Europe/Warsaw"),
					},
					CPUClass: CPUClass{
						Value: ptrToString("Alpha"),
					},
				},
			},
			want: 151. / 182,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarity(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("similarity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExperimentalInMemory_Do(t *testing.T) {
	var (
		fZRMtpX = Fingerprint{
			VisitorID: "fZRMtpX",
			OsCPU: OsCPU{
				Value: ptrToString("Windows NT 6.2"),
			},
		}
		screenResolution = [2]int{2560, 1080}
		truth            = true
	)
	tests := []struct {
		name         string
		fingerprints map[string][]Fingerprint
		f            Fingerprint
		want         Fingerprint
		want1        bool
	}{
		{
			name:  "fingerprints and f are empty",
			want:  Fingerprint{},
			want1: false,
		},
		{
			name: "fingerprints is empty",
			f: Fingerprint{
				OsCPU: OsCPU{
					Value: ptrToString("Linux x86_64"),
				},
			},
			want:  Fingerprint{},
			want1: false,
		},
		{
			name: `f is similar to fingerprints["fZRMtpX"]`,
			fingerprints: map[string][]Fingerprint{
				"fZRMtpX": {
					fZRMtpX,
				},
			},
			f: Fingerprint{
				OsCPU: OsCPU{
					Value: ptrToString("Windows NT 6.3"),
				},
			},
			want:  fZRMtpX,
			want1: true,
		},
		{
			name: "f isn't similar to any other fingerprint",
			fingerprints: map[string][]Fingerprint{
				"fZRMtpX": {
					fZRMtpX,
				},
			},
			f: Fingerprint{
				OsCPU: OsCPU{
					Value: ptrToString("Linux x86_64"),
				},
				Languages: Languages{
					Value: [][]string{
						{
							"en-US",
						},
						{
							"en-US",
							"en",
						},
					},
				},
				Timezone: Timezone{
					Value: ptrToString("Europe/Warsaw"),
				},
				CPUClass: CPUClass{
					Value: nil,
				},
				Platform: Platform{
					Value: "Linux x86_64",
				},
				Plugins: Plugins{
					Value: []PluginsValue{},
				},
				Fonts: Fonts{
					Value: []string{
						"Batang",
						"Bitstream Vera Sans Mono",
						"MS Mincho",
						"MS UI Gothic",
						"Meiryo UI",
						"PMingLiU",
					},
				},
				ProductSub: ProductSub{
					Value: "20100101",
				},
				Vendor: Vendor{
					Value: "",
				},
			},
			want:  Fingerprint{},
			want1: false,
		},
		{
			name: "f isn't similar to any other fingerprint (incompatible HW)",
			fingerprints: map[string][]Fingerprint{
				"fZRMtpX": {
					fZRMtpX,
				},
			},
			f: Fingerprint{
				ColorDepth: ColorDepth{
					Value: 24,
				},
				DeviceMemory: DeviceMemory{
					Value: nil,
				},
				ScreenResolution: ScreenResolution{
					Value: screenResolution,
				},
				AvailableScreenResolution: AvailableScreenResolution{
					Value: &screenResolution,
				},
				HardwareConcurrency: HardwareConcurrency{
					Value: 8,
				},
				TimezoneOffset: TimezoneOffset{
					Value: -60,
				},
				SessionStorage: SessionStorage{
					Value: true,
				},
				LocalStorage: LocalStorage{
					Value: true,
				},
				IndexedDB: IndexedDB{
					Value: &truth,
				},
				OpenDatabase: OpenDatabase{
					Value: false,
				},
				Canvas: Canvas{
					Value: CanvasValue{
						Winding: true,
						Data:    "a347f8567cbc22e94f1424935cba5c3f",
					},
				},
				TouchSupport: TouchSupport{
					Value: TouchSupportValue{
						MaxTouchPoints: 0,
						TouchEvent:     false,
						TouchStart:     false,
					},
				},
				Audio: Audio{
					Value: 35.73833402246237,
				},
				PluginsSupport: PluginsSupport{
					Value: true,
				},
				EmptyEvalLength: EmptyEvalLength{
					Value: 37,
				},
				ErrorFF: ErrorFF{
					Value: false,
				},
				Chrome: Chrome{
					Value: false,
				},
				CookiesEnabled: CookiesEnabled{
					Value: true,
				},
			},
			want:  Fingerprint{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ExperimentalInMemory{
				fingerprints: tt.fingerprints,
			}
			got, got1 := e.Do(tt.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Do() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestExperimentalInMemory_Store(t *testing.T) {
	e := NewExperimentalInMemory()

	keyA, keyB := "a", "b"

	e.Store(keyA, Fingerprint{VisitorID: "1"})
	e.Store(keyA, Fingerprint{VisitorID: "2"})

	e.Store(keyB, Fingerprint{VisitorID: "3"})

	if len(e.fingerprints) != 2 {
		t.Error("storage inconsistency")
	}

	if len(e.fingerprints[keyA]) != 2 {
		t.Errorf("storage inconsistency under key = %s", keyA)
	}

	if len(e.fingerprints[keyB]) != 1 {
		t.Errorf("storage inconsistency under key = %s", keyB)
	}
}

func TestNewExperimentalInMemory(t *testing.T) {
	tests := []struct {
		name string
		want ExperimentalInMemory
	}{
		{
			name: "basic",
			want: ExperimentalInMemory{
				fingerprints: make(map[string][]Fingerprint),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExperimentalInMemory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExperimentalInMemory() = %v, want %v", got, tt.want)
			}
		})
	}
}
