package logger_test

import (
	"bytes"
	"encoding/json"
	"go-base-end/logger"
	"os"
	"testing"
)

func TestDefaultLoggerPanicf(t *testing.T) {
	os.Setenv("IS_PROD_ENV", "true")

	lg := logger.Default()

	buf := new(bytes.Buffer)
	lg.SetOutput(buf)

	const msg = "THIS IS A VERY IMPORTANT INFO"

	defer func() {
		os.Setenv("IS_PROD_ENV", "false")

		if err := recover(); err == nil {
			t.Fatalf("failed under production environment:" +
				" cannot catch panic caused by Panicf")
		}

		output := struct {
			Level string `json:"level"`
			Msg   string `json:"msg"`
			Time  string `json:"time"`
		}{}
		if err := json.Unmarshal(buf.Bytes(), &output); err != nil {
			t.Fatalf("failed under production environment (not in json format): %v", err)
		}

		level := "panic"
		if output.Level != level {
			t.Fatalf("failed under production environment: unexpected level %s, expected %s",
				output.Level, level)
		}
		if output.Msg != msg {
			t.Fatalf("failed under production environment: unexpected msg %s, expected %s", output.Msg, msg)
		}

		defer func() {
			if err := recover(); err == nil {
				t.Fatalf("failed under debug environment:" +
					" cannot catch panic caused by Panicf")
			}

			if buf.Len() <= 0 {
				t.Fatalf("failed under debug environment (no output at all)")
			}
		}()
		lg = logger.Default()
		buf.Reset()

		lg.SetOutput(buf)

		lg.Panicf(msg)

		t.Fatalf("failed under debug environment: Panicf didn't cause panic")
	}()

	lg.Panicf(msg)

	t.Fatalf("failed under production environment: Panicf didn't cause panic")
}

func TestDefaultLoggerWarnf(t *testing.T) {
	os.Setenv("IS_PROD_ENV", "true")

	lg := logger.Default()

	buf := new(bytes.Buffer)
	lg.SetOutput(buf)

	const msg = "THIS IS A VERY IMPORTANT INFO"
	lg.Warnf(msg)

	output := struct {
		Level string `json:"level"`
		Msg   string `json:"msg"`
		Time  string `json:"time"`
	}{}
	if err := json.Unmarshal(buf.Bytes(), &output); err != nil {
		t.Fatalf("failed under production environment (not in json format): %v", err)
	}

	level := "warning"
	if output.Level != level {
		t.Fatalf("failed under production environment: unexpected level %s, expected %s", output.Level, level)
	}
	if output.Msg != msg {
		t.Fatalf("failed under production environment: unexpected msg %s, expected %s", output.Msg, msg)
	}

	os.Setenv("IS_PROD_ENV", "false")

	lg = logger.Default()

	buf.Reset()
	lg.SetOutput(buf)

	lg.Warnf(msg)

	if buf.Len() <= 0 {
		t.Fatalf("failed under debug environment (no output at all)")
	}
}

func TestDefaultLoggerInfof(t *testing.T) {
	os.Setenv("IS_PROD_ENV", "true")

	lg := logger.Default()

	buf := new(bytes.Buffer)
	lg.SetOutput(buf)

	const msg = "THIS IS A VERY IMPORTANT INFO"
	lg.Infof(msg)

	if buf.Len() <= 0 {
		t.Fatalf("failed under production environment (no output at all)")
	}

	output := struct {
		Level string `json:"level"`
		Msg   string `json:"msg"`
		Time  string `json:"time"`
	}{}
	if err := json.Unmarshal(buf.Bytes(), &output); err != nil {
		t.Fatalf("failed under production environment (not in json format): %v", err)
	}

	level := "info"
	if output.Level != level {
		t.Fatalf("failed under production environment: unexpected level %s, expected %s", output.Level, level)
	}
	if output.Msg != msg {
		t.Fatalf("failed under production environment: unexpected msg %s, expected %s", output.Msg, msg)
	}

	os.Setenv("IS_PROD_ENV", "false")

	lg = logger.Default()

	buf.Reset()
	lg.SetOutput(buf)

	lg.Infof(msg)

	if buf.Len() <= 0 {
		t.Fatalf("failed under debug environment (no output at all)")
	}
}

func TestDefaultLoggerDebugf(t *testing.T) {
	lg := logger.Default()

	buf := new(bytes.Buffer)
	lg.SetOutput(buf)

	const msg = "THIS IS A VERY IMPORTANT INFO"
	lg.Debugf(msg)

	if buf.Len() <= 0 {
		t.Fatalf("failed under debug environment (no output at all)")
	}
}
