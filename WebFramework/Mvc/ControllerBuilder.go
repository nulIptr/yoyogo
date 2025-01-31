package Mvc

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/WebFramework/View"
	"github.com/yoyofxteam/reflectx"
	"strings"
)

// ControllerBuilder: controller builder
type ControllerBuilder struct {
	configuration    Abstractions.IConfiguration
	mvcRouterHandler *RouterHandler
}

// NewControllerBuilder new controller builder
func NewControllerBuilder() *ControllerBuilder {
	return &ControllerBuilder{mvcRouterHandler: NewMvcRouterHandler()}
}

// add views to mvc
func (builder *ControllerBuilder) AddViews(option *View.Option) {
	XLog.GetXLogger("ControllerBuilder").Debug("add mvc views: %s", option.Path)
	builder.mvcRouterHandler.Options.ViewOption = option
}

func (builder *ControllerBuilder) AddViewsByConfig() {
	XLog.GetXLogger("ControllerBuilder").Debug("add mvc views: %s")
	if builder.configuration != nil {
		section := builder.configuration.GetSection("yoyogo.application.server.views")
		option := &View.Option{}
		section.Unmarshal(option)
		builder.mvcRouterHandler.Options.ViewOption = option
	}

	//builder.mvcRouterHandler.Options.ViewOption =

}

func (builder *ControllerBuilder) SetViewEngine(viewEngine View.IViewEngine) {
	builder.mvcRouterHandler.ViewEngine = viewEngine
}

func (builder *ControllerBuilder) SetConfiguration(configuration Abstractions.IConfiguration) {
	builder.configuration = configuration
}

// add filter to mvc
func (builder *ControllerBuilder) AddFilter(pattern string, actionFilter IActionFilter) {
	XLog.GetXLogger("ControllerBuilder").Debug("add mvc filter: %s", pattern)
	chain := NewActionFilterChain(pattern, actionFilter)
	builder.mvcRouterHandler.ControllerFilters = append(builder.mvcRouterHandler.ControllerFilters, chain)
}

// SetupOptions , setup mvc builder options
func (builder *ControllerBuilder) SetupOptions(configOption func(options Options)) {
	configOption(builder.mvcRouterHandler.Options)
}

// AddController add controller (ctor) to ioc.
func (builder *ControllerBuilder) AddController(controllerCtor interface{}) {
	logger := XLog.GetXLogger("ControllerBuilder")

	controllerName, controllerType := reflectx.GetCtorFuncOutTypeName(controllerCtor)
	controllerName = strings.ToLower(controllerName)
	// Create Controller and Action descriptors
	descriptor := NewControllerDescriptor(controllerName, controllerType, controllerCtor)
	builder.mvcRouterHandler.ControllerDescriptors[controllerName] = descriptor

	logger.Debug("add mvc controller: %s", controllerName)
	for _, desc := range descriptor.GetActionDescriptors() {
		logger.Debug("add mvc controller action: %s", desc.ActionName)
	}
}

// GetControllerDescriptorList is get controller descriptor array
func (builder *ControllerBuilder) GetControllerDescriptorList() []ControllerDescriptor {
	values := make([]ControllerDescriptor, 0, len(builder.mvcRouterHandler.ControllerDescriptors))
	for _, value := range builder.mvcRouterHandler.ControllerDescriptors {
		values = append(values, value)
	}
	return values
}

// GetControllerDescriptorByName get controller descriptor by controller name
func (builder *ControllerBuilder) GetControllerDescriptorByName(name string) ControllerDescriptor {
	return builder.mvcRouterHandler.ControllerDescriptors[name]
}

// GetMvcOptions get mvc options
func (builder *ControllerBuilder) GetMvcOptions() Options {
	return builder.mvcRouterHandler.Options
}

// GetRouterHandler is get mvc router handler.
func (builder *ControllerBuilder) GetRouterHandler() *RouterHandler {
	return builder.mvcRouterHandler
}
