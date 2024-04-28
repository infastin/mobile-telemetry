package fxzerolog

import (
	"github.com/rs/zerolog"
	"go.uber.org/fx/fxevent"
)

type Logger struct {
	logger zerolog.Logger
}

func New(lg zerolog.Logger) *Logger {
	return &Logger{
		logger: lg,
	}
}

func (l *Logger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.logger.Trace().
			Str("callee", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStart hook executing")
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.logger.Err(e.Err).
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Msg("OnStart hook failed")
		} else {
			l.logger.Trace().
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStart hook executed")
		}
	case *fxevent.OnStopExecuting:
		l.logger.Trace().
			Str("callee", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStop hook executing")
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.logger.Err(e.Err).
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Msg("OnStop hook failed")
		} else {
			l.logger.Trace().
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStop hook executed")
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.logger.Err(e.Err).
				Str("type", e.TypeName).
				Strs("moduletrace", e.ModuleTrace).
				Strs("stacktrace", e.StackTrace).
				Str("module", e.ModuleName).
				Msg("error encountered while applying options")
		} else {
			l.logger.Trace().
				Str("type", e.TypeName).
				Strs("moduletrace", e.ModuleTrace).
				Strs("stacktrace", e.StackTrace).
				Str("module", e.ModuleName).
				Msg("supplied")
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.logger.Trace().
				Str("constructor", e.ConstructorName).
				Strs("stacktrace", e.StackTrace).
				Strs("moduletrace", e.ModuleTrace).
				Str("module", e.ModuleName).
				Str("type", rtype).
				Bool("private", e.Private).
				Msg("provided")
		}
		if e.Err != nil {
			l.logger.Err(e.Err).
				Strs("moduletrace", e.ModuleTrace).
				Strs("stacktrace", e.StackTrace).
				Str("module", e.ModuleName).
				Msg("error encountered while applying options")
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			l.logger.Trace().
				Strs("stacktrace", e.StackTrace).
				Strs("moduletrace", e.ModuleTrace).
				Str("module", e.ModuleName).
				Str("type", rtype).
				Msg("replaced")
		}
		if e.Err != nil {
			l.logger.Err(e.Err).
				Strs("moduletrace", e.ModuleTrace).
				Strs("stacktrace", e.StackTrace).
				Str("module", e.ModuleName).
				Msg("error encountered while replacing")
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.logger.Trace().
				Str("decorator", e.DecoratorName).
				Strs("stacktrace", e.StackTrace).
				Strs("moduletrace", e.ModuleTrace).
				Str("module", e.ModuleName).
				Str("type", rtype).
				Msg("decorated")
		}
		if e.Err != nil {
			l.logger.Err(e.Err).
				Strs("moduletrace", e.ModuleTrace).
				Strs("stacktrace", e.StackTrace).
				Str("module", e.ModuleName).
				Msg("error encountered while applying options")
		}
	case *fxevent.Run:
		if e.Err != nil {
			l.logger.Err(e.Err).
				Str("name", e.Name).
				Str("kind", e.Kind).
				Str("module", e.ModuleName).
				Msg("error returned")
		} else {
			l.logger.Trace().
				Str("name", e.Name).
				Str("kind", e.Kind).
				Str("module", e.ModuleName).
				Msg("run")
		}
	case *fxevent.Invoking:
		l.logger.Trace().
			Str("function", e.FunctionName).
			Str("module", e.ModuleName).
			Msg("invoking")
	case *fxevent.Invoked:
		if e.Err != nil {
			l.logger.Err(e.Err).
				Str("stack", e.Trace).
				Str("function", e.FunctionName).
				Str("module", e.ModuleName).
				Msg("invoke failed")
		}
	case *fxevent.Stopping:
		l.logger.Trace().
			Str("signal", e.Signal.String()).
			Msg("received signal")
	case *fxevent.Stopped:
		if e.Err != nil {
			l.logger.Err(e.Err).Msg("stop failed")
		}
	case *fxevent.RollingBack:
		l.logger.Err(e.StartErr).Msg("start failed, rolling back")
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.logger.Err(e.Err).Msg("rollback failed")
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.logger.Err(e.Err).Msg("start failed")
		} else {
			l.logger.Trace().Msg("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.logger.Err(e.Err).Msg("custom logger initialization failed")
		} else {
			l.logger.Trace().
				Str("function", e.ConstructorName).
				Msg("initialized custom fxevent.Logger")
		}
	}
}
