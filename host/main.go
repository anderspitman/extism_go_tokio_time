package main

import (
	"context"
	"path/filepath"
	"runtime"

	"github.com/extism/go-sdk"
	"github.com/tetratelabs/wazero"
)

func main() {
	extism.SetLogLevel(extism.LogLevelDebug)

	_, curFilePath, _, ok := runtime.Caller(0)
	if !ok {
		panic("failed to get path")
	}

	dir := filepath.Dir(curFilePath)
	wasmPath := filepath.Join(dir, "../plugin/target/wasm32-wasip1/release/plugin.wasm")

	manifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmPath,
			},
		},
	}

	moduleConfig := wazero.NewModuleConfig().
		WithSysWalltime().
		WithSysNanotime().
		WithSysNanosleep()

	ctx := context.Background()
	config := extism.PluginConfig{
		ModuleConfig: moduleConfig,
		EnableWasi:   true,
	}
	hostFunctions := []extism.HostFunction{}
	plugin, err := extism.NewPlugin(ctx, manifest, config, hostFunctions)
	if err != nil {
		panic(err)
	}

	_, _, err = plugin.Call("test", nil)
	if err != nil {
		panic(err)
	}
}
