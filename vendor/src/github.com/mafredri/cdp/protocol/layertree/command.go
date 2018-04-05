// Code generated by cdpgen. DO NOT EDIT.

package layertree

import (
	"encoding/json"

	"github.com/mafredri/cdp/protocol/dom"
)

// CompositingReasonsArgs represents the arguments for CompositingReasons in the LayerTree domain.
type CompositingReasonsArgs struct {
	LayerID LayerID `json:"layerId"` // The id of the layer for which we want to get the reasons it was composited.
}

// NewCompositingReasonsArgs initializes CompositingReasonsArgs with the required arguments.
func NewCompositingReasonsArgs(layerID LayerID) *CompositingReasonsArgs {
	args := new(CompositingReasonsArgs)
	args.LayerID = layerID
	return args
}

// CompositingReasonsReply represents the return values for CompositingReasons in the LayerTree domain.
type CompositingReasonsReply struct {
	CompositingReasons []string `json:"compositingReasons"` // A list of strings specifying reasons for the given layer to become composited.
}

// LoadSnapshotArgs represents the arguments for LoadSnapshot in the LayerTree domain.
type LoadSnapshotArgs struct {
	Tiles []PictureTile `json:"tiles"` // An array of tiles composing the snapshot.
}

// NewLoadSnapshotArgs initializes LoadSnapshotArgs with the required arguments.
func NewLoadSnapshotArgs(tiles []PictureTile) *LoadSnapshotArgs {
	args := new(LoadSnapshotArgs)
	args.Tiles = tiles
	return args
}

// LoadSnapshotReply represents the return values for LoadSnapshot in the LayerTree domain.
type LoadSnapshotReply struct {
	SnapshotID SnapshotID `json:"snapshotId"` // The id of the snapshot.
}

// MakeSnapshotArgs represents the arguments for MakeSnapshot in the LayerTree domain.
type MakeSnapshotArgs struct {
	LayerID LayerID `json:"layerId"` // The id of the layer.
}

// NewMakeSnapshotArgs initializes MakeSnapshotArgs with the required arguments.
func NewMakeSnapshotArgs(layerID LayerID) *MakeSnapshotArgs {
	args := new(MakeSnapshotArgs)
	args.LayerID = layerID
	return args
}

// MakeSnapshotReply represents the return values for MakeSnapshot in the LayerTree domain.
type MakeSnapshotReply struct {
	SnapshotID SnapshotID `json:"snapshotId"` // The id of the layer snapshot.
}

// ProfileSnapshotArgs represents the arguments for ProfileSnapshot in the LayerTree domain.
type ProfileSnapshotArgs struct {
	SnapshotID     SnapshotID `json:"snapshotId"`               // The id of the layer snapshot.
	MinRepeatCount *int       `json:"minRepeatCount,omitempty"` // The maximum number of times to replay the snapshot (1, if not specified).
	MinDuration    *float64   `json:"minDuration,omitempty"`    // The minimum duration (in seconds) to replay the snapshot.
	ClipRect       *dom.Rect  `json:"clipRect,omitempty"`       // The clip rectangle to apply when replaying the snapshot.
}

// NewProfileSnapshotArgs initializes ProfileSnapshotArgs with the required arguments.
func NewProfileSnapshotArgs(snapshotID SnapshotID) *ProfileSnapshotArgs {
	args := new(ProfileSnapshotArgs)
	args.SnapshotID = snapshotID
	return args
}

// SetMinRepeatCount sets the MinRepeatCount optional argument. The maximum number of times to replay the snapshot (1, if not specified).
func (a *ProfileSnapshotArgs) SetMinRepeatCount(minRepeatCount int) *ProfileSnapshotArgs {
	a.MinRepeatCount = &minRepeatCount
	return a
}

// SetMinDuration sets the MinDuration optional argument. The minimum duration (in seconds) to replay the snapshot.
func (a *ProfileSnapshotArgs) SetMinDuration(minDuration float64) *ProfileSnapshotArgs {
	a.MinDuration = &minDuration
	return a
}

// SetClipRect sets the ClipRect optional argument. The clip rectangle to apply when replaying the snapshot.
func (a *ProfileSnapshotArgs) SetClipRect(clipRect dom.Rect) *ProfileSnapshotArgs {
	a.ClipRect = &clipRect
	return a
}

// ProfileSnapshotReply represents the return values for ProfileSnapshot in the LayerTree domain.
type ProfileSnapshotReply struct {
	Timings []PaintProfile `json:"timings"` // The array of paint profiles, one per run.
}

// ReleaseSnapshotArgs represents the arguments for ReleaseSnapshot in the LayerTree domain.
type ReleaseSnapshotArgs struct {
	SnapshotID SnapshotID `json:"snapshotId"` // The id of the layer snapshot.
}

// NewReleaseSnapshotArgs initializes ReleaseSnapshotArgs with the required arguments.
func NewReleaseSnapshotArgs(snapshotID SnapshotID) *ReleaseSnapshotArgs {
	args := new(ReleaseSnapshotArgs)
	args.SnapshotID = snapshotID
	return args
}

// ReplaySnapshotArgs represents the arguments for ReplaySnapshot in the LayerTree domain.
type ReplaySnapshotArgs struct {
	SnapshotID SnapshotID `json:"snapshotId"`         // The id of the layer snapshot.
	FromStep   *int       `json:"fromStep,omitempty"` // The first step to replay from (replay from the very start if not specified).
	ToStep     *int       `json:"toStep,omitempty"`   // The last step to replay to (replay till the end if not specified).
	Scale      *float64   `json:"scale,omitempty"`    // The scale to apply while replaying (defaults to 1).
}

// NewReplaySnapshotArgs initializes ReplaySnapshotArgs with the required arguments.
func NewReplaySnapshotArgs(snapshotID SnapshotID) *ReplaySnapshotArgs {
	args := new(ReplaySnapshotArgs)
	args.SnapshotID = snapshotID
	return args
}

// SetFromStep sets the FromStep optional argument. The first step to replay from (replay from the very start if not specified).
func (a *ReplaySnapshotArgs) SetFromStep(fromStep int) *ReplaySnapshotArgs {
	a.FromStep = &fromStep
	return a
}

// SetToStep sets the ToStep optional argument. The last step to replay to (replay till the end if not specified).
func (a *ReplaySnapshotArgs) SetToStep(toStep int) *ReplaySnapshotArgs {
	a.ToStep = &toStep
	return a
}

// SetScale sets the Scale optional argument. The scale to apply while replaying (defaults to 1).
func (a *ReplaySnapshotArgs) SetScale(scale float64) *ReplaySnapshotArgs {
	a.Scale = &scale
	return a
}

// ReplaySnapshotReply represents the return values for ReplaySnapshot in the LayerTree domain.
type ReplaySnapshotReply struct {
	DataURL string `json:"dataURL"` // A data: URL for resulting image.
}

// SnapshotCommandLogArgs represents the arguments for SnapshotCommandLog in the LayerTree domain.
type SnapshotCommandLogArgs struct {
	SnapshotID SnapshotID `json:"snapshotId"` // The id of the layer snapshot.
}

// NewSnapshotCommandLogArgs initializes SnapshotCommandLogArgs with the required arguments.
func NewSnapshotCommandLogArgs(snapshotID SnapshotID) *SnapshotCommandLogArgs {
	args := new(SnapshotCommandLogArgs)
	args.SnapshotID = snapshotID
	return args
}

// SnapshotCommandLogReply represents the return values for SnapshotCommandLog in the LayerTree domain.
type SnapshotCommandLogReply struct {
	CommandLog []json.RawMessage `json:"commandLog"` // The array of canvas function calls.
}
