package main

import (
	"fmt"
	options "homework/proto_test/api"
	"strconv"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func main() {
	protoregistry.GlobalFiles.RangeFiles(func(fd protoreflect.FileDescriptor) bool {
		services := fd.Services()
		for i := 0; i < services.Len(); i++ {
			service := services.Get(i)
			if serviceHandler, _ := proto.GetExtension(service.Options(), options.E_ServiceHandler).(*options.ServiceHandler); serviceHandler != nil {
				fmt.Println()
				fmt.Println("--- service ---")
				fmt.Println("service name: " + string(service.FullName()))

				if serviceHandler.Authorization != "" {
					fmt.Println("use interceptor authorization: " + serviceHandler.Authorization)
				}
				fmt.Println("--- service ---")
			}

			methods := service.Methods()
			for k := 0; k < methods.Len(); k++ {
				method := methods.Get(k)
				if methodHandler, _ := proto.GetExtension(method.Options(), options.E_MethodHandler).(*options.MethodHandler); methodHandler != nil {
					fmt.Println()
					fmt.Println("--- method ---")
					fmt.Println("method name: " + string(method.FullName()))
					if methodHandler.Whitelist != "" {
						fmt.Println("use interceptor whitelist: " + methodHandler.Whitelist)
					}

					fmt.Println("use interceptor logger: " + strconv.FormatBool(methodHandler.Logger))

					fmt.Println("--- method ---")
				}
			}
		}

		return true
	})
}
