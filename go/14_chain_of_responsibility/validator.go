package main

type IValidator interface {
	IsValid(d *InputData) bool
	Done()
	Fail()
}

type Validator struct {
	validator IValidator
	next Handler
}

func NewValidator(v IValidator) *Validator {
	return &Validator{validator: v}
}

// テンプレートメソッド
func (v *Validator) Validate(d *InputData) bool {
	isValid := v.validator.IsValid(d)
	if isValid {
		v.validator.Done()
	} else {
		v.validator.Fail()
	}

	// バリデーションチェックの結果に関わらず次のハンドラーのバリデーション処理を呼び出している。
	// 一つでも失敗した時点でチェックをやめることも可能。
	if v.next != nil {
		return v.next.Validate(d) && isValid
	} else {
		return isValid
	}
}

func (v *Validator) SetNext(next Handler) Handler {
	v.next = next
	// 引数で受け取ったhandlerを返せば、handler1.SetNext(handler2).SetNext(handler3)のように連鎖できる。
	return next
}