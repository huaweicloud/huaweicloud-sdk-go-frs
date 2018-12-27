package service

const (
	BASE64 = iota
	OBSURL
	FACEID
)

const (
	_FACE_DETECT_URI             = "/v1/%s/face-detect"
	_FACE_COMPARE_URI            = "/v1/%s/face-compare"
	_FACE_SEARCH_URI             = "/v1/%s/face-sets/%s/search"
	_FACE_ADD_URI                = "/v1/%s/face-sets/%s/faces"
	_FACE_GET_RANGE_URI          = "/v1/%s/face-sets/%s/faces?offset=%d&limit=%d"
	_FACE_GET_ONE_URI            = "/v1/%s/face-sets/%s/faces?face_id=%s"
	_FACE_DELETE_BY_EXT_ID_URI   = "/v1/%s/face-sets/%s/faces?external_image_id=%s"
	_FACE_DELETE_BY_FACE_ID_URI  = "/v1/%s/face-sets/%s/faces?face_id=%s"
	_FACE_DELETE_BY_FIELD_ID_URI = "/v1/%s/face-sets/%s/faces?%s=%s"
	_FACE_DELETE_BY_FILTER       = "/v1/%s/face-sets/%s/faces/batch"
	_FACE_SET_CREATE_URI         = "/v1/%s/face-sets"
	_FACE_SET_GET_ALL_URI        = "/v1/%s/face-sets"
	_FACE_SET_GET_ONE_URI        = "/v1/%s/face-sets/%s"
	_FACE_SET_DELETE_URI         = "/v1/%s/face-sets/%s"
	_LIVE_DETECT_URI             = "/v1/%s/live-detect"
	_FACE_QUALITY_URI            = "/v1/%s/face/quality/face-quality"
	_BLUR_CLASSIFY_URI           = "/v1/%s/face/quality/blur-classify"
	_HEAD_POSE_ESTIMATE          = "/v1/%s/face/quality/head-pose-estimate"
)
