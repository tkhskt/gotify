package models

// AudioAnalysis : the struct for GET https://api.spotify.com/v1/audio-analysis/{id}
type AudioAnalysis struct {
	Bars []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"bars"`
	Beats []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"beats"`
	Meta struct {
		AnalyzerVersion string  `json:"analyzer_version"`
		Platform        string  `json:"platform"`
		DetailedStatus  string  `json:"detailed_status"`
		StatusCode      int     `json:"status_code"`
		Timestamp       int     `json:"timestamp"`
		AnalysisTime    float64 `json:"analysis_time"`
		InputProcess    string  `json:"input_process"`
	} `json:"meta"`
	Sections []struct {
		Start                   float64 `json:"start"`
		Duration                float64 `json:"duration"`
		Confidence              float64 `json:"confidence"`
		Loudness                float64 `json:"loudness"`
		Tempo                   float64 `json:"tempo"`
		TempoConfidence         float64 `json:"tempo_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float64 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float64 `json:"mode_confidence"`
		TimeSignature           int     `json:"time_signature"`
		TimeSignatureConfidence float64 `json:"time_signature_confidence"`
	} `json:"sections"`
	Segments []struct {
		Start           float64   `json:"start"`
		Duration        float64   `json:"duration"`
		Confidence      float64   `json:"confidence"`
		LoudnessStart   float64   `json:"loudness_start"`
		LoudnessMaxTime float64   `json:"loudness_max_time"`
		LoudnessMax     float64   `json:"loudness_max"`
		LoudnessEnd     float64   `json:"loudness_end"`
		Pitches         []float64 `json:"pitches"`
		Timbre          []float64 `json:"timbre"`
	} `json:"segments"`
	Tatums []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"tatums"`
	Track struct {
		NumSamples              int     `json:"num_samples"`
		Duration                float64 `json:"duration"`
		SampleMd5               string  `json:"sample_md5"`
		OffsetSeconds           int     `json:"offset_seconds"`
		WindowSeconds           int     `json:"window_seconds"`
		AnalysisSampleRate      int     `json:"analysis_sample_rate"`
		AnalysisChannels        int     `json:"analysis_channels"`
		EndOfFadeIn             float64 `json:"end_of_fade_in"`
		StartOfFadeOut          float64 `json:"start_of_fade_out"`
		Loudness                float64 `json:"loudness"`
		Tempo                   float64 `json:"tempo"`
		TempoConfidence         float64 `json:"tempo_confidence"`
		TimeSignature           float64 `json:"time_signature"`
		TimeSignatureConfidence float64 `json:"time_signature_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float64 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float64 `json:"mode_confidence"`
		Codestring              string  `json:"codestring"`
		CodeVersion             float64 `json:"code_version"`
		Echoprintstring         string  `json:"echoprintstring"`
		EchoprintVersion        float64 `json:"echoprint_version"`
		Synchstring             string  `json:"synchstring"`
		SynchVersion            float64 `json:"synch_version"`
		Rhythmstring            string  `json:"rhythmstring"`
		RhythmVersion           float64 `json:"rhythm_version"`
	} `json:"track"`
}
