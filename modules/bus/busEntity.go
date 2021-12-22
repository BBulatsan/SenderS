package bus

const (
	Operation = "*.message.operation.#"
	Sale      = "*.message.sale.#"
)

func Bindings() []string {
	Bindings := []string{
		Operation,
		Sale,
	}
	return Bindings
}
