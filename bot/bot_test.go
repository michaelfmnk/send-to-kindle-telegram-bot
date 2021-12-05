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
			b := &UnkindleBot{
				Token:     tt.fields.Token,
				EmailFrom: tt.fields.EmailFrom,
				EmailTo:   tt.fields.EmailTo,
				SmtpHost:  tt.fields.SmtpHost,
				SmtpPort:  tt.fields.SmtpPort,
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
