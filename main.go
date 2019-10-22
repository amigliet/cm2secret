package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "log"
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
  return ConfigMap{"v1", map[string]string{}, "ConfigMap", map[string]string{}}
}

func NewSecret() Secret {
  return Secret{"v1", map[string][]byte{}, "Secret", map[string]string{}}
}

func (c *ConfigMap) loadConfigMap(b []byte) error {
  if json.Valid(b) {
    return json.Unmarshal(b, c)
  }
  return yaml.Unmarshal(b, c)
}

func (c *ConfigMap) CM2Secret() Secret {
  // Create a new Secret object
  s := NewSecret()

  s.ApiVersion = c.ApiVersion
  s.Metadata["name"] = c.Metadata["name"]

  for key, value := range c.Data {
    s.Data[key] = []byte(value)
  }
  return s
}

func JSONToYAML(j []byte) ([]byte, error) {
	var jsonObj interface{}

	err := yaml.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(jsonObj)
}

func fileExists(file string) bool {
  _, err := os.Stat(file)
  return err == nil
}

func main(){

  var (
    inputFile    = flag.String("f", "", "ConfigMap filename to convert.")
    outputFormat = flag.String("o", "json", "Output format. Choose 'json' or 'yaml'.")
  )
  flag.Parse()

  // TODO: Set up initial checks.
  if !fileExists(*inputFile) {
    log.Fatalf("ERROR: Input file does not exist.")
  }

  bytes, err := ioutil.ReadFile(*inputFile)
  if err != nil {
    log.Fatalf("ERROR: Cannot read input file: %s", err)
  }

  var cm ConfigMap
  if err = cm.loadConfigMap(bytes); err != nil {
    log.Fatalf("ERROR: Invalid input format: %s", err)
  }

  // Create a new Secret object from loaded ConfigMap
  sec := cm.CM2Secret()

  // NOTE: Even if the desired output is in YAML, this is a mandatory step.
  //       Because when a Secret object is directly marshaled into YAML,
  //       map[string][]byte data are not serialized and encoded properly.
  var jsonData []byte
  jsonData, err = json.Marshal(sec)
  if err != nil {
    log.Fatalf("ERROR: Failed to marshal secret: %s", err)
  }

  switch *outputFormat {
    case "json":
      fmt.Println(string(jsonData))
    case "yaml":
      yamlData, err := JSONToYAML(jsonData)
      if err != nil {
        log.Fatalf("ERROR: cannot marshaling YAML: %s", err)
      }
      fmt.Println(string(yamlData))
    default:
      fmt.Println(string(jsonData))
  }

}

