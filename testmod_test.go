package testmod

import "testing"

func TestEnglish(t *testing.T) {
	expected := "Hi, Test!"
	actual, _ := Hi("Test", "en")
	if actual != expected {
		t.Error("-> Test Failed. Expected: ", expected)
	}
}

func TestPortugese(t *testing.T) {
	expected := "Oi, Test!"
	actual, _ := Hi("Test", "pt")
	if actual != expected {
		t.Error("-> Test Failed. Expected: ", expected)
	}
}

func TestSpanish(t *testing.T) {
	expected := "¡Hola, Test!"
	actual, _ := Hi("Test", "es")
	if actual != expected {
		t.Error("-> Test Failed. Expected: ", expected)
	}
}

func TestFrench(t *testing.T) {
	expected := "Bonjour, Test!"
	actual, _ := Hi("Test", "fr")
	if actual != expected {
		t.Error("-> Test Failed. Expected: ", expected)
	}
}

func TestError(t *testing.T) {
	_, err := Hi("Test", "na")
	if err == nil {
		t.Error("-> Test Failed. Should Have Thrown Error")
	}
}
