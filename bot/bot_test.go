package bot

import "testing"

func TestUnkindleBot_verifyConfig(t *testing.T) {
	type fields struct {
		Token     string
		EmailFrom string
		EmailTo   string
		SmtpHost  string
		SmtpPort  string
		Password  string
	}
	tests := []struct {
		name         string
		fields       fields
		wantErr      bool
		errorMessage string
	}{
		{
			name: "should pass validation",
			fields: fields{
				Token:     "hello",
				EmailFrom: "from",
				EmailTo:   "to",
				SmtpHost:  "host",
				SmtpPort:  "port",
				Password:  "pass",
			},
			wantErr: false,
		},
		{
			name: "should not pass validation if Token empty",
			fields: fields{
				Token:     "",
				EmailFrom: "from",
				EmailTo:   "to",
				SmtpHost:  "host",
				SmtpPort:  "port",
				Password:  "pass",
			},
			wantErr:      true,
			errorMessage: "token for telegram bot not set",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &SendToKindleBot{
				Token:     tt.fields.Token,
				EmailFrom: tt.fields.EmailFrom,
				EmailTo:   tt.fields.EmailTo,
				SMTPHost:  tt.fields.SmtpHost,
				SMTPPort:  tt.fields.SmtpPort,
				Password:  tt.fields.Password,
			}
			err := b.verifyConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyConfig() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil && err.Error() != tt.errorMessage {
				t.Errorf("verifyConfig() error '%v' not matches '%s'", err, tt.errorMessage)
			}
		})
	}
}

func TestSplitFileName(t *testing.T) {
	tests := []struct {
		name string
		in   string
		base string
		ext  string
	}{
		{name: "simple", in: "book.pdf", base: "book", ext: "pdf"},
		{name: "multi dot", in: "book.part1.pdf", base: "book.part1", ext: "pdf"},
		{name: "no extension", in: "README", base: "README", ext: ""},
		{name: "uppercase extension", in: "FILE.PDF", base: "FILE", ext: "pdf"},
		{name: "hidden file", in: ".hidden", base: "", ext: "hidden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			base, ext := splitFileName(tt.in)
			if base != tt.base {
				t.Errorf("splitFileName(%q) base = %q, want %q", tt.in, base, tt.base)
			}
			if ext != tt.ext {
				t.Errorf("splitFileName(%q) ext = %q, want %q", tt.in, ext, tt.ext)
			}
		})
	}
}

func TestNeedToConvert(t *testing.T) {
	tests := []struct {
		ext  string
		want bool
	}{
		{ext: "pdf", want: false},
		{ext: "epub", want: false},
		{ext: "mobi", want: true},
		{ext: "xlsx", want: true},
		{ext: "", want: true},
	}

	for _, tt := range tests {
		t.Run(tt.ext, func(t *testing.T) {
			if got := needToConvert(tt.ext); got != tt.want {
				t.Errorf("needToConvert(%q) = %v, want %v", tt.ext, got, tt.want)
			}
		})
	}
}

func TestVerifyConfigDefaults(t *testing.T) {
	b := &SendToKindleBot{
		Token:     "token",
		EmailFrom: "from",
		EmailTo:   "to",
		SMTPHost:  "host",
		Password:  "pass",
	}

	if err := b.verifyConfig(); err != nil {
		t.Fatalf("verifyConfig() error = %v, want nil", err)
	}

	if b.SMTPPort != "587" {
		t.Errorf("verifyConfig() SMTPPort = %q, want %q", b.SMTPPort, "587")
	}
}
