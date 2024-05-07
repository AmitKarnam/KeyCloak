package secretenginevalidator

import (
	"testing"

	"github.com/AmitKarnam/KeyCloak/models"
)

func Test_secretEngineValidator_Validate(t *testing.T) {
	type args struct {
		data models.SecretEngine
	}
	tests := []struct {
		name    string
		sev     *secretEngineValidator
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sev := &secretEngineValidator{}
			if err := sev.Validate(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("secretEngineValidator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
