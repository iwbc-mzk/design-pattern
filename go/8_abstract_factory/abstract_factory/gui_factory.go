package abstract_factory

type IGuiFactory interface {
	CreateCheckBox() ICheckBox
	CreateButton() IButton
}