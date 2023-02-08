package alipay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-pay/gopay"
)

//ComplainQuery alipay.merchant.tradecomplain.query(查询单条交易投诉详情)
// 文档地址：https://opendocs.alipay.com/pre-open/035jkc
func (a *Client) ComplainQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ComplainQueryResponse, err error) {
	err = bm.CheckEmptyError("complain_event_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.tradecomplain.query"); err != nil {
		return nil, err
	}
	aliRsp = new(ComplainQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//ComplainBatchQuery alipay.merchant.tradecomplain.batchquery(查询交易投诉列表)
// 文档地址：https://opendocs.alipay.com/pre-open/035kov
func (a *Client) ComplainBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *ComplainBatchQueryResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.tradecomplain.batchquery"); err != nil {
		return nil, err
	}
	aliRsp = new(ComplainBatchQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//MerchantImageUpload alipay.merchant.image.upload(商户上传处理图片)
// 文档地址：https://opendocs.alipay.com/pre-open/035ia1
func (a *Client) MerchantImageUpload(ctx context.Context, bm gopay.BodyMap) (aliRsp *ImageUploadResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.image.upload"); err != nil {
		return nil, err
	}
	aliRsp = new(ImageUploadResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//ComplainFeedbackSubmit alipay.merchant.tradecomplain.feedback.submit(商家处理交易投诉)
// 文档地址：https://opendocs.alipay.com/pre-open/035jkd
func (a *Client) ComplainFeedbackSubmit(ctx context.Context, bm gopay.BodyMap) (aliRsp *ComplainFeedbackSubmitResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.tradecomplain.feedback.submit"); err != nil {
		return nil, err
	}
	aliRsp = new(ComplainFeedbackSubmitResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//ComplainReplySubmit alipay.merchant.tradecomplain.reply.submit(商家留言回复)
// 文档地址：https://opendocs.alipay.com/pre-open/03sy7b
func (a *Client) ComplainReplySubmit(ctx context.Context, bm gopay.BodyMap) (aliRsp *ComplainReplySubmitResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.tradecomplain.reply.submit"); err != nil {
		return nil, err
	}
	aliRsp = new(ComplainReplySubmitResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

//ComplainSupplementSubmit alipay.merchant.tradecomplain.supplement.submit(商家补充凭证)
// 文档地址：https://opendocs.alipay.com/pre-open/03t2ys
func (a *Client) ComplainSupplementSubmit(ctx context.Context, bm gopay.BodyMap) (aliRsp *ComplainSupplementSubmitResponse, err error) {
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.merchant.tradecomplain.supplement.submit"); err != nil {
		return nil, err
	}
	aliRsp = new(ComplainSupplementSubmitResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
