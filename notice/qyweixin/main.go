package main

import (
	"fmt"
	"os"
)

var envList = []string{
	"CORP_ID",
	"APP_SECRET",
	"AGENT_ID",
	"MESSAGE",
	"USERS",
	"PARTYS",
	"TAGS",
	"_WORKFLOW_TASK_DETAIL",
	"_WORKFLOW_FLAG_PAUSE_NOTICE",
	"_WORKFLOW_FLOW_PAUSE_HOOK_RESUME_API",
	"_WORKFLOW_FLOW_PAUSE_HOOK_STOP_API",
}

func main() {

	envs := make(map[string]string)

	for _, name := range envList {
		envs[name] = os.Getenv(name)
	}

	builder, err := NewBuilder(envs)
	if err != nil {
		fmt.Println("build failed: ", err)
		os.Exit(1)
	}
	if err := builder.run(); err != nil {
		fmt.Println("build failed: ", err)
		os.Exit(1)
	}
	fmt.Println("build success!")
}
