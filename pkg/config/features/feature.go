/*
Copyright © 2023 Codoworks
Author:  Dexter Codo
Contact: dexter.codo@gmail.com
*/
package features

import (
	"fmt"
	"reflect"

	"github.com/codoworks/go-boilerplate/pkg/clients/logger"

	"github.com/jedib0t/go-pretty/v6/table"
)

var Features = SFeatures{}

// func Init(v reflect.Value) {
// 	for _, feature := range Features.f {
// 		feature.ResolveFeatureEnabledState(v)
// 		feature.Configure(v)
// 		feature.ValidateReadiness()
// 	}
// }

type SFeatures struct {
	f []*Feature
}

func (f *SFeatures) Init(v reflect.Value) {
	logger.Debug("Initializing features ...")
	for _, feature := range f.f {
		feature.ResolveFeatureEnabledState(v)
		feature.Configure(v)
		feature.ValidateReadiness()
	}
}

func (f *SFeatures) Add(feature *Feature) {
	f.f = append(f.f, feature)
}

func (f *SFeatures) GetFeatures() []*Feature {
	return f.f
}

func (f *SFeatures) GetFeatureByName(name string) *Feature {
	for _, feature := range f.f {
		if feature.Name == name {
			return feature
		}
	}
	return nil
}

type IFeature interface {
	IsEnabled() bool
	IsConfigured() bool
	IsReady() bool
	CanStart() bool
	ValidateReadiness()
	Configure(v reflect.Value)
	ResolveFeatureEnabledState(v reflect.Value)
	PrintFeature()
	AppendFeatureToTable(t table.Writer)
}

type Feature struct {
	IFeature
	Name         string
	enabled      bool
	configured   bool
	ready        bool
	requirements []string
	Config       interface{}
}

func (f *Feature) IsEnabled() bool {
	return f.enabled
}

func (f *Feature) IsConfigured() bool {
	return f.configured
}

func (f *Feature) IsReady() bool {
	return f.ready
}

func (f *Feature) CanStart() bool {
	result := true
	for _, requirement := range f.requirements {
		val := f.GetConfigByName(requirement)
		if val == "" {
			result = false
			break
		}
	}
	return result
}

func (f *Feature) ValidateReadiness() {
	if !f.IsEnabled() {
		logger.Debug("Feature: %s - is not enabled, skipping ...", f.Name)
		return
	}
	if !f.IsConfigured() {
		logger.Debug("Feature: %s - is not configured correctly.", f.Name)
		return
	}
	if !f.CanStart() {
		logger.Debug("Feature: %s - is not ready. Config not satisfied. %s. %s. Use cmd: 'info feature %s' for more.",
			f.Name,
			f.getRequirementsString(),
			f.getCurrentConfigString(),
			f.Name)
		return
	}
	f.ready = true
	logger.Debug("Feature: %s - is ready.", f.Name)

}

func (f *Feature) getRequirementsString() string {
	str := "["
	for i, requirement := range f.requirements {
		str += fmt.Sprintf("'%s'", requirement)
		if i < len(f.requirements)-1 {
			str += ", "
		}
	}
	str += "]"
	return fmt.Sprintf("Requirements: %s", str)
}

func (f *Feature) getCurrentConfigString() string {
	str := "{"
	for i := 0; i < f.Config.(reflect.Value).NumField(); i++ {
		fieldValue := f.Config.(reflect.Value).Field(i).Interface()
		if fieldValue.(string) == "" {
			fieldValue = "_<MISSING>_"
		}
		str += fmt.Sprintf("'%s': '%s'",
			f.Config.(reflect.Value).Type().Field(i).Name,
			fieldValue.(string))
		if i < f.Config.(reflect.Value).NumField()-1 {
			str += ", "
		}

		// field := f.Config.(reflect.Value).Field(i)
		// fmt.Println(field)
	}
	str += "}"
	return fmt.Sprintf("Has: %s", str)
}

func (f *Feature) GetConfigByName(name string) string {
	return f.Config.(reflect.Value).FieldByName(name).String()
}

func (f *Feature) Configure(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	cType := reflect.TypeOf(f.Config)
	if cType.Kind() == reflect.Ptr {
		cType = cType.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		configField := v.Field(i)
		if configField.Type() == cType {
			f.Config = configField
			f.configured = true
			break
		}
	}
}

func (f *Feature) ResolveFeatureEnabledState(v reflect.Value) {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	disabled := v.FieldByName("DisableFeatures")
	for _, v := range disabled.Interface().([]string) {
		if v == f.Name {
			f.enabled = false
			return
		}
	}
	f.enabled = true
}

func (f *Feature) GetConfigByNameForPrint(name string) string {
	config, _ := f.Config.(map[string]string)[name]
	if config == "" {
		return " <-------------------------------------- MISSING"
	}
	return config
}

func (f *Feature) PrintFeature() {
	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)

	t.AppendRow(table.Row{"Feature", f.Name})
	t.AppendSeparator()
	t.AppendRow(table.Row{"Enabled", f.IsEnabled()})
	t.AppendRow(table.Row{"Required Env Vars", f.requirements})
	t.AppendRow(table.Row{"Requirments Satisfied", f.IsReady()})
	t.AppendRow(table.Row{"Config:"})
	for _, requirement := range f.requirements {
		t.AppendRow(table.Row{
			fmt.Sprintf("- %s", requirement),
			f.GetConfigByNameForPrint(requirement)})
	}
	fmt.Println(t.Render())
}

func (f *Feature) AppendFeatureToTable(t table.Writer) {
	t.AppendRow(table.Row{"Feature", f.Name})
	t.AppendRow(table.Row{"Enabled", f.IsEnabled()})
	t.AppendRow(table.Row{"Required Env Vars", f.requirements})
	t.AppendRow(table.Row{"Requirments Satisfied", f.IsReady()})
	t.AppendRow(table.Row{"Config:"})
	for _, requirement := range f.requirements {
		t.AppendRow(table.Row{
			fmt.Sprintf("- %s", requirement),
			f.GetConfigByNameForPrint(requirement)})
	}
}
