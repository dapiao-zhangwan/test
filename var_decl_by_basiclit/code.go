package var_decl_by_basiclit

func test() {
	rawArgs := make([]string, 10)

	execIdx, execExist := 0, false

	for execIdx = range rawArgs {
		_ = execIdx
		_ = execExist
	}
}
