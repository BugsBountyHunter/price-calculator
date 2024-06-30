package io_manager
type IOManager interface{
	 ReadLines()([]string, error)
	 WriteLines(data interface{})(error)
}