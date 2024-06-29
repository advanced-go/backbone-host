package register

import "fmt"

func ExampleIngressExchanges() {
	err := IngressExchange()

	fmt.Printf("test: IngressExchanges() -> [err:%v]\n", err)

	//Output:
	//fail
}
