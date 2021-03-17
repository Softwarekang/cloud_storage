package app_config

type MinioConfig struct {
	EndPoint        string `ini:"endpoint"`
	AccessKeyID     string `ini:"accessKeyID"`
	SecretAccessKey string `ini:"secretAccessKey"`
	SSL             bool   `ini:"ssl"`
	Region          string `ini:"region"`
}
