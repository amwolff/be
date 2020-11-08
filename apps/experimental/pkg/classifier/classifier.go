package classifier

import (
	"math"
	"strings"
	"sync"

	"github.com/amwolff/be/apps/experimental/internal/levenshtein"
)

const (
	Undefined = "f1a626ef-1003-44d8-9cff-f88019975c5c"
	sep       = " "
)

func stringValueOrError(err, val *string) string {
	if err != nil {
		return *err
	}
	if val != nil {
		return *val
	}
	return Undefined
}

func standardStringValueOrError(err *string, val string) string {
	if err != nil {
		return *err
	}
	return val
}

type OsCPU struct {
	Value    *string `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

func (o OsCPU) String() string {
	return stringValueOrError(o.Error, o.Value)
}

type Languages struct {
	Value    [][]string `json:"value"`
	Error    *string    `json:"error"`
	Duration int        `json:"duration"`
}

func (l Languages) String() string {
	if l.Error != nil {
		return *l.Error
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

type ColorDepth struct {
	Value    int     `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type DeviceMemory struct {
	Value    *int    `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type ScreenResolution struct {
	Value    [2]int  `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type AvailableScreenResolution struct {
	Value    *[2]int `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type HardwareConcurrency struct {
	Value    int     `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type TimezoneOffset struct {
	Value    int     `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type Timezone struct {
	Value    *string `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

func (t Timezone) String() string {
	return stringValueOrError(t.Error, t.Value)
}

type SessionStorage struct {
	Value    bool    `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type LocalStorage struct {
	Value    bool    `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type IndexedDB struct {
	Value    *bool   `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type OpenDatabase struct {
	Value    bool    `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type CPUClass struct {
	Value    *string `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

func (c CPUClass) String() string {
	return stringValueOrError(c.Error, c.Value)
}

type Platform struct {
	Value    string  `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

func (p Platform) String() string {
	return standardStringValueOrError(p.Error, p.Value)
}

type (
	PluginsValueMimeTypes struct {
		Type     string `json:"type"`
		Suffixes string `json:"suffixes"`
	}
	PluginsValue struct {
		Name        string                  `json:"name"`
		Description string                  `json:"description"`
		MimeTypes   []PluginsValueMimeTypes `json:"mimeTypes"`
	}
	Plugins struct {
		Value    []PluginsValue `json:"value"`
		Error    *string        `json:"error"`
		Duration int            `json:"duration"`
	}
)

func (p Plugins) String() string {
	if p.Error != nil {
		return *p.Error
	}

	if p.Value == nil {
		return Undefined
	}

	var b strings.Builder

	for i, v := range p.Value {
		if i > 0 {
			b.WriteString(sep)
		}

		b.WriteString(v.Name)
		b.WriteString(sep)
		b.WriteString(v.Description)

		for _, t := range v.MimeTypes {
			b.WriteString(sep)
			b.WriteString(t.Type)
			b.WriteString(sep)
			b.WriteString(t.Suffixes)
		}
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
		Error    *string     `json:"error"`
		Duration int         `json:"duration"`
	}
)

type (
	TouchSupportValue struct {
		MaxTouchPoints int  `json:"maxTouchPoints"`
		TouchEvent     bool `json:"touchEvent"`
		TouchStart     bool `json:"touchStart"`
	}
	TouchSupport struct {
		Value    TouchSupportValue `json:"value"`
		Error    *string           `json:"error"`
		Duration int               `json:"duration"`
	}
)

type Fonts struct {
	Value    []string `json:"value"`
	Error    *string  `json:"error"`
	Duration int      `json:"duration"`
}

func (f Fonts) String() string {
	if f.Error != nil {
		return *f.Error
	}
	return strings.Join(f.Value, sep)
}

type Audio struct {
	Value    float64 `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type PluginsSupport struct {
	Value    bool    `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type ProductSub struct {
	Value    string  `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

func (p ProductSub) String() string {
	return standardStringValueOrError(p.Error, p.Value)
}

type EmptyEvalLength struct {
	Value    int     `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type ErrorFF struct {
	Value    bool    `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type Vendor struct {
	Value    string  `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

func (v Vendor) String() string {
	return standardStringValueOrError(v.Error, v.Value)
}

type Chrome struct {
	Value    bool    `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type CookiesEnabled struct {
	Value    bool    `json:"value"`
	Error    *string `json:"error"`
	Duration int     `json:"duration"`
}

type Fingerprint struct {
	VisitorID                 string                    `json:"visitorID"`
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
	Store(key string, f Fingerprint) float64
}

func hardwareRelatedCompatibility(a, b Fingerprint) float64 {
	var count float64

	if a.ColorDepth.Error == nil && b.ColorDepth.Error == nil &&
		a.ColorDepth.Value == b.ColorDepth.Value {
		count++
	} else if a.ColorDepth.Error != nil && b.ColorDepth.Error != nil &&
		*a.ColorDepth.Error == *b.ColorDepth.Error {
		count++
	}

	if a.DeviceMemory.Error == nil && b.DeviceMemory.Error == nil {
		if a.DeviceMemory.Value == nil && b.DeviceMemory.Value == nil {
			count++
		} else if a.DeviceMemory.Value != nil && b.DeviceMemory.Value != nil &&
			*a.DeviceMemory.Value == *b.DeviceMemory.Value {
			count++
		}
	} else if a.DeviceMemory.Error != nil && b.DeviceMemory.Error != nil &&
		*a.DeviceMemory.Error == *b.DeviceMemory.Error {
		count++
	}

	if a.ScreenResolution.Error == nil && b.ScreenResolution.Error == nil &&
		a.ScreenResolution.Value == b.ScreenResolution.Value {
		count++
	} else if a.ScreenResolution.Error != nil && b.ScreenResolution.Error != nil &&
		*a.ScreenResolution.Error == *b.ScreenResolution.Error {
		count++
	}

	if a.AvailableScreenResolution.Error == nil && b.AvailableScreenResolution.Error == nil {
		if a.AvailableScreenResolution.Value == nil && b.AvailableScreenResolution.Value == nil {
			count++
		} else if a.AvailableScreenResolution.Value != nil && b.AvailableScreenResolution.Value != nil &&
			*a.AvailableScreenResolution.Value == *b.AvailableScreenResolution.Value {
			count++
		}
	} else if a.AvailableScreenResolution.Error != nil && b.AvailableScreenResolution.Error != nil &&
		*a.AvailableScreenResolution.Error == *b.AvailableScreenResolution.Error {
		count++
	}

	if a.HardwareConcurrency.Error == nil && b.HardwareConcurrency.Error == nil &&
		a.HardwareConcurrency.Value == b.HardwareConcurrency.Value {
		count++
	} else if a.HardwareConcurrency.Error != nil && b.HardwareConcurrency.Error != nil &&
		*a.HardwareConcurrency.Error == *b.HardwareConcurrency.Error {
		count++
	}

	if a.TimezoneOffset.Error == nil && b.TimezoneOffset.Error == nil &&
		a.TimezoneOffset.Value == b.TimezoneOffset.Value {
		count++
	} else if a.TimezoneOffset.Error != nil && b.TimezoneOffset.Error != nil &&
		*a.TimezoneOffset.Error == *b.TimezoneOffset.Error {
		count++
	}

	if a.SessionStorage.Error == nil && b.SessionStorage.Error == nil &&
		a.SessionStorage.Value == b.SessionStorage.Value {
		count++
	} else if a.SessionStorage.Error != nil && b.SessionStorage.Error != nil &&
		*a.SessionStorage.Error == *b.SessionStorage.Error {
		count++
	}

	if a.LocalStorage.Error == nil && b.LocalStorage.Error == nil &&
		a.LocalStorage.Value == b.LocalStorage.Value {
		count++
	} else if a.LocalStorage.Error != nil && b.LocalStorage.Error != nil &&
		*a.LocalStorage.Error == *b.LocalStorage.Error {
		count++
	}

	if a.IndexedDB.Error == nil && b.IndexedDB.Error == nil {
		if a.IndexedDB.Value == nil && b.IndexedDB.Value == nil {
			count++
		} else if a.IndexedDB.Value != nil && b.IndexedDB.Value != nil &&
			*a.IndexedDB.Value == *b.IndexedDB.Value {
			count++
		}
	} else if a.IndexedDB.Error != nil && b.IndexedDB.Error != nil &&
		*a.IndexedDB.Error == *b.IndexedDB.Error {
		count++
	}

	if a.OpenDatabase.Error == nil && b.OpenDatabase.Error == nil &&
		a.OpenDatabase.Value == b.OpenDatabase.Value {
		count++
	} else if a.OpenDatabase.Error != nil && b.OpenDatabase.Error != nil &&
		*a.OpenDatabase.Error == *b.OpenDatabase.Error {
		count++
	}

	if a.Canvas.Error == nil && b.Canvas.Error == nil {
		if a.Canvas.Value.Winding == b.Canvas.Value.Winding {
			count++
		}
		if a.Canvas.Value.Data == b.Canvas.Value.Data {
			count++
		}
	} else if a.Canvas.Error != nil && b.Canvas.Error != nil &&
		*a.Canvas.Error == *b.Canvas.Error {
		count += 2
	}

	if a.TouchSupport.Error == nil && b.TouchSupport.Error == nil {
		if a.TouchSupport.Value.MaxTouchPoints == b.TouchSupport.Value.MaxTouchPoints {
			count++
		}
		if a.TouchSupport.Value.TouchEvent == b.TouchSupport.Value.TouchEvent {
			count++
		}
		if a.TouchSupport.Value.TouchStart == b.TouchSupport.Value.TouchStart {
			count++
		}
	} else if a.TouchSupport.Error != nil && b.TouchSupport.Error != nil &&
		*a.TouchSupport.Error == *b.TouchSupport.Error {
		count += 3
	}

	if a.Audio.Error == nil && b.Audio.Error == nil &&
		a.Audio.Value == b.Audio.Value {
		count++
	} else if a.Audio.Error != nil && b.Audio.Error != nil &&
		*a.Audio.Error == *b.Audio.Error {
		count++
	}

	if a.PluginsSupport.Error == nil && b.PluginsSupport.Error == nil &&
		a.PluginsSupport.Value == b.PluginsSupport.Value {
		count++
	} else if a.PluginsSupport.Error != nil && b.PluginsSupport.Error != nil &&
		*a.PluginsSupport.Error == *b.PluginsSupport.Error {
		count++
	}

	if a.EmptyEvalLength.Error == nil && b.EmptyEvalLength.Error == nil &&
		a.EmptyEvalLength.Value == b.EmptyEvalLength.Value {
		count++
	} else if a.EmptyEvalLength.Error != nil && b.EmptyEvalLength.Error != nil &&
		*a.EmptyEvalLength.Error == *b.EmptyEvalLength.Error {
		count++
	}

	if a.ErrorFF.Error == nil && b.ErrorFF.Error == nil &&
		a.ErrorFF.Value == b.ErrorFF.Value {
		count++
	} else if a.ErrorFF.Error != nil && b.ErrorFF.Error != nil &&
		*a.ErrorFF.Error == *b.ErrorFF.Error {
		count++
	}

	if a.Chrome.Error == nil && b.Chrome.Error == nil &&
		a.Chrome.Value == b.Chrome.Value {
		count++
	} else if a.Chrome.Error != nil && b.Chrome.Error != nil &&
		*a.Chrome.Error == *b.Chrome.Error {
		count++
	}

	if a.CookiesEnabled.Error == nil && b.CookiesEnabled.Error == nil &&
		a.CookiesEnabled.Value == b.CookiesEnabled.Value {
		count++
	} else if a.CookiesEnabled.Error != nil && b.CookiesEnabled.Error != nil &&
		*a.CookiesEnabled.Error == *b.CookiesEnabled.Error {
		count++
	}

	return count / 21
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
			a.CPUClass.String(),
			b.CPUClass.String(),
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
		if c[0] == Undefined || c[1] == Undefined {
			if c[0] == Undefined && c[1] == Undefined {
				ratiosSum++
			}
			continue
		}
		d, l := levenshtein.LinearSpace(c[0], c[1]), max(len(c[0]), len(c[1]))
		// We are subtracting from 1 here because it then means that these two
		// strings are compatible in this ratio.
		if ratiosSum++; l > 0 {
			ratiosSum -= float64(d) / float64(l)
		}
	}

	return ratiosSum / float64(len(components))
}

type ExperimentalInMemory struct {
	fingerprints map[string][]Fingerprint
	mtx          *sync.RWMutex
}

func (e ExperimentalInMemory) Do(f Fingerprint) (Fingerprint, bool) {
	var (
		max       float64
		candidate Fingerprint
	)

	e.mtx.RLock()
	for _, fingerprints := range e.fingerprints {
		for _, g := range fingerprints {
			if hardwareRelatedCompatibility(f, g) < 0.85 {
				continue
			}

			if ratio := similarity(f, g); ratio > max {
				max, candidate = ratio, g
			}
		}
	}
	e.mtx.RUnlock()

	if max > 0.9 {
		return candidate, true
	}

	return Fingerprint{}, false
}

func (e ExperimentalInMemory) calculateEntropy() float64 {
	var lengths []int

	e.mtx.RLock()
	for _, fingerprints := range e.fingerprints {
		lengths = append(lengths, len(fingerprints))
	}
	e.mtx.RUnlock()

	var sum float64

	for _, l := range lengths {
		sum += float64(l)
	}

	var entropy float64

	for _, l := range lengths {
		if freq := float64(l) / sum; freq > 0 {
			entropy -= freq * math.Log2(freq)
		}
	}

	return entropy
}

func (e ExperimentalInMemory) Store(key string, f Fingerprint) float64 {
	e.mtx.Lock()
	if _, ok := e.fingerprints[f.VisitorID]; !ok {
		e.fingerprints[f.VisitorID] = make([]Fingerprint, 0)
		e.fingerprints[key] = append(e.fingerprints[key], f)
	}
	e.mtx.Unlock()

	return e.calculateEntropy()
}

func NewExperimentalInMemory() ExperimentalInMemory {
	return ExperimentalInMemory{
		fingerprints: make(map[string][]Fingerprint),
		mtx:          &sync.RWMutex{},
	}
}
