package settings

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
)

var configPath = "/Users/milesdowe/config.json"

//var environments = map[string]string{
//  "production":    "settings/prod.json",
//  //"preproduction": "settings/pre.json",
//  "preproduction": "~/config.json",
//  "tests":         "../../settings/tests.json",
//}

type Settings struct {
  PrivateKeyPath     string
  PublicKeyPath      string
  JWTExpirationDelta int
}

// Settings for each stage of deployment
type Staging struct {
  Preproduction Settings
  Production Settings
}

var staging Staging = Staging{}
var settings Settings = Settings{}
var env = "preproduction"

func Init() {
  env = os.Getenv("GO_ENV")
  if env == "" {
    fmt.Println("Warning: Setting preproduction environment due to lack of GO_ENV value")
    env = "preproduction"
  }
  LoadSettingsByEnv(env)
}

func LoadSettingsByEnv(env string) {
  content, err := ioutil.ReadFile(configPath)
  if err != nil {
    fmt.Println("Error while reading config file", err)
  }
  staging = Staging{}
  settings = Settings{}
  jsonErr := json.Unmarshal(content, &staging)
  if jsonErr != nil {
    fmt.Println("Error while parsing config file", jsonErr)
  }

  if env == "production" {
    settings = staging.Production
  } else {
    settings = staging.Preproduction
  }
}

func GetEnvironment() string {
  return env
}

func Get() Settings {
  if &settings == nil {
    Init()
  }
  return settings
}

func IsTestEnvironment() bool {
  return env == "tests"
}