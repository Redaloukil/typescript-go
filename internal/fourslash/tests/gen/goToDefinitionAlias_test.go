package fourslash_test

import (
	"testing"

	"github.com/microsoft/typescript-go/internal/fourslash"
	"github.com/microsoft/typescript-go/internal/testutil"
)

func TestGoToDefinitionAlias(t *testing.T) {
	t.Parallel()

	defer testutil.RecoverAndFail(t, "Panic on fourslash test")
	const content = `// @Filename: b.ts
import /*alias1Definition*/alias1 = require("fileb");
module Module {
    export import /*alias2Definition*/alias2 = alias1;
}

// Type position
var t1: [|/*alias1Type*/alias1|].IFoo;
var t2: Module.[|/*alias2Type*/alias2|].IFoo;

// Value posistion
var v1 = new [|/*alias1Value*/alias1|].Foo();
var v2 = new Module.[|/*alias2Value*/alias2|].Foo();
// @Filename: a.ts
export class Foo {
    private f;
}
export interface IFoo {
    x;
}`
	f := fourslash.NewFourslash(t, nil /*capabilities*/, content)
	f.VerifyBaselineGoToDefinition(t, "alias1Type", "alias1Value", "alias2Type", "alias2Value")
}
