/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package tasks

import (
	"fmt"

	"github.com/codoworks/go-boilerplate/pkg/config"

	"github.com/jedib0t/go-pretty/v6/table"
)

var Tasks = &AppTaskList{}

type AppTaskList struct {
	tasks []*Task
}

func (atl *AppTaskList) Init() {
	atl.tasks = []*Task{}
}

func (atl *AppTaskList) GetTasks() []*Task {
	return atl.tasks
}

func (atl *AppTaskList) GetTask(name string) *Task {
	for _, t := range atl.tasks {
		if t.Name == name {
			return t
		}
	}
	return nil
}

func (atl *AppTaskList) AddTask(task *Task) {
	if atl.tasks == nil {
		atl.Init()
	}
	atl.tasks = append(atl.tasks, task)
}

func (atl *AppTaskList) PrintTasks() {
	t := table.NewWriter()
	t.SetTitle("Available Tasks List")
	t.AppendHeader(table.Row{"Task Name", "Description"})

	for _, task := range atl.tasks {
		t.AppendRow(table.Row{task.Name, task.Description})
	}
	t.SortBy([]table.SortBy{
		{Name: "Task Name", Mode: table.Asc},
		{Name: "Description", Mode: table.Asc},
	})
	t.AppendFooter(table.Row{"TOTAL", len(atl.tasks)})
	t.SetStyle(table.StyleRounded)

	fmt.Println(t.Render())
}

type TaskEnv struct {
	Env *config.SEnv
}

func (te *TaskEnv) SetEnv(env *config.SEnv) {
	te.Env = env
}

type Task struct {
	Name         string
	Description  string
	RequiredArgs []string
	Run          func(env *TaskEnv, args map[string]string) error
}

func (t *Task) Execute(args map[string]string) error {
	env := &TaskEnv{}
	env.SetEnv(config.Env)
	if err := validateArgs(args, t.RequiredArgs); err != nil {
		return err
	}
	if err := t.Run(env, args); err != nil {
		return err
	}
	return nil
}

func validateArgs(args map[string]string, requiredArgs []string) error {
	for _, arg := range requiredArgs {
		if val, ok := args[arg]; !ok || val == "" {
			return fmt.Errorf("missing required argument: %s", arg)
		}
	}
	return nil
}
