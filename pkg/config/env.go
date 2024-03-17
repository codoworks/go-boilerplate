/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package config

import (
	"fmt"
	"reflect"

	"github.com/codoworks/go-boilerplate/pkg/clients/logger"
	"github.com/codoworks/go-boilerplate/pkg/config/features"
	"github.com/codoworks/go-boilerplate/pkg/utils"

	"github.com/jedib0t/go-pretty/v6/table"
)

var Env *SEnv

func init() {
	Env = &SEnv{
		ServiceName: "go-boilerplate",
		Features:    &features.Features,
		EnvVars:     &EnvVars{},
		Version:     "0.0.1",
	}
}

type IEnv interface {
}

type SEnv struct {
	IEnv
	ServiceName string
	Features    *features.SFeatures
	EnvVars     *EnvVars
	Version     string
}

func (e *SEnv) Init() {
	e.EnvVars.Init()
}

func (e *SEnv) InitFeatures() {
	e.Features.Init(reflect.ValueOf(e.EnvVars))
}

func (e *SEnv) OverrideUsingFlags() {
	e.EnvVars.OverrideUsingFlags()
}

func (e *SEnv) OverrideLoggerUsingFlags() {
	e.EnvVars.OverrideLoggerUsingFlags()
}

func (e *SEnv) CheckAndSetDevMode() {
	if !DevModeFlag {
		return
	}
	// fmt.Println("Running in development mode. don't do this in production!")
	logger.Warn("Running in development mode. don't do this in production!")

	e.EnvVars.SetDevMode()
}

func (e *SEnv) PrintEnvironment() {

	v := reflect.ValueOf(*e.EnvVars)
	typeOfS := v.Type()
	t := table.NewWriter()
	t.SetTitle("Environment")
	t.AppendHeader(table.Row{"ENV", "Value"})
	for i := 0; i < v.NumField(); i++ {
		vOfFeature := v.Field(i)
		if vOfFeature.Kind() == reflect.Ptr {
			vOfFeature = vOfFeature.Elem()
		}
		if vOfFeature.Kind() == reflect.Slice {
			t.AppendRow(table.Row{
				typeOfS.Field(i).Tag.Get("mapstructure"),
				v.Field(i).Interface(),
			})
			continue
		}
		typeOfF := vOfFeature.Type()
		for j := 0; j < vOfFeature.NumField(); j++ {
			t.AppendRow(table.Row{
				typeOfF.Field(j).Tag.Get("mapstructure"),
				vOfFeature.Field(j).Interface(),
			})
		}
	}
	utils.SetTableBorderStyle(t, NoBorderFlag)
	fmt.Println(t.Render())
	fmt.Printf("\n")
}

func (e *SEnv) PrintServiceFeatures() {
	t := table.NewWriter()
	rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
	t.AppendHeader(table.Row{"Service Features", "Configuration"}, rowConfigAutoMerge)
	features := e.Features.GetFeatures()
	for _, feature := range features {
		feature.AppendFeatureToTable(t)
		t.AppendSeparator()
	}
	utils.SetTableBorderStyle(t, NoBorderFlag)
	fmt.Println(t.Render())
	fmt.Printf("\n")
}
