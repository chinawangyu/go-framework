#!/bin/bash

GO ?= go

WORK_DIR=$(shell pwd)

export GOPROXY=https://mirrors.aliyun.com/goproxy/
export GO111MODULE=on

all:
	$(GO) install github.com/google/wire
	rm -rf $(WORK_DIR)/internal/di/wire_gen.go
	wire $(WORK_DIR)/internal/di/*
	$(GO) build -o http  $(WORK_DIR)/cmd/main.go


clean:
		rm -rf http
		rm -rf ./logs/*

.PHONY : clean
