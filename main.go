package main

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)


// ConfigMap object to store configuration data.
// This rough version does not support bynary data.
type ConfigMap struct {
  ApiVersion string             `json:"apiVersion" yaml:"apiVersion"`
  Data       map[string]string  `json:"data"       yaml:"data"`
  Kind       string             `json:"kind"       yaml:"kind"`
  Metadata   map[string]string  `json:"metadata"   yaml:"metadata"`
}

// Secret object to store secret data.
// The serialized form of secret data is base64 encoded string.
type Secret struct {
  ApiVersion string             `json:"apiVersion" yaml:"apiVersion"`
  Data       map[string][]byte  `json:"data"       yaml:"data"`
  Kind       string             `json:"kind"       yaml:"kind"`
  Metadata   map[string]string  `json:"metadata"   yaml:"metadata"`
}

func NewConfigMap() ConfigMap {
  return ConfigMap{"", map[string]string{}, "ConfigMap", map[string]string{}}
}

func NewSecret() Secret {
  return Secret{"v1", map[string][]byte{}, "Secret", map[string]string{}}
}


func main(){

  // FIX: blocking line if no stdin is provided
  bytes, err := ioutil.ReadAll(os.Stdin)
  if err != nil {
    panic(err)
  }

  var parsed map[string]interface{}
  if err := json.Unmarshal(bytes, &parsed); err != nil {
    panic(err)
  }

  // Create a new Secret object
  sec := NewSecret()

  for key, value := range parsed["data"].(map[string]interface{}) {
    sec.Data[key] = []byte(value.(string))
  }

  sec.Metadata["name"] = parsed["metadata"].(map[string]interface{})["name"].(string)

  var jsonData []byte
  jsonData, err = json.Marshal(sec)
  if err != nil {
    fmt.Println(err)
  }

  // Print the results to standard output
  fmt.Println(string(jsonData))

}
