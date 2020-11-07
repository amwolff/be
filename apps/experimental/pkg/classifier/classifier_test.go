package classifier

import (
	"reflect"
	"testing"
)

func TestExperimentalInMemory_Do(t *testing.T) {
	type fields struct {
		fingerprints map[string][]Fingerprint
	}
	type args struct {
		f Fingerprint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Fingerprint
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := ExperimentalInMemory{
				fingerprints: tt.fields.fingerprints,
			}
			got, got1 := e.Do(tt.args.f)
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
		// TODO: Add test cases.
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

func TestWorthless_Do(t *testing.T) {
	type args struct {
		f Fingerprint
	}
	tests := []struct {
		name  string
		args  args
		want  Fingerprint
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := Worthless{}
			got, got1 := w.Do(tt.args.f)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Do() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Do() got1 = %v, want %v", got1, tt.want1)
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
		// TODO: Add test cases.
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := similarity(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("similarity() = %v, want %v", got, tt.want)
			}
		})
	}
}
