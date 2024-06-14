package matcher_test

import (
	"testing"

	"github.com/jinmukeji/plat-pkg/v4/experiments/whitelist/matcher"
)

func TestMatch(t *testing.T) {
	type args struct {
		pattern string
		service string
		method  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "empty-service-1",
			args: args{pattern: "*/*.*", service: "", method: "EchoAPI.F1"},
			want: false,
		},
		{
			name: "empty-method-1",
			args: args{pattern: "*/*.*", service: "com.jinmuhealth.platform.srv.s1", method: ""},
			want: false,
		},
		{
			name: "exact-1",
			args: args{pattern: "com.jinmuhealth.platform.srv.svc-1/EchoAPI.F1", service: "com.jinmuhealth.platform.srv.svc-1", method: "EchoAPI.F1"},
			want: true,
		},
		{
			name: "all-1",
			args: args{pattern: "*/*.*", service: "com.jinmuhealth.platform.srv.s1", method: "EchoAPI.F1"},
			want: true,
		},
		{
			name: "all-2",
			args: args{pattern: "*/*.*", service: "com.jinmuhealth.platform.srv.s2", method: "EchoAPI.F2"},
			want: true,
		},
		{
			name: "all-3",
			args: args{pattern: "*/*.*", service: "com.jinmuhealth.platform.srv.s1", method: "ExxAPI.F1"},
			want: true,
		},
		{
			name: "service-wildcard-1",
			args: args{pattern: "com.jinmuhealth.platform.srv.*/*.*", service: "com.jinmuhealth.platform.srv.s1", method: "EchoAPI.F1"},
			want: true,
		},
		{
			name: "service-wildcard-2",
			args: args{pattern: "com.jinmuhealth.platform.srv.*/*.*", service: "com.jinmuhealth.xxx.srv.s1", method: "EchoAPI.F1"},
			want: false,
		},
		{
			name: "service-wildcard-3",
			args: args{pattern: "com.jinmuhealth.*.srv.*/*.*", service: "com.jinmuhealth.xxx.srv.s1", method: "EchoAPI.F1"},
			want: true,
		},
		{
			name: "api-wildcard-1",
			args: args{pattern: "com.jinmuhealth.platform.srv.s1/Ec*API.*", service: "com.jinmuhealth.platform.srv.s1", method: "EchoAPI.F1"},
			want: true,
		},
		{
			name: "api-wildcard-2",
			args: args{pattern: "com.jinmuhealth.platform.srv.s1/Ec*API.*", service: "com.jinmuhealth.platform.srv.s1", method: "EcxxxxxxxAPI.F1"},
			want: true,
		},
		{
			name: "call-wildcard-1",
			args: args{pattern: "com.jinmuhealth.platform.srv.s1/EchoAPI.*", service: "com.jinmuhealth.platform.srv.s1", method: "EchoAPI.F1"},
			want: true,
		},
		{
			name: "call-wildcard-2",
			args: args{pattern: "com.jinmuhealth.platform.srv.s1/EchoAPI.M*", service: "com.jinmuhealth.platform.srv.s1", method: "EchoAPI.F1"},
			want: false,
		},
		{
			name: "call-wildcard-3",
			args: args{pattern: "com.jinmuhealth.platform.srv.s1/EchoAPI.M*", service: "com.jinmuhealth.platform.srv.s1", method: "EchoAPI.M1"},
			want: true,
		},
		{
			name: "call-wildcard-4",
			args: args{pattern: "com.jinmuhealth.platform.srv.s1/EchoAPI.M*", service: "com.jinmuhealth.platform.srv.s1", method: "EchoAPI.M2"},
			want: true,
		},
		{
			name: "multiple-patterns-1",
			args: args{pattern: "{com.jinmuhealth.platform.srv.s1/EchoAPI.M*,com.jinmuhealth.platform.srv.s2/*.*}", service: "com.jinmuhealth.platform.srv.s1", method: "EchoAPI.M2"},
			want: true,
		},
		{
			name: "multiple-patterns-2",
			args: args{pattern: "{com.jinmuhealth.platform.srv.s1/EchoAPI.M*,com.jinmuhealth.platform.srv.s2/*.*}", service: "com.jinmuhealth.platform.srv.s2", method: "XxAPI.N2"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matcher.Match(tt.args.pattern, tt.args.service, tt.args.method); got != tt.want {
				t.Errorf("Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
