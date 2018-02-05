// Package page provides the Chrome Debugging Protocol
// commands, types, and events for the Page domain.
//
// Actions and events related to the inspected page belong to the page
// domain.
//
// Generated by the chromedp-gen command.
package page

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"context"
	"encoding/base64"

	"github.com/jbegley1995/cdproto/cdp"
	"github.com/jbegley1995/cdproto/debugger"
	"github.com/jbegley1995/cdproto/dom"
	"github.com/jbegley1995/cdproto/runtime"
)

// AddScriptToEvaluateOnNewDocumentParams evaluates given script in every
// frame upon creation (before loading frame's scripts).
type AddScriptToEvaluateOnNewDocumentParams struct {
	Source string `json:"source"`
}

// AddScriptToEvaluateOnNewDocument evaluates given script in every frame
// upon creation (before loading frame's scripts).
//
// parameters:
//   source
func AddScriptToEvaluateOnNewDocument(source string) *AddScriptToEvaluateOnNewDocumentParams {
	return &AddScriptToEvaluateOnNewDocumentParams{
		Source: source,
	}
}

// AddScriptToEvaluateOnNewDocumentReturns return values.
type AddScriptToEvaluateOnNewDocumentReturns struct {
	Identifier ScriptIdentifier `json:"identifier,omitempty"` // Identifier of the added script.
}

// Do executes Page.addScriptToEvaluateOnNewDocument against the provided context.
//
// returns:
//   identifier - Identifier of the added script.
func (p *AddScriptToEvaluateOnNewDocumentParams) Do(ctxt context.Context, h cdp.Executor) (identifier ScriptIdentifier, err error) {
	// execute
	var res AddScriptToEvaluateOnNewDocumentReturns
	err = h.Execute(ctxt, CommandAddScriptToEvaluateOnNewDocument, p, &res)
	if err != nil {
		return "", err
	}

	return res.Identifier, nil
}

// BringToFrontParams brings page to front (activates tab).
type BringToFrontParams struct{}

// BringToFront brings page to front (activates tab).
func BringToFront() *BringToFrontParams {
	return &BringToFrontParams{}
}

// Do executes Page.bringToFront against the provided context.
func (p *BringToFrontParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandBringToFront, nil, nil)
}

// CaptureScreenshotParams capture page screenshot.
type CaptureScreenshotParams struct {
	Format      CaptureScreenshotFormat `json:"format,omitempty"`      // Image compression format (defaults to png).
	Quality     int64                   `json:"quality,omitempty"`     // Compression quality from range [0..100] (jpeg only).
	Clip        *Viewport               `json:"clip,omitempty"`        // Capture the screenshot of a given region only.
	FromSurface bool                    `json:"fromSurface,omitempty"` // Capture the screenshot from the surface, rather than the view. Defaults to true.
}

// CaptureScreenshot capture page screenshot.
//
// parameters:
func CaptureScreenshot() *CaptureScreenshotParams {
	return &CaptureScreenshotParams{}
}

// WithFormat image compression format (defaults to png).
func (p CaptureScreenshotParams) WithFormat(format CaptureScreenshotFormat) *CaptureScreenshotParams {
	p.Format = format
	return &p
}

// WithQuality compression quality from range [0..100] (jpeg only).
func (p CaptureScreenshotParams) WithQuality(quality int64) *CaptureScreenshotParams {
	p.Quality = quality
	return &p
}

// WithClip capture the screenshot of a given region only.
func (p CaptureScreenshotParams) WithClip(clip *Viewport) *CaptureScreenshotParams {
	p.Clip = clip
	return &p
}

// WithFromSurface capture the screenshot from the surface, rather than the
// view. Defaults to true.
func (p CaptureScreenshotParams) WithFromSurface(fromSurface bool) *CaptureScreenshotParams {
	p.FromSurface = fromSurface
	return &p
}

// CaptureScreenshotReturns return values.
type CaptureScreenshotReturns struct {
	Data string `json:"data,omitempty"` // Base64-encoded image data.
}

// Do executes Page.captureScreenshot against the provided context.
//
// returns:
//   data - Base64-encoded image data.
func (p *CaptureScreenshotParams) Do(ctxt context.Context, h cdp.Executor) (data []byte, err error) {
	// execute
	var res CaptureScreenshotReturns
	err = h.Execute(ctxt, CommandCaptureScreenshot, p, &res)
	if err != nil {
		return nil, err
	}

	// decode
	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(res.Data)
	if err != nil {
		return nil, err
	}
	return dec, nil
}

// CreateIsolatedWorldParams creates an isolated world for the given frame.
type CreateIsolatedWorldParams struct {
	FrameID             cdp.FrameID `json:"frameId"`                       // Id of the frame in which the isolated world should be created.
	WorldName           string      `json:"worldName,omitempty"`           // An optional name which is reported in the Execution Context.
	GrantUniveralAccess bool        `json:"grantUniveralAccess,omitempty"` // Whether or not universal access should be granted to the isolated world. This is a powerful option, use with caution.
}

// CreateIsolatedWorld creates an isolated world for the given frame.
//
// parameters:
//   frameID - Id of the frame in which the isolated world should be created.
func CreateIsolatedWorld(frameID cdp.FrameID) *CreateIsolatedWorldParams {
	return &CreateIsolatedWorldParams{
		FrameID: frameID,
	}
}

// WithWorldName an optional name which is reported in the Execution Context.
func (p CreateIsolatedWorldParams) WithWorldName(worldName string) *CreateIsolatedWorldParams {
	p.WorldName = worldName
	return &p
}

// WithGrantUniveralAccess whether or not universal access should be granted
// to the isolated world. This is a powerful option, use with caution.
func (p CreateIsolatedWorldParams) WithGrantUniveralAccess(grantUniveralAccess bool) *CreateIsolatedWorldParams {
	p.GrantUniveralAccess = grantUniveralAccess
	return &p
}

// CreateIsolatedWorldReturns return values.
type CreateIsolatedWorldReturns struct {
	ExecutionContextID runtime.ExecutionContextID `json:"executionContextId,omitempty"` // Execution context of the isolated world.
}

// Do executes Page.createIsolatedWorld against the provided context.
//
// returns:
//   executionContextID - Execution context of the isolated world.
func (p *CreateIsolatedWorldParams) Do(ctxt context.Context, h cdp.Executor) (executionContextID runtime.ExecutionContextID, err error) {
	// execute
	var res CreateIsolatedWorldReturns
	err = h.Execute(ctxt, CommandCreateIsolatedWorld, p, &res)
	if err != nil {
		return 0, err
	}

	return res.ExecutionContextID, nil
}

// DisableParams disables page domain notifications.
type DisableParams struct{}

// Disable disables page domain notifications.
func Disable() *DisableParams {
	return &DisableParams{}
}

// Do executes Page.disable against the provided context.
func (p *DisableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandDisable, nil, nil)
}

// EnableParams enables page domain notifications.
type EnableParams struct{}

// Enable enables page domain notifications.
func Enable() *EnableParams {
	return &EnableParams{}
}

// Do executes Page.enable against the provided context.
func (p *EnableParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandEnable, nil, nil)
}

// GetAppManifestParams [no description].
type GetAppManifestParams struct{}

// GetAppManifest [no description].
func GetAppManifest() *GetAppManifestParams {
	return &GetAppManifestParams{}
}

// GetAppManifestReturns return values.
type GetAppManifestReturns struct {
	URL    string              `json:"url,omitempty"` // Manifest location.
	Errors []*AppManifestError `json:"errors,omitempty"`
	Data   string              `json:"data,omitempty"` // Manifest content.
}

// Do executes Page.getAppManifest against the provided context.
//
// returns:
//   url - Manifest location.
//   errors
//   data - Manifest content.
func (p *GetAppManifestParams) Do(ctxt context.Context, h cdp.Executor) (url string, errors []*AppManifestError, data string, err error) {
	// execute
	var res GetAppManifestReturns
	err = h.Execute(ctxt, CommandGetAppManifest, nil, &res)
	if err != nil {
		return "", nil, "", err
	}

	return res.URL, res.Errors, res.Data, nil
}

// GetFrameTreeParams returns present frame tree structure.
type GetFrameTreeParams struct{}

// GetFrameTree returns present frame tree structure.
func GetFrameTree() *GetFrameTreeParams {
	return &GetFrameTreeParams{}
}

// GetFrameTreeReturns return values.
type GetFrameTreeReturns struct {
	FrameTree *FrameTree `json:"frameTree,omitempty"` // Present frame tree structure.
}

// Do executes Page.getFrameTree against the provided context.
//
// returns:
//   frameTree - Present frame tree structure.
func (p *GetFrameTreeParams) Do(ctxt context.Context, h cdp.Executor) (frameTree *FrameTree, err error) {
	// execute
	var res GetFrameTreeReturns
	err = h.Execute(ctxt, CommandGetFrameTree, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.FrameTree, nil
}

// GetLayoutMetricsParams returns metrics relating to the layouting of the
// page, such as viewport bounds/scale.
type GetLayoutMetricsParams struct{}

// GetLayoutMetrics returns metrics relating to the layouting of the page,
// such as viewport bounds/scale.
func GetLayoutMetrics() *GetLayoutMetricsParams {
	return &GetLayoutMetricsParams{}
}

// GetLayoutMetricsReturns return values.
type GetLayoutMetricsReturns struct {
	LayoutViewport *LayoutViewport `json:"layoutViewport,omitempty"` // Metrics relating to the layout viewport.
	VisualViewport *VisualViewport `json:"visualViewport,omitempty"` // Metrics relating to the visual viewport.
	ContentSize    *dom.Rect       `json:"contentSize,omitempty"`    // Size of scrollable area.
}

// Do executes Page.getLayoutMetrics against the provided context.
//
// returns:
//   layoutViewport - Metrics relating to the layout viewport.
//   visualViewport - Metrics relating to the visual viewport.
//   contentSize - Size of scrollable area.
func (p *GetLayoutMetricsParams) Do(ctxt context.Context, h cdp.Executor) (layoutViewport *LayoutViewport, visualViewport *VisualViewport, contentSize *dom.Rect, err error) {
	// execute
	var res GetLayoutMetricsReturns
	err = h.Execute(ctxt, CommandGetLayoutMetrics, nil, &res)
	if err != nil {
		return nil, nil, nil, err
	}

	return res.LayoutViewport, res.VisualViewport, res.ContentSize, nil
}

// GetNavigationHistoryParams returns navigation history for the current
// page.
type GetNavigationHistoryParams struct{}

// GetNavigationHistory returns navigation history for the current page.
func GetNavigationHistory() *GetNavigationHistoryParams {
	return &GetNavigationHistoryParams{}
}

// GetNavigationHistoryReturns return values.
type GetNavigationHistoryReturns struct {
	CurrentIndex int64              `json:"currentIndex,omitempty"` // Index of the current navigation history entry.
	Entries      []*NavigationEntry `json:"entries,omitempty"`      // Array of navigation history entries.
}

// Do executes Page.getNavigationHistory against the provided context.
//
// returns:
//   currentIndex - Index of the current navigation history entry.
//   entries - Array of navigation history entries.
func (p *GetNavigationHistoryParams) Do(ctxt context.Context, h cdp.Executor) (currentIndex int64, entries []*NavigationEntry, err error) {
	// execute
	var res GetNavigationHistoryReturns
	err = h.Execute(ctxt, CommandGetNavigationHistory, nil, &res)
	if err != nil {
		return 0, nil, err
	}

	return res.CurrentIndex, res.Entries, nil
}

// GetResourceContentParams returns content of the given resource.
type GetResourceContentParams struct {
	FrameID cdp.FrameID `json:"frameId"` // Frame id to get resource for.
	URL     string      `json:"url"`     // URL of the resource to get content for.
}

// GetResourceContent returns content of the given resource.
//
// parameters:
//   frameID - Frame id to get resource for.
//   url - URL of the resource to get content for.
func GetResourceContent(frameID cdp.FrameID, url string) *GetResourceContentParams {
	return &GetResourceContentParams{
		FrameID: frameID,
		URL:     url,
	}
}

// GetResourceContentReturns return values.
type GetResourceContentReturns struct {
	Content       string `json:"content,omitempty"`       // Resource content.
	Base64encoded bool   `json:"base64Encoded,omitempty"` // True, if content was served as base64.
}

// Do executes Page.getResourceContent against the provided context.
//
// returns:
//   content - Resource content.
func (p *GetResourceContentParams) Do(ctxt context.Context, h cdp.Executor) (content []byte, err error) {
	// execute
	var res GetResourceContentReturns
	err = h.Execute(ctxt, CommandGetResourceContent, p, &res)
	if err != nil {
		return nil, err
	}

	// decode
	var dec []byte
	if res.Base64encoded {
		dec, err = base64.StdEncoding.DecodeString(res.Content)
		if err != nil {
			return nil, err
		}
	} else {
		dec = []byte(res.Content)
	}
	return dec, nil
}

// GetResourceTreeParams returns present frame / resource tree structure.
type GetResourceTreeParams struct{}

// GetResourceTree returns present frame / resource tree structure.
func GetResourceTree() *GetResourceTreeParams {
	return &GetResourceTreeParams{}
}

// GetResourceTreeReturns return values.
type GetResourceTreeReturns struct {
	FrameTree *FrameResourceTree `json:"frameTree,omitempty"` // Present frame / resource tree structure.
}

// Do executes Page.getResourceTree against the provided context.
//
// returns:
//   frameTree - Present frame / resource tree structure.
func (p *GetResourceTreeParams) Do(ctxt context.Context, h cdp.Executor) (frameTree *FrameResourceTree, err error) {
	// execute
	var res GetResourceTreeReturns
	err = h.Execute(ctxt, CommandGetResourceTree, nil, &res)
	if err != nil {
		return nil, err
	}

	return res.FrameTree, nil
}

// HandleJavaScriptDialogParams accepts or dismisses a JavaScript initiated
// dialog (alert, confirm, prompt, or onbeforeunload).
type HandleJavaScriptDialogParams struct {
	Accept     bool   `json:"accept"`               // Whether to accept or dismiss the dialog.
	PromptText string `json:"promptText,omitempty"` // The text to enter into the dialog prompt before accepting. Used only if this is a prompt dialog.
}

// HandleJavaScriptDialog accepts or dismisses a JavaScript initiated dialog
// (alert, confirm, prompt, or onbeforeunload).
//
// parameters:
//   accept - Whether to accept or dismiss the dialog.
func HandleJavaScriptDialog(accept bool) *HandleJavaScriptDialogParams {
	return &HandleJavaScriptDialogParams{
		Accept: accept,
	}
}

// WithPromptText the text to enter into the dialog prompt before accepting.
// Used only if this is a prompt dialog.
func (p HandleJavaScriptDialogParams) WithPromptText(promptText string) *HandleJavaScriptDialogParams {
	p.PromptText = promptText
	return &p
}

// Do executes Page.handleJavaScriptDialog against the provided context.
func (p *HandleJavaScriptDialogParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandHandleJavaScriptDialog, p, nil)
}

// NavigateParams navigates current page to the given URL.
type NavigateParams struct {
	URL            string         `json:"url"`                      // URL to navigate the page to.
	Referrer       string         `json:"referrer,omitempty"`       // Referrer URL.
	TransitionType TransitionType `json:"transitionType,omitempty"` // Intended transition type.
	FrameID        cdp.FrameID    `json:"frameId,omitempty"`        // Frame id to navigate, if not specified navigates the top frame.
}

// Navigate navigates current page to the given URL.
//
// parameters:
//   url - URL to navigate the page to.
func Navigate(url string) *NavigateParams {
	return &NavigateParams{
		URL: url,
	}
}

// WithReferrer referrer URL.
func (p NavigateParams) WithReferrer(referrer string) *NavigateParams {
	p.Referrer = referrer
	return &p
}

// WithTransitionType intended transition type.
func (p NavigateParams) WithTransitionType(transitionType TransitionType) *NavigateParams {
	p.TransitionType = transitionType
	return &p
}

// WithFrameID frame id to navigate, if not specified navigates the top
// frame.
func (p NavigateParams) WithFrameID(frameID cdp.FrameID) *NavigateParams {
	p.FrameID = frameID
	return &p
}

// NavigateReturns return values.
type NavigateReturns struct {
	FrameID   cdp.FrameID  `json:"frameId,omitempty"`   // Frame id that has navigated (or failed to navigate)
	LoaderID  cdp.LoaderID `json:"loaderId,omitempty"`  // Loader identifier.
	ErrorText string       `json:"errorText,omitempty"` // User friendly error message, present if and only if navigation has failed.
}

// Do executes Page.navigate against the provided context.
//
// returns:
//   frameID - Frame id that has navigated (or failed to navigate)
//   loaderID - Loader identifier.
//   errorText - User friendly error message, present if and only if navigation has failed.
func (p *NavigateParams) Do(ctxt context.Context, h cdp.Executor) (frameID cdp.FrameID, loaderID cdp.LoaderID, errorText string, err error) {
	// execute
	var res NavigateReturns
	err = h.Execute(ctxt, CommandNavigate, p, &res)
	if err != nil {
		return "", "", "", err
	}

	return res.FrameID, res.LoaderID, res.ErrorText, nil
}

// NavigateToHistoryEntryParams navigates current page to the given history
// entry.
type NavigateToHistoryEntryParams struct {
	EntryID int64 `json:"entryId"` // Unique id of the entry to navigate to.
}

// NavigateToHistoryEntry navigates current page to the given history entry.
//
// parameters:
//   entryID - Unique id of the entry to navigate to.
func NavigateToHistoryEntry(entryID int64) *NavigateToHistoryEntryParams {
	return &NavigateToHistoryEntryParams{
		EntryID: entryID,
	}
}

// Do executes Page.navigateToHistoryEntry against the provided context.
func (p *NavigateToHistoryEntryParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandNavigateToHistoryEntry, p, nil)
}

// PrintToPDFParams print page as PDF.
type PrintToPDFParams struct {
	Landscape               bool    `json:"landscape,omitempty"`               // Paper orientation. Defaults to false.
	DisplayHeaderFooter     bool    `json:"displayHeaderFooter,omitempty"`     // Display header and footer. Defaults to false.
	PrintBackground         bool    `json:"printBackground,omitempty"`         // Print background graphics. Defaults to false.
	Scale                   float64 `json:"scale,omitempty"`                   // Scale of the webpage rendering. Defaults to 1.
	PaperWidth              float64 `json:"paperWidth,omitempty"`              // Paper width in inches. Defaults to 8.5 inches.
	PaperHeight             float64 `json:"paperHeight,omitempty"`             // Paper height in inches. Defaults to 11 inches.
	MarginTop               float64 `json:"marginTop,omitempty"`               // Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom            float64 `json:"marginBottom,omitempty"`            // Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft              float64 `json:"marginLeft,omitempty"`              // Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight             float64 `json:"marginRight,omitempty"`             // Right margin in inches. Defaults to 1cm (~0.4 inches).
	PageRanges              string  `json:"pageRanges,omitempty"`              // Paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to the empty string, which means print all pages.
	IgnoreInvalidPageRanges bool    `json:"ignoreInvalidPageRanges,omitempty"` // Whether to silently ignore invalid but successfully parsed page ranges, such as '3-2'. Defaults to false.
	HeaderTemplate          string  `json:"headerTemplate,omitempty"`          // HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date - formatted print date - title - document title - url - document location - pageNumber - current page number - totalPages - total pages in the document  For example, <span class=title></span> would generate span containing the title.
	FooterTemplate          string  `json:"footerTemplate,omitempty"`          // HTML template for the print footer. Should use the same format as the headerTemplate.
	PreferCSSPageSize       bool    `json:"preferCSSPageSize,omitempty"`       // Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
}

// PrintToPDF print page as PDF.
//
// parameters:
func PrintToPDF() *PrintToPDFParams {
	return &PrintToPDFParams{}
}

// WithLandscape paper orientation. Defaults to false.
func (p PrintToPDFParams) WithLandscape(landscape bool) *PrintToPDFParams {
	p.Landscape = landscape
	return &p
}

// WithDisplayHeaderFooter display header and footer. Defaults to false.
func (p PrintToPDFParams) WithDisplayHeaderFooter(displayHeaderFooter bool) *PrintToPDFParams {
	p.DisplayHeaderFooter = displayHeaderFooter
	return &p
}

// WithPrintBackground print background graphics. Defaults to false.
func (p PrintToPDFParams) WithPrintBackground(printBackground bool) *PrintToPDFParams {
	p.PrintBackground = printBackground
	return &p
}

// WithScale scale of the webpage rendering. Defaults to 1.
func (p PrintToPDFParams) WithScale(scale float64) *PrintToPDFParams {
	p.Scale = scale
	return &p
}

// WithPaperWidth paper width in inches. Defaults to 8.5 inches.
func (p PrintToPDFParams) WithPaperWidth(paperWidth float64) *PrintToPDFParams {
	p.PaperWidth = paperWidth
	return &p
}

// WithPaperHeight paper height in inches. Defaults to 11 inches.
func (p PrintToPDFParams) WithPaperHeight(paperHeight float64) *PrintToPDFParams {
	p.PaperHeight = paperHeight
	return &p
}

// WithMarginTop top margin in inches. Defaults to 1cm (~0.4 inches).
func (p PrintToPDFParams) WithMarginTop(marginTop float64) *PrintToPDFParams {
	p.MarginTop = marginTop
	return &p
}

// WithMarginBottom bottom margin in inches. Defaults to 1cm (~0.4 inches).
func (p PrintToPDFParams) WithMarginBottom(marginBottom float64) *PrintToPDFParams {
	p.MarginBottom = marginBottom
	return &p
}

// WithMarginLeft left margin in inches. Defaults to 1cm (~0.4 inches).
func (p PrintToPDFParams) WithMarginLeft(marginLeft float64) *PrintToPDFParams {
	p.MarginLeft = marginLeft
	return &p
}

// WithMarginRight right margin in inches. Defaults to 1cm (~0.4 inches).
func (p PrintToPDFParams) WithMarginRight(marginRight float64) *PrintToPDFParams {
	p.MarginRight = marginRight
	return &p
}

// WithPageRanges paper ranges to print, e.g., '1-5, 8, 11-13'. Defaults to
// the empty string, which means print all pages.
func (p PrintToPDFParams) WithPageRanges(pageRanges string) *PrintToPDFParams {
	p.PageRanges = pageRanges
	return &p
}

// WithIgnoreInvalidPageRanges whether to silently ignore invalid but
// successfully parsed page ranges, such as '3-2'. Defaults to false.
func (p PrintToPDFParams) WithIgnoreInvalidPageRanges(ignoreInvalidPageRanges bool) *PrintToPDFParams {
	p.IgnoreInvalidPageRanges = ignoreInvalidPageRanges
	return &p
}

// WithHeaderTemplate HTML template for the print header. Should be valid
// HTML markup with following classes used to inject printing values into them:
// - date - formatted print date - title - document title - url - document
// location - pageNumber - current page number - totalPages - total pages in the
// document For example, <span class=title></span> would generate span
// containing the title.
func (p PrintToPDFParams) WithHeaderTemplate(headerTemplate string) *PrintToPDFParams {
	p.HeaderTemplate = headerTemplate
	return &p
}

// WithFooterTemplate HTML template for the print footer. Should use the same
// format as the headerTemplate.
func (p PrintToPDFParams) WithFooterTemplate(footerTemplate string) *PrintToPDFParams {
	p.FooterTemplate = footerTemplate
	return &p
}

// WithPreferCSSPageSize whether or not to prefer page size as defined by
// css. Defaults to false, in which case the content will be scaled to fit the
// paper size.
func (p PrintToPDFParams) WithPreferCSSPageSize(preferCSSPageSize bool) *PrintToPDFParams {
	p.PreferCSSPageSize = preferCSSPageSize
	return &p
}

// PrintToPDFReturns return values.
type PrintToPDFReturns struct {
	Data string `json:"data,omitempty"` // Base64-encoded pdf data.
}

// Do executes Page.printToPDF against the provided context.
//
// returns:
//   data - Base64-encoded pdf data.
func (p *PrintToPDFParams) Do(ctxt context.Context, h cdp.Executor) (data []byte, err error) {
	// execute
	var res PrintToPDFReturns
	err = h.Execute(ctxt, CommandPrintToPDF, p, &res)
	if err != nil {
		return nil, err
	}

	// decode
	var dec []byte
	dec, err = base64.StdEncoding.DecodeString(res.Data)
	if err != nil {
		return nil, err
	}
	return dec, nil
}

// ReloadParams reloads given page optionally ignoring the cache.
type ReloadParams struct {
	IgnoreCache            bool   `json:"ignoreCache,omitempty"`            // If true, browser cache is ignored (as if the user pressed Shift+refresh).
	ScriptToEvaluateOnLoad string `json:"scriptToEvaluateOnLoad,omitempty"` // If set, the script will be injected into all frames of the inspected page after reload. Argument will be ignored if reloading dataURL origin.
}

// Reload reloads given page optionally ignoring the cache.
//
// parameters:
func Reload() *ReloadParams {
	return &ReloadParams{}
}

// WithIgnoreCache if true, browser cache is ignored (as if the user pressed
// Shift+refresh).
func (p ReloadParams) WithIgnoreCache(ignoreCache bool) *ReloadParams {
	p.IgnoreCache = ignoreCache
	return &p
}

// WithScriptToEvaluateOnLoad if set, the script will be injected into all
// frames of the inspected page after reload. Argument will be ignored if
// reloading dataURL origin.
func (p ReloadParams) WithScriptToEvaluateOnLoad(scriptToEvaluateOnLoad string) *ReloadParams {
	p.ScriptToEvaluateOnLoad = scriptToEvaluateOnLoad
	return &p
}

// Do executes Page.reload against the provided context.
func (p *ReloadParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandReload, p, nil)
}

// RemoveScriptToEvaluateOnNewDocumentParams removes given script from the
// list.
type RemoveScriptToEvaluateOnNewDocumentParams struct {
	Identifier ScriptIdentifier `json:"identifier"`
}

// RemoveScriptToEvaluateOnNewDocument removes given script from the list.
//
// parameters:
//   identifier
func RemoveScriptToEvaluateOnNewDocument(identifier ScriptIdentifier) *RemoveScriptToEvaluateOnNewDocumentParams {
	return &RemoveScriptToEvaluateOnNewDocumentParams{
		Identifier: identifier,
	}
}

// Do executes Page.removeScriptToEvaluateOnNewDocument against the provided context.
func (p *RemoveScriptToEvaluateOnNewDocumentParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandRemoveScriptToEvaluateOnNewDocument, p, nil)
}

// RequestAppBannerParams [no description].
type RequestAppBannerParams struct{}

// RequestAppBanner [no description].
func RequestAppBanner() *RequestAppBannerParams {
	return &RequestAppBannerParams{}
}

// Do executes Page.requestAppBanner against the provided context.
func (p *RequestAppBannerParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandRequestAppBanner, nil, nil)
}

// ScreencastFrameAckParams acknowledges that a screencast frame has been
// received by the frontend.
type ScreencastFrameAckParams struct {
	SessionID int64 `json:"sessionId"` // Frame number.
}

// ScreencastFrameAck acknowledges that a screencast frame has been received
// by the frontend.
//
// parameters:
//   sessionID - Frame number.
func ScreencastFrameAck(sessionID int64) *ScreencastFrameAckParams {
	return &ScreencastFrameAckParams{
		SessionID: sessionID,
	}
}

// Do executes Page.screencastFrameAck against the provided context.
func (p *ScreencastFrameAckParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandScreencastFrameAck, p, nil)
}

// SearchInResourceParams searches for given string in resource content.
type SearchInResourceParams struct {
	FrameID       cdp.FrameID `json:"frameId"`                 // Frame id for resource to search in.
	URL           string      `json:"url"`                     // URL of the resource to search in.
	Query         string      `json:"query"`                   // String to search for.
	CaseSensitive bool        `json:"caseSensitive,omitempty"` // If true, search is case sensitive.
	IsRegex       bool        `json:"isRegex,omitempty"`       // If true, treats string parameter as regex.
}

// SearchInResource searches for given string in resource content.
//
// parameters:
//   frameID - Frame id for resource to search in.
//   url - URL of the resource to search in.
//   query - String to search for.
func SearchInResource(frameID cdp.FrameID, url string, query string) *SearchInResourceParams {
	return &SearchInResourceParams{
		FrameID: frameID,
		URL:     url,
		Query:   query,
	}
}

// WithCaseSensitive if true, search is case sensitive.
func (p SearchInResourceParams) WithCaseSensitive(caseSensitive bool) *SearchInResourceParams {
	p.CaseSensitive = caseSensitive
	return &p
}

// WithIsRegex if true, treats string parameter as regex.
func (p SearchInResourceParams) WithIsRegex(isRegex bool) *SearchInResourceParams {
	p.IsRegex = isRegex
	return &p
}

// SearchInResourceReturns return values.
type SearchInResourceReturns struct {
	Result []*debugger.SearchMatch `json:"result,omitempty"` // List of search matches.
}

// Do executes Page.searchInResource against the provided context.
//
// returns:
//   result - List of search matches.
func (p *SearchInResourceParams) Do(ctxt context.Context, h cdp.Executor) (result []*debugger.SearchMatch, err error) {
	// execute
	var res SearchInResourceReturns
	err = h.Execute(ctxt, CommandSearchInResource, p, &res)
	if err != nil {
		return nil, err
	}

	return res.Result, nil
}

// SetAdBlockingEnabledParams enable Chrome's experimental ad filter on all
// sites.
type SetAdBlockingEnabledParams struct {
	Enabled bool `json:"enabled"` // Whether to block ads.
}

// SetAdBlockingEnabled enable Chrome's experimental ad filter on all sites.
//
// parameters:
//   enabled - Whether to block ads.
func SetAdBlockingEnabled(enabled bool) *SetAdBlockingEnabledParams {
	return &SetAdBlockingEnabledParams{
		Enabled: enabled,
	}
}

// Do executes Page.setAdBlockingEnabled against the provided context.
func (p *SetAdBlockingEnabledParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetAdBlockingEnabled, p, nil)
}

// SetDocumentContentParams sets given markup as the document's HTML.
type SetDocumentContentParams struct {
	FrameID cdp.FrameID `json:"frameId"` // Frame id to set HTML for.
	HTML    string      `json:"html"`    // HTML content to set.
}

// SetDocumentContent sets given markup as the document's HTML.
//
// parameters:
//   frameID - Frame id to set HTML for.
//   html - HTML content to set.
func SetDocumentContent(frameID cdp.FrameID, html string) *SetDocumentContentParams {
	return &SetDocumentContentParams{
		FrameID: frameID,
		HTML:    html,
	}
}

// Do executes Page.setDocumentContent against the provided context.
func (p *SetDocumentContentParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetDocumentContent, p, nil)
}

// SetDownloadBehaviorParams set the behavior when downloading a file.
type SetDownloadBehaviorParams struct {
	Behavior     SetDownloadBehaviorBehavior `json:"behavior"`               // Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny).
	DownloadPath string                      `json:"downloadPath,omitempty"` // The default path to save downloaded files to. This is required if behavior is set to 'allow'
}

// SetDownloadBehavior set the behavior when downloading a file.
//
// parameters:
//   behavior - Whether to allow all or deny all download requests, or use default Chrome behavior if available (otherwise deny).
func SetDownloadBehavior(behavior SetDownloadBehaviorBehavior) *SetDownloadBehaviorParams {
	return &SetDownloadBehaviorParams{
		Behavior: behavior,
	}
}

// WithDownloadPath the default path to save downloaded files to. This is
// required if behavior is set to 'allow'.
func (p SetDownloadBehaviorParams) WithDownloadPath(downloadPath string) *SetDownloadBehaviorParams {
	p.DownloadPath = downloadPath
	return &p
}

// Do executes Page.setDownloadBehavior against the provided context.
func (p *SetDownloadBehaviorParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetDownloadBehavior, p, nil)
}

// SetLifecycleEventsEnabledParams controls whether page will emit lifecycle
// events.
type SetLifecycleEventsEnabledParams struct {
	Enabled bool `json:"enabled"` // If true, starts emitting lifecycle events.
}

// SetLifecycleEventsEnabled controls whether page will emit lifecycle
// events.
//
// parameters:
//   enabled - If true, starts emitting lifecycle events.
func SetLifecycleEventsEnabled(enabled bool) *SetLifecycleEventsEnabledParams {
	return &SetLifecycleEventsEnabledParams{
		Enabled: enabled,
	}
}

// Do executes Page.setLifecycleEventsEnabled against the provided context.
func (p *SetLifecycleEventsEnabledParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandSetLifecycleEventsEnabled, p, nil)
}

// StartScreencastParams starts sending each frame using the screencastFrame
// event.
type StartScreencastParams struct {
	Format        ScreencastFormat `json:"format,omitempty"`        // Image compression format.
	Quality       int64            `json:"quality,omitempty"`       // Compression quality from range [0..100].
	MaxWidth      int64            `json:"maxWidth,omitempty"`      // Maximum screenshot width.
	MaxHeight     int64            `json:"maxHeight,omitempty"`     // Maximum screenshot height.
	EveryNthFrame int64            `json:"everyNthFrame,omitempty"` // Send every n-th frame.
}

// StartScreencast starts sending each frame using the screencastFrame event.
//
// parameters:
func StartScreencast() *StartScreencastParams {
	return &StartScreencastParams{}
}

// WithFormat image compression format.
func (p StartScreencastParams) WithFormat(format ScreencastFormat) *StartScreencastParams {
	p.Format = format
	return &p
}

// WithQuality compression quality from range [0..100].
func (p StartScreencastParams) WithQuality(quality int64) *StartScreencastParams {
	p.Quality = quality
	return &p
}

// WithMaxWidth maximum screenshot width.
func (p StartScreencastParams) WithMaxWidth(maxWidth int64) *StartScreencastParams {
	p.MaxWidth = maxWidth
	return &p
}

// WithMaxHeight maximum screenshot height.
func (p StartScreencastParams) WithMaxHeight(maxHeight int64) *StartScreencastParams {
	p.MaxHeight = maxHeight
	return &p
}

// WithEveryNthFrame send every n-th frame.
func (p StartScreencastParams) WithEveryNthFrame(everyNthFrame int64) *StartScreencastParams {
	p.EveryNthFrame = everyNthFrame
	return &p
}

// Do executes Page.startScreencast against the provided context.
func (p *StartScreencastParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandStartScreencast, p, nil)
}

// StopLoadingParams force the page stop all navigations and pending resource
// fetches.
type StopLoadingParams struct{}

// StopLoading force the page stop all navigations and pending resource
// fetches.
func StopLoading() *StopLoadingParams {
	return &StopLoadingParams{}
}

// Do executes Page.stopLoading against the provided context.
func (p *StopLoadingParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandStopLoading, nil, nil)
}

// CrashParams crashes renderer on the IO thread, generates minidumps.
type CrashParams struct{}

// Crash crashes renderer on the IO thread, generates minidumps.
func Crash() *CrashParams {
	return &CrashParams{}
}

// Do executes Page.crash against the provided context.
func (p *CrashParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandCrash, nil, nil)
}

// StopScreencastParams stops sending each frame in the screencastFrame.
type StopScreencastParams struct{}

// StopScreencast stops sending each frame in the screencastFrame.
func StopScreencast() *StopScreencastParams {
	return &StopScreencastParams{}
}

// Do executes Page.stopScreencast against the provided context.
func (p *StopScreencastParams) Do(ctxt context.Context, h cdp.Executor) (err error) {
	return h.Execute(ctxt, CommandStopScreencast, nil, nil)
}

// Command names.
const (
	CommandAddScriptToEvaluateOnNewDocument    = "Page.addScriptToEvaluateOnNewDocument"
	CommandBringToFront                        = "Page.bringToFront"
	CommandCaptureScreenshot                   = "Page.captureScreenshot"
	CommandCreateIsolatedWorld                 = "Page.createIsolatedWorld"
	CommandDisable                             = "Page.disable"
	CommandEnable                              = "Page.enable"
	CommandGetAppManifest                      = "Page.getAppManifest"
	CommandGetFrameTree                        = "Page.getFrameTree"
	CommandGetLayoutMetrics                    = "Page.getLayoutMetrics"
	CommandGetNavigationHistory                = "Page.getNavigationHistory"
	CommandGetResourceContent                  = "Page.getResourceContent"
	CommandGetResourceTree                     = "Page.getResourceTree"
	CommandHandleJavaScriptDialog              = "Page.handleJavaScriptDialog"
	CommandNavigate                            = "Page.navigate"
	CommandNavigateToHistoryEntry              = "Page.navigateToHistoryEntry"
	CommandPrintToPDF                          = "Page.printToPDF"
	CommandReload                              = "Page.reload"
	CommandRemoveScriptToEvaluateOnNewDocument = "Page.removeScriptToEvaluateOnNewDocument"
	CommandRequestAppBanner                    = "Page.requestAppBanner"
	CommandScreencastFrameAck                  = "Page.screencastFrameAck"
	CommandSearchInResource                    = "Page.searchInResource"
	CommandSetAdBlockingEnabled                = "Page.setAdBlockingEnabled"
	CommandSetDocumentContent                  = "Page.setDocumentContent"
	CommandSetDownloadBehavior                 = "Page.setDownloadBehavior"
	CommandSetLifecycleEventsEnabled           = "Page.setLifecycleEventsEnabled"
	CommandStartScreencast                     = "Page.startScreencast"
	CommandStopLoading                         = "Page.stopLoading"
	CommandCrash                               = "Page.crash"
	CommandStopScreencast                      = "Page.stopScreencast"
)
