package builder

// インスタンスを作成するためのインターフェース(API)を宣言する
// インスタンスの各ステップのメソッドを用意する。
type iBuilder interface {
	SetHandle()
	SetTireSize()
	SetFrame()
	Build() interface{} // 様々な型に対応するため戻り値はinterface{}で定義
}

func GetBuilder(builderType string) iBuilder {
	if builderType == "bike" {
		return NewBikeBuilder()
	}
	if builderType == "manual" {
		return NewManualBuilder()
	}
	return nil
}