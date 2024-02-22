package model

import (
	"testing"
)

func TestEmployeeRole_String(t *testing.T) {
	tests := []struct {
		name string
		e    EmployeeRole
		want string
	}{
		{
			name: "admin",
			e:    EmployeeRoleAdministrator,
			want: "Administrator",
		},
		{
			name: "manager",
			e:    EmployeeRoleManager,
			want: "Manager",
		},
		{
			name: "technician",
			e:    EmployeeRoleTechnician,
			want: "Technician",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashedPassword_Compare(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		hp   HashedPassword
		args args
		want bool
	}{
		{
			name: "not compare",
			hp:   HashedPassword("test"),
			args: args{
				password: "test",
			},
			want: false,
		},
		{
			name: "compare",
			hp:   HashedPassword("$2a$10$MvXY1LBlPnIE5tyHa5uKNegIVuAi8DVMIgnBHsuwRmKaPT2h7ydym"),
			args: args{
				password: "test",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.hp.Compare(tt.args.password); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashedPassword_FromPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		hp   HashedPassword
		args args
	}{
		{
			name: "from password",
			args: args{
				password: "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.hp.FromPassword(tt.args.password)
			if len(tt.hp) == 0 {
				t.Errorf("FromPassword() = %v, want %v", tt.hp, tt.hp)
			}
		})
	}
}
