package level2_test

import (
	"testing"

	"github.com/ar-sandbox3/level2"
)

func TestCreateNIP(t *testing.T) {
	tests := []struct {
		ikhwanOrAkhwat level2.IkhwanOrAkhwat
		year           string
		month          int
		id             int
		want           string
	}{
		{
			level2.IKHWAN, "2019", 1, 1, "ARN191-00001",
		},
		{
			level2.IKHWAN, "2019", 8, 1, "ARN192-00001",
		},
		{
			level2.IKHWAN, "2019", 8, 100, "ARN192-00100",
		},
		{
			level2.AKHWAT, "2019", 1, 1, "ART191-00001",
		},
	}

	for _, test := range tests {
		got := level2.CreateNIP(test.ikhwanOrAkhwat, test.year, test.month, test.id)
		if got != test.want {
			t.Errorf("createNIP(%v, %v, %v, %v) = %v; want %v", test.ikhwanOrAkhwat, test.year, test.month, test.id, got, test.want)
		}
	}
}

func TestGenerateNIPs(t *testing.T) {
	tests := []struct {
		ikhwanOrAkhwat level2.IkhwanOrAkhwat
		year           string
		month          int
		count          int
		start          int
		want           []string
	}{
		{
			level2.IKHWAN, "2019", 1, 3, 1, []string{"ARN191-00001", "ARN191-00002", "ARN191-00003"},
		},
		{
			level2.IKHWAN, "2019", 8, 3, 1, []string{"ARN192-00001", "ARN192-00002", "ARN192-00003"},
		},
		{
			level2.AKHWAT, "2019", 1, 3, 1, []string{"ART191-00001", "ART191-00002", "ART191-00003"},
		},
	}

	for _, test := range tests {
		got := level2.GenerateNIPs(test.ikhwanOrAkhwat, test.year, test.month, test.count, test.start)
		if len(got) != len(test.want) {
			t.Errorf("generateNIPs(%v, %v, %v, %v, %v) = %v; want %v", test.ikhwanOrAkhwat, test.year, test.month, test.count, test.start, got, test.want)
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("generateNIPs(%v, %v, %v, %v, %v) = %v; want %v", test.ikhwanOrAkhwat, test.year, test.month, test.count, test.start, got, test.want)
			}
		}
	}
}

func TestCreateNextNIP(t *testing.T) {
	tests := []struct {
		nip  string
		next int
		want string
	}{
		{"ARN191-00001", 1, "ARN191-00002"},
		{"ARN191-00001", 100, "ARN191-00101"},
		{"ARN191-00001", 1000, "ARN191-01001"},
		{"ART191-00001", 1, "ART191-00002"},
		{"ART191-00001", 100, "ART191-00101"},
		{"ART191-00001", 1000, "ART191-01001"},
	}

	for _, test := range tests {
		got := level2.CreateNextNIP(test.nip, test.next)
		if got != test.want {
			t.Errorf("createNextNIP(%v, %v) = %v; want %v", test.nip, test.next, got, test.want)
		}
	}
}

func TestGenerateNextNIPs(t *testing.T) {
	tests := []struct {
		nip   string
		count int
		want  []string
	}{
		{"ARN191-00001", 3, []string{"ARN191-00002", "ARN191-00003", "ARN191-00004"}},
		{"ART191-00001", 3, []string{"ART191-00002", "ART191-00003", "ART191-00004"}},
	}

	for _, test := range tests {
		got := level2.GenerateNextNIPs(test.nip, test.count)
		if len(got) != len(test.want) {
			t.Errorf("generateNextNIPs(%v, %v) = %v; want %v", test.nip, test.count, got, test.want)
		}
		for i := range got {
			if got[i] != test.want[i] {
				t.Errorf("generateNextNIPs(%v, %v) = %v; want %v", test.nip, test.count, got, test.want)
			}
		}
	}
}
