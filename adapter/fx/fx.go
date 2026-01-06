package fx

import (
	tea "github.com/structx/teapot"
	"go.uber.org/fx/fxevent"
)

type eventLogger struct {
	l *tea.Logger
}

// interface compliance
var _ fxevent.Logger = (*eventLogger)(nil)

// New constructor fx event logger using tea logger
func New(logger *tea.Logger) fxevent.Logger {
	return &eventLogger{
		l: logger,
	}
}

// LogEvent implements [fxevent.Logger].
func (el *eventLogger) LogEvent(evt fxevent.Event) {
	switch e := evt.(type) {
	case *fxevent.BeforeRun:
		el.l.Debug(
			"[Fx] before run",
			tea.String("kind", e.Kind),
			tea.String("module_name", e.ModuleName),
			tea.String("name", e.Name),
		)
	case *fxevent.Decorated:
		if e.Err != nil {
			el.l.Error("[Fx] decorate", tea.Error(e.Err))
			return
		}
	case *fxevent.Invoked:
		if e.Err != nil {
			el.l.Error("[Fx] invoked", tea.Error(e.Err))
			return
		}
	case *fxevent.Invoking:
		el.l.Debug("[Fx] invoking",
			tea.String("function_name", e.FunctionName),
			tea.String("module_name", e.ModuleName),
		)
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			el.l.Error("[Fx] logger initialized", tea.Error(e.Err))
			return
		}
		el.l.Debug("[Fx] logger initialized", tea.String("constructor_name", e.ConstructorName))
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			el.l.Error("[Fx] on start executed", tea.Error(e.Err))
			return
		}
		el.l.Debug("[Fx] on start executed",
			tea.String("caller_name", e.CallerName),
			tea.String("function_name", e.FunctionName),
			tea.String("method", e.Method),
			tea.Int64("runtime", int64(e.Runtime)),
		)
	case *fxevent.OnStartExecuting:
		el.l.Debug("[Fx] on start executing",
			tea.String("caller_name", e.CallerName),
			tea.String("function_name", e.FunctionName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			el.l.Error("[Fx] on stop executed", tea.Error(e.Err))
			return
		}
		el.l.Debug("[Fx] on stop executed",
			tea.String("caller_name", e.CallerName),
			tea.String("function_name", e.FunctionName),
		)
	case *fxevent.OnStopExecuting:
		el.l.Debug("[Fx] on stop executing",
			tea.String("caller_name", e.CallerName),
			tea.String("function_name", e.FunctionName),
		)
	case *fxevent.Provided:
		if e.Err != nil {
			el.l.Error("[Fx] provided", tea.Error(e.Err))
			return
		}
		el.l.Debug("[Fx] provided",
			tea.String("constructor_name", e.ConstructorName),
			tea.String("module_name", e.ModuleName),
			tea.Bool("private", e.Private),
		)
	case *fxevent.Replaced:
		if e.Err != nil {
			el.l.Error("[Fx] replaced", tea.Error(e.Err))
			return
		}
		el.l.Debug("[Fx] replaced",
			tea.String("module_name", e.ModuleName),
			tea.StringSlice("module_trace", e.ModuleTrace),
			tea.StringSlice("output_type_names", e.OutputTypeNames),
			tea.StringSlice("stacktrace", e.StackTrace),
		)
	case *fxevent.RolledBack:
		if e.Err != nil {
			el.l.Error("[Fx] rolled back", tea.Error(e.Err))
		}
	case *fxevent.RollingBack:
		if e.StartErr != nil {
			el.l.Error("[Fx] rolling back", tea.Error(e.StartErr))
		}
	case *fxevent.Run:
		if e.Err != nil {
			el.l.Error("[Fx] run", tea.Error(e.Err))
			return
		}
		el.l.Debug("[Fx] run",
			tea.String("kind", e.Kind),
			tea.String("module_name", e.ModuleName),
			tea.String("name", e.Name),
		)
	case *fxevent.Started:
		if e.Err != nil {
			el.l.Error("[Fx] started", tea.Error(e.Err))
		}
	case *fxevent.Stopped:
		if e.Err != nil {
			el.l.Error("[Fx] stopped", tea.Error(e.Err))
		}
	case *fxevent.Stopping:
		el.l.Debug("[Fx] stopping", tea.String("signal", e.Signal.String()))
	case *fxevent.Supplied:
		if e.Err != nil {
			el.l.Debug("[Fx] supplied",
				tea.String("module_name", e.ModuleName),
				tea.String("type_name", e.TypeName),
				tea.StringSlice("module_trace", e.ModuleTrace),
				tea.StringSlice("stacktrace", e.StackTrace),
			)
		}
	}
}
