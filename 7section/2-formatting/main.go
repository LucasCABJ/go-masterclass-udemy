package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value interface{}
	IsSet bool
}

// %v - default formatting
// %+v - adds fields for struct
// %T - get type
// %s - string
// %d - digit (int)
// %f - float (ex: %.2f)
// %t - boolean
// %q - double quote string

func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %s, IsSet %t", c.Key, c.Value, c.IsSet)
}

func main() {
	appName := "EnvParser"
	version := 1.2
	port := 8080
	isEnabled := true

	status := fmt.Sprintf("App: %s, Version: %.1f, Port: %d, Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)
}
