package main

import (
    "io/ioutil"
    "log"
    "os/exec"
    "time"
    "fmt"
    "gopkg.in/robfig/cron.v2"
    "github.com/spf13/viper"
)

var config map[int]string = map [int]string{}
var slice =  make([]string, 0)

func init() {
  viper.SetConfigFile(`config.json`)
  err := viper.ReadInConfig()
  if err != nil {
    panic(err)
  }
}
func main() {
  config[0] = fmt.Sprintf("-h%s", viper.GetString(`host`))
  config[1] = fmt.Sprintf("-P%s", viper.GetString(`port`))
  config[2] = fmt.Sprintf("-u%s", viper.GetString(`user`))

  pass := viper.GetString(`password`)
  if pass != "" {
    config[3] = fmt.Sprintf("-p%s", pass)
  }

  config[4] = viper.GetString("path")
  config[5] = viper.GetString("db_name")

  every := viper.GetString("every")

  for i := 0; i <= 5; i++ {
    if i == 4 {
      continue
    }
    if val, ok := config[i]; ok {
      slice = append(slice, val)
    }
  }

  c := cron.New()

  c.AddFunc(every, backupData)

  c.Start()

  fmt.Printf("Cron Job Started. \n")
  select {}
  c.Stop()
}

func backupData() {
  cmd := exec.Command("mysqldump", slice...)
  stdout, err := cmd.StdoutPipe()
  if err != nil {
    log.Fatal(err)
  }

  if err := cmd.Start(); err != nil {
    log.Fatal(err)
  }

  bytes, err := ioutil.ReadAll(stdout)
  if err != nil {
    log.Fatal(err)
  }

  name := fmt.Sprintf("%s/%s_%s.sql",config[4], config[5], time.Now().Format("2006_01_02_15_04_05"))

  err = ioutil.WriteFile(name, bytes, 0644)
  if err != nil {
    panic(err)
  }

  fmt.Printf("Backuping database %s file %s\n",config[5], name)
}