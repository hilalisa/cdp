// Code generated by cdpgen. DO NOT EDIT.

// Package debugger implements the Debugger domain. Debugger domain exposes JavaScript debugging capabilities. It allows setting and removing breakpoints, stepping through execution, exploring stack traces, etc.
package debugger

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Debugger domain. Debugger domain exposes JavaScript debugging capabilities. It allows setting and removing breakpoints, stepping through execution, exploring stack traces, etc.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Debugger domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Enable invokes the Debugger method. Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Debugger.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "Enable", Err: err}
	}
	return
}

// Disable invokes the Debugger method. Disables debugger for given page.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Debugger.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "Disable", Err: err}
	}
	return
}

// SetBreakpointsActive invokes the Debugger method. Activates / deactivates all breakpoints on the page.
func (d *domainClient) SetBreakpointsActive(ctx context.Context, args *SetBreakpointsActiveArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setBreakpointsActive", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setBreakpointsActive", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetBreakpointsActive", Err: err}
	}
	return
}

// SetSkipAllPauses invokes the Debugger method. Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
func (d *domainClient) SetSkipAllPauses(ctx context.Context, args *SetSkipAllPausesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setSkipAllPauses", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setSkipAllPauses", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetSkipAllPauses", Err: err}
	}
	return
}

// SetBreakpointByURL invokes the Debugger method. Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in locations property. Further matching script parsing will result in subsequent breakpointResolved events issued. This logical breakpoint will survive page reloads.
func (d *domainClient) SetBreakpointByURL(ctx context.Context, args *SetBreakpointByURLArgs) (reply *SetBreakpointByURLReply, err error) {
	reply = new(SetBreakpointByURLReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setBreakpointByUrl", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setBreakpointByUrl", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetBreakpointByURL", Err: err}
	}
	return
}

// SetBreakpoint invokes the Debugger method. Sets JavaScript breakpoint at a given location.
func (d *domainClient) SetBreakpoint(ctx context.Context, args *SetBreakpointArgs) (reply *SetBreakpointReply, err error) {
	reply = new(SetBreakpointReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setBreakpoint", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setBreakpoint", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetBreakpoint", Err: err}
	}
	return
}

// RemoveBreakpoint invokes the Debugger method. Removes JavaScript breakpoint.
func (d *domainClient) RemoveBreakpoint(ctx context.Context, args *RemoveBreakpointArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.removeBreakpoint", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.removeBreakpoint", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "RemoveBreakpoint", Err: err}
	}
	return
}

// GetPossibleBreakpoints invokes the Debugger method. Returns possible locations for breakpoint. scriptId in start and end range locations should be the same.
func (d *domainClient) GetPossibleBreakpoints(ctx context.Context, args *GetPossibleBreakpointsArgs) (reply *GetPossibleBreakpointsReply, err error) {
	reply = new(GetPossibleBreakpointsReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.getPossibleBreakpoints", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.getPossibleBreakpoints", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "GetPossibleBreakpoints", Err: err}
	}
	return
}

// ContinueToLocation invokes the Debugger method. Continues execution until specific location is reached.
func (d *domainClient) ContinueToLocation(ctx context.Context, args *ContinueToLocationArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.continueToLocation", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.continueToLocation", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "ContinueToLocation", Err: err}
	}
	return
}

// PauseOnAsyncTask invokes the Debugger method.
func (d *domainClient) PauseOnAsyncTask(ctx context.Context, args *PauseOnAsyncTaskArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.pauseOnAsyncTask", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.pauseOnAsyncTask", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "PauseOnAsyncTask", Err: err}
	}
	return
}

// StepOver invokes the Debugger method. Steps over the statement.
func (d *domainClient) StepOver(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Debugger.stepOver", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "StepOver", Err: err}
	}
	return
}

// StepInto invokes the Debugger method. Steps into the function call.
func (d *domainClient) StepInto(ctx context.Context, args *StepIntoArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.stepInto", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.stepInto", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "StepInto", Err: err}
	}
	return
}

// StepOut invokes the Debugger method. Steps out of the function call.
func (d *domainClient) StepOut(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Debugger.stepOut", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "StepOut", Err: err}
	}
	return
}

// Pause invokes the Debugger method. Stops on the next JavaScript statement.
func (d *domainClient) Pause(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Debugger.pause", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "Pause", Err: err}
	}
	return
}

// ScheduleStepIntoAsync invokes the Debugger method. This method is deprecated - use Debugger.stepInto with breakOnAsyncCall and Debugger.pauseOnAsyncTask instead. Steps into next scheduled async task if any is scheduled before next pause. Returns success when async task is actually scheduled, returns error if no task were scheduled or another scheduleStepIntoAsync was called.
func (d *domainClient) ScheduleStepIntoAsync(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Debugger.scheduleStepIntoAsync", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "ScheduleStepIntoAsync", Err: err}
	}
	return
}

// Resume invokes the Debugger method. Resumes JavaScript execution.
func (d *domainClient) Resume(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Debugger.resume", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "Resume", Err: err}
	}
	return
}

// SearchInContent invokes the Debugger method. Searches for given string in script content.
func (d *domainClient) SearchInContent(ctx context.Context, args *SearchInContentArgs) (reply *SearchInContentReply, err error) {
	reply = new(SearchInContentReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.searchInContent", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.searchInContent", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SearchInContent", Err: err}
	}
	return
}

// SetScriptSource invokes the Debugger method. Edits JavaScript source live.
func (d *domainClient) SetScriptSource(ctx context.Context, args *SetScriptSourceArgs) (reply *SetScriptSourceReply, err error) {
	reply = new(SetScriptSourceReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setScriptSource", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setScriptSource", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetScriptSource", Err: err}
	}
	return
}

// RestartFrame invokes the Debugger method. Restarts particular call frame from the beginning.
func (d *domainClient) RestartFrame(ctx context.Context, args *RestartFrameArgs) (reply *RestartFrameReply, err error) {
	reply = new(RestartFrameReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.restartFrame", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.restartFrame", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "RestartFrame", Err: err}
	}
	return
}

// GetScriptSource invokes the Debugger method. Returns source for the script with given id.
func (d *domainClient) GetScriptSource(ctx context.Context, args *GetScriptSourceArgs) (reply *GetScriptSourceReply, err error) {
	reply = new(GetScriptSourceReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.getScriptSource", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.getScriptSource", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "GetScriptSource", Err: err}
	}
	return
}

// SetPauseOnExceptions invokes the Debugger method. Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is none.
func (d *domainClient) SetPauseOnExceptions(ctx context.Context, args *SetPauseOnExceptionsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setPauseOnExceptions", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setPauseOnExceptions", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetPauseOnExceptions", Err: err}
	}
	return
}

// EvaluateOnCallFrame invokes the Debugger method. Evaluates expression on a given call frame.
func (d *domainClient) EvaluateOnCallFrame(ctx context.Context, args *EvaluateOnCallFrameArgs) (reply *EvaluateOnCallFrameReply, err error) {
	reply = new(EvaluateOnCallFrameReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.evaluateOnCallFrame", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.evaluateOnCallFrame", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "EvaluateOnCallFrame", Err: err}
	}
	return
}

// SetVariableValue invokes the Debugger method. Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
func (d *domainClient) SetVariableValue(ctx context.Context, args *SetVariableValueArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setVariableValue", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setVariableValue", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetVariableValue", Err: err}
	}
	return
}

// SetReturnValue invokes the Debugger method. Changes return value in top frame. Available only at return break position.
func (d *domainClient) SetReturnValue(ctx context.Context, args *SetReturnValueArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setReturnValue", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setReturnValue", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetReturnValue", Err: err}
	}
	return
}

// SetAsyncCallStackDepth invokes the Debugger method. Enables or disables async call stacks tracking.
func (d *domainClient) SetAsyncCallStackDepth(ctx context.Context, args *SetAsyncCallStackDepthArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setAsyncCallStackDepth", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setAsyncCallStackDepth", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetAsyncCallStackDepth", Err: err}
	}
	return
}

// SetBlackboxPatterns invokes the Debugger method. Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
func (d *domainClient) SetBlackboxPatterns(ctx context.Context, args *SetBlackboxPatternsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setBlackboxPatterns", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setBlackboxPatterns", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetBlackboxPatterns", Err: err}
	}
	return
}

// SetBlackboxedRanges invokes the Debugger method. Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
func (d *domainClient) SetBlackboxedRanges(ctx context.Context, args *SetBlackboxedRangesArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Debugger.setBlackboxedRanges", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Debugger.setBlackboxedRanges", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Debugger", Op: "SetBlackboxedRanges", Err: err}
	}
	return
}

func (d *domainClient) ScriptParsed(ctx context.Context) (ScriptParsedClient, error) {
	s, err := rpcc.NewStream(ctx, "Debugger.scriptParsed", d.conn)
	if err != nil {
		return nil, err
	}
	return &scriptParsedClient{Stream: s}, nil
}

type scriptParsedClient struct{ rpcc.Stream }

func (c *scriptParsedClient) Recv() (*ScriptParsedReply, error) {
	event := new(ScriptParsedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Debugger", Op: "ScriptParsed Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ScriptFailedToParse(ctx context.Context) (ScriptFailedToParseClient, error) {
	s, err := rpcc.NewStream(ctx, "Debugger.scriptFailedToParse", d.conn)
	if err != nil {
		return nil, err
	}
	return &scriptFailedToParseClient{Stream: s}, nil
}

type scriptFailedToParseClient struct{ rpcc.Stream }

func (c *scriptFailedToParseClient) Recv() (*ScriptFailedToParseReply, error) {
	event := new(ScriptFailedToParseReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Debugger", Op: "ScriptFailedToParse Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) BreakpointResolved(ctx context.Context) (BreakpointResolvedClient, error) {
	s, err := rpcc.NewStream(ctx, "Debugger.breakpointResolved", d.conn)
	if err != nil {
		return nil, err
	}
	return &breakpointResolvedClient{Stream: s}, nil
}

type breakpointResolvedClient struct{ rpcc.Stream }

func (c *breakpointResolvedClient) Recv() (*BreakpointResolvedReply, error) {
	event := new(BreakpointResolvedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Debugger", Op: "BreakpointResolved Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) Paused(ctx context.Context) (PausedClient, error) {
	s, err := rpcc.NewStream(ctx, "Debugger.paused", d.conn)
	if err != nil {
		return nil, err
	}
	return &pausedClient{Stream: s}, nil
}

type pausedClient struct{ rpcc.Stream }

func (c *pausedClient) Recv() (*PausedReply, error) {
	event := new(PausedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Debugger", Op: "Paused Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) Resumed(ctx context.Context) (ResumedClient, error) {
	s, err := rpcc.NewStream(ctx, "Debugger.resumed", d.conn)
	if err != nil {
		return nil, err
	}
	return &resumedClient{Stream: s}, nil
}

type resumedClient struct{ rpcc.Stream }

func (c *resumedClient) Recv() (*ResumedReply, error) {
	event := new(ResumedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Debugger", Op: "Resumed Recv", Err: err}
	}
	return event, nil
}
