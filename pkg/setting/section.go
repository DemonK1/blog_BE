package setting

import (
  "time"
)

// config配置文件对应的结构体

type ServerSettings struct {
  RunMode      string
  HttpPort     string
  ReadTimeout  time.Duration
  WriteTimeout time.Duration
}

type AppSettings struct {
  DefaultPageSize int
  MaxPageSize     int
  LogSavePath     string
  LogFileName     string
  LogFileExt      string
}

type DatabaseSettings struct {
  DBType       string
  Username     string
  Password     string
  Host         string
  DBName       string
  TablePrefix  string
  Charset      string
  ParseTime    bool
  MaxIdleConns int
  MaxOpenConns int
}

// ReadSection 将配置文件的值解析到对应的结构体中
func (s *Setting) ReadSection(k string, v interface{}) error {
  err := s.vp.UnmarshalKey(k, v)
  if err != nil {
    return err
  }
  return nil
}
