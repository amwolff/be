package classifier

import (
	"math"
	"reflect"
	"testing"
)

var fZRMtpX = Fingerprint{
	VisitorID: "fZRMtpX",
	OsCPU: OsCPU{
		Value: "Windows NT 6.2",
	},
}

func TestExperimentalInMemory_Do(t *testing.T) {
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
					Value: "Linux x86_64",
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
					Value: "Windows NT 6.3",
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
					Value: "Linux x86_64",
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
					Value: "Europe/Warsaw",
				},
				Platform: Platform{
					Value: "Linux x86_64",
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

func TestFonts_String(t *testing.T) {
	type fields struct {
		Value    []string
		Error    string
		Duration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Fonts{
				Value:    tt.fields.Value,
				Error:    tt.fields.Error,
				Duration: tt.fields.Duration,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguages_String(t *testing.T) {
	type fields struct {
		Value    [][]string
		Error    string
		Duration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Languages{
				Value:    tt.fields.Value,
				Error:    tt.fields.Error,
				Duration: tt.fields.Duration,
			}
			if got := l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
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

func TestOsCPU_String(t *testing.T) {
	type fields struct {
		Value    string
		Error    string
		Duration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := OsCPU{
				Value:    tt.fields.Value,
				Error:    tt.fields.Error,
				Duration: tt.fields.Duration,
			}
			if got := o.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlatform_String(t *testing.T) {
	type fields struct {
		Value    string
		Error    string
		Duration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Platform{
				Value:    tt.fields.Value,
				Error:    tt.fields.Error,
				Duration: tt.fields.Duration,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlugins_String(t *testing.T) {
	type fields struct {
		Value    []PluginsValue
		Error    string
		Duration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plugins{
				Value:    tt.fields.Value,
				Error:    tt.fields.Error,
				Duration: tt.fields.Duration,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductSub_String(t *testing.T) {
	type fields struct {
		Value    string
		Error    string
		Duration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ProductSub{
				Value:    tt.fields.Value,
				Error:    tt.fields.Error,
				Duration: tt.fields.Duration,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimezone_String(t1 *testing.T) {
	type fields struct {
		Value    string
		Error    string
		Duration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Timezone{
				Value:    tt.fields.Value,
				Error:    tt.fields.Error,
				Duration: tt.fields.Duration,
			}
			if got := t.String(); got != tt.want {
				t1.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVendor_String(t *testing.T) {
	type fields struct {
		Value    string
		Error    string
		Duration int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Vendor{
				Value:    tt.fields.Value,
				Error:    tt.fields.Error,
				Duration: tt.fields.Duration,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hardwareRelatedCompatibility(t *testing.T) {
	type args struct {
		a Fingerprint
		b Fingerprint
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
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
			name: "basic, a is empty",
			args: args{
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: "Linux x86_64",
					},
				},
			},
			want: 7. / 8,
		},
		{
			name: "basic, b is empty",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: "Linux x86_64",
					},
				},
			},
			want: 7. / 8,
		},
		{
			name: "OsCPU",
			args: args{
				a: Fingerprint{
					OsCPU: OsCPU{
						Value: "Windows NT 6.2",
					},
				},
				b: Fingerprint{
					OsCPU: OsCPU{
						Value: "Windows NT 6.3",
					},
				},
			},
			want: 111. / 112,
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
