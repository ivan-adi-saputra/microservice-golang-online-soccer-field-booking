package util

import (
	"os"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func BindFromJSON(dest any, filename string, path string) error {
	v := viper.New()

	v.SetConfigType("json")
	v.AddConfigPath(path)
	v.SetConfigName(filename)

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&dest); err != nil {
		logrus.Errorf("failed to unmarshal: %v", err)
		return err
	}

	return nil
}

func SetEnvFromConsulKV(v *viper.Viper) error {
	env := make(map[string]string)

	if err := v.Unmarshal(&env); err != nil {
		logrus.Errorf("failed to unmarshal environment: %v", err)
		return err
	}

	for k, v := range env {
		var (
			valOf = reflect.ValueOf(v)
			val   string
		)

		switch valOf.Kind() {
		case reflect.String:
			val = valOf.String()
		case reflect.Int:
			val = strconv.Itoa(int(valOf.Int()))
		case reflect.Uint:
			val = strconv.Itoa(int(valOf.Uint()))
		case reflect.Float32:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Float64:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Bool:
			val = strconv.FormatBool(valOf.Bool())
		default:
			panic("unsupported type")
		}

		if err := os.Setenv(k, val); err != nil {
			logrus.Errorf("failed to set environment variable: %v", err)
			return err
		}
	}

	return nil
}

func BindFromConsul(dest any, endPoint string, path string) error {
	v := viper.New()
	v.SetConfigType("json")
	if err := v.AddRemoteProvider("consul", endPoint, path); err != nil {
		logrus.Errorf("failed to add consul provider: %v", err)
		return err
	}

	if err := v.ReadRemoteConfig(); err != nil {
		logrus.Errorf("failed to read consul config: %v", err)
		return err
	}

	if err := v.Unmarshal(&dest); err != nil {
		logrus.Errorf("failed to unmarshal: %v", err)
		return err
	}

	if err := SetEnvFromConsulKV(v); err != nil {
		logrus.Errorf("failed to set environment variables from consul: %v", err)
		return err
	}

	return nil
}
