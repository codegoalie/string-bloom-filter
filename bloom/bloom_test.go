package bloom

import "testing"

func TestAdd(t *testing.T) {
	filter := Filter{}
	filter.Add("These pretzels are making me thirsty.")
	if !filter.bitmap[176] {
		t.Errorf("Expected to set bitmap 176 to true")
	}
	if !filter.bitmap[128] {
		t.Errorf("Expected to set bitmap 128 to true")
	}
	if !filter.bitmap[78] {
		t.Errorf("Expected to set bitmap 78 to true")
	}
	if filter.bitmap[1] {
		t.Errorf("Expected not to set bitmap 1 to true")
	}
	if filter.bitmap[155] {
		t.Errorf("Expected not to set bitmap 155 to true")
	}
	if filter.bitmap[255] {
		t.Errorf("Expected not to set bitmap 255 to true")
	}
}

func TestCheck(t *testing.T) {
	filter := Filter{}
	str := "These pretzels are making me thirsty."
	filter.Add(str)

	if !filter.Check(str) {
		t.Errorf("Expected Check to find the string")
	}

	if filter.Check("Something else") {
		t.Errorf("Expected Check not to find another string")
	}
}
