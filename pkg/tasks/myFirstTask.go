/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package tasks

import "fmt"

func init() {

	var t = &Task{
		Name:         "myFirstTask",
		Description:  "This is my first task",
		RequiredArgs: []string{"firstName", "lastName"},
		Run:          execMyFirstTask,
	}

	Tasks.AddTask(t)

}

func execMyFirstTask(env *TaskEnv, args map[string]string) error {

	// This is an example task
	// You can write your own tasks here
	// ...
	// the idea behind tasks is to create a one off process that can be run from the command line
	// client or from a cron job etc, without having to go through the trouble of creating a new
	// cmd, learning how to use the cobra library, going through the initialization process every
	// time you want to create a new process that is not a part of the main cli.
	//
	// imagine having a task that sends an email to all users who have not logged in for 30 days
	// or a task to clean up the database, setting up a new user, etc

	fmt.Printf("Hello %s %s\n", args["firstName"], args["lastName"])
	fmt.Println("...")
	fmt.Println("This is an example task")
	fmt.Println("You can write your own tasks here")
	fmt.Println("...")
	fmt.Println(`The idea behind tasks is to create a one off process that can be run from the command line 
client or from a cron job etc, without having to go through the trouble of creating a new cmd, learning
how to use the cobra library, going through the initialization process every time you want to create a new
process that is not a part of the main cli.`)
	fmt.Println("...")
	fmt.Println(`Imagine having a task that sends an email to all users who have not logged in for 30 days
or a task to clean up the database, setting up a new user, resetting a password, etc`)
	fmt.Println("...")
	fmt.Println(`You can pass arguments to this task, re-run the same command and add the arguments at the end of it,
arguments are passed as key value pairs, for example: myFirstTask key1=value1 key2=value2 key3=value3`)
	fmt.Printf("Arguments: %v\n", args)
	fmt.Println("...")

	return nil

}
