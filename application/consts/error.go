package consts

import "errors"

var (
	ErrForbidden        = errors.New("권한 없음")
	ErrUnauthorized     = errors.New("인증 실패")
	ErrParameterMissing = errors.New("필수 파라미터 누락")
	ErrParameterInvalid = errors.New("파라미터 오류")
	ErrNotFound         = errors.New("조회결과 없음")
	ErrInternalServer   = errors.New("내부 오류")
)
