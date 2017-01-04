/*
go vm External API Interface
 */
package vm

type IApiService interface {
	Invoke(method string, engine *ScriptEngine) (bool, error)
}
