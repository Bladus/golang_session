package config

type Config struct {
    Server  struct {
        Addr         int
        ReadTimeout  int
        WriteTimeout int
    }
}