package classifier

import (
	"strings"

	"github.com/amwolff/be/apps/experimental/internal/levenshtein"
)

const sep = " "

type OsCPU struct {
	Value    string `json:"value"`
	Error    string `json:"error"`
	Duration int    `json:"duration"`
}

func (o OsCPU) String() string {
	if len(o.Error) > 0 {
		return o.Error
	}
	return o.Value
}

type Languages struct {
	Value    [][]string `json:"value"`
	Error    string     `json:"error"`
	Duration int        `json:"duration"`
}

func (l Languages) String() string {
	if len(l.Error) > 0 {
		return l.Error
	}

	var b strings.Builder

	for i, v := range l.Value {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(strings.Join(v, sep))
	}

	return b.String()
}

type (
	ColorDepth struct {
		Value    int    `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	DeviceMemory struct {
		Value    int    `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	ScreenResolution struct {
		Value    [2]int `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	AvailableScreenResolution struct {
		Value    [2]int `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	HardwareConcurrency struct {
		Value    int    `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	TimezoneOffset struct {
		Value    int    `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
)

type Timezone struct {
	Value    string `json:"value"`
	Error    string `json:"error"`
	Duration int    `json:"duration"`
}

func (t Timezone) String() string {
	if len(t.Error) > 0 {
		return t.Error
	}
	return t.Value
}

type (
	SessionStorage struct {
		Value    bool   `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	LocalStorage struct {
		Value    bool   `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	IndexedDB struct {
		Value    bool   `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	OpenDatabase struct {
		Value    bool   `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	CPUClass struct {
		Value    string `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
)

type Platform struct {
	Value    string `json:"value"`
	Error    string `json:"error"`
	Duration int    `json:"duration"`
}

func (p Platform) String() string {
	if len(p.Error) > 0 {
		return p.Error
	}
	return p.Value
}

type (
	PluginsValueMimeTypes struct {
		Type     string `json:"type"`
		Suffixes string `json:"suffixes"`
	}
	PluginsValue struct {
		Name        string   `json:"name"`
		Description string   `json:"description"`
		MimeTypes   []string `json:"mimeTypes"`
	}
	Plugins struct {
		Value    []PluginsValue `json:"value"`
		Error    string         `json:"error"`
		Duration int            `json:"duration"`
	}
)

func (p Plugins) String() string {
	if len(p.Error) > 0 {
		return p.Error
	}

	var b strings.Builder

	for i, v := range p.Value {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(v.Name)
		b.WriteString(sep)
		b.WriteString(v.Description)
		b.WriteString(sep)
		b.WriteString(strings.Join(v.MimeTypes, sep))
	}

	return b.String()
}

type (
	CanvasValue struct {
		Winding bool   `json:"winding"`
		Data    string `json:"data"`
	}
	Canvas struct {
		Value    CanvasValue `json:"value"`
		Error    string      `json:"error"`
		Duration int         `json:"duration"`
	}
	TouchSupportValue struct {
		MaxTouchPoints int  `json:"maxTouchPoints"`
		TouchEvent     bool `json:"touchEvent"`
		TouchStart     bool `json:"touchStart"`
	}
	TouchSupport struct {
		Value    TouchSupportValue `json:"value"`
		Error    string            `json:"error"`
		Duration int               `json:"duration"`
	}
)

type Fonts struct {
	Value    []string `json:"value"`
	Error    string   `json:"error"`
	Duration int      `json:"duration"`
}

func (f Fonts) String() string {
	if len(f.Error) > 0 {
		return f.Error
	}
	return strings.Join(f.Value, sep)
}

type (
	Audio struct {
		Value    float64 `json:"value"`
		Error    string  `json:"error"`
		Duration int     `json:"duration"`
	}
	PluginsSupport struct {
		Value    bool   `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
)

type ProductSub struct {
	Value    string `json:"value"`
	Error    string `json:"error"`
	Duration int    `json:"duration"`
}

func (p ProductSub) String() string {
	if len(p.Error) > 0 {
		return p.Error
	}
	return p.Value
}

type (
	EmptyEvalLength struct {
		Value    int    `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	ErrorFF struct {
		Value    bool   `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
)

type Vendor struct {
	Value    string `json:"value"`
	Error    string `json:"error"`
	Duration int    `json:"duration"`
}

func (v Vendor) String() string {
	if len(v.Error) > 0 {
		return v.Error
	}
	return v.Value
}

type (
	Chrome struct {
		Value    bool   `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
	CookiesEnabled struct {
		Value    bool   `json:"value"`
		Error    string `json:"error"`
		Duration int    `json:"duration"`
	}
)

type Fingerprint struct {
	VisitorID                 string                    `json:"-"`
	OsCPU                     OsCPU                     `json:"osCpu"`
	Languages                 Languages                 `json:"languages"`
	ColorDepth                ColorDepth                `json:"colorDepth"`
	DeviceMemory              DeviceMemory              `json:"deviceMemory"`
	ScreenResolution          ScreenResolution          `json:"screenResolution"`
	AvailableScreenResolution AvailableScreenResolution `json:"availableScreenResolution"`
	HardwareConcurrency       HardwareConcurrency       `json:"hardwareConcurrency"`
	TimezoneOffset            TimezoneOffset            `json:"timezoneOffset"`
	Timezone                  Timezone                  `json:"timezone"`
	SessionStorage            SessionStorage            `json:"sessionStorage"`
	LocalStorage              LocalStorage              `json:"localStorage"`
	IndexedDB                 IndexedDB                 `json:"indexedDB"`
	OpenDatabase              OpenDatabase              `json:"openDatabase"`
	CPUClass                  CPUClass                  `json:"cpuClass"`
	Platform                  Platform                  `json:"platform"`
	Plugins                   Plugins                   `json:"plugins"`
	Canvas                    Canvas                    `json:"canvas"`
	TouchSupport              TouchSupport              `json:"touchSupport"`
	Fonts                     Fonts                     `json:"fonts"`
	Audio                     Audio                     `json:"audio"`
	PluginsSupport            PluginsSupport            `json:"pluginsSupport"`
	ProductSub                ProductSub                `json:"productSub"`
	EmptyEvalLength           EmptyEvalLength           `json:"emptyEvalLength"`
	ErrorFF                   ErrorFF                   `json:"errorFF"`
	Vendor                    Vendor                    `json:"vendor"`
	Chrome                    Chrome                    `json:"chrome"`
	CookiesEnabled            CookiesEnabled            `json:"cookiesEnabled"`
}

type Classifier interface {
	Do(f Fingerprint) (Fingerprint, bool)
}

type Worthless struct{}

func (w Worthless) Do(f Fingerprint) (Fingerprint, bool) {
	return Fingerprint{}, false
}

func hardwareRelatedCompatibility(a, b Fingerprint) float64 {
	return 0.9
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func similarity(a, b Fingerprint) float64 {
	components := [][2]string{
		{
			a.OsCPU.String(),
			b.OsCPU.String(),
		},
		{
			a.Languages.String(),
			b.Languages.String(),
		},
		{
			a.Timezone.String(),
			b.Timezone.String(),
		},
		{
			a.Platform.String(),
			b.Platform.String(),
		},
		{
			a.Plugins.String(),
			b.Plugins.String(),
		},
		{
			a.Fonts.String(),
			b.Fonts.String(),
		},
		{
			a.ProductSub.String(),
			b.ProductSub.String(),
		},
		{
			a.Vendor.String(),
			b.Vendor.String(),
		},
	}

	var ratiosSum float64

	for _, c := range components {
		d, l := levenshtein.LinearSpace(c[0], c[1]), max(len(c[0]), len(c[1]))
		// We are subtracting from 1 here because it then means that these two
		// strings are compatible in this ratio.
		ratiosSum += 1 - float64(d)/float64(l)
	}

	return ratiosSum / float64(len(components))
}

type ExperimentalInMemory struct {
	fingerprints map[string][]Fingerprint
}

func (e ExperimentalInMemory) Do(f Fingerprint) (Fingerprint, bool) {
	var (
		max       float64
		candidate Fingerprint
	)

	for _, fingerprints := range e.fingerprints {
		for _, g := range fingerprints {
			if hardwareRelatedCompatibility(f, g) < 0.9 {
				continue
			}

			if ratio := similarity(f, g); ratio > max {
				max, candidate = ratio, g
			}
		}
	}

	if max > 0.8 {
		return candidate, true
	}

	return Fingerprint{}, false
}

func NewExperimentalInMemory() ExperimentalInMemory {
	return ExperimentalInMemory{
		fingerprints: make(map[string][]Fingerprint),
	}
}
