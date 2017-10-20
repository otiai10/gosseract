package gosseract

// PageSegMode represents tesseract::PageSegMode.
// See https://github.com/tesseract-ocr/tesseract/blob/a18620cfea33d03032b71fe1b9fc424777e34252/ccstruct/publictypes.h#L158-L183 for more information.
type PageSegMode int

const (
	PSM_OSD_ONLY PageSegMode = iota
	PSM_AUTO_OSD

	PSM_AUTO_ONLY
	PSM_AUTO
	PSM_SINGLE_COLUMN
	PSM_SINGLE_BLOCK_VERT_TEXT

	PSM_SINGLE_BLOCK
	PSM_SINGLE_LINE
	PSM_SINGLE_WORD
	PSM_CIRCLE_WORD
	PSM_SINGLE_CHAR

	PSM_COUNT
)
