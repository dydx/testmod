package testmod

import (
	"errors"
	"fmt"
)

// Hi accepts a String for `name` and a String for `lang`
// Accepted values of `lang` are limited to:
// * `en`
// * `pt`
// * `es`
// * `fr`
func Hi(name, lang string) (string, error) {
	switch lang {
	case "en":
		return fmt.Sprintf("Hi, %s!", name), nil
	case "pt":
		return fmt.Sprintf("Oi, %s!", name), nil
	case "es":
		return fmt.Sprintf("¡Hola, %s!", name), nil
	case "fr":
		return fmt.Sprintf("Bonjour, %s!", name), nil
	default:
		return "", errors.New("Unknown language")
	}
}
