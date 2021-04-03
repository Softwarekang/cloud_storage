/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"time"
)

import (
	"user-client/app/com/wpbs/DTO"
	"user-client/app/com/wpbs/router"
	"user-client/app/com/wpbs/service"
)

import (
	hessian "github.com/apache/dubbo-go-hessian2"
	_ "github.com/apache/dubbo-go/cluster/cluster_impl"
	_ "github.com/apache/dubbo-go/cluster/loadbalance"
	_ "github.com/apache/dubbo-go/common/proxy/proxy_factory"
	"github.com/apache/dubbo-go/config"
	_ "github.com/apache/dubbo-go/filter/filter_impl"
	_ "github.com/apache/dubbo-go/protocol/dubbo"
	_ "github.com/apache/dubbo-go/registry/protocol"
	_ "github.com/apache/dubbo-go/registry/zookeeper"
	"github.com/gin-gonic/gin"
)

var (
	survivalTimeout int = 10e9
)

// they are necessary:
// 		export CONF_CONSUMER_FILE_PATH="xxx"
// 		export APP_LOG_CONF_FILE="xxx"
func main() {
	// register DTO to hessian
	registerPOJO()
	// register service
	registerConsumerService()
	// dubbo config load
	config.Load()
	time.Sleep(3e9)
	// gin start
	r := gin.Default()
	// load router
	router.LoadRouters(r)
	r.Run(":9001")
}

func registerConsumerService() {
	userService := new(service.UserService)
	serverCheckProvider := new(service.ServerCheckService)
	fileService := new(service.FileService)
	memoryService := new(service.MemoryService)
	config.SetConsumerService(userService)
	config.SetConsumerService(serverCheckProvider)
	config.SetConsumerService(fileService)
	config.SetConsumerService(memoryService)

}

func registerPOJO() {
	hessian.RegisterPOJO(&DTO.User{})
	hessian.RegisterPOJO(&DTO.ServerCheck{})
	hessian.RegisterPOJO(&DTO.MonoFile{})
	hessian.RegisterPOJO(&DTO.FileList{})
	hessian.RegisterPOJO(&DTO.GetFileList{})
	hessian.RegisterPOJO(&DTO.DeleteFile{})
	hessian.RegisterPOJO(&DTO.Memory{})
}
