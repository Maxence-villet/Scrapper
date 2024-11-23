package config

import "time"

type ConfigHandler interface {
	GetKey()
	GetToken()
}

type Config struct {
	key        []byte
	token      []byte
	channelId  string
	image      string
	timeReload time.Duration
}

func (c *Config) GetKey() []byte {
	return c.key
}

func (c *Config) GetToken() []byte {

	return c.token
}

func (c *Config) GetChannelId() string {
	return c.channelId
}

func (c *Config) GetImage() string {
	return c.image
}

func (c *Config) GetTimeReload() time.Duration {
	return c.timeReload
}

func NewConfig() *Config {
	c := &Config{
		key:        []byte("AiQBJDc2XMsGX5tDMXoQV44EC7tFbm4J"),
		token:      []byte("Bot MTMwNDU5Mzg3ODQ1MTYyMTkxOA.G3Zl5P.oH3BQ-ZpoqMLj_y5RthniuRh0NwY7ulkndIVWo"),
		channelId:  "1304593207732080685",
		image:      "https://upload.wikimedia.org/wikipedia/commons/thumb/e/e3/Udemy_logo.svg/1200px-Udemy_logo.svg.png",
		timeReload: time.Duration(23 * time.Hour),
	}

	return c
}
