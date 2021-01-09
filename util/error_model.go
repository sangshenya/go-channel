package util

import "errors"

type ChannelErrorType int

const(
	ChannelNameError = -1
	ChannelRequestFailError = -2
	ChannelRequestTimeoutError = -3
	ChannelRequestNoError = -4
	ChannelNoImageError = -5
	ChannelNoUrlError = -6

)

type ChannelErrorModel struct {
	Type ChannelErrorType `json:"type"`
	Error error `json:"error"`
}

type ChannelErrorProtocol interface {
	ErrorString() string
	ErrorCode() int
}

// 实现协议
func (channelErrorModel *ChannelErrorModel) ErrorString() string {
	return channelErrorModel.Error.Error()
}

func (channelErrorModel *ChannelErrorModel) ErrorCode() int {
	return int(channelErrorModel.Type)
}

// 生成渠道名称错误
func NewChannelNameError(err error) ChannelErrorProtocol {
	if err == nil {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelNameError,
		Error: err,
	}
}

func NewChannelNameErrorWithText(text string) ChannelErrorProtocol {
	if len(text) == 0 {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelNameError,
		Error: errors.New(text),
	}
}
// 生成渠道请求错误
func NewChannelRequestFailErrorError(err error) ChannelErrorProtocol {
	if err == nil {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelRequestFailError,
		Error: err,
	}
}

func NewChannelRequestFailErrorWithText(text string) ChannelErrorProtocol {
	if len(text) == 0 {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelRequestFailError,
		Error: errors.New(text),
	}
}

// 生成请求超时错误
func NewChannelRequestTimeoutError(err error) ChannelErrorProtocol {
	if err == nil {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelRequestTimeoutError,
		Error: err,
	}
}

func NewChannelRequestTimeoutErrorWithText(text string) ChannelErrorProtocol {
	if len(text) == 0 {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelRequestTimeoutError,
		Error: errors.New(text),
	}
}

// 生成请求无填充错误
func NewChannelRequestNoError(err error) ChannelErrorProtocol {
	if err == nil {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelRequestNoError,
		Error: err,
	}
}

func NewChannelRequestNoErrorWithText(text string) ChannelErrorProtocol {
	if len(text) == 0 {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelRequestNoError,
		Error: errors.New(text),
	}
}
// 生成请求无图片错误
func NewChannelNoImageError(err error) ChannelErrorProtocol {
	if err == nil {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelNoImageError,
		Error: err,
	}
}

func NewChannelNoImageErrorWithText(text string) ChannelErrorProtocol {
	if len(text) == 0 {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelNoImageError,
		Error: errors.New(text),
	}
}

// 生成请求无落地页错误
func NewChannelNoUrlError(err error) ChannelErrorProtocol {
	if err == nil {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelNoUrlError,
		Error: err,
	}
}

func NewChannelNoUrlErrorWithText(text string) ChannelErrorProtocol {
	if len(text) == 0 {
		return nil
	}
	return &ChannelErrorModel{
		Type:  ChannelNoUrlError,
		Error: errors.New(text),
	}
}