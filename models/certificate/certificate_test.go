package certificate

import "testing"

func TestCertiOperators(t *testing.T) {
	var _ CertiOperator = &CertificateStrategy{}
	var _ CertiOperator = &AddressCertiOperator{}
	var _ CertiOperator = &PhoneCertiOperator{}
	var _ CertiOperator = &DodoCertiOperator{}
	var _ CertiOperator = &ContractCertiOperator{}
	var _ CertiOperator = &GaslessCertiOperator{}
}
