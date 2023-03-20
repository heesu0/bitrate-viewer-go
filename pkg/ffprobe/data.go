package ffprobe

import (
	"math"
	"strconv"
	"strings"
)

type ProbeData struct {
	Frames  []*Frame  `json:"frames"`
	Streams []*Stream `json:"streams"`
	Format  *Format   `json:"format"`
}

type Format struct {
	Filename       string      `json:"filename"`
	NBStreams      int         `json:"nb_streams"`
	NBPrograms     int         `json:"nb_programs"`
	FormatName     string      `json:"format_name"`
	FormatLongName string      `json:"format_long_name"`
	StartTime      float64     `json:"start_time,string"`
	Duration       float64     `json:"duration,string"`
	Size           string      `json:"size"`
	BitRate        string      `json:"bit_rate"`
	ProbeScore     int         `json:"probe_score"`
	Tags           *FormatTags `json:"tags"`
}

type FormatTags struct {
	MajorBrand       string `json:"major_brand"`
	MinorVersion     string `json:"minor_version"`
	CompatibleBrands string `json:"compatible_brands"`
	CreationTime     string `json:"creation_time"`
}

type Stream struct {
	Index              int                `json:"index"`
	CodecName          string             `json:"codec_name"`
	CodecLongName      string             `json:"codec_long_name"`
	Profile            string             `json:"profile,omitempty"`
	CodecType          string             `json:"codec_type"`
	CodecTagString     string             `json:"codec_tag_string"`
	CodecTag           string             `json:"codec_tag"`
	Width              int                `json:"width"`
	Height             int                `json:"height"`
	CodedWidth         int                `json:"coded_width"`
	CodedHeight        int                `json:"coded_height"`
	ClosedCaptions     int                `json:"closed_captions"`
	FilmGrain          int                `json:"film_grain"`
	HasBFrames         int                `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string             `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string             `json:"display_aspect_ratio,omitempty"`
	PixFmt             string             `json:"pix_fmt,omitempty"`
	Level              int                `json:"level,omitempty"`
	ChromaLocation     string             `json:"chroma_location"`
	FieldOrder         string             `json:"field_order,omitempty"`
	Refs               int                `json:"refs"`
	IsAvc              string             `json:"is_avc"`
	NalLengthSize      string             `json:"nal_length_size"`
	SampleFmt          string             `json:"sample_fmt,omitempty"`
	SampleRate         string             `json:"sample_rate,omitempty"`
	Channels           int                `json:"channels,omitempty"`
	ChannelLayout      string             `json:"channel_layout,omitempty"`
	BitsPerSample      int                `json:"bits_per_sample,omitempty"`
	ID                 string             `json:"id"`
	RFrameRate         string             `json:"r_frame_rate"`
	AvgFrameRate       string             `json:"avg_frame_rate"`
	TimeBase           string             `json:"time_base"`
	StartPts           int                `json:"start_pts"`
	StartTime          string             `json:"start_time"`
	DurationTs         uint64             `json:"duration_ts"`
	Duration           string             `json:"duration"`
	BitRate            string             `json:"bit_rate"`
	BitsPerRawSample   string             `json:"bits_per_raw_sample"`
	NbFrames           string             `json:"nb_frames"`
	NbReadFrames       string             `json:"nb_read_frames"`
	ExtradataSize      int                `json:"extradata_size"`
	Disposition        *StreamDisposition `json:"disposition,omitempty"`
	Tags               *StreamTags        `json:"tags"`
}

type StreamDisposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails"`
	Captions        int `json:"captions"`
	Descriptions    int `json:"descriptions"`
	Metadata        int `json:"metadata"`
	Dependent       int `json:"dependent"`
	StillImage      int `json:"still_image"`
}

type StreamTags struct {
	CreationTime string `json:"creation_time"`
	Language     string `json:"language"`
	HandlerName  string `json:"handler_name"`
	VendorID     string `json:"vendor_id"`
}

type Frame struct {
	MediaType               string `json:"media_type"`
	StreamIndex             int    `json:"stream_index"`
	KeyFrame                int    `json:"key_frame"`
	Pts                     int    `json:"pts"`
	PtsTime                 string `json:"pts_time"`
	PktDts                  int    `json:"pkt_dts"`
	PktDtsTime              string `json:"pkt_dts_time"`
	BestEffortTimestamp     int    `json:"best_effort_timestamp"`
	BestEffortTimestampTime string `json:"best_effort_timestamp_time"`
	PktDuration             int    `json:"pkt_duration"`
	PktDurationTime         string `json:"pkt_duration_time"`
	PktPos                  string `json:"pkt_pos"`
	PktSize                 int64  `json:"pkt_size,string"`
	Width                   int    `json:"width"`
	Height                  int    `json:"height"`
	PixFmt                  string `json:"pix_fmt"`
	SampleAspectRatio       string `json:"sample_aspect_ratio"`
	PictType                string `json:"pict_type"`
	CodedPictureNumber      int    `json:"coded_picture_number"`
	DisplayPictureNumber    int    `json:"display_picture_number"`
	InterlacedFrame         int    `json:"interlaced_frame"`
	TopFieldFirst           int    `json:"top_field_first"`
	RepeatPict              int    `json:"repeat_pict"`
	ChromaLocation          string `json:"chroma_location"`
}

func (s *Stream) FrameRate() (int, error) {
	parts := strings.Split(s.RFrameRate, "/")

	numerator, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, err
	}

	denominator, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return 0, err
	}

	return int(math.Round(numerator / denominator)), nil
}
